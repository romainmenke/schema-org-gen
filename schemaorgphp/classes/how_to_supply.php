<?php

class HowToSupply extends HowToItem implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToSupply';
	
	/**
	 * The estimated cost of the supply or supplies consumed when performing instructions.
	 * see : https://schema.org/estimatedCost
	 * @var \MonetaryAmount|\MonetaryAmount[]|string|string[]
	 */
	public var $estimated_cost;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToSupply'
		);
		
		$serialized = so_json_serialize( $this->estimated_cost );
		if ( ! empty( $serialized ) ) {
			$out['estimatedCost'] = $serialized;
		}
		
		return $out;
	}
}
