package schemaorggo

import "encoding/json"

// ScholarlyArticle see : https://schema.org/ScholarlyArticle
type ScholarlyArticle struct {
	Article

	typeContext

	// ArticleBody see : https://schema.org/articleBody
	// The actual body of the article.
	// types : Text
	ArticleBody string `json:"articleBody,omitempty"`

	// ArticleSection see : https://schema.org/articleSection
	// Articles may belong to one or more &#39;sections&#39; in a magazine or newspaper, such as Sports, Lifestyle, etc.
	// types : Text
	ArticleSection string `json:"articleSection,omitempty"`

	// PageEnd see : https://schema.org/pageEnd
	// The page on which the work ends; for example &quot;138&quot; or &quot;xvi&quot;.
	// types : Integer Text
	PageEnd interface{} `json:"pageEnd,omitempty"`

	// PageStart see : https://schema.org/pageStart
	// The page on which the work starts; for example &quot;135&quot; or &quot;xiii&quot;.
	// types : Integer Text
	PageStart interface{} `json:"pageStart,omitempty"`

	// Pagination see : https://schema.org/pagination
	// Any description of pages that is not separated into pageStart and pageEnd; for example, &quot;1-6, 9, 55&quot; or &quot;10-12, 46-49&quot;.
	// types : Text
	Pagination string `json:"pagination,omitempty"`

	// Speakable see : http://pending.schema.org/speakable
	// Indicates sections of a Web page that are particularly &#39;speakable&#39; in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the &#39;speakable&#39; property serves to indicate the parts most likely to be generally useful for speech.
	//
	// The speakable property can be repeated an arbitrary number of times, with three kinds of possible &#39;content-locator&#39; values:
	//
	// 1.) id-value URL references - uses id-value of an element in the page being annotated. The simplest use of speakable has (potentially relative) URL values, referencing identified sections of the document concerned.
	//
	// 2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the cssSelector (see: https://schema.org/cssSelector) property.
	//
	// 3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the xpath (see: https://schema.org/xpath) property.
	//
	// For more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this
	// we define a supporting type, SpeakableSpecification (see: https://schema.org/SpeakableSpecification)  which is defined to be a possible value of the speakable property.
	// types : SpeakableSpecification URL
	Speakable interface{} `json:"speakable,omitempty"`

	// WordCount see : https://schema.org/wordCount
	// The number of words in the text of the Article.
	// types : Integer
	WordCount float64 `json:"wordCount,omitempty"`
}

func (v ScholarlyArticle) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScholarlyArticle"

	return json.Marshal(v)
}

func (v *ScholarlyArticle) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ScholarlyArticle"

	return json.Marshal(*v)
}
