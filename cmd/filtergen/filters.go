package main

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Filter struct {
	Name                string
	FlagTimelineSupport bool
	FlagSliceThreading  bool
	FlagCommandSupport  bool
	Description         string

	NumInputs  *int
	NumOutputs *int
	Options    []FilterOption
}

type FilterOption struct {
	Name        string
	Type        OptionType
	Description string

	Expression             bool
	InferredExpressionType OptionType

	FlagEncodingParam  bool
	FlagDecodingParam  bool
	FlagFilteringParam bool
	FlagVideoParam     bool
	FlagAudioParam     bool
	FlagSubtitleParam  bool
	FlagExport         bool
	FlagReadOnly       bool
	FlagBSFParam       bool
	FlagRuntimeParam   bool
	FlagDeprecated     bool
}

func listFilters(ctx context.Context) ([]Filter, error) {
	filterLineRegex := regexp.MustCompile(`^\s+(T|\.)(S|\.)(C|\.)\s+(\w+)\s+(\w+|\|)\->(\w+|\|)\s+(.*)$`)

	output, err := exec.CommandContext(ctx, "ffmpeg", "-filters").Output()
	if err != nil {
		return nil, err
	}

	var filters []Filter
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		matches := filterLineRegex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		match := matches[0]
		var filter Filter
		filter.FlagTimelineSupport = match[1] == "T"
		filter.FlagSliceThreading = match[2] == "S"
		filter.FlagCommandSupport = match[3] == "C"
		filter.Name = match[4]
		filter.Description = match[7]
		filters = append(filters, filter)
	}

	// Query for more info about each filter
	for i := range filters {
		if err := getFilterInfo(ctx, &filters[i]); err != nil {
			return nil, err
		}
	}

	return filters, nil
}

func getFilterInfo(ctx context.Context, filter *Filter) error {
	// Get more info about the filter
	output, err := exec.CommandContext(ctx, "ffmpeg", "--help", fmt.Sprintf("filter=%s", filter.Name)).Output()
	if err != nil {
		return err
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")
	sectionLines := make(map[string][]string)
	var section string
	for _, line := range lines {
		if line == "    Inputs:" {
			section = "inputs"
		} else if line == "    Outputs:" {
			section = "outputs"
		} else if strings.HasSuffix(line, " AVOptions:") {
			section = "options"
		} else if !strings.HasPrefix(line, "   ") {
			section = ""
		} else if section != "" && strings.HasPrefix(line, "   ") {
			sectionLines[section] = append(sectionLines[section], line)
		}
	}

	for _, line := range sectionLines["inputs"] {
		if strings.TrimSpace(line) == "dynamic (depending on the options)" {
			filter.NumInputs = nil
			break
		} else {
			if filter.NumInputs == nil {
				filter.NumInputs = new(int)
			}
			*filter.NumInputs++
		}
	}

	for _, line := range sectionLines["outputs"] {
		if strings.TrimSpace(line) == "dynamic (depending on the options)" {
			filter.NumOutputs = nil
			break
		} else if strings.TrimSpace(line) == "none (sink filter)" {
			filter.NumOutputs = new(int)
			*filter.NumOutputs = 0
			break
		} else {
			if filter.NumOutputs == nil {
				filter.NumOutputs = new(int)
			}
			*filter.NumOutputs++
		}
	}

	filterOptionRegex := regexp.MustCompile(`^\s{3}(\w+)\s+<(\w+)>\s+([A-Z\.]{11})\s+(.*)$`)

	optionsMap := make(map[FilterOption]struct{})
	optionNamesMap := make(map[string]struct{})

	for _, line := range sectionLines["options"] {
		matches := filterOptionRegex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		match := matches[0]
		var option FilterOption
		option.Name = match[1]
		option.Type = OptionType(match[2])
		option.Description = match[4]

		option.Expression = option.Type == "string" && strings.Contains(option.Description, "expression")
		if option.Expression {
			inferred := inferExpressionType(option)
			if inferred != "" {
				option.InferredExpressionType = OptionType(inferred)
			}
		}

		flags := match[3]
		option.FlagEncodingParam = flags[0] == 'E'
		option.FlagDecodingParam = flags[1] == 'D'
		option.FlagFilteringParam = flags[2] == 'F'
		option.FlagVideoParam = flags[3] == 'V'
		option.FlagAudioParam = flags[4] == 'A'
		option.FlagSubtitleParam = flags[5] == 'S'
		option.FlagExport = flags[6] == 'X'
		option.FlagReadOnly = flags[7] == 'R'
		option.FlagBSFParam = flags[8] == 'B'
		option.FlagRuntimeParam = flags[9] == 'T'
		option.FlagDeprecated = flags[10] == 'P'

		if _, ok := optionNamesMap[option.Name]; ok {
			continue
		}

		optionDedupeCopy := option
		optionDedupeCopy.Name = strings.ToLower(optionDedupeCopy.Name)

		if _, ok := optionsMap[optionDedupeCopy]; !ok {
			filter.Options = append(filter.Options, option)
			optionsMap[optionDedupeCopy] = struct{}{}
			optionNamesMap[option.Name] = struct{}{}
		}
	}

	if filter.FlagTimelineSupport {
		filter.Options = append(filter.Options, FilterOption{
			Name:        "enable",
			Type:        "boolean",
			Description: "expression to enable or disable the filter",
			Expression:  true,
		})
	}

	// Parse the inputs
	return nil
}

var (
	defaultValueRegex = regexp.MustCompile(`\s+\(default "(.+)"\)$`)
	integerRegex      = regexp.MustCompile(`^\d+$`)
	floatRegex        = regexp.MustCompile(`^\d+\.\d+$`)
	colorRegex        = regexp.MustCompile(`^0x[0-9a-fA-F]{6,8}$`)
)

func inferExpressionType(option FilterOption) string {
	defaultValueMatches := defaultValueRegex.FindStringSubmatch(option.Description)
	if len(defaultValueMatches) == 0 {
		return ""
	}
	defaultValue := defaultValueMatches[1]

	if integerRegex.MatchString(defaultValue) {
		return "int"
	}
	if floatRegex.MatchString(defaultValue) {
		return "float"
	}
	if colorRegex.MatchString(defaultValue) {
		return "color"
	}

	switch defaultValue {
	case "iw", "ih", "(in_w-out_w)/2", "(in_h-out_h)/2":
		return "int"
	}
	return ""
}
