<?php

class TypeAndQuantityNode extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TypeAndQuantityNode';
	
	/**
	 * The quantity of the goods included in the offer.
	 * see : https://schema.org/amountOfThisGood
	 * @var float|float[]
	 */
	public var $amount_of_this_good;
	
	/**
	 * The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	 * see : https://schema.org/businessFunction
	 * @var \BusinessFunction|\BusinessFunction[]
	 */
	public var $business_function;
	
	/**
	 * The product that this structured value is referring to.
	 * see : https://schema.org/typeOfGood
	 * @var \Product|\Product[]|\Service|\Service[]
	 */
	public var $type_of_good;
	
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
			'@type' => 'TypeAndQuantityNode'
		);
		
		$serialized = so_json_serialize( $this->amount_of_this_good );
		if ( ! empty( $serialized ) ) {
			$out['amountOfThisGood'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->business_function );
		if ( ! empty( $serialized ) ) {
			$out['businessFunction'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->type_of_good );
		if ( ! empty( $serialized ) ) {
			$out['typeOfGood'] = $serialized;
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
