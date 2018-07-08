<?php

class SingleFamilyResidence extends House implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SingleFamilyResidence';
	
	/**
	 * The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.
	 * Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	 * see : https://schema.org/numberOfRooms
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_rooms;
	
	/**
	 * The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
	 * Typical unit code(s): C62 for person
	 * see : https://schema.org/occupancy
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $occupancy;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SingleFamilyResidence'
		);
		
		$serialized = so_json_serialize( $this->number_of_rooms );
		if ( ! empty( $serialized ) ) {
			$out['numberOfRooms'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->occupancy );
		if ( ! empty( $serialized ) ) {
			$out['occupancy'] = $serialized;
		}
		
		return $out;
	}
}
