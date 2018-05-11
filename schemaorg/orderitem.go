package schemaorg

import "encoding/json"

// OrderItem see : https://schema.org/OrderItem
type OrderItem struct {

typeContext

Intangible

// OrderDelivery see : /orderDelivery
// The delivery of the parcel related to this order or order item.
OrderDelivery *ParcelDelivery `json:"orderDelivery"`

// OrderItemNumber see : /orderItemNumber
// The identifier of the order item.
OrderItemNumber string `json:"orderItemNumber"`

// OrderItemStatus see : /orderItemStatus
// The current status of the order item.
OrderItemStatus *OrderStatus `json:"orderItemStatus"`

// OrderQuantity see : /orderQuantity
// The number of the item ordered. If the property is not set, assume the quantity is one.
OrderQuantity float64 `json:"orderQuantity"`

// OrderedItem see : /orderedItem
// The item ordered.
OrderedItem interface{} `json:"orderedItem"` // types : OrderItem Product

}

func (v *OrderItem) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderItem"

	return json.Marshal(v)
}
