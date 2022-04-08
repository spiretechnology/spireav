package memfs

import (
	"bytes"
	"io"
	"io/fs"
	"strings"
)

var _ fs.ReadDirFS = (*InMemoryFS)(nil)

type InMemoryFS struct {
	rootEntry dirEntry
}

func NewInMemoryFS() *InMemoryFS {
	return &InMemoryFS{
		rootEntry: dirEntry{
			name:     "",
			isdir:    true,
			children: make(map[string]*dirEntry),
		},
	}
}

func (f *InMemoryFS) AddFile(filename string, data []byte) {
	steps := getPathSteps(filename)
	node := &f.rootEntry
	for i, step := range steps {
		nextNode, ok := node.children[step]
		if !ok {
			nextNode = &dirEntry{
				name:     step,
				children: make(map[string]*dirEntry),
				isdir:    i < len(steps)-1,
			}
			node.children[step] = nextNode
		}
		// If this is the last step, it's the file
		if i == len(steps)-1 {
			node.data = data
		}
		node = nextNode
	}
}

func (f *InMemoryFS) Open(name string) (fs.File, error) {
	steps := getPathSteps(name)
	entry := seekEntry(steps, &f.rootEntry)
	if entry == nil {
		return nil, fs.ErrNotExist
	}
	return &inMemFile{
		reader: io.NopCloser(bytes.NewReader(entry.data)),
		entry:  entry,
	}, nil
}

func (f *InMemoryFS) ReadDir(name string) ([]fs.DirEntry, error) {
	steps := getPathSteps(name)
	entry := seekEntry(steps, &f.rootEntry)
	if entry == nil {
		return nil, fs.ErrNotExist
	}
	if !entry.isdir {
		return nil, fs.ErrInvalid
	}
	var entries []fs.DirEntry
	for _, child := range entry.children {
		entries = append(entries, child)
	}
	return entries, nil
}

func getPathSteps(name string) []string {
	trimmed := strings.Trim(name, "/")
	if len(trimmed) == 0 {
		return nil
	}
	return strings.Split(trimmed, "/")
}

func seekEntry(steps []string, entry *dirEntry) *dirEntry {
	if entry == nil {
		return nil
	}
	if len(steps) == 0 {
		return entry
	}
	nextEntry, ok := entry.children[steps[0]]
	if !ok {
		return nil
	}
	return seekEntry(steps[1:], nextEntry)
}
