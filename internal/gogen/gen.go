package gogen

import (
	"context"
	"os/exec"

	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

func Generate(ctx context.Context, tm *typemap.TypeMap, dir string, packageName string) error {
	err := writeGoDataTypes(ctx, dir, packageName)
	if err != nil {
		return err
	}

	goTypes := []string{}
	err = tm.Walk(ctx, listGoTypes(&goTypes))
	if err != nil {
		return err
	}

	err = tm.Walk(ctx, goTypeFile(goTypes, dir, packageName))
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "go", "fmt", dir)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
