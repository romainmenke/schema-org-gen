<?php

class RentalCarReservation extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'RentalCarReservation';
	
	/**
	 * Where a rental car can be dropped off.
	 * see : https://schema.org/dropoffLocation
	 * @var \Place|\Place[]
	 */
	public var $dropoff_location;
	
	/**
	 * When a rental car can be dropped off.
	 * see : https://schema.org/dropoffTime
	 * @var string|string[]
	 */
	public var $dropoff_time;
	
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
			'@type' => 'RentalCarReservation'
		);
		
		$serialized = so_json_serialize( $this->dropoff_location );
		if ( ! empty( $serialized ) ) {
			$out['dropoffLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->dropoff_time );
		if ( ! empty( $serialized ) ) {
			$out['dropoffTime'] = $serialized;
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
