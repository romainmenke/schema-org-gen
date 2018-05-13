package schemaorggo

import "encoding/json"

// Periodical see : https://schema.org/Periodical
type Periodical struct {
	CreativeWorkSeries

	typeContext

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate []interface{} `json:"endDate,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn []string `json:"issn,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate []interface{} `json:"startDate,omitempty"`
}

func (v Periodical) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWorkSeries.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.EndDate
		if len(v.EndDate) == 1 {
			value = v.EndDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Issn
		if len(v.Issn) == 1 {
			value = v.Issn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartDate
		if len(v.StartDate) == 1 {
			value = v.StartDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startDate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Periodical) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Periodical"

	return data, nil
}

func (v Periodical) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
