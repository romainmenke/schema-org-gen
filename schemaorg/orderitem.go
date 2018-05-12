package schemaorg

import "encoding/json"

// OrderItem see : https://schema.org/OrderItem
type OrderItem struct {

Intangible

typeContext

// OrderDelivery see : https://schema.org/orderDelivery
// The delivery of the parcel related to this order or order item.
OrderDelivery *ParcelDelivery `json:"orderDelivery"`

// OrderItemNumber see : https://schema.org/orderItemNumber
// The identifier of the order item.
OrderItemNumber string `json:"orderItemNumber"`

// OrderItemStatus see : https://schema.org/orderItemStatus
// The current status of the order item.
OrderItemStatus *OrderStatus `json:"orderItemStatus"`

// OrderQuantity see : https://schema.org/orderQuantity
// The number of the item ordered. If the property is not set, assume the quantity is one.
OrderQuantity float64 `json:"orderQuantity"`

// OrderedItem see : https://schema.org/orderedItem
// The item ordered.
OrderedItem interface{} `json:"orderedItem"` // types : OrderItem Product

}

func (v *OrderItem) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OrderItem"

	return json.Marshal(v)
}
