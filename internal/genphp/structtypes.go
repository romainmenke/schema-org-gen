package genphp

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"github.com/tyler-sommer/stick"
)

func phpTypeFile(dir string, packageName string) func(ctx context.Context, o *typemap.ObjectSource, err error) error {
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

		f, err := newObjectFile(path.Join(dir, "classes"), object.Name)
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
			dataField["name"] = strcase.ToSnake(field.Name)
			dataField["json_name"] = field.Name

			phpTypes := []string{}
			for _, fieldType := range field.Types {
				phpTypes = append(phpTypes, phpTypesForSchemaDataType(fieldType.Type)...)
			}

			commentLines := []string{}
			commentLines = append(commentLines, newLineRegex.Split(field.Comment, -1)...)
			commentLines = append(commentLines, fmt.Sprintf("see : %s", seeUrl(field.URL)))
			commentLines = append(commentLines, fmt.Sprintf("@var %s", strings.Join(phpTypes, "|")))

			dataField["comment_lines"] = commentLines

			dataFields = append(dataFields, dataField)
		}

		data["fields"] = dataFields

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
