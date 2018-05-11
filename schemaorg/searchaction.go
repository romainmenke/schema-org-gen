package schemaorg

import "encoding/json"

// SearchAction see : https://schema.org/SearchAction
type SearchAction struct {

typeContext

Action

// Query see : https://schema.org/query
// A sub property of instrument. The query used on this action.
Query string `json:"query"`

}

func (v *SearchAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SearchAction"

	return json.Marshal(v)
}
