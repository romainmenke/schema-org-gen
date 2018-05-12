package schemaorggo

import "encoding/json"

// ParcelDelivery see : https://schema.org/ParcelDelivery
type ParcelDelivery struct {
	Intangible

	typeContext

	// DeliveryAddress see : https://schema.org/deliveryAddress
	// Destination address.
	// types : PostalAddress
	DeliveryAddress *PostalAddress `json:"deliveryAddress,omitempty"`

	// DeliveryStatus see : https://schema.org/deliveryStatus
	// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	// types : DeliveryEvent
	DeliveryStatus *DeliveryEvent `json:"deliveryStatus,omitempty"`

	// ExpectedArrivalFrom see : https://schema.org/expectedArrivalFrom
	// The earliest date the package may arrive.
	// types : DateTime
	ExpectedArrivalFrom DateTime `json:"expectedArrivalFrom,omitempty"`

	// ExpectedArrivalUntil see : https://schema.org/expectedArrivalUntil
	// The latest date the package may arrive.
	// types : DateTime
	ExpectedArrivalUntil DateTime `json:"expectedArrivalUntil,omitempty"`

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	// types : DeliveryMethod
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty"`

	// ItemShipped see : https://schema.org/itemShipped
	// Item(s) being shipped.
	// types : Product
	ItemShipped *Product `json:"itemShipped,omitempty"`

	// OriginAddress see : https://schema.org/originAddress
	// Shipper&#39;s address.
	// types : PostalAddress
	OriginAddress *PostalAddress `json:"originAddress,omitempty"`

	// PartOfOrder see : https://schema.org/partOfOrder
	// The overall order the items in this delivery were included in.
	// types : Order
	PartOfOrder *Order `json:"partOfOrder,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider interface{} `json:"provider,omitempty"`

	// TrackingNumber see : https://schema.org/trackingNumber
	// Shipper tracking number.
	// types : Text
	TrackingNumber string `json:"trackingNumber,omitempty"`

	// TrackingUrl see : https://schema.org/trackingUrl
	// Tracking url for the parcel delivery.
	// types : URL
	TrackingUrl string `json:"trackingUrl,omitempty"`
}

func (v ParcelDelivery) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ParcelDelivery"

	return json.Marshal(v)
}

func (v *ParcelDelivery) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ParcelDelivery"

	return json.Marshal(*v)
}
