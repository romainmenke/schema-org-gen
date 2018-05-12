package schemaorg

import "encoding/json"

// Brand see : https://schema.org/Brand
type Brand struct {

Intangible

typeContext

}

func (v *Brand) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Brand"

	return json.Marshal(v)
}
