package schemaorggo

import "encoding/json"

// BookSeries see : https://schema.org/BookSeries
type BookSeries struct {
	CreativeWorkSeries

	typeContext

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate interface{} `json:"endDate,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn string `json:"issn,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate interface{} `json:"startDate,omitempty"`
}

func (v BookSeries) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BookSeries"

	return json.Marshal(v)
}

func (v *BookSeries) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BookSeries"

	return json.Marshal(*v)
}
