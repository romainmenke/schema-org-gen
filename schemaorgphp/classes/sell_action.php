<?php

class SellAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SellAction';
	
	/**
	 * A sub property of participant. The participant/person/organization that bought the object.
	 * see : https://schema.org/buyer
	 * @var \Person|\Person[]
	 */
	public var $buyer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SellAction'
		);
		
		$serialized = so_json_serialize( $this->buyer );
		if ( ! empty( $serialized ) ) {
			$out['buyer'] = $serialized;
		}
		
		return $out;
	}
}
