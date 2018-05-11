package schemaorg

import "encoding/json"

// EmployeeRole see : https://schema.org/EmployeeRole
type EmployeeRole struct {

typeContext

OrganizationRole

// BaseSalary see : /baseSalary
// The base salary of the job or of an employee in an EmployeeRole.
BaseSalary interface{} `json:"baseSalary"` // types : MonetaryAmount Number PriceSpecification

// SalaryCurrency see : /salaryCurrency
// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
SalaryCurrency string `json:"salaryCurrency"`

}

func (v *EmployeeRole) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EmployeeRole"

	return json.Marshal(v)
}
