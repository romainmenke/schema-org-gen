<?php

class ReservationPackage extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReservationPackage';
	
	/**
	 * The individual reservations included in the package. Typically a repeated property.
	 * see : https://schema.org/subReservation
	 * @var \Reservation|\Reservation[]
	 */
	public var $sub_reservation;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReservationPackage'
		);
		
		$serialized = so_json_serialize( $this->sub_reservation );
		if ( ! empty( $serialized ) ) {
			$out['subReservation'] = $serialized;
		}
		
		return $out;
	}
}
