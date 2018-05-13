<?php

class UnitPriceSpecification extends PriceSpecification implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'UnitPriceSpecification';
	
	/**
	 * This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
	 * see : https://schema.org/billingIncrement
	 * @var float|float[]
	 */
	public var $billing_increment;
	
	/**
	 * A short text or acronym indicating multiple price specifications for the same offer, e.g. SRP for the suggested retail price or INVOICE for the invoice price, mostly used in the car industry.
	 * see : https://schema.org/priceType
	 * @var string|string[]
	 */
	public var $price_type;
	
	/**
	 * The reference quantity for which a certain price applies, e.g. 1 EUR per 4 kWh of electricity. This property is a replacement for unitOfMeasurement for the advanced cases where the price does not relate to a standard unit.
	 * see : https://schema.org/referenceQuantity
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $reference_quantity;
	
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
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'UnitPriceSpecification'
		);
		
		$serialized = so_json_serialize( $this->billing_increment );
		if ( ! empty( $serialized ) ) {
			$out['billingIncrement'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_type );
		if ( ! empty( $serialized ) ) {
			$out['priceType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reference_quantity );
		if ( ! empty( $serialized ) ) {
			$out['referenceQuantity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_code );
		if ( ! empty( $serialized ) ) {
			$out['unitCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_text );
		if ( ! empty( $serialized ) ) {
			$out['unitText'] = $serialized;
		}
		
		return $out;
	}
}
