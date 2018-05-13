<?php

class RentAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'RentAction';
	
	/**
	 * A sub property of participant. The owner of the real estate property.
	 * see : https://schema.org/landlord
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $landlord;
	
	/**
	 * A sub property of participant. The real estate agent involved in the action.
	 * see : https://schema.org/realEstateAgent
	 * @var \RealEstateAgent|\RealEstateAgent[]
	 */
	public var $real_estate_agent;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'RentAction'
		);
		
		$serialized = so_json_serialize( $this->landlord );
		if ( ! empty( $serialized ) ) {
			$out['landlord'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->real_estate_agent );
		if ( ! empty( $serialized ) ) {
			$out['realEstateAgent'] = $serialized;
		}
		
		return $out;
	}
}
