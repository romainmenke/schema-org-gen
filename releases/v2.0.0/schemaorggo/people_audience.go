package schemaorggo

import "encoding/json"

// PeopleAudience see : https://schema.org/PeopleAudience
type PeopleAudience struct {

	// With properties from Audience see : https://schema.org/Audience
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AudienceType see : https://schema.org/audienceType
	// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	// types : Text
	AudienceType []string `json:"audienceType,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// GeographicArea see : https://schema.org/geographicArea
	// The geographic area associated with the audience.
	// types : AdministrativeArea
	GeographicArea []*AdministrativeArea `json:"geographicArea,omitempty"`

	// HealthCondition see : https://schema.org/healthCondition
	// Expectations for health conditions of target audience.
	// types : MedicalCondition
	HealthCondition []*MedicalCondition `json:"healthCondition,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RequiredGender see : https://schema.org/requiredGender
	// Audiences defined by a person&#39;s gender.
	// types : Text
	RequiredGender []string `json:"requiredGender,omitempty"`

	// RequiredMaxAge see : https://schema.org/requiredMaxAge
	// Audiences defined by a person&#39;s maximum age.
	// types : Integer
	RequiredMaxAge []float64 `json:"requiredMaxAge,omitempty"`

	// RequiredMinAge see : https://schema.org/requiredMinAge
	// Audiences defined by a person&#39;s minimum age.
	// types : Integer
	RequiredMinAge []float64 `json:"requiredMinAge,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SuggestedGender see : https://schema.org/suggestedGender
	// The gender of the person or audience.
	// types : Text
	SuggestedGender []string `json:"suggestedGender,omitempty"`

	// SuggestedMaxAge see : https://schema.org/suggestedMaxAge
	// Maximal age recommended for viewing content.
	// types : Number
	SuggestedMaxAge []float64 `json:"suggestedMaxAge,omitempty"`

	// SuggestedMinAge see : https://schema.org/suggestedMinAge
	// Minimal age recommended for viewing content.
	// types : Number
	SuggestedMinAge []float64 `json:"suggestedMinAge,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v PeopleAudience) MarshalJSON() ([]byte, error) {
	type Alias PeopleAudience

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"PeopleAudience\","), b[1:]...), nil
}