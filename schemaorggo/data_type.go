package schemaorggo

import "encoding/json"

// DataType see : https://schema.org/DataType
type DataType struct {

	// With properties from Rdfs:Class see : https://schema.org
	//

}

func (v DataType) MarshalJSON() ([]byte, error) {
	type Alias DataType

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"DataType\","), b[1:]...), nil
}
