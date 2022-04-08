package mediadir

import (
	"errors"
	"io/fs"
)

func MultiFormatClipParser(formats ...ClipParser) ClipParser {
	return ClipParserFunc(func(fsys fs.FS, filename string) (*Clip, error) {
		var clip *Clip
		var err error
		for _, format := range formats {
			clip, err = format.ParseClip(fsys, filename)
			if clip != nil && err == nil {
				return clip, nil
			}
		}
		if err == nil {
			err = errors.New("no clip resolved")
		}
		return nil, err
	})
}
