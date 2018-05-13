package schemaorggo

import "encoding/json"

// ParcelDelivery see : https://schema.org/ParcelDelivery
type ParcelDelivery struct {
	Intangible

	typeContext

	// DeliveryAddress see : https://schema.org/deliveryAddress
	// Destination address.
	// types : PostalAddress
	DeliveryAddress []*PostalAddress `json:"deliveryAddress,omitempty"`

	// DeliveryStatus see : https://schema.org/deliveryStatus
	// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	// types : DeliveryEvent
	DeliveryStatus []*DeliveryEvent `json:"deliveryStatus,omitempty"`

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

	// ItemShipped see : https://schema.org/itemShipped
	// Item(s) being shipped.
	// types : Product
	ItemShipped []*Product `json:"itemShipped,omitempty"`

	// OriginAddress see : https://schema.org/originAddress
	// Shipper&#39;s address.
	// types : PostalAddress
	OriginAddress []*PostalAddress `json:"originAddress,omitempty"`

	// PartOfOrder see : https://schema.org/partOfOrder
	// The overall order the items in this delivery were included in.
	// types : Order
	PartOfOrder []*Order `json:"partOfOrder,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// TrackingNumber see : https://schema.org/trackingNumber
	// Shipper tracking number.
	// types : Text
	TrackingNumber []string `json:"trackingNumber,omitempty"`

	// TrackingUrl see : https://schema.org/trackingUrl
	// Tracking url for the parcel delivery.
	// types : URL
	TrackingUrl []string `json:"trackingUrl,omitempty"`
}

func (v ParcelDelivery) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.DeliveryAddress
		if len(v.DeliveryAddress) == 1 {
			value = v.DeliveryAddress[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deliveryAddress"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DeliveryStatus
		if len(v.DeliveryStatus) == 1 {
			value = v.DeliveryStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deliveryStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExpectedArrivalFrom
		if len(v.ExpectedArrivalFrom) == 1 {
			value = v.ExpectedArrivalFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["expectedArrivalFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExpectedArrivalUntil
		if len(v.ExpectedArrivalUntil) == 1 {
			value = v.ExpectedArrivalUntil[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["expectedArrivalUntil"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasDeliveryMethod
		if len(v.HasDeliveryMethod) == 1 {
			value = v.HasDeliveryMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasDeliveryMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ItemShipped
		if len(v.ItemShipped) == 1 {
			value = v.ItemShipped[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemShipped"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OriginAddress
		if len(v.OriginAddress) == 1 {
			value = v.OriginAddress[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["originAddress"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfOrder
		if len(v.PartOfOrder) == 1 {
			value = v.PartOfOrder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfOrder"] = json.RawMessage(b)
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
		var value interface{} = v.TrackingNumber
		if len(v.TrackingNumber) == 1 {
			value = v.TrackingNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trackingNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TrackingUrl
		if len(v.TrackingUrl) == 1 {
			value = v.TrackingUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trackingUrl"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ParcelDelivery) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ParcelDelivery"

	return data, nil
}

func (v ParcelDelivery) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
