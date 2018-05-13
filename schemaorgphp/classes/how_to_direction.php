<?php

class HowToDirection extends ListItem implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToDirection';
	
	/**
	 * A media object representing the circumstances after performing this direction.
	 * see : https://schema.org/afterMedia
	 * @var \MediaObject|\MediaObject[]
	 */
	public var $after_media;
	
	/**
	 * A media object representing the circumstances before performing this direction.
	 * see : https://schema.org/beforeMedia
	 * @var \MediaObject|\MediaObject[]
	 */
	public var $before_media;
	
	/**
	 * A media object representing the circumstances while performing this direction.
	 * see : https://schema.org/duringMedia
	 * @var \MediaObject|\MediaObject[]
	 */
	public var $during_media;
	
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
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToDirection'
		);
		
		$serialized = so_json_serialize( $this->after_media );
		if ( ! empty( $serialized ) ) {
			$out['afterMedia'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->before_media );
		if ( ! empty( $serialized ) ) {
			$out['beforeMedia'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->during_media );
		if ( ! empty( $serialized ) ) {
			$out['duringMedia'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->perform_time );
		if ( ! empty( $serialized ) ) {
			$out['performTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->prep_time );
		if ( ! empty( $serialized ) ) {
			$out['prepTime'] = $serialized;
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
		
		return $out;
	}
}
