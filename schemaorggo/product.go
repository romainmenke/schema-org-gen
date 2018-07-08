package schemaorggo

import "encoding/json"

// Product see : https://schema.org/Product
type Product struct {
	Thing

	typeContext

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// Category see : https://pending.schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : PhysicalActivityCategory Text Thing
	Category []interface{} `json:"category,omitempty"`

	// Color see : https://schema.org/color
	// The color of the product.
	// types : Text
	Color []string `json:"color,omitempty"`

	// Gtin8 see : https://schema.org/gtin8
	// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin8 []string `json:"gtin8,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// IsConsumableFor see : https://schema.org/isConsumableFor
	// A pointer to another product (or multiple products) for which this product is a consumable.
	// types : Product
	IsConsumableFor []*Product `json:"isConsumableFor,omitempty"`

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

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn []string `json:"mpn,omitempty"`

	// ProductID see : https://schema.org/productID
	// The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	// types : Text
	ProductID []string `json:"productID,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

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
