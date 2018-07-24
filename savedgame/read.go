package savedgame

import (
	"archive/zip"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

func readSaveJSON(savePath, filePath string, v interface{}) (err error) {
	zr, err := zip.OpenReader(savePath)
	if err != nil {
		return
	}
	defer func() {
		if e := zr.Close(); err == nil {
			err = e
		}
	}()

	for _, f := range zr.File {
		if f.Name == filePath {
			var rc io.ReadCloser
			rc, err = f.Open()
			if err != nil {
				return
			}
			defer func() {
				if e := rc.Close(); err == nil {
					err = e
				}
			}()

			return json.NewDecoder(rc).Decode(v)
		}
	}

	return errors.Errorf("file %q was not found in archive", filePath)
}
