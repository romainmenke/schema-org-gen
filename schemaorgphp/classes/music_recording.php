<?php

class MusicRecording extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicRecording';
	
	/**
	 * The artist that performed this album or recording.
	 * see : https://schema.org/byArtist
	 * @var \MusicGroup|\MusicGroup[]
	 */
	public var $by_artist;
	
	/**
	 * The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/duration
	 * @var \Duration|\Duration[]
	 */
	public var $duration;
	
	/**
	 * The album to which this recording belongs.
	 * see : https://schema.org/inAlbum
	 * @var \MusicAlbum|\MusicAlbum[]
	 */
	public var $in_album;
	
	/**
	 * The playlist to which this recording belongs.
	 * see : https://schema.org/inPlaylist
	 * @var \MusicPlaylist|\MusicPlaylist[]
	 */
	public var $in_playlist;
	
	/**
	 * The International Standard Recording Code for the recording.
	 * see : https://schema.org/isrcCode
	 * @var string|string[]
	 */
	public var $isrc_code;
	
	/**
	 * The composition this track is a recording of. Inverse property: recordedAs (see: https://schema.org/recordedAs).
	 * see : https://schema.org/recordingOf
	 * @var \MusicComposition|\MusicComposition[]
	 */
	public var $recording_of;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicRecording'
		);
		
		$serialized = so_json_serialize( $this->by_artist );
		if ( ! empty( $serialized ) ) {
			$out['byArtist'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duration );
		if ( ! empty( $serialized ) ) {
			$out['duration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->in_album );
		if ( ! empty( $serialized ) ) {
			$out['inAlbum'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->in_playlist );
		if ( ! empty( $serialized ) ) {
			$out['inPlaylist'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->isrc_code );
		if ( ! empty( $serialized ) ) {
			$out['isrcCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recording_of );
		if ( ! empty( $serialized ) ) {
			$out['recordingOf'] = $serialized;
		}
		
		return $out;
	}
}
