<?php

class Map extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Map';
	
	/**
	 * Indicates the kind of Map, from the MapCategoryType Enumeration.
	 * see : https://schema.org/mapType
	 * @var \MapCategoryType|\MapCategoryType[]
	 */
	public var $map_type;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Map'
		);
		
		$serialized = so_json_serialize( $this->map_type );
		if ( ! empty( $serialized ) ) {
			$out['mapType'] = $serialized;
		}
		
		return $out;
	}
}
