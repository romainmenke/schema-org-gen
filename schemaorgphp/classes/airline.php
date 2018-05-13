<?php

class Airline extends Organization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Airline';
	
	/**
	 * The type of boarding policy used by the airline (e.g. zone-based or group-based).
	 * see : https://schema.org/boardingPolicy
	 * @var \BoardingPolicyType|\BoardingPolicyType[]
	 */
	public var $boarding_policy;
	
	/**
	 * IATA identifier for an airline or airport.
	 * see : https://schema.org/iataCode
	 * @var string|string[]
	 */
	public var $iata_code;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Airline'
		);
		
		$serialized = so_json_serialize( $this->boarding_policy );
		if ( ! empty( $serialized ) ) {
			$out['boardingPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->iata_code );
		if ( ! empty( $serialized ) ) {
			$out['iataCode'] = $serialized;
		}
		
		return $out;
	}
}
