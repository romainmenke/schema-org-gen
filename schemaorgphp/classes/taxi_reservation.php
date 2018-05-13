<?php

class TaxiReservation extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TaxiReservation';
	
	/**
	 * Number of people the reservation should accommodate.
	 * see : https://schema.org/partySize
	 * @var integer|integer[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $party_size;
	
	/**
	 * Where a taxi will pick up a passenger or a rental car can be picked up.
	 * see : https://schema.org/pickupLocation
	 * @var \Place|\Place[]
	 */
	public var $pickup_location;
	
	/**
	 * When a taxi will pickup a passenger or a rental car can be picked up.
	 * see : https://schema.org/pickupTime
	 * @var string|string[]
	 */
	public var $pickup_time;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TaxiReservation'
		);
		
		$serialized = so_json_serialize( $this->party_size );
		if ( ! empty( $serialized ) ) {
			$out['partySize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->pickup_location );
		if ( ! empty( $serialized ) ) {
			$out['pickupLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->pickup_time );
		if ( ! empty( $serialized ) ) {
			$out['pickupTime'] = $serialized;
		}
		
		return $out;
	}
}
