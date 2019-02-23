package genphp

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/romainmenke/schema-org-gen/internal/ast"
)

var newLineRegex = regexp.MustCompile(`\r?\n`)

func seeUrl(str string) string {
	if strings.HasPrefix(str, "https://") {
		return str
	}

	if strings.HasPrefix(str, "http://") {
		return "https://" + strings.TrimPrefix(str, "http://")
	}

	return "https://schema.org" + str
}

func listPhpFiles(tm ast.Typemap) []string {
	out := []string{}

	for _, o := range tm {
		out = append(out, strcase.ToSnake(o.Name))
	}

	return out
}
