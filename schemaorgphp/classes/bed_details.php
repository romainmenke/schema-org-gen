<?php

class BedDetails extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BedDetails';
	
	/**
	 * The quantity of the given bed type available in the HotelRoom, Suite, House, or Apartment.
	 * see : https://schema.org/numberOfBeds
	 * @var float|float[]
	 */
	public var $number_of_beds;
	
	/**
	 * The type of bed to which the BedDetail refers, i.e. the type of bed available in the quantity indicated by quantity.
	 * see : https://schema.org/typeOfBed
	 * @var \BedType|\BedType[]|string|string[]
	 */
	public var $type_of_bed;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BedDetails'
		);
		
		$serialized = so_json_serialize( $this->number_of_beds );
		if ( ! empty( $serialized ) ) {
			$out['numberOfBeds'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->type_of_bed );
		if ( ! empty( $serialized ) ) {
			$out['typeOfBed'] = $serialized;
		}
		
		return $out;
	}
}
