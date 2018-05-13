<?php

class PaymentChargeSpecification extends PriceSpecification implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PaymentChargeSpecification';
	
	/**
	 * The delivery method(s) to which the delivery charge or payment charge specification applies.
	 * see : https://schema.org/appliesToDeliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $applies_to_delivery_method;
	
	/**
	 * The payment method(s) to which the payment charge specification applies.
	 * see : https://schema.org/appliesToPaymentMethod
	 * @var \PaymentMethod|\PaymentMethod[]
	 */
	public var $applies_to_payment_method;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PaymentChargeSpecification'
		);
		
		$serialized = so_json_serialize( $this->applies_to_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['appliesToDeliveryMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->applies_to_payment_method );
		if ( ! empty( $serialized ) ) {
			$out['appliesToPaymentMethod'] = $serialized;
		}
		
		return $out;
	}
}
