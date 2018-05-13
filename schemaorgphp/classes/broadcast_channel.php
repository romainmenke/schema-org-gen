<?php

class BroadcastChannel extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BroadcastChannel';
	
	/**
	 * The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	 * see : https://schema.org/broadcastChannelId
	 * @var string|string[]
	 */
	public var $broadcast_channel_id;
	
	/**
	 * The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	 * see : http://pending.schema.org/broadcastFrequency
	 * @var \BroadcastFrequencySpecification|\BroadcastFrequencySpecification[]|string|string[]
	 */
	public var $broadcast_frequency;
	
	/**
	 * The type of service required to have access to the channel (e.g. Standard or Premium).
	 * see : https://schema.org/broadcastServiceTier
	 * @var string|string[]
	 */
	public var $broadcast_service_tier;
	
	/**
	 * Genre of the creative work, broadcast channel or group.
	 * see : https://schema.org/genre
	 * @var string|string[]|string|string[]
	 */
	public var $genre;
	
	/**
	 * The CableOrSatelliteService offering the channel.
	 * see : https://schema.org/inBroadcastLineup
	 * @var \CableOrSatelliteService|\CableOrSatelliteService[]
	 */
	public var $in_broadcast_lineup;
	
	/**
	 * The BroadcastService offered on this channel. Inverse property: hasBroadcastChannel (see: https://schema.orghttp://pending.schema.org/hasBroadcastChannel).
	 * see : https://schema.org/providesBroadcastService
	 * @var \BroadcastService|\BroadcastService[]
	 */
	public var $provides_broadcast_service;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BroadcastChannel'
		);
		
		$serialized = so_json_serialize( $this->broadcast_channel_id );
		if ( ! empty( $serialized ) ) {
			$out['broadcastChannelId'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcast_frequency );
		if ( ! empty( $serialized ) ) {
			$out['broadcastFrequency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broadcast_service_tier );
		if ( ! empty( $serialized ) ) {
			$out['broadcastServiceTier'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->genre );
		if ( ! empty( $serialized ) ) {
			$out['genre'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->in_broadcast_lineup );
		if ( ! empty( $serialized ) ) {
			$out['inBroadcastLineup'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provides_broadcast_service );
		if ( ! empty( $serialized ) ) {
			$out['providesBroadcastService'] = $serialized;
		}
		
		return $out;
	}
}
