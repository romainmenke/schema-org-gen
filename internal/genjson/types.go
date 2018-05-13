package genjson

import (
	"context"
	"encoding/json"
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

func typeFile(types []*ast.Object, dir string) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
	return func(ctx context.Context, o *typemap.ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil {
			return nil
		}

		object := o.Object

		if object.Name == "" {
			return nil
		}

		f, err := newObjectFile(dir, object.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")

		data := map[string]interface{}{
			"@context": "http://schema.org",
			"@type":    object.Name,
		}

		parentObject := object.ParentObject
		for parentObject != nil {
			setFields(types, data, parentObject, 0)
			parentObject = parentObject.ParentObject
		}

		data = setFields(types, data, object, 0)

		err = enc.Encode(data)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}

		return nil
	}
}

func setFields(types []*ast.Object, data map[string]interface{}, object *ast.Object, nestingLevel int) map[string]interface{} {
	for _, field := range object.Fields {
		if field.Name == "" {
			continue
		}
		if len(field.Types) == 0 {
			continue
		}

		fieldTypes := normalizeTypes(field.Types)
		fieldType := fieldTypes[0].Type

		data[field.Name] = jsonValueForSchemaDataType(types, fieldType, nestingLevel)
	}

	return data
}

func jsonValueForSchemaDataType(types []*ast.Object, schemaDataType string, nestingLevel int) interface{} {
	switch schemaDataType {
	case "Boolean":
		return randomdata.Boolean()
	case "Number":
		return randomdata.Number(100)
	case "Float":
		return randomdata.Decimal(0, 20, 3)
	case "Integer":
		return randomdata.Number(10)
	case "Text":
		return randomdata.SillyName()
	case "URL":
		return fmt.Sprintf("https://%s.com/%s", randomdata.SillyName(), randomdata.SillyName())
	case "Date":
		return randomDate().Format("2006-01-02")
	case "DateTime":
		return randomDate().Format("2006-01-02T15:04")
	case "Time":
		return randomDate().Format("2006-01-02T15:04")
	default:
		if nestingLevel > 0 {
			return nil
		}
		for _, t := range types {
			if schemaDataType == t.Name {
				return setFields(types, map[string]interface{}{
					"@type": t.Name,
				}, t, nestingLevel+1)
			}
		}
	}

	return nil
}
