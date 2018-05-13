<?php

class BroadcastEvent extends PublicationEvent implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BroadcastEvent';
	
	/**
	 * The event being broadcast such as a sporting event or awards ceremony.
	 * see : https://schema.org/broadcastOfEvent
	 * @var \Event|\Event[]
	 */
	public var $broadcast_of_event;
	
	/**
	 * True is the broadcast is of a live event.
	 * see : https://schema.org/isLiveBroadcast
	 * @var boolean|boolean[]
	 */
	public var $is_live_broadcast;
	
	/**
	 * The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	 * see : https://schema.org/videoFormat
	 * @var string|string[]
	 */
	public var $video_format;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BroadcastEvent'
		);
		
		$serialized = so_json_serialize( $this->broadcast_of_event );
		if ( ! empty( $serialized ) ) {
			$out['broadcastOfEvent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_live_broadcast );
		if ( ! empty( $serialized ) ) {
			$out['isLiveBroadcast'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->video_format );
		if ( ! empty( $serialized ) ) {
			$out['videoFormat'] = $serialized;
		}
		
		return $out;
	}
}
