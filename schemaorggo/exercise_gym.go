package schemaorggo

import "encoding/json"

// ExerciseGym see : https://schema.org/ExerciseGym
type ExerciseGym struct {
	SportsActivityLocation

	typeContext
}

func (v ExerciseGym) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SportsActivityLocation.intoMap(intop)

	into := *intop

	*intop = into

	return nil
}

func (v ExerciseGym) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ExerciseGym"

	return data, nil
}

func (v ExerciseGym) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
