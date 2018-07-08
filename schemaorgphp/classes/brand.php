<?php

class Brand extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Brand';
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating|\AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject|\ImageObject[]|string|string[]
	 */
	public var $logo;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Brand'
		);
		
		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		return $out;
	}
}
