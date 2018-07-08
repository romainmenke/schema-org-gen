<?php

class EngineSpecification extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EngineSpecification';
	
	/**
	 * The volume swept by all of the pistons inside the cylinders of an internal combustion engine in a single movement. 
	 * 
	 * Typical unit code(s): CMQ for cubic centimeter, LTR for liters, INQ for cubic inches
	 * * Note 1: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	 * * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * see : https://auto.schema.org/engineDisplacement
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $engine_displacement;
	
	/**
	 * The power of the vehicle&#39;s engine.
	 *     Typical unit code(s): KWT for kilowatt, BHP for brake horsepower, N12 for metric horsepower (PS, with 1 PS = 735,49875 W)
	 * 
	 * 
	 * Note 1: There are many different ways of measuring an engine&#39;s power. For an overview, see  http://en.wikipedia.org/wiki/Horsepower#Engine (see: https://schema.orghttp://en.wikipedia.org/wiki/Horsepower#Engine_power_test_codes)powertest_codes.
	 * Note 2: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	 * Note 3: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/enginePower
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $engine_power;
	
	/**
	 * The type of engine or engines powering the vehicle.
	 * see : https://auto.schema.org/engineType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
	 */
	public var $engine_type;
	
	/**
	 * The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	 * see : https://schema.org/fuelType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
	 */
	public var $fuel_type;
	
	/**
	 * The torque (turning force) of the vehicle&#39;s engine.
	 * 
	 * Typical unit code(s): NU for newton metre (N m), F17 for pound-force per foot, or F48 for pound-force per inch
	 * 
	 * 
	 * Note 1: You can link to information about how the given value has been determined (e.g. reference RPM) using the valueReference (see: https://schema.org/valueReference) property.
	 * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/torque
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $torque;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EngineSpecification'
		);
		
		$serialized = so_json_serialize( $this->engine_displacement );
		if ( ! empty( $serialized ) ) {
			$out['engineDisplacement'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->engine_power );
		if ( ! empty( $serialized ) ) {
			$out['enginePower'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->engine_type );
		if ( ! empty( $serialized ) ) {
			$out['engineType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fuel_type );
		if ( ! empty( $serialized ) ) {
			$out['fuelType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->torque );
		if ( ! empty( $serialized ) ) {
			$out['torque'] = $serialized;
		}
		
		return $out;
	}
}
