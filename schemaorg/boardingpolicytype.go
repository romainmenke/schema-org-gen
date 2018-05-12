package schemaorg

import "encoding/json"

// BoardingPolicyType see : https://schema.org/BoardingPolicyType
type BoardingPolicyType struct {

Enumeration

typeContext

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *BoardingPolicyType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BoardingPolicyType"

	return json.Marshal(v)
}
