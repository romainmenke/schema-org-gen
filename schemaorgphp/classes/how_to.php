<?php

class HowTo extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowTo';
	
	/**
	 * The estimated cost of the supply or supplies consumed when performing instructions.
	 * see : https://schema.org/estimatedCost
	 * @var \MonetaryAmount|\MonetaryAmount[]|string|string[]
	 */
	public var $estimated_cost;
	
	/**
	 * The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/performTime
	 * @var \Duration|\Duration[]
	 */
	public var $perform_time;
	
	/**
	 * The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/prepTime
	 * @var \Duration|\Duration[]
	 */
	public var $prep_time;
	
	/**
	 * The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	 * see : https://schema.org/steps
	 * @var \CreativeWork|\CreativeWork[]|\ItemList|\ItemList[]|string|string[]
	 */
	public var $steps;
	
	/**
	 * A sub-property of instrument. A supply consumed when performing instructions or a direction.
	 * see : https://schema.org/supply
	 * @var \HowToSupply|\HowToSupply[]|string|string[]
	 */
	public var $supply;
	
	/**
	 * A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	 * see : https://schema.org/tool
	 * @var \HowToTool|\HowToTool[]|string|string[]
	 */
	public var $tool;
	
	/**
	 * The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/totalTime
	 * @var \Duration|\Duration[]
	 */
	public var $total_time;
	
	/**
	 * The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
	 * see : https://schema.org/yield
	 * @var \QuantitativeValue|\QuantitativeValue[]|string|string[]
	 */
	public var $yield;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowTo'
		);
		
		$serialized = so_json_serialize( $this->estimated_cost );
		if ( ! empty( $serialized ) ) {
			$out['estimatedCost'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->perform_time );
		if ( ! empty( $serialized ) ) {
			$out['performTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->prep_time );
		if ( ! empty( $serialized ) ) {
			$out['prepTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->steps );
		if ( ! empty( $serialized ) ) {
			$out['steps'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->supply );
		if ( ! empty( $serialized ) ) {
			$out['supply'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->tool );
		if ( ! empty( $serialized ) ) {
			$out['tool'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->total_time );
		if ( ! empty( $serialized ) ) {
			$out['totalTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->yield );
		if ( ! empty( $serialized ) ) {
			$out['yield'] = $serialized;
		}
		
		return $out;
	}
}
