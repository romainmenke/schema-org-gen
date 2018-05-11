package schemaorg

// BroadcastService see : https://schema.org/BroadcastService
type BroadcastService struct {

Service// BroadcastAffiliateOf see : /broadcastAffiliateOf
// The media network(s) whose content is broadcast on this station.
BroadcastAffiliateOf string `json:"broadcastAffiliateOf"`

// BroadcastDisplayName see : /broadcastDisplayName
// The name displayed in the channel guide. For many US affiliates, it is the network name.
BroadcastDisplayName string `json:"broadcastDisplayName"`

// BroadcastFrequency see : http://pending.schema.org/broadcastFrequency
// The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. "87 FM".
BroadcastFrequency interface{} `json:"broadcastFrequency"` // types : BroadcastFrequencySpecification Text

// BroadcastTimezone see : /broadcastTimezone
// The timezone in ISO 8601 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts
BroadcastTimezone string `json:"broadcastTimezone"`

// Broadcaster see : /broadcaster
// The organization owning or operating the broadcast service.
Broadcaster string `json:"broadcaster"`

// HasBroadcastChannel see : http://pending.schema.org/hasBroadcastChannel
// A broadcast channel of a broadcast service. Inverse property: providesBroadcastService (see: https://schema.org/providesBroadcastService).
HasBroadcastChannel string `json:"hasBroadcastChannel"`

// ParentService see : /parentService
// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
ParentService string `json:"parentService"`

// VideoFormat see : /videoFormat
// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
VideoFormat string `json:"videoFormat"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

// AvailableChannel see : /availableChannel
// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
AvailableChannel string `json:"availableChannel"`

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// HasOfferCatalog see : /hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog string `json:"hasOfferCatalog"`

// HoursAvailable see : /hoursAvailable
// The hours during which this service or contact is available.
HoursAvailable string `json:"hoursAvailable"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products).
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products).
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers string `json:"offers"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// ProviderMobility see : /providerMobility
// Indicates the mobility of a provided service (e.g. 'static', 'dynamic').
ProviderMobility string `json:"providerMobility"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// ServiceOutput see : /serviceOutput
// The tangible thing generated by the service, e.g. a passport, permit, etc. Supersedes produces (see: https://schema.org/produces).
ServiceOutput string `json:"serviceOutput"`

// ServiceType see : /serviceType
// The type of service being offered, e.g. veterans' benefits, emergency relief, etc.
ServiceType string `json:"serviceType"`

// TermsOfService see : http://pending.schema.org/termsOfService
// Human-readable terms of service documentation.
TermsOfService interface{} `json:"termsOfService"` // types : Text URL

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// ParentService see : /parentService
// A broadcast service to which the broadcast service may belong to such as regional variations of a national channel. 
ParentService string `json:"parentService"`

// ProvidesBroadcastService see : /providesBroadcastService
// The BroadcastService offered on this channel.  inverse property: hasBroadcastChannel (see: https://schema.orghttp://pending.schema.org/hasBroadcastChannel).
ProvidesBroadcastService string `json:"providesBroadcastService"`

// PublishedOn see : /publishedOn
// A broadcast service associated with the publication event. 
PublishedOn string `json:"publishedOn"`

}

