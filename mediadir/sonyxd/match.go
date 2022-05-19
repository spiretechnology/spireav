package sonyxd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spiretechnology/spireav/mediadir"
)

// extOptions are the valid filename extension values for Sony XD clips
var extOptions = []string{
	".mxf",
	".mp4",
	".mov",
}

// xdrootNames are the acceptible names for the "ROOT" dir found in XD-type directories
var xdrootNames = []string{
	"XDROOT",
	"M4ROOT",
}

func isExtValid(filename string) bool {
	ext := strings.ToUpper(filepath.Ext(filename))
	for _, extOption := range extOptions {
		if strings.ToUpper(extOption) == ext {
			return true
		}
	}
	return false
}

func isXdrootDirname(part string) bool {
	partUpper := strings.ToUpper(part)
	for _, xdrootName := range xdrootNames {
		if partUpper == strings.ToUpper(xdrootName) {
			return true
		}
	}
	return false
}

func ParseXDPath(fsys fs.FS, filename string) (*mediadir.Clip, error) {

	// GoPro
	// .../SDRR08GA/DCIM/100GOPRO/SDRR08GA_GH010192.MP4

	// XDCam
	// .../RR08KC/Clip/C0001.MXF
	// .../SA7RR08MA/PRIVATE/M4ROOT/CLIP/SA7RR08MA_C0001.MP4
	// .../TC0110JB/XDROOT/Clip/J006C001_200111KE.MXF
	// .../RR08ZEB/PRIVATE/XDROOT/Clip/Clip0001.MXF

	// Parent directory should always be CLIP
	// If the parent directory is XDROOT, look up
	// If parent is XDROOT or M4ROOT, tape ID is the grandparent dir. Otherwise, tape ID is the parent dir.
	//

	// Check that the file extension is valid for the path
	if !isExtValid(filename) {
		return nil, fmt.Errorf("invalid extension for Sony XD clip: %q", filepath.Ext(filename))
	}

	// Split the path parts
	pathParts := strings.Split(filename, "/")
	Reverse(pathParts)

	// We require at least the file, a parent directory, and a grandparent directory
	if len(pathParts) < 3 {
		return nil, errors.New("insufficient path depth for Sony XD clip")
	}

	// The parent directory MUST be "CLIP"
	if strings.ToUpper(pathParts[1]) != "CLIP" {
		return nil, errors.New("sony xd clips must be in CLIP parent directory")
	}

	// Get the tape ID from the tape
	tapeIDPartIndex := 2
	var validTapeID bool
	for i := tapeIDPartIndex; i < len(pathParts); i++ {
		part := pathParts[i]
		if !(isXdrootDirname(part) || strings.ToUpper(part) == "PRIVATE") {
			tapeIDPartIndex = i
			validTapeID = true
			break
		}
	}
	if !validTapeID {
		return nil, errors.New("no valid tape ID found for Sony XD path")
	}
	tapeName := pathParts[tapeIDPartIndex]

	// Resolve all of the other files with valid extensions in the CLIP directory
	dir := strings.ReplaceAll(filepath.Dir(filename), string(os.PathSeparator), "/")
	allFiles, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, fmt.Errorf("error reading CLIP dir \"%s\": %s", dir, err)
	}
	var validFiles []string
	for _, f := range allFiles {
		if f.IsDir() {
			continue
		}
		if !isExtValid(f.Name()) {
			continue
		}
		validFiles = append(validFiles, f.Name())
	}

	// Sort all the valid files alphabetically
	sort.Strings(validFiles)

	// Find the index of the filename
	clipIndex := indexOfString(validFiles, filepath.Base(filename))
	if clipIndex < 0 {
		return nil, fmt.Errorf("error finding file in CLIP dir")
	}

	// Return the clip object
	return &mediadir.Clip{
		TapeName:  tapeName,
		ClipIndex: clipIndex,
	}, nil

}

func indexOfString(strs []string, needle string) int {
	for i, str := range strs {
		if str == needle {
			return i
		}
	}
	return -1
}

func Reverse[T any](arr []T) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}
