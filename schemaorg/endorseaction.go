package schemaorg

import "encoding/json"

// EndorseAction see : https://schema.org/EndorseAction
type EndorseAction struct {

typeContext

ReactAction

// Endorsee see : https://schema.org/endorsee
// A sub property of participant. The person/organization being supported.
Endorsee interface{} `json:"endorsee"` // types : Organization Person

}

func (v *EndorseAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EndorseAction"

	return json.Marshal(v)
}
