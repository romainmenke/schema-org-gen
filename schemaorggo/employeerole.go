package schemaorggo

import "encoding/json"

// EmployeeRole see : https://schema.org/EmployeeRole
type EmployeeRole struct {
	OrganizationRole

	typeContext

	// BaseSalary see : https://schema.org/baseSalary
	// The base salary of the job or of an employee in an EmployeeRole.
	// types : MonetaryAmount Number PriceSpecification
	BaseSalary []interface{} `json:"baseSalary,omitempty"`

	// SalaryCurrency see : https://schema.org/salaryCurrency
	// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	// types : Text
	SalaryCurrency []string `json:"salaryCurrency,omitempty"`
}

func (v EmployeeRole) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.OrganizationRole.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.BaseSalary
		if len(v.BaseSalary) == 1 {
			value = v.BaseSalary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["baseSalary"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SalaryCurrency
		if len(v.SalaryCurrency) == 1 {
			value = v.SalaryCurrency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["salaryCurrency"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v EmployeeRole) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "EmployeeRole"

	return data, nil
}

func (v EmployeeRole) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
