package schemaorggo

import "encoding/json"

// MovieTheater see : https://schema.org/MovieTheater
type MovieTheater struct {
	CivicStructure

	typeContext

	// ScreenCount see : https://schema.org/screenCount
	// The number of screens in the movie theater.
	// types : Number
	ScreenCount []float64 `json:"screenCount,omitempty"`
}

func (v MovieTheater) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CivicStructure.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ScreenCount
		if len(v.ScreenCount) == 1 {
			value = v.ScreenCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["screenCount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MovieTheater) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MovieTheater"

	return data, nil
}

func (v MovieTheater) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
