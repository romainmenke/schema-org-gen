<?php

class Car extends Vehicle implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Car';
	
	/**
	 * The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
	 * see : https://schema.org/acrissCode
	 * @var string|string[]
	 */
	public var $acriss_code;
	
	/**
	 * The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * 
	 * 
	 * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference)
	 * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://schema.org/roofLoad
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $roof_load;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Car'
		);
		
		$serialized = so_json_serialize( $this->acriss_code );
		if ( ! empty( $serialized ) ) {
			$out['acrissCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->roof_load );
		if ( ! empty( $serialized ) ) {
			$out['roofLoad'] = $serialized;
		}
		
		return $out;
	}
}
