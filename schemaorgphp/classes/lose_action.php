<?php

class LoseAction extends AchieveAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LoseAction';
	
	/**
	 * A sub property of participant. The winner of the action.
	 * see : https://schema.org/winner
	 * @var \Person|\Person[]
	 */
	public var $winner;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LoseAction'
		);
		
		$serialized = so_json_serialize( $this->winner );
		if ( ! empty( $serialized ) ) {
			$out['winner'] = $serialized;
		}
		
		return $out;
	}
}
