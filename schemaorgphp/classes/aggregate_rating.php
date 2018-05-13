<?php

class AggregateRating extends Rating implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AggregateRating';
	
	/**
	 * The item that is being reviewed/rated.
	 * see : https://schema.org/itemReviewed
	 * @var \Thing|\Thing[]
	 */
	public var $item_reviewed;
	
	/**
	 * The count of total number of ratings.
	 * see : https://schema.org/ratingCount
	 * @var integer|integer[]
	 */
	public var $rating_count;
	
	/**
	 * The count of total number of reviews.
	 * see : https://schema.org/reviewCount
	 * @var integer|integer[]
	 */
	public var $review_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AggregateRating'
		);
		
		$serialized = so_json_serialize( $this->item_reviewed );
		if ( ! empty( $serialized ) ) {
			$out['itemReviewed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->rating_count );
		if ( ! empty( $serialized ) ) {
			$out['ratingCount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review_count );
		if ( ! empty( $serialized ) ) {
			$out['reviewCount'] = $serialized;
		}
		
		return $out;
	}
}
