<?php

class House extends Accommodation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'House';
	
	/**
	 * The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	 * Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	 * see : https://schema.org/numberOfRooms
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_rooms;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'House'
		);
		
		$serialized = so_json_serialize( $this->number_of_rooms );
		if ( ! empty( $serialized ) ) {
			$out['numberOfRooms'] = $serialized;
		}
		
		return $out;
	}
}
