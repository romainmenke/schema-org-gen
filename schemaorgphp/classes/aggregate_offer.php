<?php

class AggregateOffer extends Offer implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AggregateOffer';
	
	/**
	 * The highest price of all offers available.
	 * see : https://schema.org/highPrice
	 * @var float|float[]|string|string[]
	 */
	public var $high_price;
	
	/**
	 * The lowest price of all offers available.
	 * see : https://schema.org/lowPrice
	 * @var float|float[]|string|string[]
	 */
	public var $low_price;
	
	/**
	 * The number of offers for the product.
	 * see : https://schema.org/offerCount
	 * @var integer|integer[]
	 */
	public var $offer_count;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer|\Offer[]
	 */
	public var $offers;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AggregateOffer'
		);
		
		$serialized = so_json_serialize( $this->high_price );
		if ( ! empty( $serialized ) ) {
			$out['highPrice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->low_price );
		if ( ! empty( $serialized ) ) {
			$out['lowPrice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offer_count );
		if ( ! empty( $serialized ) ) {
			$out['offerCount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		return $out;
	}
}
