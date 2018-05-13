<?php

class ReceiveAction extends TransferAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReceiveAction';
	
	/**
	 * A sub property of instrument. The method of delivery.
	 * see : https://schema.org/deliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $delivery_method;
	
	/**
	 * A sub property of participant. The participant who is at the sending end of the action.
	 * see : https://schema.org/sender
	 * @var \Audience|\Audience[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $sender;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReceiveAction'
		);
		
		$serialized = so_json_serialize( $this->delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['deliveryMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sender );
		if ( ! empty( $serialized ) ) {
			$out['sender'] = $serialized;
		}
		
		return $out;
	}
}
