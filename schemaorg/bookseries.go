package schemaorg

import "encoding/json"

// BookSeries see : https://schema.org/BookSeries
type BookSeries struct {

typeContext

CreativeWorkSeries

// EndDate see : https://schema.org/endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// Issn see : https://schema.org/issn
// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
Issn string `json:"issn"`

// StartDate see : https://schema.org/startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

}

func (v *BookSeries) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BookSeries"

	return json.Marshal(v)
}
