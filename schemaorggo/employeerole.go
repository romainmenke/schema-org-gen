package schemaorggo

import "encoding/json"

// EmployeeRole see : https://schema.org/EmployeeRole
type EmployeeRole struct {
	OrganizationRole

	typeContext

	// BaseSalary see : https://schema.org/baseSalary
	// The base salary of the job or of an employee in an EmployeeRole.
	BaseSalary interface{} `json:"baseSalary,omitempty"` // types : MonetaryAmount Number PriceSpecification

	// SalaryCurrency see : https://schema.org/salaryCurrency
	// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	SalaryCurrency string `json:"salaryCurrency,omitempty"` // types : Text

}

func (v EmployeeRole) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EmployeeRole"

	return json.Marshal(v)
}

func (v *EmployeeRole) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EmployeeRole"

	return json.Marshal(*v)
}
