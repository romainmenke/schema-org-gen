package schemaorggo

import "encoding/json"

// Flight see : https://schema.org/Flight
type Flight struct {
	typeContext

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Trip see : https://schema.org/Trip
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// Aircraft see : https://schema.org/aircraft
	// The kind of aircraft (e.g., &quot;Boeing 747&quot;).
	// types : Text Vehicle
	Aircraft []interface{} `json:"aircraft,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// ArrivalAirport see : https://schema.org/arrivalAirport
	// The airport where the flight terminates.
	// types : Airport
	ArrivalAirport []*Airport `json:"arrivalAirport,omitempty"`

	// ArrivalGate see : https://schema.org/arrivalGate
	// Identifier of the flight&#39;s arrival gate.
	// types : Text
	ArrivalGate []string `json:"arrivalGate,omitempty"`

	// ArrivalTerminal see : https://schema.org/arrivalTerminal
	// Identifier of the flight&#39;s arrival terminal.
	// types : Text
	ArrivalTerminal []string `json:"arrivalTerminal,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime []DateTime `json:"arrivalTime,omitempty"`

	// BoardingPolicy see : https://schema.org/boardingPolicy
	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	// types : BoardingPolicyType
	BoardingPolicy []*BoardingPolicyType `json:"boardingPolicy,omitempty"`

	// DepartureAirport see : https://schema.org/departureAirport
	// The airport where the flight originates.
	// types : Airport
	DepartureAirport []*Airport `json:"departureAirport,omitempty"`

	// DepartureGate see : https://schema.org/departureGate
	// Identifier of the flight&#39;s departure gate.
	// types : Text
	DepartureGate []string `json:"departureGate,omitempty"`

	// DepartureTerminal see : https://schema.org/departureTerminal
	// Identifier of the flight&#39;s departure terminal.
	// types : Text
	DepartureTerminal []string `json:"departureTerminal,omitempty"`

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

	// EstimatedFlightDuration see : https://schema.org/estimatedFlightDuration
	// The estimated time the flight will take.
	// types : Duration Text
	EstimatedFlightDuration []interface{} `json:"estimatedFlightDuration,omitempty"`

	// FlightDistance see : https://schema.org/flightDistance
	// The distance of the flight.
	// types : Distance Text
	FlightDistance []interface{} `json:"flightDistance,omitempty"`

	// FlightNumber see : https://schema.org/flightNumber
	// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is &#39;UA&#39;, the flightNumber is &#39;UA110&#39;.
	// types : Text
	FlightNumber []string `json:"flightNumber,omitempty"`

	// HasPart see : https://schema.org/hasPart
	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense). Inverse property: isPartOf (see: https://schema.org/isPartOf).
	// types : CreativeWork Trip
	HasPart []interface{} `json:"hasPart,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
	Image []interface{} `json:"image,omitempty"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	// types : CreativeWork Trip
	IsPartOf []interface{} `json:"isPartOf,omitempty"`

	// Itinerary see : https://pending.schema.org/itinerary
	// Destination(s) ( Place (see: https://schema.org/Place) ) that make up a trip. For a trip where destination order is important use ItemList (see: https://schema.org/ItemList) to specify that order (see examples).
	// types : ItemList Place
	Itinerary []interface{} `json:"itinerary,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MealService see : https://schema.org/mealService
	// Description of the meals that will be provided or available for purchase.
	// types : Text
	MealService []string `json:"mealService,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// WebCheckinTime see : https://schema.org/webCheckinTime
	// The time when a passenger can check into the flight online.
	// types : DateTime
	WebCheckinTime []DateTime `json:"webCheckinTime,omitempty"`
}

func (v Flight) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AdditionalType
		if len(v.AdditionalType) == 1 {
			value = v.AdditionalType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Aircraft
		if len(v.Aircraft) == 1 {
			value = v.Aircraft[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aircraft"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlternateName
		if len(v.AlternateName) == 1 {
			value = v.AlternateName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alternateName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalAirport
		if len(v.ArrivalAirport) == 1 {
			value = v.ArrivalAirport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalAirport"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalGate
		if len(v.ArrivalGate) == 1 {
			value = v.ArrivalGate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalGate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalTerminal
		if len(v.ArrivalTerminal) == 1 {
			value = v.ArrivalTerminal[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTerminal"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalTime
		if len(v.ArrivalTime) == 1 {
			value = v.ArrivalTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BoardingPolicy
		if len(v.BoardingPolicy) == 1 {
			value = v.BoardingPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["boardingPolicy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureAirport
		if len(v.DepartureAirport) == 1 {
			value = v.DepartureAirport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureAirport"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureGate
		if len(v.DepartureGate) == 1 {
			value = v.DepartureGate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureGate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTerminal
		if len(v.DepartureTerminal) == 1 {
			value = v.DepartureTerminal[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTerminal"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTime
		if len(v.DepartureTime) == 1 {
			value = v.DepartureTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Description
		if len(v.Description) == 1 {
			value = v.Description[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["description"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DisambiguatingDescription
		if len(v.DisambiguatingDescription) == 1 {
			value = v.DisambiguatingDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["disambiguatingDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EstimatedFlightDuration
		if len(v.EstimatedFlightDuration) == 1 {
			value = v.EstimatedFlightDuration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["estimatedFlightDuration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FlightDistance
		if len(v.FlightDistance) == 1 {
			value = v.FlightDistance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["flightDistance"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FlightNumber
		if len(v.FlightNumber) == 1 {
			value = v.FlightNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["flightNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasPart
		if len(v.HasPart) == 1 {
			value = v.HasPart[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasPart"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Identifier
		if len(v.Identifier) == 1 {
			value = v.Identifier[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["identifier"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Image
		if len(v.Image) == 1 {
			value = v.Image[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["image"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsPartOf
		if len(v.IsPartOf) == 1 {
			value = v.IsPartOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isPartOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Itinerary
		if len(v.Itinerary) == 1 {
			value = v.Itinerary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itinerary"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainEntityOfPage
		if len(v.MainEntityOfPage) == 1 {
			value = v.MainEntityOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainEntityOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MealService
		if len(v.MealService) == 1 {
			value = v.MealService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mealService"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Name
		if len(v.Name) == 1 {
			value = v.Name[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["name"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Offers
		if len(v.Offers) == 1 {
			value = v.Offers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PotentialAction
		if len(v.PotentialAction) == 1 {
			value = v.PotentialAction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["potentialAction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Provider
		if len(v.Provider) == 1 {
			value = v.Provider[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["provider"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SameAs
		if len(v.SameAs) == 1 {
			value = v.SameAs[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sameAs"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Seller
		if len(v.Seller) == 1 {
			value = v.Seller[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seller"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubjectOf
		if len(v.SubjectOf) == 1 {
			value = v.SubjectOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subjectOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Url
		if len(v.Url) == 1 {
			value = v.Url[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["url"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WebCheckinTime
		if len(v.WebCheckinTime) == 1 {
			value = v.WebCheckinTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["webCheckinTime"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Flight) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Flight"

	return data, nil
}

func (v Flight) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
