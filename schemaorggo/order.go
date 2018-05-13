package schemaorggo

import "encoding/json"

// Order see : https://schema.org/Order
type Order struct {
	Intangible

	typeContext

	// AcceptedOffer see : https://schema.org/acceptedOffer
	// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	// types : Offer
	AcceptedOffer []*Offer `json:"acceptedOffer,omitempty"`

	// BillingAddress see : https://schema.org/billingAddress
	// The billing address for the order.
	// types : PostalAddress
	BillingAddress []*PostalAddress `json:"billingAddress,omitempty"`

	// Broker see : https://schema.org/broker
	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	// types : Organization Person
	Broker []interface{} `json:"broker,omitempty"`

	// ConfirmationNumber see : https://schema.org/confirmationNumber
	// A number that confirms the given order or payment has been received.
	// types : Text
	ConfirmationNumber []string `json:"confirmationNumber,omitempty"`

	// Customer see : https://schema.org/customer
	// Party placing the order or paying the invoice.
	// types : Organization Person
	Customer []interface{} `json:"customer,omitempty"`

	// Discount see : https://schema.org/discount
	// Any discount applied (to an Order).
	// types : Number Text
	Discount []interface{} `json:"discount,omitempty"`

	// DiscountCode see : https://schema.org/discountCode
	// Code used to redeem a discount.
	// types : Text
	DiscountCode []string `json:"discountCode,omitempty"`

	// DiscountCurrency see : https://schema.org/discountCurrency
	// The currency (in 3-letter ISO 4217 format) of the discount.
	// types : Text
	DiscountCurrency []string `json:"discountCurrency,omitempty"`

	// IsGift see : https://schema.org/isGift
	// Was the offer accepted as a gift for someone other than the buyer.
	// types : Boolean
	IsGift []bool `json:"isGift,omitempty"`

	// OrderDate see : https://schema.org/orderDate
	// Date order was placed.
	// types : DateTime
	OrderDate []DateTime `json:"orderDate,omitempty"`

	// OrderDelivery see : https://schema.org/orderDelivery
	// The delivery of the parcel related to this order or order item.
	// types : ParcelDelivery
	OrderDelivery []*ParcelDelivery `json:"orderDelivery,omitempty"`

	// OrderNumber see : https://schema.org/orderNumber
	// The identifier of the transaction.
	// types : Text
	OrderNumber []string `json:"orderNumber,omitempty"`

	// OrderStatus see : https://schema.org/orderStatus
	// The current status of the order.
	// types : OrderStatus
	OrderStatus []*OrderStatus `json:"orderStatus,omitempty"`

	// OrderedItem see : https://schema.org/orderedItem
	// The item ordered.
	// types : OrderItem Product
	OrderedItem []interface{} `json:"orderedItem,omitempty"`

	// PartOfInvoice see : https://schema.org/partOfInvoice
	// The order is being paid as part of the referenced Invoice.
	// types : Invoice
	PartOfInvoice []*Invoice `json:"partOfInvoice,omitempty"`

	// PaymentDueDate see : https://schema.org/paymentDueDate
	// The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
	// types : DateTime
	PaymentDueDate []DateTime `json:"paymentDueDate,omitempty"`

	// PaymentMethod see : https://schema.org/paymentMethod
	// The name of the credit card or other method of payment for the order.
	// types : PaymentMethod
	PaymentMethod []*PaymentMethod `json:"paymentMethod,omitempty"`

	// PaymentMethodId see : https://schema.org/paymentMethodId
	// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	// types : Text
	PaymentMethodId []string `json:"paymentMethodId,omitempty"`

	// PaymentUrl see : https://schema.org/paymentUrl
	// The URL for sending a payment.
	// types : URL
	PaymentUrl []string `json:"paymentUrl,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`
}

func (v Order) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcceptedOffer
		if len(v.AcceptedOffer) == 1 {
			value = v.AcceptedOffer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acceptedOffer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BillingAddress
		if len(v.BillingAddress) == 1 {
			value = v.BillingAddress[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["billingAddress"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Broker
		if len(v.Broker) == 1 {
			value = v.Broker[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["broker"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ConfirmationNumber
		if len(v.ConfirmationNumber) == 1 {
			value = v.ConfirmationNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["confirmationNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Customer
		if len(v.Customer) == 1 {
			value = v.Customer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["customer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Discount
		if len(v.Discount) == 1 {
			value = v.Discount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["discount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DiscountCode
		if len(v.DiscountCode) == 1 {
			value = v.DiscountCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["discountCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DiscountCurrency
		if len(v.DiscountCurrency) == 1 {
			value = v.DiscountCurrency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["discountCurrency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsGift
		if len(v.IsGift) == 1 {
			value = v.IsGift[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isGift"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderDate
		if len(v.OrderDate) == 1 {
			value = v.OrderDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderDate"] = json.RawMessage(b)
		}
	}

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
		var value interface{} = v.OrderNumber
		if len(v.OrderNumber) == 1 {
			value = v.OrderNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OrderStatus
		if len(v.OrderStatus) == 1 {
			value = v.OrderStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["orderStatus"] = json.RawMessage(b)
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

	{
		var value interface{} = v.PartOfInvoice
		if len(v.PartOfInvoice) == 1 {
			value = v.PartOfInvoice[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfInvoice"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PaymentDueDate
		if len(v.PaymentDueDate) == 1 {
			value = v.PaymentDueDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["paymentDueDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PaymentMethod
		if len(v.PaymentMethod) == 1 {
			value = v.PaymentMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["paymentMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PaymentMethodId
		if len(v.PaymentMethodId) == 1 {
			value = v.PaymentMethodId[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["paymentMethodId"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PaymentUrl
		if len(v.PaymentUrl) == 1 {
			value = v.PaymentUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["paymentUrl"] = json.RawMessage(b)
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

	*intop = into

	return nil
}

func (v Order) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Order"

	return data, nil
}

func (v Order) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
