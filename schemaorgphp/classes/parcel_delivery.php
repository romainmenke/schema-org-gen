<?php

class ParcelDelivery extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ParcelDelivery';
	
	/**
	 * Destination address.
	 * see : https://schema.org/deliveryAddress
	 * @var \PostalAddress|\PostalAddress[]
	 */
	public var $delivery_address;
	
	/**
	 * New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	 * see : https://schema.org/deliveryStatus
	 * @var \DeliveryEvent|\DeliveryEvent[]
	 */
	public var $delivery_status;
	
	/**
	 * The earliest date the package may arrive.
	 * see : https://schema.org/expectedArrivalFrom
	 * @var string|string[]
	 */
	public var $expected_arrival_from;
	
	/**
	 * The latest date the package may arrive.
	 * see : https://schema.org/expectedArrivalUntil
	 * @var string|string[]
	 */
	public var $expected_arrival_until;
	
	/**
	 * Method used for delivery or shipping.
	 * see : https://schema.org/hasDeliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $has_delivery_method;
	
	/**
	 * Item(s) being shipped.
	 * see : https://schema.org/itemShipped
	 * @var \Product|\Product[]
	 */
	public var $item_shipped;
	
	/**
	 * Shipper&#39;s address.
	 * see : https://schema.org/originAddress
	 * @var \PostalAddress|\PostalAddress[]
	 */
	public var $origin_address;
	
	/**
	 * The overall order the items in this delivery were included in.
	 * see : https://schema.org/partOfOrder
	 * @var \Order|\Order[]
	 */
	public var $part_of_order;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	/**
	 * Shipper tracking number.
	 * see : https://schema.org/trackingNumber
	 * @var string|string[]
	 */
	public var $tracking_number;
	
	/**
	 * Tracking url for the parcel delivery.
	 * see : https://schema.org/trackingUrl
	 * @var string|string[]
	 */
	public var $tracking_url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ParcelDelivery'
		);
		
		$serialized = so_json_serialize( $this->delivery_address );
		if ( ! empty( $serialized ) ) {
			$out['deliveryAddress'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->delivery_status );
		if ( ! empty( $serialized ) ) {
			$out['deliveryStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->expected_arrival_from );
		if ( ! empty( $serialized ) ) {
			$out['expectedArrivalFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->expected_arrival_until );
		if ( ! empty( $serialized ) ) {
			$out['expectedArrivalUntil'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['hasDeliveryMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->item_shipped );
		if ( ! empty( $serialized ) ) {
			$out['itemShipped'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->origin_address );
		if ( ! empty( $serialized ) ) {
			$out['originAddress'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_order );
		if ( ! empty( $serialized ) ) {
			$out['partOfOrder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->tracking_number );
		if ( ! empty( $serialized ) ) {
			$out['trackingNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->tracking_url );
		if ( ! empty( $serialized ) ) {
			$out['trackingUrl'] = $serialized;
		}
		
		return $out;
	}
}
