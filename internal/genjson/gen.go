package genjson

import (
	"context"
	"log"
	"os"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

func Generate(ctx context.Context, tm *typemap.TypeMap, dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Println(err)
	}

	types := []*ast.Object{}
	err = tm.Walk(ctx, listTypes(&types))
	if err != nil {
		return err
	}

	err = tm.Walk(ctx, typeFile(types, dir))
	if err != nil {
		return err
	}

	return nil
}
