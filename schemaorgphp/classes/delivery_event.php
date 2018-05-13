<?php

class DeliveryEvent extends Event implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DeliveryEvent';
	
	/**
	 * Password, PIN, or access code needed for delivery (e.g. from a locker).
	 * see : https://schema.org/accessCode
	 * @var string|string[]
	 */
	public var $access_code;
	
	/**
	 * When the item is available for pickup from the store, locker, etc.
	 * see : https://schema.org/availableFrom
	 * @var string|string[]
	 */
	public var $available_from;
	
	/**
	 * After this date, the item will no longer be available for pickup.
	 * see : https://schema.org/availableThrough
	 * @var string|string[]
	 */
	public var $available_through;
	
	/**
	 * Method used for delivery or shipping.
	 * see : https://schema.org/hasDeliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $has_delivery_method;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DeliveryEvent'
		);
		
		$serialized = so_json_serialize( $this->access_code );
		if ( ! empty( $serialized ) ) {
			$out['accessCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_from );
		if ( ! empty( $serialized ) ) {
			$out['availableFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_through );
		if ( ! empty( $serialized ) ) {
			$out['availableThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['hasDeliveryMethod'] = $serialized;
		}
		
		return $out;
	}
}
