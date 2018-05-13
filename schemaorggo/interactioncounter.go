package schemaorggo

import "encoding/json"

// InteractionCounter see : https://schema.org/InteractionCounter
type InteractionCounter struct {
	StructuredValue

	typeContext

	// InteractionService see : https://schema.org/interactionService
	// The WebSite or SoftwareApplication where the interactions took place.
	// types : SoftwareApplication WebSite
	InteractionService []interface{} `json:"interactionService,omitempty"`

	// InteractionType see : https://schema.org/interactionType
	// The Action representing the type of interaction. For up votes, +1s, etc. use LikeAction (see: https://schema.org/LikeAction). For down votes use DislikeAction (see: https://schema.org/DislikeAction). Otherwise, use the most specific Action.
	// types : Action
	InteractionType []*Action `json:"interactionType,omitempty"`

	// UserInteractionCount see : https://schema.org/userInteractionCount
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication.
	// types : Integer
	UserInteractionCount []float64 `json:"userInteractionCount,omitempty"`
}

func (v InteractionCounter) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.InteractionService
		if len(v.InteractionService) == 1 {
			value = v.InteractionService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["interactionService"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InteractionType
		if len(v.InteractionType) == 1 {
			value = v.InteractionType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["interactionType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UserInteractionCount
		if len(v.UserInteractionCount) == 1 {
			value = v.UserInteractionCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["userInteractionCount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v InteractionCounter) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "InteractionCounter"

	return data, nil
}

func (v InteractionCounter) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
