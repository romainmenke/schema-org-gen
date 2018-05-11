package schemaorg

import "encoding/json"

// PeopleAudience see : https://schema.org/PeopleAudience
type PeopleAudience struct {

typeContext

Audience

// HealthCondition see : http://health-lifesci.schema.org/healthCondition
// Specifying the health condition(s) of a patient, medical study, or other target audience.
HealthCondition interface{} `json:"healthCondition"`

// RequiredGender see : /requiredGender
// Audiences defined by a person's gender.
RequiredGender string `json:"requiredGender"`

// RequiredMaxAge see : /requiredMaxAge
// Audiences defined by a person's maximum age.
RequiredMaxAge int `json:"requiredMaxAge"`

// RequiredMinAge see : /requiredMinAge
// Audiences defined by a person's minimum age.
RequiredMinAge int `json:"requiredMinAge"`

// SuggestedGender see : /suggestedGender
// The gender of the person or audience.
SuggestedGender string `json:"suggestedGender"`

// SuggestedMaxAge see : /suggestedMaxAge
// Maximal age recommended for viewing content.
SuggestedMaxAge float64 `json:"suggestedMaxAge"`

// SuggestedMinAge see : /suggestedMinAge
// Minimal age recommended for viewing content.
SuggestedMinAge float64 `json:"suggestedMinAge"`

}

func (v *PeopleAudience) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PeopleAudience"

	return json.Marshal(v)
}
