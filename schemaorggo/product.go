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
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award []string `json:"award,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category []interface{} `json:"category,omitempty"`

	// Color see : https://schema.org/color
	// The color of the product.
	// types : Text
	Color []string `json:"color,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	// types : Distance QuantitativeValue
	Depth []interface{} `json:"depth,omitempty"`

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

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// IsAccessoryOrSparePartFor see : https://schema.org/isAccessoryOrSparePartFor
	// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	// types : Product
	IsAccessoryOrSparePartFor []*Product `json:"isAccessoryOrSparePartFor,omitempty"`

	// IsConsumableFor see : https://schema.org/isConsumableFor
	// A pointer to another product (or multiple products) for which this product is a consumable.
	// types : Product
	IsConsumableFor []*Product `json:"isConsumableFor,omitempty"`

	// IsRelatedTo see : https://schema.org/isRelatedTo
	// A pointer to another, somehow related product (or multiple products).
	// types : Product Service
	IsRelatedTo []interface{} `json:"isRelatedTo,omitempty"`

	// IsSimilarTo see : https://schema.org/isSimilarTo
	// A pointer to another, functionally similar product (or multiple products).
	// types : Product Service
	IsSimilarTo []interface{} `json:"isSimilarTo,omitempty"`

	// ItemCondition see : https://schema.org/itemCondition
	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	// types : OfferItemCondition
	ItemCondition []*OfferItemCondition `json:"itemCondition,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo []interface{} `json:"logo,omitempty"`

	// Manufacturer see : https://schema.org/manufacturer
	// The manufacturer of the product.
	// types : Organization
	Manufacturer []*Organization `json:"manufacturer,omitempty"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	// types : Product Text URL
	Material []interface{} `json:"material,omitempty"`

	// Model see : https://schema.org/model
	// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	// types : ProductModel Text
	Model []interface{} `json:"model,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn []string `json:"mpn,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// ProductID see : https://schema.org/productID
	// The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	// types : Text
	ProductID []string `json:"productID,omitempty"`

	// ProductionDate see : https://schema.org/productionDate
	// The date of production of the item, e.g. vehicle.
	// types : Date
	ProductionDate []Date `json:"productionDate,omitempty"`

	// PurchaseDate see : https://schema.org/purchaseDate
	// The date the item e.g. vehicle was purchased by the current owner.
	// types : Date
	PurchaseDate []Date `json:"purchaseDate,omitempty"`

	// ReleaseDate see : https://schema.org/releaseDate
	// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	// types : Date
	ReleaseDate []Date `json:"releaseDate,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Sku see : https://schema.org/sku
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	// types : Text
	Sku []string `json:"sku,omitempty"`

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	// types : QuantitativeValue
	Weight []*QuantitativeValue `json:"weight,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width []interface{} `json:"width,omitempty"`
}

func (v Product) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Thing.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AdditionalProperty
		if len(v.AdditionalProperty) == 1 {
			value = v.AdditionalProperty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalProperty"] = json.RawMessage(b)
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
		var value interface{} = v.Audience
		if len(v.Audience) == 1 {
			value = v.Audience[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audience"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Award
		if len(v.Award) == 1 {
			value = v.Award[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["award"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Brand
		if len(v.Brand) == 1 {
			value = v.Brand[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["brand"] = json.RawMessage(b)
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
		var value interface{} = v.Color
		if len(v.Color) == 1 {
			value = v.Color[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["color"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Depth
		if len(v.Depth) == 1 {
			value = v.Depth[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["depth"] = json.RawMessage(b)
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
		var value interface{} = v.Height
		if len(v.Height) == 1 {
			value = v.Height[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["height"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsAccessoryOrSparePartFor
		if len(v.IsAccessoryOrSparePartFor) == 1 {
			value = v.IsAccessoryOrSparePartFor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAccessoryOrSparePartFor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsConsumableFor
		if len(v.IsConsumableFor) == 1 {
			value = v.IsConsumableFor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isConsumableFor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsRelatedTo
		if len(v.IsRelatedTo) == 1 {
			value = v.IsRelatedTo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isRelatedTo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsSimilarTo
		if len(v.IsSimilarTo) == 1 {
			value = v.IsSimilarTo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isSimilarTo"] = json.RawMessage(b)
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
		var value interface{} = v.Logo
		if len(v.Logo) == 1 {
			value = v.Logo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["logo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Manufacturer
		if len(v.Manufacturer) == 1 {
			value = v.Manufacturer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["manufacturer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Material
		if len(v.Material) == 1 {
			value = v.Material[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["material"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Model
		if len(v.Model) == 1 {
			value = v.Model[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["model"] = json.RawMessage(b)
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
		var value interface{} = v.Offers
		if len(v.Offers) == 1 {
			value = v.Offers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductID
		if len(v.ProductID) == 1 {
			value = v.ProductID[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productID"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductionDate
		if len(v.ProductionDate) == 1 {
			value = v.ProductionDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productionDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PurchaseDate
		if len(v.PurchaseDate) == 1 {
			value = v.PurchaseDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["purchaseDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReleaseDate
		if len(v.ReleaseDate) == 1 {
			value = v.ReleaseDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["releaseDate"] = json.RawMessage(b)
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
		var value interface{} = v.Weight
		if len(v.Weight) == 1 {
			value = v.Weight[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["weight"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Width
		if len(v.Width) == 1 {
			value = v.Width[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["width"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Product) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Product"

	return data, nil
}

func (v Product) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
