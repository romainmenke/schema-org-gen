<?php

class ReserveAction extends PlanAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReserveAction';
	
	/**
	 * The time the object is scheduled to.
	 * see : https://schema.org/scheduledTime
	 * @var string|string[]
	 */
	public var $scheduled_time;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReserveAction'
		);
		
		$serialized = so_json_serialize( $this->scheduled_time );
		if ( ! empty( $serialized ) ) {
			$out['scheduledTime'] = $serialized;
		}
		
		return $out;
	}
}
