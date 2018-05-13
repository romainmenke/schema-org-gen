<?php

class VoteAction extends ChooseAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'VoteAction';
	
	/**
	 * A sub property of object. The candidate subject of this action.
	 * see : https://schema.org/candidate
	 * @var \Person|\Person[]
	 */
	public var $candidate;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'VoteAction'
		);
		
		$serialized = so_json_serialize( $this->candidate );
		if ( ! empty( $serialized ) ) {
			$out['candidate'] = $serialized;
		}
		
		return $out;
	}
}
