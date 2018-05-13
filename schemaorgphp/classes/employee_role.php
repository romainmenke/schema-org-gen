<?php

class EmployeeRole extends OrganizationRole implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EmployeeRole';
	
	/**
	 * The base salary of the job or of an employee in an EmployeeRole.
	 * see : https://schema.org/baseSalary
	 * @var \MonetaryAmount|\MonetaryAmount[]|float|float[]|\PriceSpecification|\PriceSpecification[]
	 */
	public var $base_salary;
	
	/**
	 * The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	 * see : https://schema.org/salaryCurrency
	 * @var string|string[]
	 */
	public var $salary_currency;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EmployeeRole'
		);
		
		$serialized = so_json_serialize( $this->base_salary );
		if ( ! empty( $serialized ) ) {
			$out['baseSalary'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->salary_currency );
		if ( ! empty( $serialized ) ) {
			$out['salaryCurrency'] = $serialized;
		}
		
		return $out;
	}
}
