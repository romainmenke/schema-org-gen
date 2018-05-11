package schemaorg

import "encoding/json"

// GovernmentService see : https://schema.org/GovernmentService
type GovernmentService struct {

typeContext

Service

// ServiceOperator see : /serviceOperator
// The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor.
ServiceOperator *Organization `json:"serviceOperator"`

}

func (v *GovernmentService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GovernmentService"

	return json.Marshal(v)
}
