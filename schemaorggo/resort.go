package schemaorggo

import "encoding/json"

// Resort see : https://schema.org/Resort
type Resort struct {
	LodgingBusiness

	typeContext

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature []*LocationFeatureSpecification `json:"amenityFeature,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	// types : Language Text
	AvailableLanguage []interface{} `json:"availableLanguage,omitempty"`

	// CheckinTime see : https://schema.org/checkinTime
	// The earliest someone may check into a lodging establishment.
	// types : DateTime
	CheckinTime []DateTime `json:"checkinTime,omitempty"`

	// CheckoutTime see : https://schema.org/checkoutTime
	// The latest someone may check out of a lodging establishment.
	// types : DateTime
	CheckoutTime []DateTime `json:"checkoutTime,omitempty"`

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	// types : Boolean Text
	PetsAllowed []interface{} `json:"petsAllowed,omitempty"`

	// StarRating see : https://schema.org/starRating
	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	// types : Rating
	StarRating []*Rating `json:"starRating,omitempty"`
}

func (v Resort) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.LodgingBusiness.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AmenityFeature
		if len(v.AmenityFeature) == 1 {
			value = v.AmenityFeature[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["amenityFeature"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Audience
		if len(v.Audience) == 1 {
			value = v.Audience[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audience"] = json.RawMessage(b)
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
		var value interface{} = v.CheckinTime
		if len(v.CheckinTime) == 1 {
			value = v.CheckinTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["checkinTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CheckoutTime
		if len(v.CheckoutTime) == 1 {
			value = v.CheckoutTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["checkoutTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PetsAllowed
		if len(v.PetsAllowed) == 1 {
			value = v.PetsAllowed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["petsAllowed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StarRating
		if len(v.StarRating) == 1 {
			value = v.StarRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["starRating"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Resort) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Resort"

	return data, nil
}

func (v Resort) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
