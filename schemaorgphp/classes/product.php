<?php

class Product extends Thing implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Product';
	
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
	 * An award won by or for this item.
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
	 * see : https://pending.schema.org/category
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
	 * A pointer to another product (or multiple products) for which this product is a consumable.
	 * see : https://schema.org/isConsumableFor
	 * @var \Product|\Product[]
	 */
	public var $is_consumable_for;
	
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
	 * The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	 * see : https://schema.org/mpn
	 * @var string|string[]
	 */
	public var $mpn;
	
	/**
	 * The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	 * see : https://schema.org/productID
	 * @var string|string[]
	 */
	public var $producti_d;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
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
		
		$serialized = so_json_serialize( $this->gtin_8 );
		if ( ! empty( $serialized ) ) {
			$out['gtin8'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_consumable_for );
		if ( ! empty( $serialized ) ) {
			$out['isConsumableFor'] = $serialized;
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
		
		$serialized = so_json_serialize( $this->mpn );
		if ( ! empty( $serialized ) ) {
			$out['mpn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->producti_d );
		if ( ! empty( $serialized ) ) {
			$out['productID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}
		
		return $out;
	}
}
