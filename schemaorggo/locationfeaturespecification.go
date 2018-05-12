package schemaorggo

import "encoding/json"

// LocationFeatureSpecification see : https://schema.org/LocationFeatureSpecification
type LocationFeatureSpecification struct {
	PropertyValue

	typeContext

	// HoursAvailable see : https://schema.org/hoursAvailable
	// The hours during which this service or contact is available.
	HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	ValidFrom DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	ValidThrough DateTime `json:"validThrough,omitempty"`
}

func (v LocationFeatureSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LocationFeatureSpecification"

	return json.Marshal(v)
}

func (v *LocationFeatureSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LocationFeatureSpecification"

	return json.Marshal(*v)
}
