package schemaorggo

import "encoding/json"

// ParentAudience see : https://schema.org/ParentAudience
type ParentAudience struct {
	PeopleAudience

	typeContext

	// ChildMaxAge see : https://schema.org/childMaxAge
	// Maximal age of the child.
	// types : Number
	ChildMaxAge []float64 `json:"childMaxAge,omitempty"`

	// ChildMinAge see : https://schema.org/childMinAge
	// Minimal age of the child.
	// types : Number
	ChildMinAge []float64 `json:"childMinAge,omitempty"`
}

func (v ParentAudience) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PeopleAudience.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ChildMaxAge
		if len(v.ChildMaxAge) == 1 {
			value = v.ChildMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["childMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ChildMinAge
		if len(v.ChildMinAge) == 1 {
			value = v.ChildMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["childMinAge"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ParentAudience) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ParentAudience"

	return data, nil
}

func (v ParentAudience) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
