package transform

import (
	"fmt"
	"strings"
)

// formatTransform formats a filter transform so that it conforms to the FFMpeg format
func formatTransform(filterName string, opts map[string]string) string {
	optsStr := formatTransformOptions(opts)
	if len(optsStr) == 0 {
		return filterName
	}
	return fmt.Sprintf("%s=%s", filterName, optsStr)
}

// formatTransformOptions converts a map of keys and values into a string of options
func formatTransformOptions(opts map[string]string) string {
	var keyvals []string
	for k, v := range opts {
		pair := fmt.Sprintf("%s=%s", k, formatTransformOptionValue(v))
		keyvals = append(keyvals, pair)
	}
	return strings.Join(keyvals, ":")
}

// formatTransformOptionValue formats a single value to ensure it's properly escaped and quote-wrapped
func formatTransformOptionValue(value string) string {
	if strings.Contains(value, ":") && !strings.HasPrefix(value, "'") {
		return fmt.Sprintf("'%s'", value)
	}
	return value
}
