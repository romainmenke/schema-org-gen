package schemaorggo

import "encoding/json"

// ParcelDelivery see : https://schema.org/ParcelDelivery
type ParcelDelivery struct {

	// With properties from Intangible see : https://schema.org/Intangible
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

	// Carrier see : https://schema.org/carrier
	// &#39;carrier&#39; is an out-dated term indicating the &#39;provider&#39; for parcel delivery and flights.
	// types : Organization
	Carrier []*Organization `json:"carrier,omitempty"`

	// DeliveryAddress see : https://schema.org/deliveryAddress
	// Destination address.
	// types : PostalAddress
	DeliveryAddress []*PostalAddress `json:"deliveryAddress,omitempty"`

	// DeliveryStatus see : https://schema.org/deliveryStatus
	// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	// types : DeliveryEvent
	DeliveryStatus []*DeliveryEvent `json:"deliveryStatus,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// ExpectedArrivalFrom see : https://schema.org/expectedArrivalFrom
	// The earliest date the package may arrive.
	// types : DateTime
	ExpectedArrivalFrom []DateTime `json:"expectedArrivalFrom,omitempty"`

	// ExpectedArrivalUntil see : https://schema.org/expectedArrivalUntil
	// The latest date the package may arrive.
	// types : DateTime
	ExpectedArrivalUntil []DateTime `json:"expectedArrivalUntil,omitempty"`

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	// types : DeliveryMethod
	HasDeliveryMethod []*DeliveryMethod `json:"hasDeliveryMethod,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// ItemShipped see : https://schema.org/itemShipped
	// Item(s) being shipped.
	// types : Product
	ItemShipped []*Product `json:"itemShipped,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// OriginAddress see : https://schema.org/originAddress
	// Shipper&#39;s address.
	// types : PostalAddress
	OriginAddress []*PostalAddress `json:"originAddress,omitempty"`

	// PartOfOrder see : https://schema.org/partOfOrder
	// The overall order the items in this delivery were included in.
	// types : Order
	PartOfOrder []*Order `json:"partOfOrder,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
	Provider []interface{} `json:"provider,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// TrackingNumber see : https://schema.org/trackingNumber
	// Shipper tracking number.
	// types : Text
	TrackingNumber []string `json:"trackingNumber,omitempty"`

	// TrackingUrl see : https://schema.org/trackingUrl
	// Tracking url for the parcel delivery.
	// types : URL
	TrackingUrl []string `json:"trackingUrl,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v ParcelDelivery) MarshalJSON() ([]byte, error) {
	type Alias ParcelDelivery

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"ParcelDelivery\","), b[1:]...), nil
}
