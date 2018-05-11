package schemaorg

import "encoding/json"

// TaxiService see : https://schema.org/TaxiService
type TaxiService struct {

typeContext

Service

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating *AggregateRating `json:"aggregateRating"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience *Audience `json:"audience"`

// AvailableChannel see : /availableChannel
// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
AvailableChannel *ServiceChannel `json:"availableChannel"`

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
HasOfferCatalog *OfferCatalog `json:"hasOfferCatalog"`

// HoursAvailable see : /hoursAvailable
// The hours during which this service or contact is available.
HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable"`

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
Offers *Offer `json:"offers"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// ProviderMobility see : /providerMobility
// Indicates the mobility of a provided service (e.g. 'static', 'dynamic').
ProviderMobility string `json:"providerMobility"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review *Review `json:"review"`

// ServiceOutput see : /serviceOutput
// The tangible thing generated by the service, e.g. a passport, permit, etc. Supersedes produces (see: https://schema.org/produces).
ServiceOutput *Thing `json:"serviceOutput"`

// ServiceType see : /serviceType
// The type of service being offered, e.g. veterans' benefits, emergency relief, etc.
ServiceType string `json:"serviceType"`

// TermsOfService see : http://pending.schema.org/termsOfService
// Human-readable terms of service documentation.
TermsOfService interface{} `json:"termsOfService"` // types : Text URL

}

func (v *TaxiService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TaxiService"

	return json.Marshal(v)
}
