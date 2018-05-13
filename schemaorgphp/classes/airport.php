<?php

class Airport extends CivicStructure implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Airport';
	
	/**
	 * IATA identifier for an airline or airport.
	 * see : https://schema.org/iataCode
	 * @var string|string[]
	 */
	public var $iata_code;
	
	/**
	 * ICAO identifier for an airport.
	 * see : https://schema.org/icaoCode
	 * @var string|string[]
	 */
	public var $icao_code;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Airport'
		);
		
		$serialized = so_json_serialize( $this->iata_code );
		if ( ! empty( $serialized ) ) {
			$out['iataCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->icao_code );
		if ( ! empty( $serialized ) ) {
			$out['icaoCode'] = $serialized;
		}
		
		return $out;
	}
}
