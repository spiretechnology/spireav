package avid_test

import (
	"testing"
)

const avidProxyFilename = "../../reference-media/SS0815KB.01/SS0815KBV01.D97B84568C1FF8V.mxf"

func TestParseAvidMxfOpAtomMeta(t *testing.T) {

	// // Get the metadata for the file
	// metaRaw, err := spireav.GetMetadata(context.Background(), avidProxyFilename)
	// if err != nil {
	// 	t.Fatal("Failed to get meta: ", err)
	// }

	// // Parse the Avid-specific metadata for the file
	// meta, err := avid.ParseAvidMxfOpAtomMeta(metaRaw)
	// if err != nil {
	// 	t.Fatal("Failed to parse AVID meta: ", err)
	// }

	// log.Printf("%+v", meta)

	// // Check the metadata
	// assert.Equal(t, "0x060A2B340101010501010F1013000000568C1FF887119905F515003EE1CE9497", meta.FilePackageUmid)
	// assert.Equal(t, "SS0815KB", meta.ReelName)
	// assert.Equal(t, "0x060A2B340101010501010F101300000055455298871199055DAD003EE1CE9497", meta.ReelUmid)
	// assert.Equal(t, "SS0815KB.06.new.01", meta.MaterialPackageName)
	// assert.Equal(t, "0x060A2B340101010501010F1013000000568C1FF187119905A7C0003EE1CE9497", meta.MaterialPackageUmid)
	// assert.Equal(t, "Avid Media Composer 18.12.5.51845", meta.MediaComposerVersion)
	// assert.Equal(t, "15:21:37:14", meta.Timecode)
	// assert.NotNil(t, meta.EssenceStream)

}
