<?php

class LoanOrCredit extends FinancialProduct implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LoanOrCredit';
	
	/**
	 * The amount of money.
	 * see : https://schema.org/amount
	 * @var \MonetaryAmount|\MonetaryAmount[]|float|float[]
	 */
	public var $amount;
	
	/**
	 * The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
	 * see : https://schema.org/currency
	 * @var string|string[]
	 */
	public var $currency;
	
	/**
	 * The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	 * see : http://pending.schema.org/gracePeriod
	 * @var \Duration|\Duration[]
	 */
	public var $grace_period;
	
	/**
	 * A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	 * see : http://pending.schema.org/loanRepaymentForm
	 * @var \RepaymentSpecification|\RepaymentSpecification[]
	 */
	public var $loan_repayment_form;
	
	/**
	 * The duration of the loan or credit agreement.
	 * see : https://schema.org/loanTerm
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $loan_term;
	
	/**
	 * The type of a loan or credit.
	 * see : http://pending.schema.org/loanType
	 * @var string|string[]|string|string[]
	 */
	public var $loan_type;
	
	/**
	 * The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	 * see : http://pending.schema.org/recourseLoan
	 * @var boolean|boolean[]
	 */
	public var $recourse_loan;
	
	/**
	 * Whether the terms for payment of interest can be renegotiated during the life of the loan.
	 * see : http://pending.schema.org/renegotiableLoan
	 * @var boolean|boolean[]
	 */
	public var $renegotiable_loan;
	
	/**
	 * Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	 * see : https://schema.org/requiredCollateral
	 * @var string|string[]|\Thing|\Thing[]
	 */
	public var $required_collateral;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LoanOrCredit'
		);
		
		$serialized = so_json_serialize( $this->amount );
		if ( ! empty( $serialized ) ) {
			$out['amount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->currency );
		if ( ! empty( $serialized ) ) {
			$out['currency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->grace_period );
		if ( ! empty( $serialized ) ) {
			$out['gracePeriod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->loan_repayment_form );
		if ( ! empty( $serialized ) ) {
			$out['loanRepaymentForm'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->loan_term );
		if ( ! empty( $serialized ) ) {
			$out['loanTerm'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->loan_type );
		if ( ! empty( $serialized ) ) {
			$out['loanType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recourse_loan );
		if ( ! empty( $serialized ) ) {
			$out['recourseLoan'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->renegotiable_loan );
		if ( ! empty( $serialized ) ) {
			$out['renegotiableLoan'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->required_collateral );
		if ( ! empty( $serialized ) ) {
			$out['requiredCollateral'] = $serialized;
		}
		
		return $out;
	}
}
