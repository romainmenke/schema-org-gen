package schemaorggo

import "encoding/json"

// AggregateOffer see : https://schema.org/AggregateOffer
type AggregateOffer struct {

	// With properties from Offer see : https://schema.org/Offer
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AcceptedPaymentMethod see : https://schema.org/acceptedPaymentMethod
	// The payment method(s) accepted by seller for this offer.
	// types : LoanOrCredit PaymentMethod
	AcceptedPaymentMethod []interface{} `json:"acceptedPaymentMethod,omitempty"`

	// AddOn see : https://schema.org/addOn
	// An additional offer that can only be obtained in combination with the first base offer (e.g. supplements and extensions that are available for a surcharge).
	// types : Offer
	AddOn []*Offer `json:"addOn,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AdvanceBookingRequirement see : https://schema.org/advanceBookingRequirement
	// The amount of time that is required between accepting the offer and the actual usage of the resource or service.
	// types : QuantitativeValue
	AdvanceBookingRequirement []*QuantitativeValue `json:"advanceBookingRequirement,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided.
	// types : Place AdministrativeArea GeoShape Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// Availability see : https://schema.org/availability
	// The availability of this item&amp;#x2014;for example In stock, Out of stock, Pre-order, etc.
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
	// types : Text Thing
	Category []interface{} `json:"category,omitempty"`

	// DeliveryLeadTime see : https://schema.org/deliveryLeadTime
	// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
	// types : QuantitativeValue
	DeliveryLeadTime []*QuantitativeValue `json:"deliveryLeadTime,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

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
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.\n\nSee also [[ineligibleRegion]].
	//
	// types : GeoShape Place Text
	EligibleRegion []interface{} `json:"eligibleRegion,omitempty"`

	// EligibleTransactionVolume see : https://schema.org/eligibleTransactionVolume
	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	// types : PriceSpecification
	EligibleTransactionVolume []*PriceSpecification `json:"eligibleTransactionVolume,omitempty"`

	// Gtin12 see : https://schema.org/gtin12
	// The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin12 []string `json:"gtin12,omitempty"`

	// Gtin13 see : https://schema.org/gtin13
	// The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin13 []string `json:"gtin13,omitempty"`

	// Gtin14 see : https://schema.org/gtin14
	// The GTIN-14 code of the product, or the product to which the offer refers. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin14 []string `json:"gtin14,omitempty"`

	// Gtin8 see : https://schema.org/gtin8
	// The [GTIN-8](http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin8 []string `json:"gtin8,omitempty"`

	// HighPrice see : https://schema.org/highPrice
	// The highest price of all offers available.
	// types : Number Text
	HighPrice []interface{} `json:"highPrice,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IncludesObject see : https://schema.org/includesObject
	// This links to a node or nodes indicating the exact quantity of the products included in the offer.
	// types : TypeAndQuantityNode
	IncludesObject []*TypeAndQuantityNode `json:"includesObject,omitempty"`

	// IneligibleRegion see : https://schema.org/ineligibleRegion
	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.\n\nSee also [[eligibleRegion]].
	//
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

	// LowPrice see : https://schema.org/lowPrice
	// The lowest price of all offers available.
	// types : Number Text
	LowPrice []interface{} `json:"lowPrice,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn []string `json:"mpn,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// OfferCount see : https://schema.org/offerCount
	// The number of offers for the product.
	// types : Integer
	OfferCount []float64 `json:"offerCount,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Price see : https://schema.org/price
	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.\n\nUsage guidelines:\n\n* Use the [[priceCurrency]] property (with standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for [Local Exchange Tradings Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;) instead of including [ambiguous symbols](http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as &#39;$&#39; in the value.\n* Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.\n* Note that both [RDFa](http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values alongside more human-friendly formatting.\n* Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	//
	// types : Number Text
	Price []interface{} `json:"price,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency of the price, or a price component when attached to [[PriceSpecification]] and its subtypes.\n\nUse standard formats: [ISO 4217 currency format](http://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; [Ticker symbol](https://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for [Local Exchange Tradings Systems](https://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
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
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Reviews see : https://schema.org/reviews
	// Review of the item.
	// types : Review
	Reviews []*Review `json:"reviews,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seller see : https://schema.org/seller
	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
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

	// SubjectOf see : https://schema.org/subjectOf
	// A CreativeWork or Event about this Thing..
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// Warranty see : https://schema.org/warranty
	// The warranty promise(s) included in the offer.
	// types : WarrantyPromise
	Warranty []*WarrantyPromise `json:"warranty,omitempty"`
}

func (v AggregateOffer) MarshalJSON() ([]byte, error) {
	type Alias AggregateOffer

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"AggregateOffer\","), b[1:]...), nil
}
