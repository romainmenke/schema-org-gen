<?php

class QuantitativeValue extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'QuantitativeValue';
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue|\PropertyValue[]
	 */
	public var $additional_property;
	
	/**
	 * The upper value of some characteristic or property.
	 * see : https://schema.org/maxValue
	 * @var float|float[]
	 */
	public var $max_value;
	
	/**
	 * The lower value of some characteristic or property.
	 * see : https://schema.org/minValue
	 * @var float|float[]
	 */
	public var $min_value;
	
	/**
	 * The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	 * see : https://schema.org/unitCode
	 * @var string|string[]|string|string[]
	 */
	public var $unit_code;
	
	/**
	 * A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	 * unitCode (see: https://schema.orgunitCode).
	 * see : https://schema.org/unitText
	 * @var string|string[]
	 */
	public var $unit_text;
	
	/**
	 * The value of the quantitative value or property value node.
	 * 
	 * 
	 * For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is &#39;Number&#39;.
	 * For PropertyValue (see: https://schema.org/PropertyValue), it can be &#39;Text;&#39;, &#39;Number&#39;, &#39;Boolean&#39;, or &#39;StructuredValue&#39;.
	 * 
	 * 
	 * see : https://schema.org/value
	 * @var boolean|boolean[]|float|float[]|\StructuredValue|\StructuredValue[]|string|string[]
	 */
	public var $value;
	
	/**
	 * A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	 * see : https://schema.org/valueReference
	 * @var \Enumeration|\Enumeration[]|\PropertyValue|\PropertyValue[]|\QualitativeValue|\QualitativeValue[]|\QuantitativeValue|\QuantitativeValue[]|\StructuredValue|\StructuredValue[]
	 */
	public var $value_reference;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'QuantitativeValue'
		);
		
		$serialized = so_json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->max_value );
		if ( ! empty( $serialized ) ) {
			$out['maxValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->min_value );
		if ( ! empty( $serialized ) ) {
			$out['minValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_code );
		if ( ! empty( $serialized ) ) {
			$out['unitCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_text );
		if ( ! empty( $serialized ) ) {
			$out['unitText'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value );
		if ( ! empty( $serialized ) ) {
			$out['value'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_reference );
		if ( ! empty( $serialized ) ) {
			$out['valueReference'] = $serialized;
		}
		
		return $out;
	}
}
