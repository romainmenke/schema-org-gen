package schemaorggo

import "encoding/json"

// GovernmentService see : https://schema.org/GovernmentService
type GovernmentService struct {
	Service

	typeContext

	// ServiceOperator see : https://schema.org/serviceOperator
	// The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor.
	// types : Organization
	ServiceOperator *Organization `json:"serviceOperator,omitempty"`
}

func (v GovernmentService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GovernmentService"

	return json.Marshal(v)
}

func (v *GovernmentService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GovernmentService"

	return json.Marshal(*v)
}
