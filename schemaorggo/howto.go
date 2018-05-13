package schemaorggo

import "encoding/json"

// HowTo see : https://schema.org/HowTo
type HowTo struct {
	CreativeWork

	typeContext

	// EstimatedCost see : https://schema.org/estimatedCost
	// The estimated cost of the supply or supplies consumed when performing instructions.
	// types : MonetaryAmount Text
	EstimatedCost []interface{} `json:"estimatedCost,omitempty"`

	// PerformTime see : https://schema.org/performTime
	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PerformTime []*Duration `json:"performTime,omitempty"`

	// PrepTime see : https://schema.org/prepTime
	// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PrepTime []*Duration `json:"prepTime,omitempty"`

	// Steps see : https://schema.org/steps
	// The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	// types : CreativeWork ItemList Text
	Steps []interface{} `json:"steps,omitempty"`

	// Supply see : https://schema.org/supply
	// A sub-property of instrument. A supply consumed when performing instructions or a direction.
	// types : HowToSupply Text
	Supply []interface{} `json:"supply,omitempty"`

	// Tool see : https://schema.org/tool
	// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	// types : HowToTool Text
	Tool []interface{} `json:"tool,omitempty"`

	// TotalTime see : https://schema.org/totalTime
	// The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	TotalTime []*Duration `json:"totalTime,omitempty"`

	// Yield see : https://schema.org/yield
	// The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
	// types : QuantitativeValue Text
	Yield []interface{} `json:"yield,omitempty"`
}

func (v HowTo) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EstimatedCost
		if len(v.EstimatedCost) == 1 {
			value = v.EstimatedCost[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["estimatedCost"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PerformTime
		if len(v.PerformTime) == 1 {
			value = v.PerformTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["performTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrepTime
		if len(v.PrepTime) == 1 {
			value = v.PrepTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["prepTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Steps
		if len(v.Steps) == 1 {
			value = v.Steps[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["steps"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Supply
		if len(v.Supply) == 1 {
			value = v.Supply[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["supply"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Tool
		if len(v.Tool) == 1 {
			value = v.Tool[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["tool"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TotalTime
		if len(v.TotalTime) == 1 {
			value = v.TotalTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["totalTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Yield
		if len(v.Yield) == 1 {
			value = v.Yield[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["yield"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v HowTo) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowTo"

	return data, nil
}

func (v HowTo) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
