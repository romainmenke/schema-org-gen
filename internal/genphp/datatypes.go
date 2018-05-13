package genphp

func phpTypesForSchemaDataType(schemaDataType string) []string {
	switch schemaDataType {
	case "Boolean":
		return []string{"boolean", "boolean[]"}
	case "Number":
		return []string{"float", "float[]"}
	case "Float":
		return []string{"float", "float[]"}
	case "Integer":
		return []string{"integer", "integer[]"}
	case "Text":
		return []string{"string", "string[]"}
	case "URL":
		return []string{"string", "string[]"}
	case "Date":
		return []string{"string", "string[]"}
	case "DateTime":
		return []string{"string", "string[]"}
	case "Time":
		return []string{"string", "string[]"}
	default:
		return []string{`\` + schemaDataType, `\` + schemaDataType + "[]"}
	}
}
