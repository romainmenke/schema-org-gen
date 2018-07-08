package schemaorggo

import "encoding/json"

// EatAction see : https://schema.org/EatAction
type EatAction struct {
	ConsumeAction

	typeContext

	// ActionAccessibilityRequirement see : https://pending.schema.org/actionAccessibilityRequirement
	// A set of requirements that a must be fulfilled in order to perform an Action. If more than one value is specied, fulfilling one set of requirements will allow the Action to be performed.
	// types : ActionAccessSpecification
	ActionAccessibilityRequirement []interface{} `json:"actionAccessibilityRequirement,omitempty"`

	// ExpectsAcceptanceOf see : https://pending.schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf []*Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v EatAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ConsumeAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ActionAccessibilityRequirement
		if len(v.ActionAccessibilityRequirement) == 1 {
			value = v.ActionAccessibilityRequirement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionAccessibilityRequirement"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExpectsAcceptanceOf
		if len(v.ExpectsAcceptanceOf) == 1 {
			value = v.ExpectsAcceptanceOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["expectsAcceptanceOf"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v EatAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "EatAction"

	return data, nil
}

func (v EatAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
