<?php

class BuyAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BuyAction';
	
	/**
	 * An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	 * see : https://schema.org/seller
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $seller;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BuyAction'
		);
		
		$serialized = so_json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}
		
		return $out;
	}
}
