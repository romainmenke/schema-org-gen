<?php

class BankAccount extends FinancialProduct implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BankAccount';
	
	/**
	 * A minimum amount that has to be paid in every month.
	 * see : https://pending.schema.org/accountMinimumInflow
	 * @var \MonetaryAmount|\MonetaryAmount[]
	 */
	public var $account_minimum_inflow;
	
	/**
	 * An overdraft is an extension of credit from a lending institution when an account reaches zero. An overdraft allows the individual to continue withdrawing money even if the account has no funds in it. Basically the bank allows people to borrow a set amount of money.
	 * see : https://pending.schema.org/accountOverdraftLimit
	 * @var \MonetaryAmount|\MonetaryAmount[]
	 */
	public var $account_overdraft_limit;
	
	/**
	 * The type of a bank account.
	 * see : https://pending.schema.org/bankAccountType
	 * @var string|string[]|string|string[]
	 */
	public var $bank_account_type;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BankAccount'
		);
		
		$serialized = so_json_serialize( $this->account_minimum_inflow );
		if ( ! empty( $serialized ) ) {
			$out['accountMinimumInflow'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->account_overdraft_limit );
		if ( ! empty( $serialized ) ) {
			$out['accountOverdraftLimit'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->bank_account_type );
		if ( ! empty( $serialized ) ) {
			$out['bankAccountType'] = $serialized;
		}
		
		return $out;
	}
}
