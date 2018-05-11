package schemaorg

import "encoding/json"

// Campground see : https://schema.org/Campground
type Campground struct {

typeContext

CivicStructure

// AmenityFeature see : /amenityFeature
// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
AmenityFeature *LocationFeatureSpecification `json:"amenityFeature"`

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience *Audience `json:"audience"`

// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// CheckinTime see : /checkinTime
// The earliest someone may check into a lodging establishment.
CheckinTime interface{} `json:"checkinTime"`

// CheckoutTime see : /checkoutTime
// The latest someone may check out of a lodging establishment.
CheckoutTime interface{} `json:"checkoutTime"`

// PetsAllowed see : /petsAllowed
// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
PetsAllowed interface{} `json:"petsAllowed"` // types : Boolean Text

// StarRating see : /starRating
// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
StarRating *Rating `json:"starRating"`

}

func (v *Campground) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Campground"

	return json.Marshal(v)
}
