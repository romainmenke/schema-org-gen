<?php

class FollowAction extends InteractAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'FollowAction';
	
	/**
	 * A sub property of object. The person or organization being followed.
	 * see : https://schema.org/followee
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $followee;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'FollowAction'
		);
		
		$serialized = so_json_serialize( $this->followee );
		if ( ! empty( $serialized ) ) {
			$out['followee'] = $serialized;
		}
		
		return $out;
	}
}
