package genphp

import (
	"context"
	"errors"
	"sort"

	"github.com/tyler-sommer/stick"
)

func writeLoad(ctx context.Context, dir string, packageName string, phpFiles []string) error {
	f, err := newObjectFile(dir, packageName)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpls := templatesFromContext(ctx)
	if tmpls == nil {
		return errors.New("no template found")
	}

	sort.Strings(phpFiles)

	err = tmpls.Execute("/templates/load.twig", f,
		map[string]stick.Value{
			"files": phpFiles,
		},
	)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
