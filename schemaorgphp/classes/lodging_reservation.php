<?php

class LodgingReservation extends Reservation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LodgingReservation';
	
	/**
	 * The earliest someone may check into a lodging establishment.
	 * see : https://schema.org/checkinTime
	 * @var string|string[]
	 */
	public var $checkin_time;
	
	/**
	 * The latest someone may check out of a lodging establishment.
	 * see : https://schema.org/checkoutTime
	 * @var string|string[]
	 */
	public var $checkout_time;
	
	/**
	 * A full description of the lodging unit.
	 * see : https://schema.org/lodgingUnitDescription
	 * @var string|string[]
	 */
	public var $lodging_unit_description;
	
	/**
	 * Textual description of the unit type (including suite vs. room, size of bed, etc.).
	 * see : https://schema.org/lodgingUnitType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]
	 */
	public var $lodging_unit_type;
	
	/**
	 * The number of adults staying in the unit.
	 * see : https://schema.org/numAdults
	 * @var integer|integer[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $num_adults;
	
	/**
	 * The number of children staying in the unit.
	 * see : https://schema.org/numChildren
	 * @var integer|integer[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $num_children;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LodgingReservation'
		);
		
		$serialized = so_json_serialize( $this->checkin_time );
		if ( ! empty( $serialized ) ) {
			$out['checkinTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->checkout_time );
		if ( ! empty( $serialized ) ) {
			$out['checkoutTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lodging_unit_description );
		if ( ! empty( $serialized ) ) {
			$out['lodgingUnitDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lodging_unit_type );
		if ( ! empty( $serialized ) ) {
			$out['lodgingUnitType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->num_adults );
		if ( ! empty( $serialized ) ) {
			$out['numAdults'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->num_children );
		if ( ! empty( $serialized ) ) {
			$out['numChildren'] = $serialized;
		}
		
		return $out;
	}
}
