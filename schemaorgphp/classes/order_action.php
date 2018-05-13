<?php

class OrderAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OrderAction';
	
	/**
	 * A sub property of instrument. The method of delivery.
	 * see : https://schema.org/deliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $delivery_method;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OrderAction'
		);
		
		$serialized = so_json_serialize( $this->delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['deliveryMethod'] = $serialized;
		}
		
		return $out;
	}
}
