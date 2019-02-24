package schemaorggo

import "encoding/json"

// BroadcastService see : https://schema.org/BroadcastService
type BroadcastService struct {

	// With properties from Service see : https://schema.org/Service
	//

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

	// Area see : https://schema.org/area
	// The area within which users can expect to reach the broadcast service.
	// types : Place
	Area []*Place `json:"area,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided.
	// types : Place AdministrativeArea GeoShape Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// AvailableChannel see : https://schema.org/availableChannel
	// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
	// types : ServiceChannel
	AvailableChannel []*ServiceChannel `json:"availableChannel,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// BroadcastAffiliateOf see : https://schema.org/broadcastAffiliateOf
	// The media network(s) whose content is broadcast on this station.
	// types : Organization
	BroadcastAffiliateOf []*Organization `json:"broadcastAffiliateOf,omitempty"`

	// BroadcastDisplayName see : https://schema.org/broadcastDisplayName
	// The name displayed in the channel guide. For many US affiliates, it is the network name.
	// types : Text
	BroadcastDisplayName []string `json:"broadcastDisplayName,omitempty"`

	// Broadcaster see : https://schema.org/broadcaster
	// The organization owning or operating the broadcast service.
	// types : Organization
	Broadcaster []*Organization `json:"broadcaster,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category []interface{} `json:"category,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	// types : OfferCatalog
	HasOfferCatalog []*OfferCatalog `json:"hasOfferCatalog,omitempty"`

	// HoursAvailable see : https://schema.org/hoursAvailable
	// The hours during which this service or contact is available.
	// types : OpeningHoursSpecification
	HoursAvailable []*OpeningHoursSpecification `json:"hoursAvailable,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// ParentService see : https://schema.org/parentService
	// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	// types : BroadcastService
	ParentService []*BroadcastService `json:"parentService,omitempty"`

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

	// ProviderMobility see : https://schema.org/providerMobility
	// Indicates the mobility of a provided service (e.g. &#39;static&#39;, &#39;dynamic&#39;).
	// types : Text
	ProviderMobility []string `json:"providerMobility,omitempty"`

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
	// types : Place AdministrativeArea GeoShape
	ServiceArea []interface{} `json:"serviceArea,omitempty"`

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

	// Timezone see : https://schema.org/broadcastTimezone
	// The timezone in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 format&lt;/a&gt; for which the service bases its broadcasts.
	// types : Text
	Timezone []string `json:"timezone,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// VideoFormat see : https://schema.org/videoFormat
	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	// types : Text
	VideoFormat []string `json:"videoFormat,omitempty"`
}

func (v BroadcastService) MarshalJSON() ([]byte, error) {
	type Alias BroadcastService

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"BroadcastService\","), b[1:]...), nil
}
