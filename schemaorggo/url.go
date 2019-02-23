package schemaorggo

import "encoding/json"

// URL see : https://schema.org/URL
type URL struct {

	// With properties from Text see : https://schema.org
	//

}

func (v URL) MarshalJSON() ([]byte, error) {
	type Alias URL

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"URL\","), b[1:]...), nil
}
