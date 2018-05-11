package schemaorg

import "encoding/json"

// TouristAttraction see : https://schema.org/TouristAttraction
type TouristAttraction struct {

typeContext

Place

// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// TouristType see : /touristType
// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
TouristType interface{} `json:"touristType"` // types : Audience Text

}

func (v *TouristAttraction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TouristAttraction"

	return json.Marshal(v)
}
