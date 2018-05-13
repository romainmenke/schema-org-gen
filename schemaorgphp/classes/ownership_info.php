<?php

class OwnershipInfo extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OwnershipInfo';
	
	/**
	 * The organization or person from which the product was acquired.
	 * see : https://schema.org/acquiredFrom
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $acquired_from;
	
	/**
	 * The date and time of obtaining the product.
	 * see : https://schema.org/ownedFrom
	 * @var string|string[]
	 */
	public var $owned_from;
	
	/**
	 * The date and time of giving up ownership on the product.
	 * see : https://schema.org/ownedThrough
	 * @var string|string[]
	 */
	public var $owned_through;
	
	/**
	 * The product that this structured value is referring to.
	 * see : https://schema.org/typeOfGood
	 * @var \Product|\Product[]|\Service|\Service[]
	 */
	public var $type_of_good;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OwnershipInfo'
		);
		
		$serialized = so_json_serialize( $this->acquired_from );
		if ( ! empty( $serialized ) ) {
			$out['acquiredFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->owned_from );
		if ( ! empty( $serialized ) ) {
			$out['ownedFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->owned_through );
		if ( ! empty( $serialized ) ) {
			$out['ownedThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->type_of_good );
		if ( ! empty( $serialized ) ) {
			$out['typeOfGood'] = $serialized;
		}
		
		return $out;
	}
}
