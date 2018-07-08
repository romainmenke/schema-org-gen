<?php

class Review extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Review';
	
	/**
	 * The item that is being reviewed/rated.
	 * see : https://schema.org/itemReviewed
	 * @var \Thing|\Thing[]
	 */
	public var $item_reviewed;
	
	/**
	 * This Review or Rating is relevant to this part or facet of the itemReviewed.
	 * see : https://pending.schema.org/reviewAspect
	 * @var string|string[]
	 */
	public var $review_aspect;
	
	/**
	 * The actual body of the review.
	 * see : https://schema.org/reviewBody
	 * @var string|string[]
	 */
	public var $review_body;
	
	/**
	 * The rating given in this review. Note that reviews can themselves be rated. The reviewRating applies to rating given by the review. The aggregateRating (see: https://schema.org/aggregateRating) property applies to the review itself, as a creative work.
	 * see : https://schema.org/reviewRating
	 * @var \Rating|\Rating[]
	 */
	public var $review_rating;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Review'
		);
		
		$serialized = so_json_serialize( $this->item_reviewed );
		if ( ! empty( $serialized ) ) {
			$out['itemReviewed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review_aspect );
		if ( ! empty( $serialized ) ) {
			$out['reviewAspect'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review_body );
		if ( ! empty( $serialized ) ) {
			$out['reviewBody'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review_rating );
		if ( ! empty( $serialized ) ) {
			$out['reviewRating'] = $serialized;
		}
		
		return $out;
	}
}
