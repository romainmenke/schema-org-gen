package schemaorggo

import "encoding/json"

// OfferItemCondition see : https://schema.org/OfferItemCondition
type OfferItemCondition struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v OfferItemCondition) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OfferItemCondition"

	return json.Marshal(v)
}

func (v *OfferItemCondition) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OfferItemCondition"

	return json.Marshal(*v)
}
