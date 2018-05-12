package schemaorggo

import "encoding/json"

// ContactPoint see : https://schema.org/ContactPoint
type ContactPoint struct {
	StructuredValue

	typeContext

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	AreaServed interface{} `json:"areaServed,omitempty"` // types : AdministrativeArea GeoShape Place Text

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	AvailableLanguage interface{} `json:"availableLanguage,omitempty"` // types : Language Text

	// ContactOption see : https://schema.org/contactOption
	// An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
	ContactOption *ContactPointOption `json:"contactOption,omitempty"` // types : ContactPointOption

	// ContactType see : https://schema.org/contactType
	// A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
	ContactType string `json:"contactType,omitempty"` // types : Text

	// Email see : https://schema.org/email
	// Email address.
	Email string `json:"email,omitempty"` // types : Text

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	FaxNumber string `json:"faxNumber,omitempty"` // types : Text

	// HoursAvailable see : https://schema.org/hoursAvailable
	// The hours during which this service or contact is available.
	HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable,omitempty"` // types : OpeningHoursSpecification

	// ProductSupported see : https://schema.org/productSupported
	// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. "iPhone") or a general category of products or services (e.g. "smartphones").
	ProductSupported interface{} `json:"productSupported,omitempty"` // types : Product Text

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	Telephone string `json:"telephone,omitempty"` // types : Text

}

func (v ContactPoint) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ContactPoint"

	return json.Marshal(v)
}

func (v *ContactPoint) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ContactPoint"

	return json.Marshal(*v)
}
