package schemaorggo

import "encoding/json"

// PostalAddress see : https://schema.org/PostalAddress
type PostalAddress struct {
	ContactPoint

	typeContext

	// AddressCountry see : https://schema.org/addressCountry
	// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	// types : Country Text
	AddressCountry interface{} `json:"addressCountry,omitempty"`

	// AddressLocality see : https://schema.org/addressLocality
	// The locality. For example, Mountain View.
	// types : Text
	AddressLocality string `json:"addressLocality,omitempty"`

	// AddressRegion see : https://schema.org/addressRegion
	// The region. For example, CA.
	// types : Text
	AddressRegion string `json:"addressRegion,omitempty"`

	// PostOfficeBoxNumber see : https://schema.org/postOfficeBoxNumber
	// The post office box number for PO box addresses.
	// types : Text
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode string `json:"postalCode,omitempty"`

	// StreetAddress see : https://schema.org/streetAddress
	// The street address. For example, 1600 Amphitheatre Pkwy.
	// types : Text
	StreetAddress string `json:"streetAddress,omitempty"`
}

func (v PostalAddress) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PostalAddress"

	return json.Marshal(v)
}

func (v *PostalAddress) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PostalAddress"

	return json.Marshal(*v)
}
