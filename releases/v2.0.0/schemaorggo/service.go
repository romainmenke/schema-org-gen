package schemaorggo

import "encoding/json"

// Service see : https://schema.org/Service
type Service struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AvailableChannel see : https://schema.org/availableChannel
	// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
	// types : ServiceChannel
	AvailableChannel []*ServiceChannel `json:"availableChannel,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

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

	// Produces see : https://schema.org/produces
	// The tangible thing generated by the service, e.g. a passport, permit, etc.
	// types : Thing
	Produces []*Thing `json:"produces,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
	Provider []interface{} `json:"provider,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// ServiceArea see : https://schema.org/serviceArea
	// The geographic area where the service is provided.
	// types : AdministrativeArea
	ServiceArea []*AdministrativeArea `json:"serviceArea,omitempty"`

	// ServiceAudience see : https://schema.org/serviceAudience
	// The audience eligible for this service.
	// types : Audience
	ServiceAudience []*Audience `json:"serviceAudience,omitempty"`

	// ServiceOutput see : https://schema.org/serviceOutput
	// The tangible thing generated by the service, e.g. a passport, permit, etc.
	// types : Thing
	ServiceOutput []*Thing `json:"serviceOutput,omitempty"`

	// ServiceType see : https://schema.org/serviceType
	// The type of service being offered, e.g. veterans&#39; benefits, emergency relief, etc.
	// types : Text
	ServiceType []string `json:"serviceType,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v Service) MarshalJSON() ([]byte, error) {
	type Alias Service

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Service\","), b[1:]...), nil
}
