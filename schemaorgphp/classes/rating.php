<?php

class Rating extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Rating';
	
	/**
	 * The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	 * see : https://schema.org/author
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $author;
	
	/**
	 * The highest value allowed in this rating system. If bestRating is omitted, 5 is assumed.
	 * see : https://schema.org/bestRating
	 * @var float|float[]|string|string[]
	 */
	public var $best_rating;
	
	/**
	 * The rating for the content.
	 * see : https://schema.org/ratingValue
	 * @var float|float[]|string|string[]
	 */
	public var $rating_value;
	
	/**
	 * The lowest value allowed in this rating system. If worstRating is omitted, 1 is assumed.
	 * see : https://schema.org/worstRating
	 * @var float|float[]|string|string[]
	 */
	public var $worst_rating;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Rating'
		);
		
		$serialized = so_json_serialize( $this->author );
		if ( ! empty( $serialized ) ) {
			$out['author'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->best_rating );
		if ( ! empty( $serialized ) ) {
			$out['bestRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->rating_value );
		if ( ! empty( $serialized ) ) {
			$out['ratingValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->worst_rating );
		if ( ! empty( $serialized ) ) {
			$out['worstRating'] = $serialized;
		}
		
		return $out;
	}
}
