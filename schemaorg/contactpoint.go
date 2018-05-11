package schemaorg

import "encoding/json"

// ContactPoint see : https://schema.org/ContactPoint
type ContactPoint struct {

typeContext

StructuredValue

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// ContactOption see : /contactOption
// An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
ContactOption *ContactPointOption `json:"contactOption"`

// ContactType see : /contactType
// A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
ContactType string `json:"contactType"`

// Email see : /email
// Email address.
Email string `json:"email"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// HoursAvailable see : /hoursAvailable
// The hours during which this service or contact is available.
HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable"`

// ProductSupported see : /productSupported
// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. "iPhone") or a general category of products or services (e.g. "smartphones").
ProductSupported interface{} `json:"productSupported"` // types : Product Text

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

}

func (v *ContactPoint) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ContactPoint"

	return json.Marshal(v)
}
