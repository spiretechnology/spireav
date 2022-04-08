package memfs

import "io/fs"

var _ fs.DirEntry = (*dirEntry)(nil)

type dirEntry struct {
	name     string
	isdir    bool
	data     []byte
	filemode fs.FileMode
	children map[string]*dirEntry
}

func (e *dirEntry) Name() string {
	return e.name
}

func (e *dirEntry) IsDir() bool {
	return e.isdir
}

func (e *dirEntry) Type() fs.FileMode {
	return e.filemode
}

func (e *dirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func (e *dirEntry) AddChild(child *dirEntry) {
	if e.children == nil {
		e.children = make(map[string]*dirEntry)
	}
	e.children[child.name] = child
}
