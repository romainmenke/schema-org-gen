<?php

class OrderItem extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OrderItem';
	
	/**
	 * The delivery of the parcel related to this order or order item.
	 * see : https://schema.org/orderDelivery
	 * @var \ParcelDelivery|\ParcelDelivery[]
	 */
	public var $order_delivery;
	
	/**
	 * The identifier of the order item.
	 * see : https://schema.org/orderItemNumber
	 * @var string|string[]
	 */
	public var $order_item_number;
	
	/**
	 * The current status of the order item.
	 * see : https://schema.org/orderItemStatus
	 * @var \OrderStatus|\OrderStatus[]
	 */
	public var $order_item_status;
	
	/**
	 * The number of the item ordered. If the property is not set, assume the quantity is one.
	 * see : https://schema.org/orderQuantity
	 * @var float|float[]
	 */
	public var $order_quantity;
	
	/**
	 * The item ordered.
	 * see : https://schema.org/orderedItem
	 * @var \OrderItem|\OrderItem[]|\Product|\Product[]
	 */
	public var $ordered_item;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OrderItem'
		);
		
		$serialized = so_json_serialize( $this->order_delivery );
		if ( ! empty( $serialized ) ) {
			$out['orderDelivery'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_item_number );
		if ( ! empty( $serialized ) ) {
			$out['orderItemNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_item_status );
		if ( ! empty( $serialized ) ) {
			$out['orderItemStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_quantity );
		if ( ! empty( $serialized ) ) {
			$out['orderQuantity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ordered_item );
		if ( ! empty( $serialized ) ) {
			$out['orderedItem'] = $serialized;
		}
		
		return $out;
	}
}
