package schemaorggo

import "encoding/json"

// BusTrip see : https://schema.org/BusTrip
type BusTrip struct {

	// With properties from Trip see : https://schema.org/Trip
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

	// ArrivalBusStop see : https://schema.org/arrivalBusStop
	// The stop or station from which the bus arrives.
	// types : BusStation BusStop
	ArrivalBusStop []interface{} `json:"arrivalBusStop,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime []DateTime `json:"arrivalTime,omitempty"`

	// BusName see : https://schema.org/busName
	// The name of the bus (e.g. Bolt Express).
	// types : Text
	BusName []string `json:"busName,omitempty"`

	// BusNumber see : https://schema.org/busNumber
	// The unique identifier for the bus.
	// types : Text
	BusNumber []string `json:"busNumber,omitempty"`

	// DepartureBusStop see : https://schema.org/departureBusStop
	// The stop or station from which the bus departs.
	// types : BusStation BusStop
	DepartureBusStop []interface{} `json:"departureBusStop,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	// types : DateTime
	DepartureTime []DateTime `json:"departureTime,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
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

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
	Provider []interface{} `json:"provider,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SubjectOf see : https://schema.org/subjectOf
	// A CreativeWork or Event about this Thing..
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v BusTrip) MarshalJSON() ([]byte, error) {
	type Alias BusTrip

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"BusTrip\","), b[1:]...), nil
}
