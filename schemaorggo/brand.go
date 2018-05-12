package schemaorggo

import "encoding/json"

// Brand see : https://schema.org/Brand
type Brand struct {
	Intangible

	typeContext
}

func (v Brand) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Brand"

	return json.Marshal(v)
}

func (v *Brand) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Brand"

	return json.Marshal(*v)
}
