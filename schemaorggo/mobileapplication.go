package schemaorggo

import "encoding/json"

// MobileApplication see : https://schema.org/MobileApplication
type MobileApplication struct {
	SoftwareApplication

	typeContext

	// CarrierRequirements see : https://schema.org/carrierRequirements
	// Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
	// types : Text
	CarrierRequirements string `json:"carrierRequirements,omitempty"`
}

func (v MobileApplication) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MobileApplication"

	return json.Marshal(v)
}

func (v *MobileApplication) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MobileApplication"

	return json.Marshal(*v)
}
