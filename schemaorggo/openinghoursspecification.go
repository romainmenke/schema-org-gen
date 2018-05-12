package schemaorggo

import "encoding/json"

// OpeningHoursSpecification see : https://schema.org/OpeningHoursSpecification
type OpeningHoursSpecification struct {
	StructuredValue

	typeContext

	// Closes see : https://schema.org/closes
	// The closing hour of the place or service on the given day(s) of the week.
	// types : Time
	Closes Time `json:"closes,omitempty"`

	// DayOfWeek see : https://schema.org/dayOfWeek
	// The day of the week for which these opening hours are valid.
	// types : DayOfWeek
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`

	// Opens see : https://schema.org/opens
	// The opening hour of the place or service on the given day(s) of the week.
	// types : Time
	Opens Time `json:"opens,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough DateTime `json:"validThrough,omitempty"`
}

func (v OpeningHoursSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OpeningHoursSpecification"

	return json.Marshal(v)
}

func (v *OpeningHoursSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OpeningHoursSpecification"

	return json.Marshal(*v)
}
