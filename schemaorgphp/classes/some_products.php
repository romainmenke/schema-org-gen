<?php

class SomeProducts extends Product implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SomeProducts';
	
	/**
	 * The current approximate inventory level for the item or items.
	 * see : https://schema.org/inventoryLevel
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $inventory_level;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SomeProducts'
		);
		
		$serialized = so_json_serialize( $this->inventory_level );
		if ( ! empty( $serialized ) ) {
			$out['inventoryLevel'] = $serialized;
		}
		
		return $out;
	}
}
