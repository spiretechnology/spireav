package mediadir

import "io/fs"

type Clip struct {
	TapeName  string
	ClipIndex int
}

type ClipParser interface {
	ParseClip(fs.FS, string) (*Clip, error)
}

type ClipParserFunc func(fs.FS, string) (*Clip, error)

var _ ClipParser = ClipParserFunc(nil)

func (f ClipParserFunc) ParseClip(fsys fs.FS, filename string) (*Clip, error) {
	return f(fsys, filename)
}
