package schemaorg

import "encoding/json"

// OfferItemCondition see : https://schema.org/OfferItemCondition
type OfferItemCondition struct {

typeContext

Enumeration

// SupersededBy see : https://schema.orghttp://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *OfferItemCondition) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OfferItemCondition"

	return json.Marshal(v)
}
