<?php

class Suite extends Accommodation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Suite';
	
	/**
	 * The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
	 *       If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	 * see : https://schema.org/bed
	 * @var \BedDetails|\BedDetails[]|\BedType|\BedType[]|string|string[]
	 */
	public var $bed;
	
	/**
	 * The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
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
			'@type' => 'Suite'
		);
		
		$serialized = so_json_serialize( $this->bed );
		if ( ! empty( $serialized ) ) {
			$out['bed'] = $serialized;
		}
		
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
