<?php

class FlightReservation extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'FlightReservation';
	
	/**
	 * The airline-specific indicator of boarding order / preference.
	 * see : https://schema.org/boardingGroup
	 * @var string|string[]
	 */
	public var $boarding_group;
	
	/**
	 * The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
	 * see : https://schema.org/passengerPriorityStatus
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]
	 */
	public var $passenger_priority_status;
	
	/**
	 * The passenger&#39;s sequence number as assigned by the airline.
	 * see : https://schema.org/passengerSequenceNumber
	 * @var string|string[]
	 */
	public var $passenger_sequence_number;
	
	/**
	 * The type of security screening the passenger is subject to.
	 * see : https://schema.org/securityScreening
	 * @var string|string[]
	 */
	public var $security_screening;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'FlightReservation'
		);
		
		$serialized = so_json_serialize( $this->boarding_group );
		if ( ! empty( $serialized ) ) {
			$out['boardingGroup'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->passenger_priority_status );
		if ( ! empty( $serialized ) ) {
			$out['passengerPriorityStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->passenger_sequence_number );
		if ( ! empty( $serialized ) ) {
			$out['passengerSequenceNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->security_screening );
		if ( ! empty( $serialized ) ) {
			$out['securityScreening'] = $serialized;
		}
		
		return $out;
	}
}
