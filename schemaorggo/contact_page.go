package schemaorggo

import "encoding/json"

// ContactPage see : https://schema.org/ContactPage
type ContactPage struct {
	WebPage

	typeContext

	// Breadcrumb see : https://schema.org/breadcrumb
	// A set of links that can help a user understand and navigate a website hierarchy.
	// types : BreadcrumbList Text
	Breadcrumb []interface{} `json:"breadcrumb,omitempty"`

	// LastReviewed see : https://schema.org/lastReviewed
	// Date on which the content on this web page was last reviewed for accuracy and/or completeness.
	// types : Date
	LastReviewed []Date `json:"lastReviewed,omitempty"`

	// MainContentOfPage see : https://schema.org/mainContentOfPage
	// Indicates if this web page element is the main subject of the page. Supersedes aspect (see: https://schema.orghttps://health-lifesci.schema.org/aspect).
	// types : WebPageElement
	MainContentOfPage []*WebPageElement `json:"mainContentOfPage,omitempty"`

	// PrimaryImageOfPage see : https://schema.org/primaryImageOfPage
	// Indicates the main image on the page.
	// types : ImageObject
	PrimaryImageOfPage []*ImageObject `json:"primaryImageOfPage,omitempty"`

	// RelatedLink see : https://schema.org/relatedLink
	// A link related to this web page, for example to other related web pages.
	// types : URL
	RelatedLink []string `json:"relatedLink,omitempty"`

	// ReviewedBy see : https://schema.org/reviewedBy
	// People or organizations that have reviewed the content on this web page for accuracy and/or completeness.
	// types : Organization Person
	ReviewedBy []interface{} `json:"reviewedBy,omitempty"`

	// SignificantLink see : https://schema.org/significantLink
	// One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most. Supersedes significantLinks (see: https://schema.org/significantLinks).
	// types : URL
	SignificantLink []string `json:"significantLink,omitempty"`

	// Speakable see : https://pending.schema.org/speakable
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

	// Specialty see : https://schema.org/specialty
	// One of the domain specialities to which this web page&#39;s content applies.
	// types : Specialty
	Specialty []*Specialty `json:"specialty,omitempty"`
}

func (v ContactPage) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.WebPage.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Breadcrumb
		if len(v.Breadcrumb) == 1 {
			value = v.Breadcrumb[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["breadcrumb"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LastReviewed
		if len(v.LastReviewed) == 1 {
			value = v.LastReviewed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lastReviewed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainContentOfPage
		if len(v.MainContentOfPage) == 1 {
			value = v.MainContentOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainContentOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PrimaryImageOfPage
		if len(v.PrimaryImageOfPage) == 1 {
			value = v.PrimaryImageOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["primaryImageOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RelatedLink
		if len(v.RelatedLink) == 1 {
			value = v.RelatedLink[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["relatedLink"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReviewedBy
		if len(v.ReviewedBy) == 1 {
			value = v.ReviewedBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reviewedBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SignificantLink
		if len(v.SignificantLink) == 1 {
			value = v.SignificantLink[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["significantLink"] = json.RawMessage(b)
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
		var value interface{} = v.Specialty
		if len(v.Specialty) == 1 {
			value = v.Specialty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["specialty"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ContactPage) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ContactPage"

	return data, nil
}

func (v ContactPage) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
