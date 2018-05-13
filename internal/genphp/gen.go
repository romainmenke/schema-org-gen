package genphp

import (
	"context"
	"log"
	"os"

	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

func Generate(ctx context.Context, tm *typemap.TypeMap, dir string, packageName string) error {
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

	err = tm.Walk(ctx, phpTypeFile(dir, packageName))
	if err != nil {
		return err
	}

	phpFiles := []string{}
	err = tm.Walk(ctx, listPhpFiles(&phpFiles))
	if err != nil {
		return err
	}

	err = writeLoad(ctx, dir, packageName, phpFiles)
	if err != nil {
		return err
	}

	return nil
}
