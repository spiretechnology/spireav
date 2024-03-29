package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	// FilterDir is the directory where the filtergen tool will generate the filter packages.
	FilterDir = "filter"

	// FilterConfigFile is the file where the filtergen tool will read the filter configuration from.
	FilterConfigFile = "cmd/filtergen/filters.yml"

	// MethodPrefix is the prefix that will be added to the filter option methods.
	// We've used "With" before, but no prefix is cleaner.
	MethodPrefix = ""
)

func main() {
	// Load the filters config file
	config, err := loadConfig()
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	// Generate the filter packages
	if err := generateFilters(config); err != nil {
		log.Fatal("error generating filters: ", err)
	}
}

func loadConfig() (*Config, error) {
	file, err := os.Open(FilterConfigFile)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("decoding config file: %w", err)
	}
	return &config, nil
}

func generateFilters(config *Config) error {
	for _, filter := range config.Filters {
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

	_, err = fmt.Fprintf(file, "package %s\n\n", packageName)
	if err != nil {
		return fmt.Errorf("writing package declaration: %w", err)
	}

	imports := []string{
		"github.com/spiretechnology/spireav/filter",
		"github.com/spiretechnology/spireav/filter/expr",
	}
	fmt.Fprintf(file, "import (\n")
	for _, im := range imports {
		fmt.Fprintf(file, "\t%q\n", im)
	}
	fmt.Fprintf(file, ")\n")

	builderInterfaceName := fmt.Sprintf("%sBuilder", filter.Func)
	builderImplName := fmt.Sprintf("impl%sBuilder", filter.Func)

	// Compile a slice of all the option methods
	var methods []FilterOptionMethod
	for _, option := range filter.Options {
		methods = append(methods, option.GetMethods()...)
	}

	// Define the interface for the filter builder
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "// %s %s.\n", builderInterfaceName, filter.CommentString())
	fmt.Fprintf(file, "// Documentation: %s\n", filter.GetDocsLink())
	fmt.Fprintf(file, "type %s interface {\n", builderInterfaceName)
	fmt.Fprintf(file, "\tfilter.Filter\n")
	for _, method := range methods {
		fmt.Fprintf(file, "\t// %s %s.\n", method.Name, method.Option.CommentString())
		fmt.Fprintf(file, "\t%s(%s) %s\n", method.Name, method.Args, builderInterfaceName)
	}
	for _, method := range filter.ExtraMethods {
		fmt.Fprintf(file, "\t%s\n", method)
	}
	fmt.Fprintf(file, "}\n")

	// If the filter has a configurable number of outputs
	outputs, err := filter.DecodeOutputs()
	if err != nil {
		return fmt.Errorf("decoding filter outputs: %w", err)
	}
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "// %s creates a new %s to configure the %q filter.\n", filter.Func, builderInterfaceName, filter.Name)
	if outputs.Configurable {
		// Define the function to create the filter builder
		fmt.Fprintf(file, "func %s(outputs int, opts ...filter.Option) %s {\n", filter.Func, builderInterfaceName)
		fmt.Fprintf(file, "\tf := filter.New(%q, outputs, opts...)\n", filter.Name)
		if filter.OutputsOption != nil {
			fmt.Fprintf(file, "\tf = f.With(%q, expr.Int(outputs))\n", *filter.OutputsOption)
		}
		fmt.Fprintf(file, "\treturn &%s{f: f}\n", builderImplName)
		fmt.Fprintf(file, "}\n")

	} else {
		// Define the function to create the filter builder
		fmt.Fprintf(file, "func %s(opts ...filter.Option) %s {\n", filter.Func, builderInterfaceName)
		fmt.Fprintf(file, "\treturn &%s{\n", builderImplName)
		fmt.Fprintf(file, "\t\tf: filter.New(%q, %d, opts...),\n", filter.Name, outputs.Outputs)
		fmt.Fprintf(file, "\t}\n")
		fmt.Fprintf(file, "}\n")
	}

	// Define the implementation of the filter builder
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "type %s struct {\n", builderImplName)
	fmt.Fprintf(file, "\tf filter.Filter\n")
	fmt.Fprintf(file, "}\n")

	// Define methods to satisfy filter.Filter interface
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (b *%s) String() string {\n", builderImplName)
	fmt.Fprintf(file, "\treturn b.f.String()\n")
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (b *%s) Outputs() int {\n", builderImplName)
	fmt.Fprintf(file, "\treturn b.f.Outputs()\n")
	fmt.Fprintf(file, "}\n")

	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (b *%s) With(key string, value expr.Expr) filter.Filter {\n", builderImplName)
	fmt.Fprintf(file, "\treturn b.withOption(key, value)\n")
	fmt.Fprintf(file, "}\n")

	// Helper method to reduce the amount of code duplication
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (b *%s) withOption(key string, value expr.Expr) %s {\n", builderImplName, builderInterfaceName)
	fmt.Fprintf(file, "\tbCopy := *b\n")
	fmt.Fprintf(file, "\tbCopy.f = b.f.With(key, value)\n")
	fmt.Fprintf(file, "\treturn &bCopy\n")
	fmt.Fprintf(file, "}\n")

	// Define methods to satisfy the filter builder interface (all the options)
	for _, method := range methods {
		fmt.Fprintf(file, "\n")
		fmt.Fprintf(file, "func (b *%s) %s(%s) %s {\n",
			builderImplName,
			method.Name,
			method.Args,
			builderInterfaceName,
		)
		fmt.Fprintf(file, "\treturn b.withOption(%q, %s)\n", method.Option.Key, method.ExprValue)
		fmt.Fprintf(file, "}\n")
	}

	return nil
}
