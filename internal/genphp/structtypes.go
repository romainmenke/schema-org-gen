package genphp

import (
	"context"
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/tyler-sommer/stick"
)

func phpTypeFile(ctx context.Context, dir string, packageName string, object ast.Object) error {
	tmpls := templatesFromContext(ctx)
	if tmpls == nil {
		return errors.New("no template found")
	}

	if object.Name == "" {
		return nil
	}

	f, err := newObjectFile(path.Join(dir, "classes"), object.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	sort.Slice(object.Fields, func(i int, j int) bool {
		return object.Fields[i].Name < object.Fields[j].Name
	})

	data := map[string]stick.Value{
		"package_name": packageName,
		"type_name":    strings.Title(object.Name),
		"type_url":     seeUrl(object.URL),
		"fields":       objectFields(object),
	}

	if len(object.SubClassOf) > 0 {
		parents := []map[string]interface{}{}
		for _, p := range object.SubClassOf {
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

func objectFields(object ast.Object) []map[string]interface{} {
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
		dataField["name"] = strcase.ToSnake(field.Name)
		dataField["json_name"] = field.Name

		phpTypes := []string{}
		for _, fieldType := range field.Types {
			phpTypes = append(phpTypes, phpTypesForSchemaDataType(fieldType.Type)...)
		}

		phpTypes = shakePhpTypesForSchemaDataType(phpTypes)

		commentLines := []string{}
		commentLines = append(commentLines, newLineRegex.Split(field.Comment, -1)...)
		commentLines = append(commentLines, fmt.Sprintf("see : %s", seeUrl(field.URL)))
		commentLines = append(commentLines, fmt.Sprintf("@var %s", strings.Join(phpTypes, " | ")))

		dataField["comment_lines"] = commentLines

		dataFields = append(dataFields, dataField)
	}

	return dataFields
}
