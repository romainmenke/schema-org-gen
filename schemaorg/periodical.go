package schemaorg

import "encoding/json"

// Periodical see : https://schema.org/Periodical
type Periodical struct {

typeContext

CreativeWorkSeries

// EndDate see : /endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// Issn see : /issn
// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
Issn string `json:"issn"`

// StartDate see : /startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

}

func (v *Periodical) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Periodical"

	return json.Marshal(v)
}
