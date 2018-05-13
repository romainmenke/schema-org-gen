package schemaorggo

import "encoding/json"

// NewsArticle see : https://schema.org/NewsArticle
type NewsArticle struct {
	Article

	typeContext

	// Dateline see : https://schema.org/dateline
	// A dateline (see: https://schema.orghttps://en.wikipedia.org/wiki/Dateline) is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.
	// types : Text
	Dateline []string `json:"dateline,omitempty"`

	// PrintColumn see : https://schema.org/printColumn
	// The number of the column in which the NewsArticle appears in the print edition.
	// types : Text
	PrintColumn []string `json:"printColumn,omitempty"`

	// PrintEdition see : https://schema.org/printEdition
	// The edition of the print product in which the NewsArticle appears.
	// types : Text
	PrintEdition []string `json:"printEdition,omitempty"`

	// PrintPage see : https://schema.org/printPage
	// If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
	// types : Text
	PrintPage []string `json:"printPage,omitempty"`

	// PrintSection see : https://schema.org/printSection
	// If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
	// types : Text
	PrintSection []string `json:"printSection,omitempty"`
}

func (v NewsArticle) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Article.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Dateline
		if len(v.Dateline) == 1 {
			value = v.Dateline[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateline"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrintColumn
		if len(v.PrintColumn) == 1 {
			value = v.PrintColumn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["printColumn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrintEdition
		if len(v.PrintEdition) == 1 {
			value = v.PrintEdition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["printEdition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrintPage
		if len(v.PrintPage) == 1 {
			value = v.PrintPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["printPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrintSection
		if len(v.PrintSection) == 1 {
			value = v.PrintSection[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["printSection"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v NewsArticle) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "NewsArticle"

	return data, nil
}

func (v NewsArticle) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
