<?php

class SportsOrganization extends Organization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SportsOrganization';
	
	/**
	 * A type of sport (e.g. Baseball).
	 * see : https://schema.org/sport
	 * @var string|string[]|string|string[]
	 */
	public var $sport;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SportsOrganization'
		);
		
		$serialized = so_json_serialize( $this->sport );
		if ( ! empty( $serialized ) ) {
			$out['sport'] = $serialized;
		}
		
		return $out;
	}
}
