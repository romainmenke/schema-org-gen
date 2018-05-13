<?php

class AccountingService extends FinancialService implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AccountingService';
	
	/**
	 * Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	 * see : https://schema.org/feesAndCommissionsSpecification
	 * @var string|string[]|string|string[]
	 */
	public var $fees_and_commissions_specification;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AccountingService'
		);
		
		$serialized = so_json_serialize( $this->fees_and_commissions_specification );
		if ( ! empty( $serialized ) ) {
			$out['feesAndCommissionsSpecification'] = $serialized;
		}
		
		return $out;
	}
}
