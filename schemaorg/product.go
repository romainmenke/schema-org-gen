package schemaorg

import "encoding/json"

// Product see : https://schema.org/Product
type Product struct {

typeContext

Thing

// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty *PropertyValue `json:"additionalProperty"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating *AggregateRating `json:"aggregateRating"`

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience *Audience `json:"audience"`

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// Color see : /color
// The color of the product.
Color string `json:"color"`

// Depth see : /depth
// The depth of the item.
Depth interface{} `json:"depth"` // types : Distance QuantitativeValue

// Gtin12 see : /gtin12
// The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin12 string `json:"gtin12"`

// Gtin13 see : /gtin13
// The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin13 string `json:"gtin13"`

// Gtin14 see : /gtin14
// The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin14 string `json:"gtin14"`

// Gtin8 see : /gtin8
// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin8 string `json:"gtin8"`

// Height see : /height
// The height of the item.
Height interface{} `json:"height"` // types : Distance QuantitativeValue

// IsAccessoryOrSparePartFor see : /isAccessoryOrSparePartFor
// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
IsAccessoryOrSparePartFor *Product `json:"isAccessoryOrSparePartFor"`

// IsConsumableFor see : /isConsumableFor
// A pointer to another product (or multiple products) for which this product is a consumable.
IsConsumableFor *Product `json:"isConsumableFor"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products).
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products).
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// ItemCondition see : /itemCondition
// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
ItemCondition *OfferItemCondition `json:"itemCondition"`

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// Manufacturer see : /manufacturer
// The manufacturer of the product.
Manufacturer *Organization `json:"manufacturer"`

// Material see : /material
// A material that something is made from, e.g. leather, wool, cotton, paper.
Material interface{} `json:"material"` // types : Product Text URL

// Model see : /model
// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
Model interface{} `json:"model"` // types : ProductModel Text

// Mpn see : /mpn
// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
Mpn string `json:"mpn"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers *Offer `json:"offers"`

// ProductID see : /productID
// The product identifier, such as ISBN. For example: meta itemprop="productID" content="isbn:123-456-789".
ProductID string `json:"productID"`

// ProductionDate see : /productionDate
// The date of production of the item, e.g. vehicle.
ProductionDate interface{} `json:"productionDate"`

// PurchaseDate see : /purchaseDate
// The date the item e.g. vehicle was purchased by the current owner.
PurchaseDate interface{} `json:"purchaseDate"`

// ReleaseDate see : /releaseDate
// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
ReleaseDate interface{} `json:"releaseDate"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review *Review `json:"review"`

// Sku see : /sku
// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
Sku string `json:"sku"`

// Weight see : /weight
// The weight of the product or person.
Weight *QuantitativeValue `json:"weight"`

// Width see : /width
// The width of the item.
Width interface{} `json:"width"` // types : Distance QuantitativeValue

}

func (v *Product) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Product"

	return json.Marshal(v)
}
