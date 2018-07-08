<?php

class PaymentCard extends FinancialProduct implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PaymentCard';
	
	/**
	 * A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
	 * see : https://pending.schema.org/cashBack
	 * @var boolean|boolean[]|float|float[]
	 */
	public var $cash_back;
	
	/**
	 * A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
	 * see : https://pending.schema.org/contactlessPayment
	 * @var boolean|boolean[]
	 */
	public var $contactless_payment;
	
	/**
	 * A floor limit is the amount of money above which credit card transactions must be authorized.
	 * see : https://pending.schema.org/floorLimit
	 * @var \MonetaryAmount|\MonetaryAmount[]
	 */
	public var $floor_limit;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PaymentCard'
		);
		
		$serialized = so_json_serialize( $this->cash_back );
		if ( ! empty( $serialized ) ) {
			$out['cashBack'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contactless_payment );
		if ( ! empty( $serialized ) ) {
			$out['contactlessPayment'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->floor_limit );
		if ( ! empty( $serialized ) ) {
			$out['floorLimit'] = $serialized;
		}
		
		return $out;
	}
}
