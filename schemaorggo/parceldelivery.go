package schemaorggo

import "encoding/json"

// ParcelDelivery see : https://schema.org/ParcelDelivery
type ParcelDelivery struct {
	Intangible

	typeContext

	// DeliveryAddress see : https://schema.org/deliveryAddress
	// Destination address.
	DeliveryAddress *PostalAddress `json:"deliveryAddress,omitempty"` // types : PostalAddress

	// DeliveryStatus see : https://schema.org/deliveryStatus
	// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	DeliveryStatus *DeliveryEvent `json:"deliveryStatus,omitempty"` // types : DeliveryEvent

	// ExpectedArrivalFrom see : https://schema.org/expectedArrivalFrom
	// The earliest date the package may arrive.
	ExpectedArrivalFrom DateTime `json:"expectedArrivalFrom,omitempty"` // types : DateTime

	// ExpectedArrivalUntil see : https://schema.org/expectedArrivalUntil
	// The latest date the package may arrive.
	ExpectedArrivalUntil DateTime `json:"expectedArrivalUntil,omitempty"` // types : DateTime

	// HasDeliveryMethod see : https://schema.org/hasDeliveryMethod
	// Method used for delivery or shipping.
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty"` // types : DeliveryMethod

	// ItemShipped see : https://schema.org/itemShipped
	// Item(s) being shipped.
	ItemShipped *Product `json:"itemShipped,omitempty"` // types : Product

	// OriginAddress see : https://schema.org/originAddress
	// Shipper's address.
	OriginAddress *PostalAddress `json:"originAddress,omitempty"` // types : PostalAddress

	// PartOfOrder see : https://schema.org/partOfOrder
	// The overall order the items in this delivery were included in.
	PartOfOrder *Order `json:"partOfOrder,omitempty"` // types : Order

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider,omitempty"` // types : Organization Person

	// TrackingNumber see : https://schema.org/trackingNumber
	// Shipper tracking number.
	TrackingNumber string `json:"trackingNumber,omitempty"` // types : Text

	// TrackingUrl see : https://schema.org/trackingUrl
	// Tracking url for the parcel delivery.
	TrackingUrl string `json:"trackingUrl,omitempty"` // types : URL

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
