package sonyxd

import (
	"testing"

	"github.com/spiretechnology/spireav/mediadir/testcases"
)

func TestParseXDPath(t *testing.T) {

	memFS, cases := testcases.LoadTestCases()
	for _, tc := range cases {
		clip, err := ParseXDPath(memFS, tc.Filename)
		expectedNil := tc.TapeType != "sonyxd"
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
