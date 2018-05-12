package gogen

import (
	"context"
	"fmt"
)

func writeGoDataTypes(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_data-types")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", packageName))
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
