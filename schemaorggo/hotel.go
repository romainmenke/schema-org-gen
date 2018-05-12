package schemaorggo

import "encoding/json"

// Hotel see : https://schema.org/Hotel
type Hotel struct {
	LodgingBusiness

	typeContext

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty"` // types : LocationFeatureSpecification

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	Audience *Audience `json:"audience,omitempty"` // types : Audience

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	AvailableLanguage interface{} `json:"availableLanguage,omitempty"` // types : Language Text

	// CheckinTime see : https://schema.org/checkinTime
	// The earliest someone may check into a lodging establishment.
	CheckinTime DateTime `json:"checkinTime,omitempty"` // types : DateTime

	// CheckoutTime see : https://schema.org/checkoutTime
	// The latest someone may check out of a lodging establishment.
	CheckoutTime DateTime `json:"checkoutTime,omitempty"` // types : DateTime

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	PetsAllowed interface{} `json:"petsAllowed,omitempty"` // types : Boolean Text

	// StarRating see : https://schema.org/starRating
	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	StarRating *Rating `json:"starRating,omitempty"` // types : Rating

}

func (v Hotel) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Hotel"

	return json.Marshal(v)
}

func (v *Hotel) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Hotel"

	return json.Marshal(*v)
}
