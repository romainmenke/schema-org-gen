package schemaorggo

import "encoding/json"

// PrependAction see : https://schema.org/PrependAction
type PrependAction struct {
	InsertAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	ToLocation *Place `json:"toLocation,omitempty"`
}

func (v PrependAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PrependAction"

	return json.Marshal(v)
}

func (v *PrependAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PrependAction"

	return json.Marshal(*v)
}
