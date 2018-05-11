package schemaorg

import "encoding/json"

// OpeningHoursSpecification see : https://schema.org/OpeningHoursSpecification
type OpeningHoursSpecification struct {

typeContext

StructuredValue

// Closes see : /closes
// The closing hour of the place or service on the given day(s) of the week.
Closes interface{} `json:"closes"`

// DayOfWeek see : /dayOfWeek
// The day of the week for which these opening hours are valid.
DayOfWeek *DayOfWeek `json:"dayOfWeek"`

// Opens see : /opens
// The opening hour of the place or service on the given day(s) of the week.
Opens interface{} `json:"opens"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom interface{} `json:"validFrom"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough interface{} `json:"validThrough"`

}

func (v *OpeningHoursSpecification) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OpeningHoursSpecification"

	return json.Marshal(v)
}
