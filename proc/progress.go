package proc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ProgressEstimate struct {
	Percent          float64
	SecondsRemaining float64
}

type Progress struct {
	Frame    int
	FPS      int
	SizeKB   int
	Time     string
	Bitrate  int
	Speed    float64
	Done     bool
	Estimate *ProgressEstimate
}

func EmptyProgress() *Progress {
	return &Progress{
		Frame:   0,
		FPS:     0,
		SizeKB:  0,
		Time:    "00:00:00.00",
		Bitrate: 0,
		Speed:   0,
		Done:    false,
	}
}

func ParseProgressLine(
	line string,
	estimatedLengthFrames int,
) *Progress {

	// If the line doesn't start with frame=, it's not a progress line
	if !strings.HasPrefix(line, "frame=") {
		return nil
	}

	// Parse the progress line into a map of keys and values
	values := progressLineToMap(line)

	// Check if there is a Lsize (last size) key
	// The final progress output uses key Lsize instead of size, so we need to
	// handle this situation
	sizeKey := "size"
	done := false
	if _, ok := values["Lsize"]; ok {
		sizeKey = "Lsize"
		done = true
	}

	// Return the progress value
	prog := &Progress{
		Frame:   int(parseProgressValue(values, "frame", "", 0)),
		FPS:     int(parseProgressValue(values, "fps", "", 0)),
		SizeKB:  int(parseProgressValue(values, sizeKey, "kB", 0)),
		Time:    values["time"],
		Bitrate: int(parseProgressValue(values, "bitrate", "kbits/s", 0) * 1000),
		Speed:   parseProgressValue(values, "speed", "x", 0),
		Done:    done,
	}

	// If there is an estimated duration
	if estimatedLengthFrames > 0 {
		prog.Estimate = &ProgressEstimate{
			Percent:          float64(prog.Frame) / float64(estimatedLengthFrames),
			SecondsRemaining: float64(estimatedLengthFrames-prog.Frame) / math.Max(1, float64(prog.FPS)),
		}
	}

	// Return the progress
	return prog

}

func parseProgressValue(
	values map[string]string,
	key string,
	suffix string,
	defaultValue float64,
) float64 {
	val, ok := values[key]
	if !ok {
		return defaultValue
	}
	val = strings.TrimSuffix(val, suffix)
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Println(err.Error())
		return defaultValue
	}
	return f
}

func progressLineToMap(line string) map[string]string {
	result := map[string]string{}
	simplify := regexp.MustCompile(`=\s+`)
	line = simplify.ReplaceAllString(line, "=")
	spaces := regexp.MustCompile(`\s+`)
	parts := spaces.Split(line, -1)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) == 0 {
			continue
		}
		keyval := strings.SplitN(part, "=", 2)
		if len(keyval) < 2 {
			continue
		}
		result[keyval[0]] = keyval[1]
	}
	return result
}
