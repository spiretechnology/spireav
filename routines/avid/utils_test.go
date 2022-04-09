package avid_test

import "path/filepath"

const testFilesDir = "../../test-files"

func testMedia(relPath string) string {
	return filepath.Join(testFilesDir, relPath)
}
