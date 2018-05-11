package schemaorg

import "encoding/json"

// LoseAction see : https://schema.org/LoseAction
type LoseAction struct {

typeContext

AchieveAction

// Winner see : /winner
// A sub property of participant. The winner of the action.
Winner *Person `json:"winner"`

}

func (v *LoseAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LoseAction"

	return json.Marshal(v)
}
