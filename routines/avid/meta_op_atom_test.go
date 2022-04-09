package avid_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/routines/avid"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	filename string
	meta     avid.AvidMxfMeta
}

var testCases = []testCase{
	{
		filename: "../../test-files/avid/V01.DE767F09_179DE179DE134V.mxf",
		meta: avid.AvidMxfMeta{
			FilePackageUmid:      "0x060A2B340101010501010F1013000000179DE1349678990527A6147DDAC0824E",
			ReelName:             "output.mp4",
			ReelUmid:             "0x060A2B340101010501010F10130000001780D0799678990545D7147DDAC0824E",
			MaterialPackageName:  "output.new.01",
			MaterialPackageUmid:  "0x060A2B340101010501010F1013000000179DE12C96789905C680147DDAC0824E",
			MediaComposerVersion: "Avid Media Composer 21.12.0.55552",
			Timecode:             "00:00:00:00",
			EssenceStream: &spireav.StreamMeta{
				Index:     0,
				CodecType: "video",
			},
		},
	},
	{
		filename: "../../test-files/avid/A01.DE767F0A_179DE179DE147A.mxf",
		meta: avid.AvidMxfMeta{
			FilePackageUmid:     "0x060A2B340101010501010F1013000000179DE14796789905971C147DDAC0824E",
			ReelName:            "output.mp4",
			ReelUmid:            "0x060A2B340101010501010F10130000001780D0799678990545D7147DDAC0824E",
			MaterialPackageName: "output.new.01",
			MaterialPackageUmid: "0x060A2B340101010501010F1013000000179DE12C96789905C680147DDAC0824E",
			// MediaComposerVersion: "Avid Media Composer 21.12.0.55552",
			Timecode: "00:00:00:00",
			EssenceStream: &spireav.StreamMeta{
				Index:     1,
				CodecType: "audio",
			},
		},
	},
	{
		filename: "../../test-files/avid/A02.DE767F0B_179DE179DE157A.mxf",
		meta: avid.AvidMxfMeta{
			FilePackageUmid:     "0x060A2B340101010501010F1013000000179DE15796789905004D147DDAC0824E",
			ReelName:            "output.mp4",
			ReelUmid:            "0x060A2B340101010501010F10130000001780D0799678990545D7147DDAC0824E",
			MaterialPackageName: "output.new.01",
			MaterialPackageUmid: "0x060A2B340101010501010F1013000000179DE12C96789905C680147DDAC0824E",
			// MediaComposerVersion: "Avid Media Composer 21.12.0.55552",
			Timecode: "00:00:00:00",
			EssenceStream: &spireav.StreamMeta{
				Index:     2,
				CodecType: "audio",
			},
		},
	},
}

func TestParseAvidMxfOpAtomMeta(t *testing.T) {

	for _, tc := range testCases {

		// Get the metadata for the file
		metaRaw, err := spireav.GetMetadata(context.Background(), tc.filename)
		if err != nil {
			t.Fatal("Failed to get meta: ", err)
		}

		// Parse the Avid-specific metadata for the file
		meta, err := avid.ParseAvidMxfOpAtomMeta(metaRaw)
		if err != nil {
			t.Fatal("Failed to parse AVID meta: ", err)
		}

		assert.Equal(t, tc.meta.FilePackageUmid, meta.FilePackageUmid, "file package umid: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.ReelName, meta.ReelName, "reel name: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.ReelUmid, meta.ReelUmid, "reel umid: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.MaterialPackageName, meta.MaterialPackageName, "material package name: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.MaterialPackageUmid, meta.MaterialPackageUmid, "material package umid: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.MediaComposerVersion, meta.MediaComposerVersion, "media composer version: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.Timecode, meta.Timecode, "timecode: %s", filepath.Base(tc.filename))
		assert.NotNil(t, meta.EssenceStream, "essence stream nil: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.EssenceStream.Index, meta.EssenceStream.Index, "essence stream index: %s", filepath.Base(tc.filename))
		assert.Equal(t, tc.meta.EssenceStream.CodecType, meta.EssenceStream.CodecType, "essence stream codec: %s", filepath.Base(tc.filename))
	}

}
