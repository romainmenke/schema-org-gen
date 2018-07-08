<?php

class CreditCard extends LoanOrCredit implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CreditCard';
	
	/**
	 * The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
	 * see : https://pending.schema.org/monthlyMinimumRepaymentAmount
	 * @var \MonetaryAmount|\MonetaryAmount[]|float|float[]
	 */
	public var $monthly_minimum_repayment_amount;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CreditCard'
		);
		
		$serialized = so_json_serialize( $this->monthly_minimum_repayment_amount );
		if ( ! empty( $serialized ) ) {
			$out['monthlyMinimumRepaymentAmount'] = $serialized;
		}
		
		return $out;
	}
}
