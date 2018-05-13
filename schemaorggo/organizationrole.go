package schemaorggo

import "encoding/json"

// OrganizationRole see : https://schema.org/OrganizationRole
type OrganizationRole struct {
	Role

	typeContext

	// NumberedPosition see : https://schema.org/numberedPosition
	// A number associated with a role in an organization, for example, the number on an athlete&#39;s jersey.
	// types : Number
	NumberedPosition []float64 `json:"numberedPosition,omitempty"`
}

func (v OrganizationRole) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Role.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.NumberedPosition
		if len(v.NumberedPosition) == 1 {
			value = v.NumberedPosition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberedPosition"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v OrganizationRole) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "OrganizationRole"

	return data, nil
}

func (v OrganizationRole) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
