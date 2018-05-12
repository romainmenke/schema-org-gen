package schemaorggo

import "encoding/json"

// TouristAttraction see : https://schema.org/TouristAttraction
type TouristAttraction struct {
	Place

	typeContext

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	AvailableLanguage interface{} `json:"availableLanguage,omitempty"` // types : Language Text

	// TouristType see : https://schema.org/touristType
	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	TouristType interface{} `json:"touristType,omitempty"` // types : Audience Text

}

func (v TouristAttraction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TouristAttraction"

	return json.Marshal(v)
}

func (v *TouristAttraction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TouristAttraction"

	return json.Marshal(*v)
}
