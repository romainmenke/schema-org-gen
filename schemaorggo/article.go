package schemaorggo

import "encoding/json"

// Article see : https://schema.org/Article
type Article struct {
	CreativeWork

	typeContext

	// ArticleBody see : https://schema.org/articleBody
	// The actual body of the article.
	// types : Text
	ArticleBody []string `json:"articleBody,omitempty"`

	// ArticleSection see : https://schema.org/articleSection
	// Articles may belong to one or more &#39;sections&#39; in a magazine or newspaper, such as Sports, Lifestyle, etc.
	// types : Text
	ArticleSection []string `json:"articleSection,omitempty"`

	// PageEnd see : https://schema.org/pageEnd
	// The page on which the work ends; for example &quot;138&quot; or &quot;xvi&quot;.
	// types : Integer Text
	PageEnd []interface{} `json:"pageEnd,omitempty"`

	// PageStart see : https://schema.org/pageStart
	// The page on which the work starts; for example &quot;135&quot; or &quot;xiii&quot;.
	// types : Integer Text
	PageStart []interface{} `json:"pageStart,omitempty"`

	// Pagination see : https://schema.org/pagination
	// Any description of pages that is not separated into pageStart and pageEnd; for example, &quot;1-6, 9, 55&quot; or &quot;10-12, 46-49&quot;.
	// types : Text
	Pagination []string `json:"pagination,omitempty"`

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
	Speakable []interface{} `json:"speakable,omitempty"`

	// WordCount see : https://schema.org/wordCount
	// The number of words in the text of the Article.
	// types : Integer
	WordCount []float64 `json:"wordCount,omitempty"`
}

func (v Article) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ArticleBody
		if len(v.ArticleBody) == 1 {
			value = v.ArticleBody[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["articleBody"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArticleSection
		if len(v.ArticleSection) == 1 {
			value = v.ArticleSection[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["articleSection"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PageEnd
		if len(v.PageEnd) == 1 {
			value = v.PageEnd[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pageEnd"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PageStart
		if len(v.PageStart) == 1 {
			value = v.PageStart[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pageStart"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Pagination
		if len(v.Pagination) == 1 {
			value = v.Pagination[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pagination"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Speakable
		if len(v.Speakable) == 1 {
			value = v.Speakable[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["speakable"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WordCount
		if len(v.WordCount) == 1 {
			value = v.WordCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["wordCount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Article) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Article"

	return data, nil
}

func (v Article) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
