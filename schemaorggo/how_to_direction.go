package schemaorggo

import "encoding/json"

// HowToDirection see : https://schema.org/HowToDirection
type HowToDirection struct {
	ListItem

	typeContext

	// AfterMedia see : https://schema.org/afterMedia
	// A media object representing the circumstances after performing this direction.
	// types : MediaObject URL
	AfterMedia []interface{} `json:"afterMedia,omitempty"`

	// BeforeMedia see : https://schema.org/beforeMedia
	// A media object representing the circumstances before performing this direction.
	// types : MediaObject URL
	BeforeMedia []interface{} `json:"beforeMedia,omitempty"`

	// DuringMedia see : https://schema.org/duringMedia
	// A media object representing the circumstances while performing this direction.
	// types : MediaObject URL
	DuringMedia []interface{} `json:"duringMedia,omitempty"`

	// PerformTime see : https://schema.org/performTime
	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PerformTime []*Duration `json:"performTime,omitempty"`

	// PrepTime see : https://schema.org/prepTime
	// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PrepTime []*Duration `json:"prepTime,omitempty"`

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
}

func (v HowToDirection) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ListItem.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AfterMedia
		if len(v.AfterMedia) == 1 {
			value = v.AfterMedia[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["afterMedia"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BeforeMedia
		if len(v.BeforeMedia) == 1 {
			value = v.BeforeMedia[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["beforeMedia"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DuringMedia
		if len(v.DuringMedia) == 1 {
			value = v.DuringMedia[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["duringMedia"] = json.RawMessage(b)
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

	*intop = into

	return nil
}

func (v HowToDirection) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToDirection"

	return data, nil
}

func (v HowToDirection) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
