package gengo

import (
	"context"
	"errors"

	"github.com/tyler-sommer/stick"
)

func writeUtil(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_util")
	if err != nil {
		return err
	}
	defer f.Close()

	tmpls := templatesFromContext(ctx)
	if tmpls == nil {
		return errors.New("no template found")
	}

	err = tmpls.Execute("/templates/util.twig", f,
		map[string]stick.Value{
			"package_name": packageName,
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
