package gogen

import (
	"context"
	"fmt"
	"strings"

	"github.com/romainmenke/schema-org-gen/internal/typemap"
)

func goTypeFile(goTypes []string, dir string, packageName string) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
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

		objectTypeName := strings.Title(object.Name)

		f, err := newObjectFile(dir, object.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

		f.WriteString("import \"encoding/json\"\n\n")

		f.WriteString(fmt.Sprintf("// %s see : %s\n", objectTypeName, seeUrl(object.URL)))
		f.WriteString(fmt.Sprintf("type %s struct {\n\n", objectTypeName))

		if object.ParentObject != nil {
			f.WriteString(object.ParentObject.Name + "\n\n")
		}

		f.WriteString("typeContext\n\n")

		for _, field := range object.Fields {
			if field.Name == "" {
				continue
			}
			if len(field.Types) == 0 {
				continue
			}

			fieldTypes := normalizeTypes(field.Types)

			fieldName := strings.Title(field.Name)

			comment := newLineRegex.ReplaceAllString(field.Comment, "\n// ")

			f.WriteString(fmt.Sprintf("// %s see : %s\n", fieldName, seeUrl(field.URL)))
			f.WriteString(fmt.Sprintf("// %s\n", comment))

			fieldTypesComment := ""
			for _, fieldType := range field.Types {
				fieldTypesComment = fieldTypesComment + " " + fieldType.Type
			}
			if len(fieldTypes) > 1 {
				f.WriteString(fmt.Sprintf("%s interface{} `json:\"%s,omitempty\"` // types :%s\n\n", fieldName, field.Name, fieldTypesComment))
			} else {
				f.WriteString(fmt.Sprintf("%s %s `json:\"%s,omitempty\"` // types :%s\n\n", fieldName, goTypeForSchemaDataType(goTypes, fieldTypes[0].Type), field.Name, fieldTypesComment))
			}
		}

		f.WriteString("}\n")

		f.WriteString(fmt.Sprintf(`
func (v %s) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "%s"

	return json.Marshal(v)
}

func (v *%s) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "%s"

	return json.Marshal(*v)
}
`, objectTypeName, objectTypeName, objectTypeName, objectTypeName))

		err = f.Close()
		if err != nil {
			return err
		}

		return nil
	}
}
