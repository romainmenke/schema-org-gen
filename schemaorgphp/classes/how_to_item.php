<?php

class HowToItem extends ListItem implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToItem';
	
	/**
	 * The required quantity of the item(s).
	 * see : https://schema.org/requiredQuantity
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]|string|string[]
	 */
	public var $required_quantity;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToItem'
		);
		
		$serialized = so_json_serialize( $this->required_quantity );
		if ( ! empty( $serialized ) ) {
			$out['requiredQuantity'] = $serialized;
		}
		
		return $out;
	}
}