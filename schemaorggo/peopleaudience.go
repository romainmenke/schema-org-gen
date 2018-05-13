package schemaorggo

import "encoding/json"

// PeopleAudience see : https://schema.org/PeopleAudience
type PeopleAudience struct {
	Audience

	typeContext

	// HealthCondition see : http://health-lifesci.schema.org/healthCondition
	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	// types : MedicalCondition
	HealthCondition []interface{} `json:"healthCondition,omitempty"`

	// RequiredGender see : https://schema.org/requiredGender
	// Audiences defined by a person&#39;s gender.
	// types : Text
	RequiredGender []string `json:"requiredGender,omitempty"`

	// RequiredMaxAge see : https://schema.org/requiredMaxAge
	// Audiences defined by a person&#39;s maximum age.
	// types : Integer
	RequiredMaxAge []float64 `json:"requiredMaxAge,omitempty"`

	// RequiredMinAge see : https://schema.org/requiredMinAge
	// Audiences defined by a person&#39;s minimum age.
	// types : Integer
	RequiredMinAge []float64 `json:"requiredMinAge,omitempty"`

	// SuggestedGender see : https://schema.org/suggestedGender
	// The gender of the person or audience.
	// types : Text
	SuggestedGender []string `json:"suggestedGender,omitempty"`

	// SuggestedMaxAge see : https://schema.org/suggestedMaxAge
	// Maximal age recommended for viewing content.
	// types : Number
	SuggestedMaxAge []float64 `json:"suggestedMaxAge,omitempty"`

	// SuggestedMinAge see : https://schema.org/suggestedMinAge
	// Minimal age recommended for viewing content.
	// types : Number
	SuggestedMinAge []float64 `json:"suggestedMinAge,omitempty"`
}

func (v PeopleAudience) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Audience.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.HealthCondition
		if len(v.HealthCondition) == 1 {
			value = v.HealthCondition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["healthCondition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredGender
		if len(v.RequiredGender) == 1 {
			value = v.RequiredGender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredGender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredMaxAge
		if len(v.RequiredMaxAge) == 1 {
			value = v.RequiredMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredMinAge
		if len(v.RequiredMinAge) == 1 {
			value = v.RequiredMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredMinAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedGender
		if len(v.SuggestedGender) == 1 {
			value = v.SuggestedGender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedGender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedMaxAge
		if len(v.SuggestedMaxAge) == 1 {
			value = v.SuggestedMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedMinAge
		if len(v.SuggestedMinAge) == 1 {
			value = v.SuggestedMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedMinAge"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PeopleAudience) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PeopleAudience"

	return data, nil
}

func (v PeopleAudience) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
