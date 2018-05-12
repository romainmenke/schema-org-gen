package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

var newLineRegex = regexp.MustCompile(`\r?\n`)

func listGoTypes(sp *[]string) func(ctx context.Context, o *ObjectSource, err error) error {
	return func(ctx context.Context, o *ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil || sp == nil {
			return nil
		}

		s := *sp
		*sp = append(s, strings.Title(o.Object.Name))

		return nil
	}
}

func goTypeFile(goTypes []string) func(ctx context.Context, o *ObjectSource, err error) error {
	return func(ctx context.Context, o *ObjectSource, err error) error {
		if err != nil {
			return err
		}
		if o == nil || o.Object == nil {
			return nil
		}

		if verboseLog {
			start := time.Now()
			defer func() {
				log.Printf("generated go : %s, duration : %v\n", o.URL, time.Since(start))
			}()
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

func seeUrl(str string) string {
	if strings.HasPrefix(str, "http") {
		return str
	}

	return "https://schema.org" + str
}

func writeGoDataTypes(ctx context.Context) error {
	if verboseLog {
		start := time.Now()
		defer func() {
			log.Printf("generated go data types, duration : %v\n", time.Since(start))
		}()
	}

	f, err := newObjectFile("./schemaorg", "data-types")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package schemaorg\n\n")
	f.WriteString(`
import (
	"fmt"
	"time"
)
`)

	f.WriteString(fmt.Sprintf("// Data Types see : %s\n\n", "https://schema.org/DataType"))

	f.WriteString("// typeContext is used for fixed values")
	f.WriteString(fmt.Sprintf(`
type typeContext struct {
	C string %s
	T string %s
}
`, "`json:\"@context\"`", "`json:\"@type\"`"))

	f.WriteString("// Date https://en.wikipedia.org/wiki/ISO_8601")
	f.WriteString(`
type Date time.Time

func (v Date)MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("\"%s\"", time.Time(v).Format("2006-01-02"))
	return []byte(jsonValue), nil
}
`)

	f.WriteString("// DateTime https://en.wikipedia.org/wiki/ISO_8601")
	f.WriteString(`
type DateTime time.Time

`)

	f.WriteString("// Time https://schema.org/Time")
	f.WriteString(`
type Time string

`)

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
	default:
		for _, goType := range goTypes {
			if goType == schemaDataType {
				return "*" + schemaDataType
			}
		}
		return "interface{}"
	}
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
