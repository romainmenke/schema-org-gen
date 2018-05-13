package schemaorggo

import "encoding/json"

// OrderItem see : https://schema.org/OrderItem
type OrderItem struct {
	Intangible

	typeContext

	// OrderDelivery see : https://schema.org/orderDelivery
	// The delivery of the parcel related to this order or order item.
	// types : ParcelDelivery
	OrderDelivery []*ParcelDelivery `json:"orderDelivery,omitempty"`

	// OrderItemNumber see : https://schema.org/orderItemNumber
	// The identifier of the order item.
	// types : Text
	OrderItemNumber []string `json:"orderItemNumber,omitempty"`

	// OrderItemStatus see : https://schema.org/orderItemStatus
	// The current status of the order item.
	// types : OrderStatus
	OrderItemStatus []*OrderStatus `json:"orderItemStatus,omitempty"`

	// OrderQuantity see : https://schema.org/orderQuantity
	// The number of the item ordered. If the property is not set, assume the quantity is one.
	// types : Number
	OrderQuantity []float64 `json:"orderQuantity,omitempty"`

	// OrderedItem see : https://schema.org/orderedItem
	// The item ordered.
	// types : OrderItem Product
	OrderedItem []interface{} `json:"orderedItem,omitempty"`
}

func (v OrderItem) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.OrderDelivery
		if len(v.OrderDelivery) == 1 {
			value = v.OrderDelivery[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderDelivery"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderItemNumber
		if len(v.OrderItemNumber) == 1 {
			value = v.OrderItemNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderItemNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderItemStatus
		if len(v.OrderItemStatus) == 1 {
			value = v.OrderItemStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderItemStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderQuantity
		if len(v.OrderQuantity) == 1 {
			value = v.OrderQuantity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderQuantity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderedItem
		if len(v.OrderedItem) == 1 {
			value = v.OrderedItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderedItem"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v OrderItem) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "OrderItem"

	return data, nil
}

func (v OrderItem) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
