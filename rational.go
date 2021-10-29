package spireav

import (
	"errors"
	"strconv"
	"strings"
)

// RationalToFloat converts a rational string (ex. "24000/1001") to a float64 value
// that is equivalent to the rational string
func RationalToFloat(rational string) (float64, error) {

	// Split it up by slash
	parts := strings.SplitN(rational, "/", 2)
	if len(parts) != 2 {
		return 0, errors.New("invalid rational string")
	}

	// Parse both of the halves
	num, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}
	den, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	// Divide the numbers
	return float64(num) / float64(den), nil

}
