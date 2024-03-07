package main

import (
	"fmt"
	"log"
	"strings"
)

type Config struct {
	Filters []Filter `yaml:"filters"`
}

type Filter struct {
	Name          string         `yaml:"name"`
	Comment       *string        `yaml:"comment"`
	Link          *string        `yaml:"link"`
	Func          string         `yaml:"func"`
	Outputs       any            `yaml:"outputs"`
	OutputsOption *string        `yaml:"outputs_option"`
	Options       []FilterOption `yaml:"options"`
	ExtraMethods  []string       `yaml:"extra_methods"`
}

func (f Filter) GetDocsLink() string {
	if f.Link != nil {
		return *f.Link
	}
	return fmt.Sprintf("https://ffmpeg.org/ffmpeg-filters.html#%s", f.Name)
}

func (f Filter) CommentString() string {
	if f.Comment != nil {
		return *f.Comment
	}
	return fmt.Sprintf("corresponds to the %q FFmpeg filter", f.Name)
}

type FilterOutputs struct {
	Outputs      int
	Configurable bool
}

func (f Filter) DecodeOutputs() (*FilterOutputs, error) {
	// If no outputs are defined, default to 1 output
	if f.Outputs == nil {
		return &FilterOutputs{Outputs: 1}, nil
	}
	switch outputs := f.Outputs.(type) {
	// If it's an int, return it
	case int:
		return &FilterOutputs{Outputs: outputs}, nil
	// If it's a string...
	case string:
		switch outputs {
		case "configurable":
			return &FilterOutputs{Configurable: true}, nil
		}
	}
	return nil, fmt.Errorf("invalid outputs value: %v", f.Outputs)
}

type FilterOption struct {
	Name    string   `yaml:"name"`
	Key     string   `yaml:"key"`
	Value   *string  `yaml:"value"`
	Type    []string `yaml:"type"`
	Comment *string  `yaml:"comment"`
}

func (fo *FilterOption) CommentString() string {
	if fo.Comment != nil {
		return *fo.Comment
	}
	return fmt.Sprintf("sets the %q option on the filter", fo.Key)
}

type FilterOptionMethod struct {
	Option    FilterOption
	Name      string
	Args      string
	ExprValue string
}

type OptionType struct {
	GoType   string
	ExprCast string
	Suffix   string
}

// optionTypes maps type values from the config file to Go types and expr.Expr casting functions
// that are used in the generated Go code.
var optionTypes = map[string]OptionType{
	"expr":   {"expr.Expr", "", "Expr"},
	"int":    {"int", "expr.Int", "Int"},
	"string": {"string", "expr.String", "String"},
	"bool":   {"bool", "expr.Bool", "Bool"},
	"size":   {"", "expr.Size", "Size"},
}

func (fo FilterOption) GetMethods() []FilterOptionMethod {
	// If there is a hard-coded value for the option
	if fo.Value != nil {
		return []FilterOptionMethod{
			{
				Option:    fo,
				Name:      fmt.Sprintf("%s%s", MethodPrefix, fo.Name),
				Args:      "",
				ExprValue: fmt.Sprintf("%q", *fo.Value),
			},
		}
	}

	// Convert the name to camel case. This will be the method argument name.
	camelCaseName := strings.ToLower(fo.Name[:1]) + fo.Name[1:]

	// Create one method per type
	var methods []FilterOptionMethod
	for typeIndex, typeKey := range fo.Type {
		optionType, ok := optionTypes[typeKey]
		if !ok {
			log.Fatalf("unknown option type: %s", typeKey)
		}

		var method FilterOptionMethod
		method.Option = fo
		if typeIndex == 0 {
			method.Name = fmt.Sprintf("%s%s", MethodPrefix, fo.Name)
		} else {
			method.Name = fmt.Sprintf("%s%s%s", MethodPrefix, fo.Name, optionType.Suffix)
		}
		if typeKey == "size" {
			method.Args = "width, height int"
			method.ExprValue = "expr.Size(width, height)"
		} else {
			method.Args = fmt.Sprintf("%s %s", camelCaseName, optionType.GoType)
			method.ExprValue = camelCaseName
			if optionType.ExprCast != "" {
				method.ExprValue = fmt.Sprintf("%s(%s)", optionType.ExprCast, camelCaseName)
			}
		}
		methods = append(methods, method)
	}
	return methods
}
