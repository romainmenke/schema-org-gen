package genphp

import (
	"context"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

var newLineRegex = regexp.MustCompile(`\r?\n`)

func seeUrl(str string) string {
	if strings.HasPrefix(str, "http") {
		return str
	}

	return "https://schema.org" + str
}

func listPhpFiles(sp *[]string) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
	return func(ctx context.Context, o *typemap.ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil || o.Object.Name == "" || sp == nil {
			return nil
		}

		s := *sp
		*sp = append(s, strcase.ToSnake(o.Object.Name))

		return nil
	}
}
