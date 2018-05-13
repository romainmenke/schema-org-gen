<?php

class PrependAction extends InsertAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PrependAction';
	
	/**
	 * A sub property of location. The final location of the object or the agent after the action.
	 * see : https://schema.org/toLocation
	 * @var \Place|\Place[]
	 */
	public var $to_location;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PrependAction'
		);
		
		$serialized = so_json_serialize( $this->to_location );
		if ( ! empty( $serialized ) ) {
			$out['toLocation'] = $serialized;
		}
		
		return $out;
	}
}
