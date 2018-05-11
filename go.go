package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var newLineRegex = regexp.MustCompile(`\r?\n`)

func goTypeFile(ctx context.Context, o *ObjectSource, err error) error {
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

	f, err := newObjectFile("./schemaorg", object.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package schemaorg\n\n")

	f.WriteString(fmt.Sprintf("// %s see : %s\n", strings.Title(object.Name), object.URL))
	f.WriteString(fmt.Sprintf("type %s struct {\n\n", strings.Title(object.Name)))

	if object.ParentObject != nil {
		f.WriteString(object.ParentObject.Name)
	}

	for _, field := range object.Fields {
		if field.Name == "" {
			continue
		}
		if len(field.Types) == 0 {
			continue
		}

		comment := newLineRegex.ReplaceAllString(field.Comment, "\n// ")

		f.WriteString(fmt.Sprintf("// %s see : %s\n", strings.Title(field.Name), field.URL))
		f.WriteString(fmt.Sprintf("// %s\n", comment))

		if len(field.Types) > 1 {
			fieldTypesComment := ""
			for _, fieldType := range field.Types {
				fieldTypesComment = fieldTypesComment + " " + fieldType.Type
			}
			f.WriteString(fmt.Sprintf("%s interface{} `json:\"%s\"` // types :%s\n\n", strings.Title(field.Name), field.Name, fieldTypesComment))
		} else {
			f.WriteString(fmt.Sprintf("%s string `json:\"%s\"`\n\n", strings.Title(field.Name), field.Name))
		}
	}

	f.WriteString("}\n\n")

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func newObjectFile(dir string, typeName string) (*os.File, error) {
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, err
	}

	os.Remove(path.Join(dir, strings.ToLower(typeName+".go")))

	f, err := os.Create(path.Join(dir, strings.ToLower(typeName+".go")))
	if err != nil {
		return nil, err
	}

	return f, nil
}
