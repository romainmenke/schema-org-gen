package schemaorggo

import "encoding/json"

// PostalAddress see : https://schema.org/PostalAddress
type PostalAddress struct {
	ContactPoint

	typeContext

	// AddressCountry see : https://schema.org/addressCountry
	// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	// types : Country Text
	AddressCountry []interface{} `json:"addressCountry,omitempty"`

	// AddressLocality see : https://schema.org/addressLocality
	// The locality. For example, Mountain View.
	// types : Text
	AddressLocality []string `json:"addressLocality,omitempty"`

	// AddressRegion see : https://schema.org/addressRegion
	// The region. For example, CA.
	// types : Text
	AddressRegion []string `json:"addressRegion,omitempty"`

	// PostOfficeBoxNumber see : https://schema.org/postOfficeBoxNumber
	// The post office box number for PO box addresses.
	// types : Text
	PostOfficeBoxNumber []string `json:"postOfficeBoxNumber,omitempty"`

	// PostalCode see : https://schema.org/postalCode
	// The postal code. For example, 94043.
	// types : Text
	PostalCode []string `json:"postalCode,omitempty"`

	// StreetAddress see : https://schema.org/streetAddress
	// The street address. For example, 1600 Amphitheatre Pkwy.
	// types : Text
	StreetAddress []string `json:"streetAddress,omitempty"`
}

func (v PostalAddress) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ContactPoint.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AddressCountry
		if len(v.AddressCountry) == 1 {
			value = v.AddressCountry[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["addressCountry"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AddressLocality
		if len(v.AddressLocality) == 1 {
			value = v.AddressLocality[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["addressLocality"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AddressRegion
		if len(v.AddressRegion) == 1 {
			value = v.AddressRegion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["addressRegion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PostOfficeBoxNumber
		if len(v.PostOfficeBoxNumber) == 1 {
			value = v.PostOfficeBoxNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["postOfficeBoxNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PostalCode
		if len(v.PostalCode) == 1 {
			value = v.PostalCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["postalCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StreetAddress
		if len(v.StreetAddress) == 1 {
			value = v.StreetAddress[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["streetAddress"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PostalAddress) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PostalAddress"

	return data, nil
}

func (v PostalAddress) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
