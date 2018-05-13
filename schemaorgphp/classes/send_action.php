<?php

class SendAction extends TransferAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SendAction';
	
	/**
	 * A sub property of instrument. The method of delivery.
	 * see : https://schema.org/deliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $delivery_method;
	
	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $recipient;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SendAction'
		);
		
		$serialized = so_json_serialize( $this->delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['deliveryMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}
		
		return $out;
	}
}
