package schemaorg

import "encoding/json"

// FollowAction see : https://schema.org/FollowAction
type FollowAction struct {

typeContext

InteractAction

// Followee see : https://schema.org/followee
// A sub property of object. The person or organization being followed.
Followee interface{} `json:"followee"` // types : Organization Person

}

func (v *FollowAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FollowAction"

	return json.Marshal(v)
}
