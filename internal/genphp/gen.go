package genphp

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/romainmenke/schema-org-gen/internal/ast"
)

func Generate(ctx context.Context, tm ast.Typemap, dir string, packageName string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Println(err)
	}

	tmpls, err := loadTemplates()
	if err != nil {
		return err
	}

	ctx = contextWithTemplates(ctx, tmpls)

	err = writeUtil(ctx, dir, packageName)
	if err != nil {
		return err
	}

	for _, o := range tm {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if len(o.Fields) == 0 {
			continue
		}

		err := phpTypeFile(ctx, dir, packageName, o)
		if err != nil {
			return err
		}
	}

	phpFiles := listPhpFiles(tm)

	err = writeLoad(ctx, dir, packageName, phpFiles)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "./vendor-php/bin/phpcbf", "--standard=WordPress", "./schemaorgphp", dir)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
