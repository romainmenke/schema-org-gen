package schemaorggo

import "encoding/json"

// MedicalOrganization see : https://schema.org/MedicalOrganization
type MedicalOrganization struct {
	Organization

	typeContext

	// HealthPlanNetworkId see : http://pending.schema.org/healthPlanNetworkId
	// Name or unique ID of network. (Networks are often reused across different insurance plans).
	// types : Text
	HealthPlanNetworkId []string `json:"healthPlanNetworkId,omitempty"`

	// IsAcceptingNewPatients see : http://pending.schema.org/isAcceptingNewPatients
	// Whether the provider is accepting new patients.
	// types : Boolean
	IsAcceptingNewPatients []bool `json:"isAcceptingNewPatients,omitempty"`

	// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	// types : MedicalSpecialty
	MedicalSpecialty []interface{} `json:"medicalSpecialty,omitempty"`
}

func (v MedicalOrganization) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Organization.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.HealthPlanNetworkId
		if len(v.HealthPlanNetworkId) == 1 {
			value = v.HealthPlanNetworkId[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["healthPlanNetworkId"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsAcceptingNewPatients
		if len(v.IsAcceptingNewPatients) == 1 {
			value = v.IsAcceptingNewPatients[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAcceptingNewPatients"] = json.RawMessage(b)
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

func (v MedicalOrganization) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MedicalOrganization"

	return data, nil
}

func (v MedicalOrganization) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
