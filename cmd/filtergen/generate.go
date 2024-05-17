package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// optionMappings creates better Go method names for some filter options that are ambiguous.
var optionMappings = map[string]map[string]string{
	"hue": {
		"h": "hue_degrees",
		"H": "hue_radians",
		"s": "saturation",
		"b": "brightness",
	},
}

func generateFilters(filters []Filter) error {
	for _, filter := range filters {
		// Exclude sink filters (abuffersink, buffersink)
		if filter.NumOutputs != nil && *filter.NumOutputs == 0 {
			continue
		}
		if err := generateFilter(filter); err != nil {
			return fmt.Errorf("generating filter: %w", err)
		}
	}
	return nil
}

func generateFilter(filter Filter) error {
	// Create the filter package directory
	dir := filepath.Join(FilterDir, "filters")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating filter package directory: %w", err)
	}

	// Create the filter Go file
	filename := filepath.Join(dir, fmt.Sprintf("%s.gen.go", filter.Name))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating filter package file: %w", err)
	}
	defer file.Close()

	// Write the filter package file
	if err := writeFilterFile(file, filter); err != nil {
		return fmt.Errorf("writing filter package file: %w", err)
	}
	return nil
}

func writeFilterFile(file io.Writer, filter Filter) error {
	_, err := fmt.Fprint(file, "// Code generated by cmd/filtergen. DO NOT EDIT.\n\n")
	if err != nil {
		return fmt.Errorf("writing package declaration: %w", err)
	}

	packageName := "filters"
	funcName := ucfirst(snakeToCamel(filter.Name))

	_, err = fmt.Fprintf(file, "package %s\n\n", packageName)
	if err != nil {
		return fmt.Errorf("writing package declaration: %w", err)
	}

	// Define the imports
	localImports := []string{
		"github.com/spiretechnology/spireav/filter",
		"github.com/spiretechnology/spireav/filter/expr",
	}
	stdlibImports := []string{}

	// Compile a slice of all the option methods
	var allMethods []optionMethod
	for _, option := range filter.Options {
		if option.Name == "outputs" {
			continue
		}
		if mappings, ok := optionMappings[filter.Name]; ok {
			if mapping, ok := mappings[option.Name]; ok && mapping == "" {
				continue
			}
		}
		allMethods = append(allMethods, optionToMethods(filter, option)...)
	}

	// Deduplicate the methods
	var methods []optionMethod
	methodsAdded := make(map[string]struct{})
	for i, method := range allMethods {
		if _, ok := methodsAdded[method.funcName]; ok {
			continue
		}
		var hasHigherPriority bool
		for j, other := range allMethods {
			if i != j && other.funcName == method.funcName && other.priority > method.priority {
				hasHigherPriority = true
				break
			}
		}
		if !hasHigherPriority {
			methods = append(methods, method)
			methodsAdded[method.funcName] = struct{}{}
		}
	}

	// Combine the imports from all the methods
	for _, method := range methods {
		for _, im := range method.imports {
			var importGroup *[]string
			if im == "github.com/spiretechnology/spireav" || strings.HasPrefix(im, "github.com/spiretechnology/spireav/") {
				importGroup = &localImports
			} else {
				importGroup = &stdlibImports
			}
			if importGroup != nil && !slices.Contains(*importGroup, im) {
				*importGroup = append(*importGroup, im)
			}
		}
	}

	// Combine the import groups
	importGroups := [][]string{}
	if len(stdlibImports) > 0 {
		importGroups = append(importGroups, stdlibImports)
	}
	if len(localImports) > 0 {
		importGroups = append(importGroups, localImports)
	}

	// Write the imports
	if len(importGroups) > 0 {
		fmt.Fprintf(file, "import (\n")
		for i, group := range importGroups {
			sort.Strings(group)
			if i > 0 {
				fmt.Fprintf(file, "\n")
			}
			for _, im := range group {
				fmt.Fprintf(file, "\t%q\n", im)
			}
		}
		fmt.Fprintf(file, ")\n")
	}

	builderInterfaceName := fmt.Sprintf("%sBuilder", funcName)
	builderImplName := fmt.Sprintf("impl%sBuilder", funcName)

	// Define the interface for the filter builder
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "// %s %s.\n", builderInterfaceName, strings.TrimSuffix(filter.Description, "."))
	fmt.Fprintf(file, "// Documentation: %s\n", fmt.Sprintf("https://ffmpeg.org/ffmpeg-filters.html#%s", filter.Name))
	fmt.Fprintf(file, "type %s interface {\n", builderInterfaceName)
	fmt.Fprintf(file, "\tfilter.Filter\n")
	for _, method := range methods {
		fmt.Fprintf(file, "\t// %s\n", method.comment)
		fmt.Fprintf(file, "\t%s %s\n", method.header, builderInterfaceName)
	}
	fmt.Fprintf(file, "}\n")

	// If the filter has a configurable number of outputs
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "// %s creates a new %s to configure the %q filter.\n", funcName, builderInterfaceName, filter.Name)
	if filter.NumOutputs == nil {
		// Define the function to create the filter builder
		fmt.Fprintf(file, "func %s(outputs int, opts ...filter.Option) %s {\n", funcName, builderInterfaceName)
		fmt.Fprintf(file, "\tf := filter.New(%q, outputs, opts...)\n", filter.Name)
		fmt.Fprintf(file, "\tf = f.With(\"outputs\", expr.Int(outputs))\n")
		fmt.Fprintf(file, "\treturn &%s{f: f}\n", builderImplName)
		fmt.Fprintf(file, "}\n")
	} else {
		// Define the function to create the filter builder
		fmt.Fprintf(file, "func %s(opts ...filter.Option) %s {\n", funcName, builderInterfaceName)
		fmt.Fprintf(file, "\treturn &%s{\n", builderImplName)
		fmt.Fprintf(file, "\t\tf: filter.New(%q, %d, opts...),\n", filter.Name, *filter.NumOutputs)
		fmt.Fprintf(file, "\t}\n")
		fmt.Fprintf(file, "}\n")
	}

	receiver := fmt.Sprintf("%sBuilder", filter.Name)

	// Define the implementation of the filter builder
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "type %s struct {\n", builderImplName)
	fmt.Fprintf(file, "\tf filter.Filter\n")
	fmt.Fprintf(file, "}\n")

	// Define methods to satisfy filter.Filter interface
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (%s *%s) String() string {\n", receiver, builderImplName)
	fmt.Fprintf(file, "\treturn %s.f.String()\n", receiver)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (%s *%s) Outputs() int {\n", receiver, builderImplName)
	fmt.Fprintf(file, "\treturn %s.f.Outputs()\n", receiver)
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (%s *%s) With(key string, value expr.Expr) filter.Filter {\n", receiver, builderImplName)
	fmt.Fprintf(file, "\treturn %s.withOption(key, value)\n", receiver)
	fmt.Fprintf(file, "}\n")

	// Helper method to reduce the amount of code duplication
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (%s *%s) withOption(key string, value expr.Expr) %s {\n", receiver, builderImplName, builderInterfaceName)
	fmt.Fprintf(file, "\tbCopy := *%s\n", receiver)
	fmt.Fprintf(file, "\tbCopy.f = %s.f.With(key, value)\n", receiver)
	fmt.Fprintf(file, "\treturn &bCopy\n")
	fmt.Fprintf(file, "}\n")

	// Define methods to satisfy the filter builder interface (all the options)
	for _, method := range methods {
		fmt.Fprintf(file, "\n")
		fmt.Fprintf(file, "func (%s *%s) %s %s {\n",
			receiver,
			builderImplName,
			method.header,
			builderInterfaceName,
		)
		fmt.Fprintf(file, "\treturn %s.withOption(%q, %s)\n", receiver, method.option.Name, method.toExpr)
		fmt.Fprintf(file, "}\n")
	}

	return nil
}

func ucfirst(str string) string {
	if len(str) < 2 {
		return strings.ToUpper(str)
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

type optionMethod struct {
	funcName string
	option   FilterOption
	comment  string
	header   string
	toExpr   string
	imports  []string
	priority int
}

func optionToMethods(filter Filter, option FilterOption) []optionMethod {
	var methods []optionMethod

	// The ffmpeg CLI doesn't tell us for sure if an option supports expression evaluation or not.
	// We have to guess based on the option type and description.
	// So we have three cases:
	// 1. The option is definitely an expression.
	// 2. The option might be an expression.
	// 3. The option is definitely not an expression.

	if option.Expression {
		// If there is an inferred type
		if option.InferredExpressionType != "" {
			copiedOption := option
			copiedOption.Type = option.InferredExpressionType
			inferredExprMethod, err := optionToMethod(filter, copiedOption, "")
			if err != nil {
				log.Printf("Filter (%s): %s", filter.Name, err)
			} else {
				inferredExprMethod.priority = 0
				methods = append(methods, inferredExprMethod)
			}
		}
		// If it's definitely an expression, add one method:
		// - "x" => X(expr.Expr)
		suffix := ""
		if option.InferredExpressionType != "" {
			suffix = "Expr"
		}
		exprMethod, err := optionExpressionToMethod(filter, option, suffix)
		if err != nil {
			log.Printf("Filter (%s): %s", filter.Name, err)
		} else {
			exprMethod.priority = 1
			methods = append(methods, exprMethod)
		}
	} else if option.FlagRuntimeParam {
		// If it's maybe an expression, add both. For example:
		// - Fontcolor(str string)
		// - FontcolorExpr(expr.Expr)
		baseMethod, err := optionToMethod(filter, option, "")
		if err != nil {
			log.Printf("Filter (%s): %s", filter.Name, err)
		} else {
			baseMethod.priority = 0
			methods = append(methods, baseMethod)
		}
		exprMethod, err := optionExpressionToMethod(filter, option, "Expr")
		if err != nil {
			log.Printf("Filter (%s): %s", filter.Name, err)
		} else {
			exprMethod.priority = 0
			methods = append(methods, exprMethod)
		}
	} else {
		// If it's definitely not an expression
		baseMethod, err := optionToMethod(filter, option, "")
		if err != nil {
			log.Printf("Filter (%s): %s", filter.Name, err)
		} else {
			baseMethod.priority = 1
			methods = append(methods, baseMethod)
		}
	}
	return methods
}

func optionToMethod(filter Filter, option FilterOption, suffix string) (optionMethod, error) {
	funcName := cleanFuncName(filter.Name, option.Name) + suffix
	argName := cleanArgName(snakeToCamel(option.Name))

	goType, imports, err := option.Type.GoType()
	if err != nil {
		return optionMethod{}, err
	}
	toExpr, err := option.Type.ToExpr(argName)
	if err != nil {
		return optionMethod{}, err
	}

	var method optionMethod
	method.funcName = funcName
	method.option = option
	method.comment = fmt.Sprintf("%s %s.", funcName, option.Description)
	method.header = fmt.Sprintf("%s(%s %s)", funcName, argName, goType)
	method.toExpr = toExpr
	method.imports = imports
	return method, nil
}

func optionExpressionToMethod(filter Filter, option FilterOption, suffix string) (optionMethod, error) {
	funcName := cleanFuncName(filter.Name, option.Name) + suffix
	argName := cleanArgName(snakeToCamel(option.Name))

	goType := "expr.Expr"
	toExpr := argName

	var method optionMethod
	method.funcName = funcName
	method.option = option
	method.comment = fmt.Sprintf("%s %s.", funcName, strings.TrimSuffix(option.Description, "."))
	method.header = fmt.Sprintf("%s(%s %s)", funcName, argName, goType)
	method.toExpr = toExpr
	method.imports = []string{}
	return method, nil
}

type OptionType string

func (t OptionType) GoType() (string, []string, error) {
	switch t {
	case "int":
		return "int", nil, nil
	case "int64":
		return "int64", nil, nil
	case "float":
		return "float32", nil, nil
	case "double":
		return "float64", nil, nil
	case "duration":
		return "time.Duration", []string{"time"}, nil
	case "boolean":
		return "bool", nil, nil
	case "string":
		return "string", nil, nil
	case "image_size":
		return "expr.Size", nil, nil
	case "rational", "video_rate":
		return "expr.Rational", nil, nil
	case "color":
		return "expr.Color", nil, nil
	case "pix_fmt":
		return "expr.PixFmt", nil, nil
	case "sample_fmt":
		return "expr.SampleFmt", nil, nil
	case "channel_layout":
		return "expr.ChannelLayout", nil, nil
	case "dictionary":
		return "expr.Dictionary", nil, nil
	case "flags":
		return "...string", nil, nil
	default:
		return "", nil, fmt.Errorf("unsupported option type: %s", t)
	}
}

func (t OptionType) ToExpr(name string) (string, error) {
	switch t {
	case "int":
		return fmt.Sprintf("expr.Int(%s)", name), nil
	case "int64":
		return fmt.Sprintf("expr.Int64(%s)", name), nil
	case "float":
		return fmt.Sprintf("expr.Float(%s)", name), nil
	case "double":
		return fmt.Sprintf("expr.Double(%s)", name), nil
	case "duration":
		return fmt.Sprintf("expr.Duration(%s)", name), nil
	case "boolean":
		return fmt.Sprintf("expr.Bool(%s)", name), nil
	case "string":
		return fmt.Sprintf("expr.String(%s)", name), nil
	case "image_size":
		return name, nil
	case "rational", "video_rate":
		return name, nil
	case "color":
		return name, nil
	case "pix_fmt":
		return name, nil
	case "sample_fmt":
		return name, nil
	case "channel_layout":
		return name, nil
	case "dictionary":
		return name, nil
	case "flags":
		return fmt.Sprintf("expr.Flags(%s)", name), nil
	default:
		return "", fmt.Errorf("unsupported option type: %s", t)
	}
}

func cleanArgName(name string) string {
	if name[0] >= '0' && name[0] <= '9' {
		return fmt.Sprintf("val%s", name)
	}
	switch name {
	case "type":
		return "typ"
	case "range":
		return "rng"
	case "map":
		return "mapVal"
	case "expr":
		return "expression"
	default:
		return name
	}
}

func cleanFuncName(filterName string, name string) string {
	if mappings, ok := optionMappings[filterName]; ok {
		if newName, ok := mappings[name]; ok {
			name = newName
		}
	}
	name = snakeToCamel(name)
	if name[0] >= '0' && name[0] <= '9' {
		return fmt.Sprintf("With%s", name)
	}
	return ucfirst(name)
}

func snakeToCamel(name string) string {
	parts := strings.Split(name, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = ucfirst(parts[i])
	}
	return strings.Join(parts, "")
}
