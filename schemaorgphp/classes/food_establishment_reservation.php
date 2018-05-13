<?php

class FoodEstablishmentReservation extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'FoodEstablishmentReservation';
	
	/**
	 * The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/endTime
	 * @var string|string[]
	 */
	public var $end_time;
	
	/**
	 * Number of people the reservation should accommodate.
	 * see : https://schema.org/partySize
	 * @var integer|integer[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $party_size;
	
	/**
	 * The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/startTime
	 * @var string|string[]
	 */
	public var $start_time;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'FoodEstablishmentReservation'
		);
		
		$serialized = so_json_serialize( $this->end_time );
		if ( ! empty( $serialized ) ) {
			$out['endTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->party_size );
		if ( ! empty( $serialized ) ) {
			$out['partySize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_time );
		if ( ! empty( $serialized ) ) {
			$out['startTime'] = $serialized;
		}
		
		return $out;
	}
}
