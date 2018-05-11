package schemaorg

// Product see : https://schema.org/Product
type Product struct {

Thing// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty string `json:"additionalProperty"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

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
IsAccessoryOrSparePartFor string `json:"isAccessoryOrSparePartFor"`

// IsConsumableFor see : /isConsumableFor
// A pointer to another product (or multiple products) for which this product is a consumable.
IsConsumableFor string `json:"isConsumableFor"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products).
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products).
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// ItemCondition see : /itemCondition
// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
ItemCondition string `json:"itemCondition"`

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// Manufacturer see : /manufacturer
// The manufacturer of the product.
Manufacturer string `json:"manufacturer"`

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
Offers string `json:"offers"`

// ProductID see : /productID
// The product identifier, such as ISBN. For example: meta itemprop="productID" content="isbn:123-456-789".
ProductID string `json:"productID"`

// ProductionDate see : /productionDate
// The date of production of the item, e.g. vehicle.
ProductionDate string `json:"productionDate"`

// PurchaseDate see : /purchaseDate
// The date the item e.g. vehicle was purchased by the current owner.
PurchaseDate string `json:"purchaseDate"`

// ReleaseDate see : /releaseDate
// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
ReleaseDate string `json:"releaseDate"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Sku see : /sku
// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
Sku string `json:"sku"`

// Weight see : /weight
// The weight of the product or person.
Weight string `json:"weight"`

// Width see : /width
// The width of the item.
Width interface{} `json:"width"` // types : Distance QuantitativeValue

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// IsAccessoryOrSparePartFor see : /isAccessoryOrSparePartFor
// A pointer to another product (or multiple products) for which this product is an accessory or spare part. 
IsAccessoryOrSparePartFor string `json:"isAccessoryOrSparePartFor"`

// IsBasedOn see : /isBasedOn
// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.  Supersedes isBasedOnUrl (see: https://schema.org/isBasedOnUrl).
IsBasedOn string `json:"isBasedOn"`

// IsConsumableFor see : /isConsumableFor
// A pointer to another product (or multiple products) for which this product is a consumable. 
IsConsumableFor string `json:"isConsumableFor"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products). 
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products). 
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// ItemOffered see : /itemOffered
// The item being offered. 
ItemOffered interface{} `json:"itemOffered"` // types : Demand Offer

// ItemShipped see : /itemShipped
// Item(s) being shipped. 
ItemShipped string `json:"itemShipped"`

// Material see : /material
// A material that something is made from, e.g. leather, wool, cotton, paper. 
Material interface{} `json:"material"` // types : CreativeWork Product

// OrderedItem see : /orderedItem
// The item ordered. 
OrderedItem interface{} `json:"orderedItem"` // types : Order OrderItem

// Owns see : /owns
// Products owned by the organization or person. 
Owns interface{} `json:"owns"` // types : Organization Person

// ProductSupported see : /productSupported
// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. "iPhone") or a general category of products or services (e.g. "smartphones"). 
ProductSupported string `json:"productSupported"`

// TypeOfGood see : /typeOfGood
// The product that this structured value is referring to. 
TypeOfGood interface{} `json:"typeOfGood"` // types : OwnershipInfo TypeAndQuantityNode

}

