package schemaorggo

import "encoding/json"

// ContactPoint see : https://schema.org/ContactPoint
type ContactPoint struct {
	StructuredValue

	typeContext

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	// types : Language Text
	AvailableLanguage []interface{} `json:"availableLanguage,omitempty"`

	// ContactOption see : https://schema.org/contactOption
	// An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
	// types : ContactPointOption
	ContactOption []*ContactPointOption `json:"contactOption,omitempty"`

	// ContactType see : https://schema.org/contactType
	// A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
	// types : Text
	ContactType []string `json:"contactType,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email []string `json:"email,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// HoursAvailable see : https://schema.org/hoursAvailable
	// The hours during which this service or contact is available.
	// types : OpeningHoursSpecification
	HoursAvailable []*OpeningHoursSpecification `json:"hoursAvailable,omitempty"`

	// ProductSupported see : https://schema.org/productSupported
	// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. &quot;iPhone&quot;) or a general category of products or services (e.g. &quot;smartphones&quot;).
	// types : Product Text
	ProductSupported []interface{} `json:"productSupported,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone []string `json:"telephone,omitempty"`
}

func (v ContactPoint) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AreaServed
		if len(v.AreaServed) == 1 {
			value = v.AreaServed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["areaServed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailableLanguage
		if len(v.AvailableLanguage) == 1 {
			value = v.AvailableLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContactOption
		if len(v.ContactOption) == 1 {
			value = v.ContactOption[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contactOption"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContactType
		if len(v.ContactType) == 1 {
			value = v.ContactType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contactType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Email
		if len(v.Email) == 1 {
			value = v.Email[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["email"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FaxNumber
		if len(v.FaxNumber) == 1 {
			value = v.FaxNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["faxNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HoursAvailable
		if len(v.HoursAvailable) == 1 {
			value = v.HoursAvailable[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hoursAvailable"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductSupported
		if len(v.ProductSupported) == 1 {
			value = v.ProductSupported[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productSupported"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Telephone
		if len(v.Telephone) == 1 {
			value = v.Telephone[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["telephone"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ContactPoint) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ContactPoint"

	return data, nil
}

func (v ContactPoint) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
