package schemaorggo

import "encoding/json"

// Product see : https://schema.org/Product
type Product struct {
	Thing

	typeContext

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	//
	// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	// types : PropertyValue
	AdditionalProperty *PropertyValue `json:"additionalProperty,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience *Audience `json:"audience,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award string `json:"award,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand interface{} `json:"brand,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category interface{} `json:"category,omitempty"`

	// Color see : https://schema.org/color
	// The color of the product.
	// types : Text
	Color string `json:"color,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	// types : Distance QuantitativeValue
	Depth interface{} `json:"depth,omitempty"`

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

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height interface{} `json:"height,omitempty"`

	// IsAccessoryOrSparePartFor see : https://schema.org/isAccessoryOrSparePartFor
	// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	// types : Product
	IsAccessoryOrSparePartFor *Product `json:"isAccessoryOrSparePartFor,omitempty"`

	// IsConsumableFor see : https://schema.org/isConsumableFor
	// A pointer to another product (or multiple products) for which this product is a consumable.
	// types : Product
	IsConsumableFor *Product `json:"isConsumableFor,omitempty"`

	// IsRelatedTo see : https://schema.org/isRelatedTo
	// A pointer to another, somehow related product (or multiple products).
	// types : Product Service
	IsRelatedTo interface{} `json:"isRelatedTo,omitempty"`

	// IsSimilarTo see : https://schema.org/isSimilarTo
	// A pointer to another, functionally similar product (or multiple products).
	// types : Product Service
	IsSimilarTo interface{} `json:"isSimilarTo,omitempty"`

	// ItemCondition see : https://schema.org/itemCondition
	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	// types : OfferItemCondition
	ItemCondition *OfferItemCondition `json:"itemCondition,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo interface{} `json:"logo,omitempty"`

	// Manufacturer see : https://schema.org/manufacturer
	// The manufacturer of the product.
	// types : Organization
	Manufacturer *Organization `json:"manufacturer,omitempty"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	// types : Product Text URL
	Material interface{} `json:"material,omitempty"`

	// Model see : https://schema.org/model
	// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	// types : ProductModel Text
	Model interface{} `json:"model,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn string `json:"mpn,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers *Offer `json:"offers,omitempty"`

	// ProductID see : https://schema.org/productID
	// The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	// types : Text
	ProductID string `json:"productID,omitempty"`

	// ProductionDate see : https://schema.org/productionDate
	// The date of production of the item, e.g. vehicle.
	// types : Date
	ProductionDate Date `json:"productionDate,omitempty"`

	// PurchaseDate see : https://schema.org/purchaseDate
	// The date the item e.g. vehicle was purchased by the current owner.
	// types : Date
	PurchaseDate Date `json:"purchaseDate,omitempty"`

	// ReleaseDate see : https://schema.org/releaseDate
	// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	// types : Date
	ReleaseDate Date `json:"releaseDate,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review *Review `json:"review,omitempty"`

	// Sku see : https://schema.org/sku
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	// types : Text
	Sku string `json:"sku,omitempty"`

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	// types : QuantitativeValue
	Weight *QuantitativeValue `json:"weight,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width interface{} `json:"width,omitempty"`
}

func (v Product) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Product"

	return json.Marshal(v)
}

func (v *Product) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Product"

	return json.Marshal(*v)
}
