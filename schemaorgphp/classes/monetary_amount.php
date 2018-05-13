<?php

class MonetaryAmount extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MonetaryAmount';
	
	/**
	 * The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
	 * see : https://schema.org/currency
	 * @var string|string[]
	 */
	public var $currency;
	
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
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 * @var string|string[]
	 */
	public var $valid_from;
	
	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 * @var string|string[]
	 */
	public var $valid_through;
	
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
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MonetaryAmount'
		);
		
		$serialized = so_json_serialize( $this->currency );
		if ( ! empty( $serialized ) ) {
			$out['currency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->max_value );
		if ( ! empty( $serialized ) ) {
			$out['maxValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->min_value );
		if ( ! empty( $serialized ) ) {
			$out['minValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value );
		if ( ! empty( $serialized ) ) {
			$out['value'] = $serialized;
		}
		
		return $out;
	}
}
