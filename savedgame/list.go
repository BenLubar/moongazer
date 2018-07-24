package savedgame

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func List() ([]*Summary, error) {
	dir, err := FindBaseDir()
	if err != nil {
		return nil, errors.Wrap(err, "could not determine base save directory")
	}
	dir = filepath.Join(dir, "saves")

	var infos []os.FileInfo
	f, err := os.Open(dir)
	if err == nil {
		infos, err = f.Readdir(0)
	}
	if e := f.Close(); err == nil {
		err = e
	}

	if err != nil {
		return nil, errors.Wrap(err, "attempting to read save directory")
	}

	list := make([]*Summary, 0, len(infos))
	for _, fi := range infos {
		s, err := getSummary(fi, dir)
		if err != nil {
			return list, errors.Wrapf(err, "error encountered while reading saved game file %q", fi.Name())
		}
		if s != nil {
			list = append(list, s)
		}
	}

	return list, nil
}
