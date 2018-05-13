<?php

class TrackAction extends FindAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TrackAction';
	
	/**
	 * A sub property of instrument. The method of delivery.
	 * see : https://schema.org/deliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $delivery_method;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TrackAction'
		);
		
		$serialized = so_json_serialize( $this->delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['deliveryMethod'] = $serialized;
		}
		
		return $out;
	}
}
