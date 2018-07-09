<?php

namespace SchemaOrg;

// DepositAccount see : https://schema.org/DepositAccount
class DepositAccount implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DepositAccount';
	
	
	/**
	 * The amount of money.
	 * see : https://schema.org/amount
	 * @var \MonetaryAmount | \MonetaryAmount[] | float | float[]
	 */
	public var $amount;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DepositAccount'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->amount );
		if ( ! empty( $serialized ) ) {
			$out['amount'] = $serialized;
		}
		
		return $out;
	}
}
