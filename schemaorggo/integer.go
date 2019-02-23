package schemaorggo

import "encoding/json"

// Integer see : https://schema.org/Integer
type Integer struct {

	// With properties from Number see : https://schema.org
	//

}

func (v Integer) MarshalJSON() ([]byte, error) {
	type Alias Integer

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Integer\","), b[1:]...), nil
}
