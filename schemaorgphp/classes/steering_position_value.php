<?php

class SteeringPositionValue extends QualitativeValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SteeringPositionValue';
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue|\PropertyValue[]
	 */
	public var $additional_property;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is equal to the object.
	 * see : https://schema.org/equal
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $equal;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is greater than the object.
	 * see : https://schema.org/greater
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $greater;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
	 * see : https://schema.org/greaterOrEqual
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $greater_or_equal;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is lesser than the object.
	 * see : https://schema.org/lesser
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $lesser;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
	 * see : https://schema.org/lesserOrEqual
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $lesser_or_equal;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is not equal to the object.
	 * see : https://schema.org/nonEqual
	 * @var \QualitativeValue|\QualitativeValue[]
	 */
	public var $non_equal;
	
	/**
	 * A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	 * see : https://schema.org/valueReference
	 * @var \Enumeration|\Enumeration[]|\PropertyValue|\PropertyValue[]|\QualitativeValue|\QualitativeValue[]|\QuantitativeValue|\QuantitativeValue[]|\StructuredValue|\StructuredValue[]
	 */
	public var $value_reference;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SteeringPositionValue'
		);
		
		$serialized = so_json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->equal );
		if ( ! empty( $serialized ) ) {
			$out['equal'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->greater );
		if ( ! empty( $serialized ) ) {
			$out['greater'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->greater_or_equal );
		if ( ! empty( $serialized ) ) {
			$out['greaterOrEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lesser );
		if ( ! empty( $serialized ) ) {
			$out['lesser'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lesser_or_equal );
		if ( ! empty( $serialized ) ) {
			$out['lesserOrEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->non_equal );
		if ( ! empty( $serialized ) ) {
			$out['nonEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_reference );
		if ( ! empty( $serialized ) ) {
			$out['valueReference'] = $serialized;
		}
		
		return $out;
	}
}
