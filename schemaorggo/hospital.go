package schemaorggo

import "encoding/json"

// Hospital see : https://schema.org/Hospital
type Hospital struct {
	typeContext

	// AvailableService see : https://health-lifesci.schema.org/availableService
	// A medical service available from this provider.
	// types : MedicalProcedure MedicalTest MedicalTherapy
	AvailableService []interface{} `json:"availableService,omitempty"`

	// MedicalSpecialty see : https://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	// types : MedicalSpecialty
	MedicalSpecialty []interface{} `json:"medicalSpecialty,omitempty"`
}

func (v Hospital) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AvailableService
		if len(v.AvailableService) == 1 {
			value = v.AvailableService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableService"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MedicalSpecialty
		if len(v.MedicalSpecialty) == 1 {
			value = v.MedicalSpecialty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["medicalSpecialty"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Hospital) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Hospital"

	return data, nil
}

func (v Hospital) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
