package schemaorggo

import "encoding/json"

// MobileApplication see : https://schema.org/MobileApplication
type MobileApplication struct {
	SoftwareApplication

	typeContext

	// CarrierRequirements see : https://schema.org/carrierRequirements
	// Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
	// types : Text
	CarrierRequirements []string `json:"carrierRequirements,omitempty"`
}

func (v MobileApplication) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SoftwareApplication.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CarrierRequirements
		if len(v.CarrierRequirements) == 1 {
			value = v.CarrierRequirements[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["carrierRequirements"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MobileApplication) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MobileApplication"

	return data, nil
}

func (v MobileApplication) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
