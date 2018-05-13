<?php

class WinAction extends AchieveAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WinAction';
	
	/**
	 * A sub property of participant. The loser of the action.
	 * see : https://schema.org/loser
	 * @var \Person|\Person[]
	 */
	public var $loser;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WinAction'
		);
		
		$serialized = so_json_serialize( $this->loser );
		if ( ! empty( $serialized ) ) {
			$out['loser'] = $serialized;
		}
		
		return $out;
	}
}
