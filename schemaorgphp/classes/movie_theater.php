<?php

namespace SchemaOrg;

// MovieTheater see : https://schema.org/MovieTheater
class MovieTheater implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MovieTheater';
	
	
	/**
	 * The number of screens in the movie theater.
	 * see : https://schema.org/screenCount
	 * @var float | float[]
	 */
	public $screen_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MovieTheater'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->screen_count );
		if ( ! empty( $serialized ) ) {
			$out['screenCount'] = $serialized;
		}
		
		return $out;
	}
}
