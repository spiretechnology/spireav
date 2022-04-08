package mediadir_test

import (
	"testing"

	"github.com/spiretechnology/spireav/mediadir"
	"github.com/spiretechnology/spireav/mediadir/sonyxd"
	"github.com/spiretechnology/spireav/mediadir/testcases"
)

func TestMultiFormat(t *testing.T) {

	// Create the multiformat
	multi := mediadir.MultiFormatClipParser(
		mediadir.ClipParserFunc(sonyxd.ParseXDPath),
	)

	memFS, cases := testcases.LoadTestCases()
	for _, tc := range cases {
		clip, err := multi.ParseClip(memFS, tc.Filename)
		expectedNil := tc.TapeType == ""
		if clip == nil && !expectedNil {
			if err != nil {
				t.Log("Got error: ", err.Error())
			}
			t.Fatalf("Got nil, expected %s.%d: %s", tc.Expected.TapeName, tc.Expected.ClipIndex, tc.Filename)
		}
		if clip == nil && expectedNil {
			continue
		}
		if clip.ClipIndex != tc.Expected.ClipIndex || clip.TapeName != tc.Expected.TapeName {
			t.Fatalf("Got %s.%d, expected %s.%d: %s", clip.TapeName, clip.ClipIndex, tc.Expected.TapeName, tc.Expected.ClipIndex, tc.Filename)
		}
	}

}
