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

		f, err := newObjectFile(dir, object.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

		f.WriteString("import \"encoding/json\"\n\n")

		f.WriteString(fmt.Sprintf("// %s see : %s\n", strings.Title(object.Name), seeUrl(object.URL)))
		f.WriteString(fmt.Sprintf("type %s struct {\n\n", strings.Title(object.Name)))

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

			comment := newLineRegex.ReplaceAllString(field.Comment, "\n// ")

			f.WriteString(fmt.Sprintf("// %s see : %s\n", strings.Title(field.Name), seeUrl(field.URL)))
			f.WriteString(fmt.Sprintf("// %s\n", comment))

			if len(field.Types) > 1 {
				fieldTypesComment := ""
				for _, fieldType := range field.Types {
					fieldTypesComment = fieldTypesComment + " " + fieldType.Type
				}
				f.WriteString(fmt.Sprintf("%s interface{} `json:\"%s\"` // types :%s\n\n", strings.Title(field.Name), field.Name, fieldTypesComment))
			} else {
				f.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n\n", strings.Title(field.Name), goTypeForSchemaDataType(goTypes, field.Types[0].Type), field.Name))
			}
		}

		f.WriteString("}\n")

		f.WriteString(fmt.Sprintf(`
func (v *%s) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "%s"

	return json.Marshal(v)
}
`, strings.Title(object.Name), strings.Title(object.Name)))

		err = f.Close()
		if err != nil {
			return err
		}

		return nil
	}
}
