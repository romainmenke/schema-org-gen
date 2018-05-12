package schemaorggo

import "encoding/json"

// NewsArticle see : https://schema.org/NewsArticle
type NewsArticle struct {
	Article

	typeContext

	// Dateline see : https://schema.org/dateline
	// A dateline (see: https://schema.orghttps://en.wikipedia.org/wiki/Dateline) is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.
	Dateline string `json:"dateline"`

	// PrintColumn see : https://schema.org/printColumn
	// The number of the column in which the NewsArticle appears in the print edition.
	PrintColumn string `json:"printColumn"`

	// PrintEdition see : https://schema.org/printEdition
	// The edition of the print product in which the NewsArticle appears.
	PrintEdition string `json:"printEdition"`

	// PrintPage see : https://schema.org/printPage
	// If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
	PrintPage string `json:"printPage"`

	// PrintSection see : https://schema.org/printSection
	// If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
	PrintSection string `json:"printSection"`
}

func (v *NewsArticle) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NewsArticle"

	return json.Marshal(v)
}
