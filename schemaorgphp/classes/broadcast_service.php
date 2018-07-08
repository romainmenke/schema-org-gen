<?php

class BroadcastService extends Service implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BroadcastService';
	
	/**
	 * The media network(s) whose content is broadcast on this station.
	 * see : https://schema.org/broadcastAffiliateOf
	 * @var \Organization|\Organization[]
	 */
	public var $broadcast_affiliate_of;
	
	/**
	 * The name displayed in the channel guide. For many US affiliates, it is the network name.
	 * see : https://schema.org/broadcastDisplayName
	 * @var string|string[]
	 */
	public var $broadcast_display_name;
	
	/**
	 * The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	 * see : https://pending.schema.org/broadcastFrequency
	 * @var \BroadcastFrequencySpecification|\BroadcastFrequencySpecification[]|string|string[]
	 */
	public var $broadcast_frequency;
	
	/**
	 * The timezone in ISO 8601 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts
	 * see : https://schema.org/broadcastTimezone
	 * @var string|string[]
	 */
	public var $broadcast_timezone;
	
	/**
	 * The organization owning or operating the broadcast service.
	 * see : https://schema.org/broadcaster
	 * @var \Organization|\Organization[]
	 */
	public var $broadcaster;
	
	/**
	 * A broadcast channel of a broadcast service. Inverse property: providesBroadcastService (see: https://schema.org/providesBroadcastService).
	 * see : https://pending.schema.org/hasBroadcastChannel
	 * @var \BroadcastChannel|\BroadcastChannel[]
	 */
	public var $has_broadcast_channel;
	
	/**
	 * A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	 * see : https://schema.org/parentService
	 * @var \BroadcastService|\BroadcastService[]
	 */
	public var $parent_service;
	
	/**
	 * The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	 * see : https://schema.org/videoFormat
	 * @var string|string[]
	 */
	public var $video_format;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BroadcastService'
		);
		
		$serialized = so_json_serialize( $this->broadcast_affiliate_of );
		if ( ! empty( $serialized ) ) {
			$out['broadcastAffiliateOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcast_display_name );
		if ( ! empty( $serialized ) ) {
			$out['broadcastDisplayName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcast_frequency );
		if ( ! empty( $serialized ) ) {
			$out['broadcastFrequency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcast_timezone );
		if ( ! empty( $serialized ) ) {
			$out['broadcastTimezone'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcaster );
		if ( ! empty( $serialized ) ) {
			$out['broadcaster'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_broadcast_channel );
		if ( ! empty( $serialized ) ) {
			$out['hasBroadcastChannel'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->parent_service );
		if ( ! empty( $serialized ) ) {
			$out['parentService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->video_format );
		if ( ! empty( $serialized ) ) {
			$out['videoFormat'] = $serialized;
		}
		
		return $out;
	}
}
