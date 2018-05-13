<?php

class InsertAction extends AddAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'InsertAction';
	
	/**
	 * A sub property of location. The final location of the object or the agent after the action.
	 * see : https://schema.org/toLocation
	 * @var \Place|\Place[]
	 */
	public var $to_location;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'InsertAction'
		);
		
		$serialized = so_json_serialize( $this->to_location );
		if ( ! empty( $serialized ) ) {
			$out['toLocation'] = $serialized;
		}
		
		return $out;
	}
}
