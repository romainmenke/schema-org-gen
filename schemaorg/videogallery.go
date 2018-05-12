package schemaorg

import "encoding/json"

// VideoGallery see : https://schema.org/VideoGallery
type VideoGallery struct {

CollectionPage

typeContext

// Breadcrumb see : https://schema.org/breadcrumb
// A set of links that can help a user understand and navigate a website hierarchy.
Breadcrumb interface{} `json:"breadcrumb"` // types : BreadcrumbList Text

// LastReviewed see : https://schema.org/lastReviewed
// Date on which the content on this web page was last reviewed for accuracy and/or completeness.
LastReviewed interface{} `json:"lastReviewed"`

// MainContentOfPage see : https://schema.org/mainContentOfPage
// Indicates if this web page element is the main subject of the page. Supersedes aspect (see: https://schema.orghttp://health-lifesci.schema.org/aspect).
MainContentOfPage *WebPageElement `json:"mainContentOfPage"`

// PrimaryImageOfPage see : https://schema.org/primaryImageOfPage
// Indicates the main image on the page.
PrimaryImageOfPage *ImageObject `json:"primaryImageOfPage"`

// RelatedLink see : https://schema.org/relatedLink
// A link related to this web page, for example to other related web pages.
RelatedLink string `json:"relatedLink"`

// ReviewedBy see : https://schema.org/reviewedBy
// People or organizations that have reviewed the content on this web page for accuracy and/or completeness.
ReviewedBy interface{} `json:"reviewedBy"` // types : Organization Person

// SignificantLink see : https://schema.org/significantLink
// One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most. Supersedes significantLinks (see: https://schema.org/significantLinks).
SignificantLink string `json:"significantLink"`

// Speakable see : http://pending.schema.org/speakable
// Indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the 'speakable' property serves to indicate the parts most likely to be generally useful for speech.
// 
// The speakable property can be repeated an arbitrary number of times, with three kinds of possible 'content-locator' values:
// 
// 1.) id-value URL references - uses id-value of an element in the page being annotated. The simplest use of speakable has (potentially relative) URL values, referencing identified sections of the document concerned.
// 
// 2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the cssSelector (see: https://schema.org/cssSelector) property.
// 
// 3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the xpath (see: https://schema.org/xpath) property.
// 
// For more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this
// we define a supporting type, SpeakableSpecification (see: https://schema.org/SpeakableSpecification)  which is defined to be a possible value of the speakable property.
Speakable interface{} `json:"speakable"` // types : SpeakableSpecification URL

// Specialty see : https://schema.org/specialty
// One of the domain specialities to which this web page's content applies.
Specialty *Specialty `json:"specialty"`

}

func (v *VideoGallery) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGallery"

	return json.Marshal(v)
}
