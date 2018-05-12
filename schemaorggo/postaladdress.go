package schemaorggo

import "encoding/json"

// PostalAddress see : https://schema.org/PostalAddress
type PostalAddress struct {
	ContactPoint

	typeContext

	// AddressCountry see : https://schema.org/addressCountry
	// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	AddressCountry interface{} `json:"addressCountry"` // types : Country Text

	// AddressLocality see : https://schema.org/addressLocality
	// The locality. For example, Mountain View.
	AddressLocality string `json:"addressLocality"`

	// AddressRegion see : https://schema.org/addressRegion
	// The region. For example, CA.
	AddressRegion string `json:"addressRegion"`

	// PostOfficeBoxNumber see : https://schema.org/postOfficeBoxNumber
	// The post office box number for PO box addresses.
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	PostalCode string `json:"postalCode"`

	// StreetAddress see : https://schema.org/streetAddress
	// The street address. For example, 1600 Amphitheatre Pkwy.
	StreetAddress string `json:"streetAddress"`
}

func (v *PostalAddress) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PostalAddress"

	return json.Marshal(v)
}
