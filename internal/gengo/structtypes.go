package gengo

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/romainmenke/schema-org-gen/internal/ast"
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

		if len(object.ParentObjects) > 0 {
			sort.Slice(object.ParentObjects, func(i int, j int) bool {
				if (object.ParentObjects[i] == nil || object.ParentObjects[i].Name == "") && (object.ParentObjects[j] == nil || object.ParentObjects[i].Name == "") {
					return false
				}

				if object.ParentObjects[i] == nil || object.ParentObjects[i].Name == "" {
					return true
				}

				if object.ParentObjects[j] == nil || object.ParentObjects[i].Name == "" {
					return true
				}

				return object.ParentObjects[i].Name < object.ParentObjects[j].Name
			})

			duplicateFields := make(map[string]struct{})
			for fIndex, f := range object.Fields {
				if _, ok := duplicateFields[f.Name]; ok {
					object.Fields[fIndex].Duplicate = true
				}

				duplicateFields[f.Name] = struct{}{}
			}
			for _, p := range object.ParentObjects {
				for _, f := range p.Fields {
					if _, ok := duplicateFields[f.Name]; ok {
						f.Duplicate = true
					}

					object.Fields = append(object.Fields, f)
					duplicateFields[f.Name] = struct{}{}
				}
			}

			sort.Slice(object.Fields, func(i int, j int) bool {
				return object.Fields[i].Name < object.Fields[j].Name
			})
		}

		data := map[string]stick.Value{
			"package_name": packageName,
			"type_name":    strings.Title(object.Name),
			"type_url":     seeUrl(object.URL),
			"fields":       objectFields(goTypes, object),
		}

		if len(object.ParentObjects) > 0 {
			parents := []map[string]interface{}{}
			for _, p := range object.ParentObjects {
				parents = append(parents, map[string]interface{}{
					"type_name": strings.Title(p.Name),
					"type_url":  seeUrl(p.URL),
				})
			}

			data["parents"] = parents
		}

		err = tmpls.Execute("/templates/structtypes.twig", f, data)
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

func objectFields(goTypes []string, object *ast.Object) []map[string]interface{} {
	dataFields := []map[string]interface{}{}
	for _, field := range object.Fields {
		if field.Name == "" {
			continue
		}
		if len(field.Types) == 0 {
			continue
		}
		if field.Duplicate {
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
			dataField["go_type"] = "[]interface{}"
		} else {
			dataField["go_type"] = goTypeForSchemaDataType(goTypes, fieldTypes[0].Type)
		}

		dataFields = append(dataFields, dataField)
	}

	return dataFields
}
