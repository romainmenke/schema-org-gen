<?php

class Product extends Thing implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Product';
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue|\PropertyValue[]
	 */
	public var $additional_property;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating|\AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience|\Audience[]
	 */
	public var $audience;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string|string[]
	 */
	public var $award;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand|\Brand[]|\Organization|\Organization[]
	 */
	public var $brand;
	
	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://schema.org/category
	 * @var \PhysicalActivityCategory|\PhysicalActivityCategory[]|string|string[]|\Thing|\Thing[]
	 */
	public var $category;
	
	/**
	 * The color of the product.
	 * see : https://schema.org/color
	 * @var string|string[]
	 */
	public var $color;
	
	/**
	 * The depth of the item.
	 * see : https://schema.org/depth
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $depth;
	
	/**
	 * The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin12
	 * @var string|string[]
	 */
	public var $gtin_12;
	
	/**
	 * The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin13
	 * @var string|string[]
	 */
	public var $gtin_13;
	
	/**
	 * The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin14
	 * @var string|string[]
	 */
	public var $gtin_14;
	
	/**
	 * The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin8
	 * @var string|string[]
	 */
	public var $gtin_8;
	
	/**
	 * The height of the item.
	 * see : https://schema.org/height
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $height;
	
	/**
	 * A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	 * see : https://schema.org/isAccessoryOrSparePartFor
	 * @var \Product|\Product[]
	 */
	public var $is_accessory_or_spare_part_for;
	
	/**
	 * A pointer to another product (or multiple products) for which this product is a consumable.
	 * see : https://schema.org/isConsumableFor
	 * @var \Product|\Product[]
	 */
	public var $is_consumable_for;
	
	/**
	 * A pointer to another, somehow related product (or multiple products).
	 * see : https://schema.org/isRelatedTo
	 * @var \Product|\Product[]|\Service|\Service[]
	 */
	public var $is_related_to;
	
	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 * @var \Product|\Product[]|\Service|\Service[]
	 */
	public var $is_similar_to;
	
	/**
	 * A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	 * see : https://schema.org/itemCondition
	 * @var \OfferItemCondition|\OfferItemCondition[]
	 */
	public var $item_condition;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject|\ImageObject[]|string|string[]
	 */
	public var $logo;
	
	/**
	 * The manufacturer of the product.
	 * see : https://schema.org/manufacturer
	 * @var \Organization|\Organization[]
	 */
	public var $manufacturer;
	
	/**
	 * A material that something is made from, e.g. leather, wool, cotton, paper.
	 * see : https://schema.org/material
	 * @var \Product|\Product[]|string|string[]|string|string[]
	 */
	public var $material;
	
	/**
	 * The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	 * see : https://schema.org/model
	 * @var \ProductModel|\ProductModel[]|string|string[]
	 */
	public var $model;
	
	/**
	 * The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	 * see : https://schema.org/mpn
	 * @var string|string[]
	 */
	public var $mpn;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer|\Offer[]
	 */
	public var $offers;
	
	/**
	 * The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	 * see : https://schema.org/productID
	 * @var string|string[]
	 */
	public var $producti_d;
	
	/**
	 * The date of production of the item, e.g. vehicle.
	 * see : https://schema.org/productionDate
	 * @var string|string[]
	 */
	public var $production_date;
	
	/**
	 * The date the item e.g. vehicle was purchased by the current owner.
	 * see : https://schema.org/purchaseDate
	 * @var string|string[]
	 */
	public var $purchase_date;
	
	/**
	 * The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	 * see : https://schema.org/releaseDate
	 * @var string|string[]
	 */
	public var $release_date;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
	/**
	 * The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	 * see : https://schema.org/sku
	 * @var string|string[]
	 */
	public var $sku;
	
	/**
	 * The weight of the product or person.
	 * see : https://schema.org/weight
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $weight;
	
	/**
	 * The width of the item.
	 * see : https://schema.org/width
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $width;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Product'
		);
		
		$serialized = so_json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->color );
		if ( ! empty( $serialized ) ) {
			$out['color'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->depth );
		if ( ! empty( $serialized ) ) {
			$out['depth'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->gtin_12 );
		if ( ! empty( $serialized ) ) {
			$out['gtin12'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->gtin_13 );
		if ( ! empty( $serialized ) ) {
			$out['gtin13'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->gtin_14 );
		if ( ! empty( $serialized ) ) {
			$out['gtin14'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->gtin_8 );
		if ( ! empty( $serialized ) ) {
			$out['gtin8'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_accessory_or_spare_part_for );
		if ( ! empty( $serialized ) ) {
			$out['isAccessoryOrSparePartFor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_consumable_for );
		if ( ! empty( $serialized ) ) {
			$out['isConsumableFor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_related_to );
		if ( ! empty( $serialized ) ) {
			$out['isRelatedTo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_similar_to );
		if ( ! empty( $serialized ) ) {
			$out['isSimilarTo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->item_condition );
		if ( ! empty( $serialized ) ) {
			$out['itemCondition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->manufacturer );
		if ( ! empty( $serialized ) ) {
			$out['manufacturer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->material );
		if ( ! empty( $serialized ) ) {
			$out['material'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->model );
		if ( ! empty( $serialized ) ) {
			$out['model'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->mpn );
		if ( ! empty( $serialized ) ) {
			$out['mpn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->producti_d );
		if ( ! empty( $serialized ) ) {
			$out['productID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_date );
		if ( ! empty( $serialized ) ) {
			$out['productionDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->purchase_date );
		if ( ! empty( $serialized ) ) {
			$out['purchaseDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->release_date );
		if ( ! empty( $serialized ) ) {
			$out['releaseDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sku );
		if ( ! empty( $serialized ) ) {
			$out['sku'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->weight );
		if ( ! empty( $serialized ) ) {
			$out['weight'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}
		
		return $out;
	}
}
