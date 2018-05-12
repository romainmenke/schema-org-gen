package schemaorggo

import "encoding/json"

// FollowAction see : https://schema.org/FollowAction
type FollowAction struct {
	InteractAction

	typeContext

	// Followee see : https://schema.org/followee
	// A sub property of object. The person or organization being followed.
	// types : Organization Person
	Followee interface{} `json:"followee,omitempty"`
}

func (v FollowAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FollowAction"

	return json.Marshal(v)
}

func (v *FollowAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "FollowAction"

	return json.Marshal(*v)
}
