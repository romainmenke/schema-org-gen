package gengo

import (
	"os"
	"path"
	"strings"
)

func newObjectFile(dir string, typeName string) (*os.File, error) {
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, err
	}

	os.Remove(path.Join(dir, strings.ToLower(typeName+".go")))

	f, err := os.Create(path.Join(dir, strings.ToLower(typeName+".go")))
	if err != nil {
		return nil, err
	}

	return f, nil
}
