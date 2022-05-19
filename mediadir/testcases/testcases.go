package testcases

import (
	_ "embed"
	"io/fs"
	"strconv"
	"strings"

	"github.com/spiretechnology/spireav/internal/memfs"
	"github.com/spiretechnology/spireav/mediadir"
)

//go:embed test_filenames.csv
var filenamesStr string

type TestCase struct {
	TapeType string
	Expected mediadir.Clip
	Filename string
}

func LoadTestCases() (fs.FS, []TestCase) {

	// Create the slice of test cases
	var cases []TestCase
	lines := strings.Split(filenamesStr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		cols := strings.Split(line, ",")
		if len(cols) < 4 {
			continue
		}
		clipIndex, _ := strconv.Atoi(cols[2])
		tc := TestCase{
			TapeType: cols[0],
			Expected: mediadir.Clip{
				TapeName:  cols[1],
				ClipIndex: clipIndex,
			},
			Filename: cols[3],
		}
		cases = append(cases, tc)
	}

	// Create the in-memory file system
	memFS := memfs.NewInMemoryFS()
	for _, tc := range cases {
		memFS.AddFile(tc.Filename, nil)
	}

	return memFS, cases

}
