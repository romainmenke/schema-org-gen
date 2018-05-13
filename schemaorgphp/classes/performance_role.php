<?php

class PerformanceRole extends Role implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PerformanceRole';
	
	/**
	 * The name of a character played in some acting or performing role, i.e. in a PerformanceRole.
	 * see : https://schema.org/characterName
	 * @var string|string[]
	 */
	public var $character_name;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PerformanceRole'
		);
		
		$serialized = so_json_serialize( $this->character_name );
		if ( ! empty( $serialized ) ) {
			$out['characterName'] = $serialized;
		}
		
		return $out;
	}
}
