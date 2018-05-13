package schemaorggo

import "encoding/json"

// Offer see : https://schema.org/Offer
type Offer struct {
	Intangible

	typeContext

	// AcceptedPaymentMethod see : https://schema.org/acceptedPaymentMethod
	// The payment method(s) accepted by seller for this offer.
	// types : LoanOrCredit PaymentMethod
	AcceptedPaymentMethod []interface{} `json:"acceptedPaymentMethod,omitempty"`

	// AddOn see : https://schema.org/addOn
	// An additional offer that can only be obtained in combination with the first base offer (e.g. supplements and extensions that are available for a surcharge).
	// types : Offer
	AddOn []*Offer `json:"addOn,omitempty"`

	// AdvanceBookingRequirement see : https://schema.org/advanceBookingRequirement
	// The amount of time that is required between accepting the offer and the actual usage of the resource or service.
	// types : QuantitativeValue
	AdvanceBookingRequirement []*QuantitativeValue `json:"advanceBookingRequirement,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// Availability see : https://schema.org/availability
	// The availability of this item—for example In stock, Out of stock, Pre-order, etc.
	// types : ItemAvailability
	Availability []*ItemAvailability `json:"availability,omitempty"`

	// AvailabilityEnds see : https://schema.org/availabilityEnds
	// The end of the availability of the product or service included in the offer.
	// types : DateTime
	AvailabilityEnds []DateTime `json:"availabilityEnds,omitempty"`

	// AvailabilityStarts see : https://schema.org/availabilityStarts
	// The beginning of the availability of the product or service included in the offer.
	// types : DateTime
	AvailabilityStarts []DateTime `json:"availabilityStarts,omitempty"`

	// AvailableAtOrFrom see : https://schema.org/availableAtOrFrom
	// The place(s) from which the offer can be obtained (e.g. store locations).
	// types : Place
	AvailableAtOrFrom []*Place `json:"availableAtOrFrom,omitempty"`

	// AvailableDeliveryMethod see : https://schema.org/availableDeliveryMethod
	// The delivery method(s) available for this offer.
	// types : DeliveryMethod
	AvailableDeliveryMethod []*DeliveryMethod `json:"availableDeliveryMethod,omitempty"`

	// BusinessFunction see : https://schema.org/businessFunction
	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	// types : BusinessFunction
	BusinessFunction []*BusinessFunction `json:"businessFunction,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category []interface{} `json:"category,omitempty"`

	// DeliveryLeadTime see : https://schema.org/deliveryLeadTime
	// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
	// types : QuantitativeValue
	DeliveryLeadTime []*QuantitativeValue `json:"deliveryLeadTime,omitempty"`

	// EligibleCustomerType see : https://schema.org/eligibleCustomerType
	// The type(s) of customers for which the given offer is valid.
	// types : BusinessEntityType
	EligibleCustomerType []*BusinessEntityType `json:"eligibleCustomerType,omitempty"`

	// EligibleDuration see : https://schema.org/eligibleDuration
	// The duration for which the given offer is valid.
	// types : QuantitativeValue
	EligibleDuration []*QuantitativeValue `json:"eligibleDuration,omitempty"`

	// EligibleQuantity see : https://schema.org/eligibleQuantity
	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	// types : QuantitativeValue
	EligibleQuantity []*QuantitativeValue `json:"eligibleQuantity,omitempty"`

	// EligibleRegion see : https://schema.org/eligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
	//
	// See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
	// types : GeoShape Place Text
	EligibleRegion []interface{} `json:"eligibleRegion,omitempty"`

	// EligibleTransactionVolume see : https://schema.org/eligibleTransactionVolume
	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	// types : PriceSpecification
	EligibleTransactionVolume []*PriceSpecification `json:"eligibleTransactionVolume,omitempty"`

	// Gtin12 see : https://schema.org/gtin12
	// The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin12 []string `json:"gtin12,omitempty"`

	// Gtin13 see : https://schema.org/gtin13
	// The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin13 []string `json:"gtin13,omitempty"`

	// Gtin14 see : https://schema.org/gtin14
	// The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin14 []string `json:"gtin14,omitempty"`

	// Gtin8 see : https://schema.org/gtin8
	// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin8 []string `json:"gtin8,omitempty"`

	// IncludesObject see : https://schema.org/includesObject
	// This links to a node or nodes indicating the exact quantity of the products included in the offer.
	// types : TypeAndQuantityNode
	IncludesObject []*TypeAndQuantityNode `json:"includesObject,omitempty"`

	// IneligibleRegion see : https://schema.org/ineligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
	//
	// See also eligibleRegion (see: https://schema.org/eligibleRegion).
	// types : GeoShape Place Text
	IneligibleRegion []interface{} `json:"ineligibleRegion,omitempty"`

	// InventoryLevel see : https://schema.org/inventoryLevel
	// The current approximate inventory level for the item or items.
	// types : QuantitativeValue
	InventoryLevel []*QuantitativeValue `json:"inventoryLevel,omitempty"`

	// ItemCondition see : https://schema.org/itemCondition
	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	// types : OfferItemCondition
	ItemCondition []*OfferItemCondition `json:"itemCondition,omitempty"`

	// ItemOffered see : https://schema.org/itemOffered
	// The item being offered.
	// types : Product Service
	ItemOffered []interface{} `json:"itemOffered,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn []string `json:"mpn,omitempty"`

	// OfferedBy see : https://schema.org/offeredBy
	// A pointer to the organization or person making the offer. Inverse property: makesOffer (see: https://schema.org/makesOffer).
	// types : Organization Person
	OfferedBy []interface{} `json:"offeredBy,omitempty"`

	// Price see : https://schema.org/price
	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
	//
	// Usage guidelines:
	//
	//
	// Use the priceCurrency (see: https://schema.org/priceCurrency) property (with ISO 4217 codes (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217#Active_codes) e.g. &quot;USD&quot;) instead of
	//   including ambiguous symbols (see: https://schema.orghttp://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as &#39;$&#39; in the value.
	// Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.
	// Note that both RDFa (see: https://schema.orghttp://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values alongside more human-friendly formatting.
	// Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	//
	//
	// types : Number Text
	Price []interface{} `json:"price,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	// types : Text
	PriceCurrency []string `json:"priceCurrency,omitempty"`

	// PriceSpecification see : https://schema.org/priceSpecification
	// One or more detailed price specifications, indicating the unit price and delivery or payment charges.
	// types : PriceSpecification
	PriceSpecification []*PriceSpecification `json:"priceSpecification,omitempty"`

	// PriceValidUntil see : https://schema.org/priceValidUntil
	// The date after which the price is no longer available.
	// types : Date
	PriceValidUntil []Date `json:"priceValidUntil,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	// types : Organization Person
	Seller []interface{} `json:"seller,omitempty"`

	// SerialNumber see : https://schema.org/serialNumber
	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	// types : Text
	SerialNumber []string `json:"serialNumber,omitempty"`

	// Sku see : https://schema.org/sku
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	// types : Text
	Sku []string `json:"sku,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// Warranty see : https://schema.org/warranty
	// The warranty promise(s) included in the offer. Supersedes warrantyPromise (see: https://schema.org/warrantyPromise).
	// types : WarrantyPromise
	Warranty []*WarrantyPromise `json:"warranty,omitempty"`
}

func (v Offer) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcceptedPaymentMethod
		if len(v.AcceptedPaymentMethod) == 1 {
			value = v.AcceptedPaymentMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acceptedPaymentMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AddOn
		if len(v.AddOn) == 1 {
			value = v.AddOn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["addOn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AdvanceBookingRequirement
		if len(v.AdvanceBookingRequirement) == 1 {
			value = v.AdvanceBookingRequirement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["advanceBookingRequirement"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AggregateRating
		if len(v.AggregateRating) == 1 {
			value = v.AggregateRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aggregateRating"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AreaServed
		if len(v.AreaServed) == 1 {
			value = v.AreaServed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["areaServed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Availability
		if len(v.Availability) == 1 {
			value = v.Availability[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availability"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailabilityEnds
		if len(v.AvailabilityEnds) == 1 {
			value = v.AvailabilityEnds[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availabilityEnds"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailabilityStarts
		if len(v.AvailabilityStarts) == 1 {
			value = v.AvailabilityStarts[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availabilityStarts"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailableAtOrFrom
		if len(v.AvailableAtOrFrom) == 1 {
			value = v.AvailableAtOrFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableAtOrFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AvailableDeliveryMethod
		if len(v.AvailableDeliveryMethod) == 1 {
			value = v.AvailableDeliveryMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableDeliveryMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BusinessFunction
		if len(v.BusinessFunction) == 1 {
			value = v.BusinessFunction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["businessFunction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Category
		if len(v.Category) == 1 {
			value = v.Category[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["category"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DeliveryLeadTime
		if len(v.DeliveryLeadTime) == 1 {
			value = v.DeliveryLeadTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deliveryLeadTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleCustomerType
		if len(v.EligibleCustomerType) == 1 {
			value = v.EligibleCustomerType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleCustomerType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleDuration
		if len(v.EligibleDuration) == 1 {
			value = v.EligibleDuration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleDuration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleQuantity
		if len(v.EligibleQuantity) == 1 {
			value = v.EligibleQuantity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleQuantity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleRegion
		if len(v.EligibleRegion) == 1 {
			value = v.EligibleRegion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleRegion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EligibleTransactionVolume
		if len(v.EligibleTransactionVolume) == 1 {
			value = v.EligibleTransactionVolume[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eligibleTransactionVolume"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Gtin12
		if len(v.Gtin12) == 1 {
			value = v.Gtin12[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gtin12"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Gtin13
		if len(v.Gtin13) == 1 {
			value = v.Gtin13[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gtin13"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Gtin14
		if len(v.Gtin14) == 1 {
			value = v.Gtin14[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gtin14"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Gtin8
		if len(v.Gtin8) == 1 {
			value = v.Gtin8[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gtin8"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IncludesObject
		if len(v.IncludesObject) == 1 {
			value = v.IncludesObject[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["includesObject"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IneligibleRegion
		if len(v.IneligibleRegion) == 1 {
			value = v.IneligibleRegion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ineligibleRegion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InventoryLevel
		if len(v.InventoryLevel) == 1 {
			value = v.InventoryLevel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inventoryLevel"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ItemCondition
		if len(v.ItemCondition) == 1 {
			value = v.ItemCondition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemCondition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ItemOffered
		if len(v.ItemOffered) == 1 {
			value = v.ItemOffered[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemOffered"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Mpn
		if len(v.Mpn) == 1 {
			value = v.Mpn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mpn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OfferedBy
		if len(v.OfferedBy) == 1 {
			value = v.OfferedBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offeredBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Price
		if len(v.Price) == 1 {
			value = v.Price[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["price"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PriceCurrency
		if len(v.PriceCurrency) == 1 {
			value = v.PriceCurrency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceCurrency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PriceSpecification
		if len(v.PriceSpecification) == 1 {
			value = v.PriceSpecification[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceSpecification"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PriceValidUntil
		if len(v.PriceValidUntil) == 1 {
			value = v.PriceValidUntil[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceValidUntil"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Review
		if len(v.Review) == 1 {
			value = v.Review[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["review"] = json.RawMessage(b)
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
		var value interface{} = v.SerialNumber
		if len(v.SerialNumber) == 1 {
			value = v.SerialNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serialNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Sku
		if len(v.Sku) == 1 {
			value = v.Sku[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sku"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidFrom
		if len(v.ValidFrom) == 1 {
			value = v.ValidFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidThrough
		if len(v.ValidThrough) == 1 {
			value = v.ValidThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Warranty
		if len(v.Warranty) == 1 {
			value = v.Warranty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["warranty"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Offer) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Offer"

	return data, nil
}

func (v Offer) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
