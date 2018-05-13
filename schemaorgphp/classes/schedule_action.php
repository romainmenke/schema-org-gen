<?php

class ScheduleAction extends PlanAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ScheduleAction';
	
	/**
	 * The time the object is scheduled to.
	 * see : https://schema.org/scheduledTime
	 * @var string|string[]
	 */
	public var $scheduled_time;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ScheduleAction'
		);
		
		$serialized = so_json_serialize( $this->scheduled_time );
		if ( ! empty( $serialized ) ) {
			$out['scheduledTime'] = $serialized;
		}
		
		return $out;
	}
}
