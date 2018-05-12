package gogen

import (
	"context"
	"errors"

	"github.com/tyler-sommer/stick"
)

func writeGoDataTypes(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_data_types")
	if err != nil {
		return err
	}
	defer f.Close()

	tmpls := templatesFromContext(ctx)
	if tmpls == nil {
		return errors.New("no template found")
	}

	tmpls.Execute("/templates/datatypes.twig", f, map[string]stick.Value{"package_name": packageName})

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func goTypeForSchemaDataType(goTypes []string, schemaDataType string) string {
	switch schemaDataType {
	case "Boolean":
		return "bool"
	case "Number":
		return "float64"
	case "Float":
		return "float64"
	case "Integer":
		return "int"
	case "Text":
		return "string"
	case "URL":
		return "string"
	case "Date":
		return "Date"
	case "DateTime":
		return "DateTime"
	case "Time":
		return "Time"
	default:
		for _, goType := range goTypes {
			if goType == schemaDataType {
				return "*" + schemaDataType
			}
		}
		return "interface{}"
	}
}
