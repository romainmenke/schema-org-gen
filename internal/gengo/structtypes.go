package gengo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"github.com/tyler-sommer/stick"
)

func goTypeFile(goTypes []string, dir string, packageName string) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
	return func(ctx context.Context, o *typemap.ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil {
			return nil
		}

		tmpls := templatesFromContext(ctx)
		if tmpls == nil {
			return errors.New("no template found")
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

		data := map[string]stick.Value{
			"package_name": packageName,
			"type_name":    strings.Title(object.Name),
			"type_url":     seeUrl(object.URL),
		}

		if object.ParentObject != nil {
			data["parent_type_name"] = object.ParentObject.Name
		}

		dataFields := []map[string]interface{}{}
		for _, field := range object.Fields {
			if field.Name == "" {
				continue
			}
			if len(field.Types) == 0 {
				continue
			}

			dataField := map[string]interface{}{}
			dataField["name"] = strings.Title(field.Name)
			dataField["json_name"] = field.Name
			dataField["external_ref"] = fmt.Sprintf("// %s see : %s", strings.Title(field.Name), seeUrl(field.URL))

			comment := newLineRegex.ReplaceAllString(field.Comment, "\n// ")
			dataField["comment"] = fmt.Sprintf("// %s", comment)

			fieldTypes := normalizeTypes(field.Types)

			fieldTypesComment := ""
			for _, fieldType := range field.Types {
				fieldTypesComment = fieldTypesComment + " " + fieldType.Type
			}
			dataField["types_comment"] = fmt.Sprintf("// types :%s", fieldTypesComment)

			if len(fieldTypes) > 1 {
				dataField["go_type"] = "interface{}"
			} else {
				dataField["go_type"] = goTypeForSchemaDataType(goTypes, fieldTypes[0].Type)
			}

			dataFields = append(dataFields, dataField)
		}

		data["fields"] = dataFields
		tmpls.Execute("/templates/structtypes.twig", f, data)

		err = f.Close()
		if err != nil {
			return err
		}

		return nil
	}
}
