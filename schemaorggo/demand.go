package schemaorggo

import "encoding/json"

// Demand see : https://schema.org/Demand
type Demand struct {
	Intangible

	typeContext

	// AcceptedPaymentMethod see : https://schema.org/acceptedPaymentMethod
	// The payment method(s) accepted by seller for this offer.
	// types : LoanOrCredit PaymentMethod
	AcceptedPaymentMethod interface{} `json:"acceptedPaymentMethod,omitempty"`

	// AdvanceBookingRequirement see : https://schema.org/advanceBookingRequirement
	// The amount of time that is required between accepting the offer and the actual usage of the resource or service.
	// types : QuantitativeValue
	AdvanceBookingRequirement *QuantitativeValue `json:"advanceBookingRequirement,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed interface{} `json:"areaServed,omitempty"`

	// Availability see : https://schema.org/availability
	// The availability of this item—for example In stock, Out of stock, Pre-order, etc.
	// types : ItemAvailability
	Availability *ItemAvailability `json:"availability,omitempty"`

	// AvailabilityEnds see : https://schema.org/availabilityEnds
	// The end of the availability of the product or service included in the offer.
	// types : DateTime
	AvailabilityEnds DateTime `json:"availabilityEnds,omitempty"`

	// AvailabilityStarts see : https://schema.org/availabilityStarts
	// The beginning of the availability of the product or service included in the offer.
	// types : DateTime
	AvailabilityStarts DateTime `json:"availabilityStarts,omitempty"`

	// AvailableAtOrFrom see : https://schema.org/availableAtOrFrom
	// The place(s) from which the offer can be obtained (e.g. store locations).
	// types : Place
	AvailableAtOrFrom *Place `json:"availableAtOrFrom,omitempty"`

	// AvailableDeliveryMethod see : https://schema.org/availableDeliveryMethod
	// The delivery method(s) available for this offer.
	// types : DeliveryMethod
	AvailableDeliveryMethod *DeliveryMethod `json:"availableDeliveryMethod,omitempty"`

	// BusinessFunction see : https://schema.org/businessFunction
	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	// types : BusinessFunction
	BusinessFunction *BusinessFunction `json:"businessFunction,omitempty"`

	// DeliveryLeadTime see : https://schema.org/deliveryLeadTime
	// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
	// types : QuantitativeValue
	DeliveryLeadTime *QuantitativeValue `json:"deliveryLeadTime,omitempty"`

	// EligibleCustomerType see : https://schema.org/eligibleCustomerType
	// The type(s) of customers for which the given offer is valid.
	// types : BusinessEntityType
	EligibleCustomerType *BusinessEntityType `json:"eligibleCustomerType,omitempty"`

	// EligibleDuration see : https://schema.org/eligibleDuration
	// The duration for which the given offer is valid.
	// types : QuantitativeValue
	EligibleDuration *QuantitativeValue `json:"eligibleDuration,omitempty"`

	// EligibleQuantity see : https://schema.org/eligibleQuantity
	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	// types : QuantitativeValue
	EligibleQuantity *QuantitativeValue `json:"eligibleQuantity,omitempty"`

	// EligibleRegion see : https://schema.org/eligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
	//
	// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
	// types : GeoShape Place Text
	EligibleRegion interface{} `json:"eligibleRegion,omitempty"`

	// EligibleTransactionVolume see : https://schema.org/eligibleTransactionVolume
	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	// types : PriceSpecification
	EligibleTransactionVolume *PriceSpecification `json:"eligibleTransactionVolume,omitempty"`

	// Gtin12 see : https://schema.org/gtin12
	// The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin12 string `json:"gtin12,omitempty"`

	// Gtin13 see : https://schema.org/gtin13
	// The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin13 string `json:"gtin13,omitempty"`

	// Gtin14 see : https://schema.org/gtin14
	// The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin14 string `json:"gtin14,omitempty"`

	// Gtin8 see : https://schema.org/gtin8
	// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin8 string `json:"gtin8,omitempty"`

	// IncludesObject see : https://schema.org/includesObject
	// This links to a node or nodes indicating the exact quantity of the products included in the offer.
	// types : TypeAndQuantityNode
	IncludesObject *TypeAndQuantityNode `json:"includesObject,omitempty"`

	// IneligibleRegion see : https://schema.org/ineligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
	//
	// See also eligibleRegion (see: https://schema.org/eligibleRegion).
	// types : GeoShape Place Text
	IneligibleRegion interface{} `json:"ineligibleRegion,omitempty"`

	// InventoryLevel see : https://schema.org/inventoryLevel
	// The current approximate inventory level for the item or items.
	// types : QuantitativeValue
	InventoryLevel *QuantitativeValue `json:"inventoryLevel,omitempty"`

	// ItemCondition see : https://schema.org/itemCondition
	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	// types : OfferItemCondition
	ItemCondition *OfferItemCondition `json:"itemCondition,omitempty"`

	// ItemOffered see : https://schema.org/itemOffered
	// The item being offered.
	// types : Product Service
	ItemOffered interface{} `json:"itemOffered,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn string `json:"mpn,omitempty"`

	// PriceSpecification see : https://schema.org/priceSpecification
	// One or more detailed price specifications, indicating the unit price and delivery or payment charges.
	// types : PriceSpecification
	PriceSpecification *PriceSpecification `json:"priceSpecification,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller interface{} `json:"seller,omitempty"`

	// SerialNumber see : https://schema.org/serialNumber
	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	// types : Text
	SerialNumber string `json:"serialNumber,omitempty"`

	// Sku see : https://schema.org/sku
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	// types : Text
	Sku string `json:"sku,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough DateTime `json:"validThrough,omitempty"`

	// Warranty see : https://schema.org/warranty
	// The warranty promise(s) included in the offer. Supersedes warrantyPromise (see: https://schema.org/warrantyPromise).
	// types : WarrantyPromise
	Warranty *WarrantyPromise `json:"warranty,omitempty"`
}

func (v Demand) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Demand"

	return json.Marshal(v)
}

func (v *Demand) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Demand"

	return json.Marshal(*v)
}
