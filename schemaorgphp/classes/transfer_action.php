<?php

class TransferAction extends Action implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TransferAction';
	
	/**
	 * A sub property of location. The original location of the object or the agent before the action.
	 * see : https://schema.org/fromLocation
	 * @var \Place|\Place[]
	 */
	public var $from_location;
	
	/**
	 * A sub property of location. The final location of the object or the agent after the action.
	 * see : https://schema.org/toLocation
	 * @var \Place|\Place[]
	 */
	public var $to_location;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TransferAction'
		);
		
		$serialized = so_json_serialize( $this->from_location );
		if ( ! empty( $serialized ) ) {
			$out['fromLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->to_location );
		if ( ! empty( $serialized ) ) {
			$out['toLocation'] = $serialized;
		}
		
		return $out;
	}
}
