<?php

class Accommodation extends Place implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Accommodation';
	
	/**
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 * @var \LocationFeatureSpecification|\LocationFeatureSpecification[]
	 */
	public var $amenity_feature;
	
	/**
	 * The size of the accommodation, e.g. in square meter or squarefoot.
	 * Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	 * see : https://schema.org/floorSize
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $floor_size;
	
	/**
	 * The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
	 * Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	 * see : https://schema.org/numberOfRooms
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_rooms;
	
	/**
	 * Indications regarding the permitted usage of the accommodation.
	 * see : https://schema.org/permittedUsage
	 * @var string|string[]
	 */
	public var $permitted_usage;
	
	/**
	 * Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	 * see : https://schema.org/petsAllowed
	 * @var boolean|boolean[]|string|string[]
	 */
	public var $pets_allowed;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Accommodation'
		);
		
		$serialized = so_json_serialize( $this->amenity_feature );
		if ( ! empty( $serialized ) ) {
			$out['amenityFeature'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->floor_size );
		if ( ! empty( $serialized ) ) {
			$out['floorSize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_rooms );
		if ( ! empty( $serialized ) ) {
			$out['numberOfRooms'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->permitted_usage );
		if ( ! empty( $serialized ) ) {
			$out['permittedUsage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->pets_allowed );
		if ( ! empty( $serialized ) ) {
			$out['petsAllowed'] = $serialized;
		}
		
		return $out;
	}
}
