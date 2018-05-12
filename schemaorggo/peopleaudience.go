package schemaorggo

import "encoding/json"

// PeopleAudience see : https://schema.org/PeopleAudience
type PeopleAudience struct {
	Audience

	typeContext

	// HealthCondition see : http://health-lifesci.schema.org/healthCondition
	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	HealthCondition interface{} `json:"healthCondition,omitempty"` // types : MedicalCondition

	// RequiredGender see : https://schema.org/requiredGender
	// Audiences defined by a person's gender.
	RequiredGender string `json:"requiredGender,omitempty"` // types : Text

	// RequiredMaxAge see : https://schema.org/requiredMaxAge
	// Audiences defined by a person's maximum age.
	RequiredMaxAge float64 `json:"requiredMaxAge,omitempty"` // types : Integer

	// RequiredMinAge see : https://schema.org/requiredMinAge
	// Audiences defined by a person's minimum age.
	RequiredMinAge float64 `json:"requiredMinAge,omitempty"` // types : Integer

	// SuggestedGender see : https://schema.org/suggestedGender
	// The gender of the person or audience.
	SuggestedGender string `json:"suggestedGender,omitempty"` // types : Text

	// SuggestedMaxAge see : https://schema.org/suggestedMaxAge
	// Maximal age recommended for viewing content.
	SuggestedMaxAge float64 `json:"suggestedMaxAge,omitempty"` // types : Number

	// SuggestedMinAge see : https://schema.org/suggestedMinAge
	// Minimal age recommended for viewing content.
	SuggestedMinAge float64 `json:"suggestedMinAge,omitempty"` // types : Number

}

func (v PeopleAudience) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PeopleAudience"

	return json.Marshal(v)
}

func (v *PeopleAudience) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PeopleAudience"

	return json.Marshal(*v)
}
