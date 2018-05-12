package schemaorggo

import "encoding/json"

// LoseAction see : https://schema.org/LoseAction
type LoseAction struct {
	AchieveAction

	typeContext

	// Winner see : https://schema.org/winner
	// A sub property of participant. The winner of the action.
	// types : Person
	Winner *Person `json:"winner,omitempty"`
}

func (v LoseAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LoseAction"

	return json.Marshal(v)
}

func (v *LoseAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LoseAction"

	return json.Marshal(*v)
}
