<?php

class TVSeason extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TVSeason';
	
	/**
	 * The country of the principal offices of the production company or individual responsible for the movie or program.
	 * see : https://schema.org/countryOfOrigin
	 * @var \Country|\Country[]
	 */
	public var $country_of_origin;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TVSeason'
		);
		
		$serialized = so_json_serialize( $this->country_of_origin );
		if ( ! empty( $serialized ) ) {
			$out['countryOfOrigin'] = $serialized;
		}
		
		return $out;
	}
}
