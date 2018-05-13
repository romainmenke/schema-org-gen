<?php

class ConfirmAction extends InformAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ConfirmAction';
	
	/**
	 * Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	 * see : https://schema.org/event
	 * @var \Event|\Event[]
	 */
	public var $event;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ConfirmAction'
		);
		
		$serialized = so_json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}
		
		return $out;
	}
}
