package schemaorggo

import "encoding/json"

// WinAction see : https://schema.org/WinAction
type WinAction struct {
	AchieveAction

	typeContext

	// Loser see : https://schema.org/loser
	// A sub property of participant. The loser of the action.
	Loser *Person `json:"loser,omitempty"` // types : Person

}

func (v WinAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WinAction"

	return json.Marshal(v)
}

func (v *WinAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WinAction"

	return json.Marshal(*v)
}
