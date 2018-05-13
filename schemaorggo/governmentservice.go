package schemaorggo

import "encoding/json"

// GovernmentService see : https://schema.org/GovernmentService
type GovernmentService struct {
	Service

	typeContext

	// ServiceOperator see : https://schema.org/serviceOperator
	// The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor.
	// types : Organization
	ServiceOperator []*Organization `json:"serviceOperator,omitempty"`
}

func (v GovernmentService) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Service.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ServiceOperator
		if len(v.ServiceOperator) == 1 {
			value = v.ServiceOperator[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serviceOperator"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v GovernmentService) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "GovernmentService"

	return data, nil
}

func (v GovernmentService) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
