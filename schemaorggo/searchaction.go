package schemaorggo

import "encoding/json"

// SearchAction see : https://schema.org/SearchAction
type SearchAction struct {
	Action

	typeContext

	// Query see : https://schema.org/query
	// A sub property of instrument. The query used on this action.
	Query string `json:"query,omitempty"`
}

func (v SearchAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SearchAction"

	return json.Marshal(v)
}

func (v *SearchAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SearchAction"

	return json.Marshal(*v)
}
