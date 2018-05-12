package schemaorggo

import "encoding/json"

// AppendAction see : https://schema.org/AppendAction
type AppendAction struct {
	InsertAction

	typeContext

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation *Place `json:"toLocation,omitempty"`
}

func (v AppendAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AppendAction"

	return json.Marshal(v)
}

func (v *AppendAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AppendAction"

	return json.Marshal(*v)
}
