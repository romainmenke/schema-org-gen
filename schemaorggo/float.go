package schemaorggo

import "encoding/json"

// Float see : https://schema.org/Float
type Float struct {

	// With properties from Number see : https://schema.org
	//

}

func (v Float) MarshalJSON() ([]byte, error) {
	type Alias Float

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Float\","), b[1:]...), nil
}
