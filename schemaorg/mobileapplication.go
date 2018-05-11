package schemaorg

import "encoding/json"

// MobileApplication see : https://schema.org/MobileApplication
type MobileApplication struct {

typeContext

SoftwareApplication

// CarrierRequirements see : https://schema.org/carrierRequirements
// Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
CarrierRequirements string `json:"carrierRequirements"`

}

func (v *MobileApplication) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MobileApplication"

	return json.Marshal(v)
}
