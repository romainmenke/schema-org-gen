package gengo

import (
	"regexp"
	"strings"

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

func listGoTypes(tm ast.Typemap) []string {
	out := []string{}

	for _, o := range tm {
		out = append(out, strings.Title(o.Name))
	}

	return out
}

func normalizeTypes(types []ast.FieldType) []ast.FieldType {
	textType := ast.FieldType{
		Type: "Text",
	}
	numberType := ast.FieldType{
		Type: "Number",
	}

	var textTypes, numberTypes int
	for _, t := range types {
		if t.Type == "Text" {
			textTypes++
			textType = t
		}
		if t.Type == "URL" {
			textTypes++
		}

		if t.Type == "Number" {
			numberTypes++
			numberType = t
		}
		if t.Type == "Float" {
			numberTypes++
		}
		if t.Type == "Integer" {
			numberTypes++
		}
	}

	if textTypes == len(types) {
		return []ast.FieldType{textType}
	}

	if numberTypes == len(types) {
		return []ast.FieldType{numberType}
	}

	return types
}
