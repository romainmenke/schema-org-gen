<?php

class PaymentService extends FinancialProduct implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PaymentService';
	
	/**
	 * The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	 * see : https://schema.org/annualPercentageRate
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $annual_percentage_rate;
	
	/**
	 * Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	 * see : https://schema.org/feesAndCommissionsSpecification
	 * @var string|string[]|string|string[]
	 */
	public var $fees_and_commissions_specification;
	
	/**
	 * The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	 * see : https://schema.org/interestRate
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $interest_rate;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PaymentService'
		);
		
		$serialized = so_json_serialize( $this->annual_percentage_rate );
		if ( ! empty( $serialized ) ) {
			$out['annualPercentageRate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fees_and_commissions_specification );
		if ( ! empty( $serialized ) ) {
			$out['feesAndCommissionsSpecification'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->interest_rate );
		if ( ! empty( $serialized ) ) {
			$out['interestRate'] = $serialized;
		}
		
		return $out;
	}
}
