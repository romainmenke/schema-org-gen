<?php

class CompoundPriceSpecification extends PriceSpecification implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CompoundPriceSpecification';
	
	/**
	 * This property links to all UnitPriceSpecification (see: https://schema.org/UnitPriceSpecification) nodes that apply in parallel for the CompoundPriceSpecification (see: https://schema.org/CompoundPriceSpecification) node.
	 * see : https://schema.org/priceComponent
	 * @var \UnitPriceSpecification|\UnitPriceSpecification[]
	 */
	public var $price_component;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CompoundPriceSpecification'
		);
		
		$serialized = so_json_serialize( $this->price_component );
		if ( ! empty( $serialized ) ) {
			$out['priceComponent'] = $serialized;
		}
		
		return $out;
	}
}
