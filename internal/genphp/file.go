package genphp

import (
	"os"
	"path"

	"github.com/iancoleman/strcase"
)

func newObjectFile(dir string, typeName string) (*os.File, error) {
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, err
	}

	os.Remove(path.Join(dir, strcase.ToSnake(typeName+".php")))

	f, err := os.Create(path.Join(dir, strcase.ToSnake(typeName+".php")))
	if err != nil {
		return nil, err
	}

	return f, nil
}
