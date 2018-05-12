package genjson

import (
	"context"
	"math/rand"
	"regexp"
	"time"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

var newLineRegex = regexp.MustCompile(`\r?\n`)

func listTypes(sp *[]*ast.Object) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
	return func(ctx context.Context, o *typemap.ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil || sp == nil {
			return nil
		}

		s := *sp
		*sp = append(s, o.Object)

		return nil
	}
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

func randomDate() time.Time {
	r := rand.Intn(int(time.Now().Unix()))
	return time.Unix(int64(r), 0)
}
