package schemaorggo

import "encoding/json"

// DeliveryChargeSpecification see : https://schema.org/DeliveryChargeSpecification
type DeliveryChargeSpecification struct {
	PriceSpecification

	typeContext

	// AppliesToDeliveryMethod see : https://schema.org/appliesToDeliveryMethod
	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliesToDeliveryMethod"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

	// EligibleRegion see : https://schema.org/eligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
	//
	// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
	EligibleRegion interface{} `json:"eligibleRegion"` // types : GeoShape Place Text

	// IneligibleRegion see : https://schema.org/ineligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
	//
	// See also eligibleRegion (see: https://schema.org/eligibleRegion).
	IneligibleRegion interface{} `json:"ineligibleRegion"` // types : GeoShape Place Text

}

func (v *DeliveryChargeSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DeliveryChargeSpecification"

	return json.Marshal(v)
}
