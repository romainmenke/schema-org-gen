package schemaorg

import "encoding/json"

// NewsArticle see : https://schema.org/NewsArticle
type NewsArticle struct {

typeContext

Article

// Dateline see : /dateline
// A dateline (see: https://schema.orghttps://en.wikipedia.org/wiki/Dateline) is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.
Dateline string `json:"dateline"`

// PrintColumn see : /printColumn
// The number of the column in which the NewsArticle appears in the print edition.
PrintColumn string `json:"printColumn"`

// PrintEdition see : /printEdition
// The edition of the print product in which the NewsArticle appears.
PrintEdition string `json:"printEdition"`

// PrintPage see : /printPage
// If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
PrintPage string `json:"printPage"`

// PrintSection see : /printSection
// If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
PrintSection string `json:"printSection"`

}

func (v *NewsArticle) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NewsArticle"

	return json.Marshal(v)
}
