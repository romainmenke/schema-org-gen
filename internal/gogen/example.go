package gogen

import (
	"context"
	"errors"

	"github.com/tyler-sommer/stick"
)

func writeExample(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_example_test")
	if err != nil {
		return err
	}
	defer f.Close()

	tmpls := templatesFromContext(ctx)
	if tmpls == nil {
		return errors.New("no template found")
	}

	tmpls.Execute("/templates/example.twig", f,
		map[string]stick.Value{
			"package_name": packageName,
		},
	)

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
