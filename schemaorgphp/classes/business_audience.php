<?php

class BusinessAudience extends Audience implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BusinessAudience';
	
	/**
	 * The number of employees in an organization e.g. business.
	 * see : https://schema.org/numberOfEmployees
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_employees;
	
	/**
	 * The size of the business in annual revenue.
	 * see : https://schema.org/yearlyRevenue
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $yearly_revenue;
	
	/**
	 * The age of the business.
	 * see : https://schema.org/yearsInOperation
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $years_in_operation;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BusinessAudience'
		);
		
		$serialized = so_json_serialize( $this->number_of_employees );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEmployees'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->yearly_revenue );
		if ( ! empty( $serialized ) ) {
			$out['yearlyRevenue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->years_in_operation );
		if ( ! empty( $serialized ) ) {
			$out['yearsInOperation'] = $serialized;
		}
		
		return $out;
	}
}
