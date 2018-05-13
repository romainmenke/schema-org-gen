<?php

class MusicGroup extends PerformingGroup implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicGroup';
	
	/**
	 * A music album. Supersedes albums (see: https://schema.org/albums).
	 * see : https://schema.org/album
	 * @var \MusicAlbum|\MusicAlbum[]
	 */
	public var $album;
	
	/**
	 * Genre of the creative work, broadcast channel or group.
	 * see : https://schema.org/genre
	 * @var string|string[]|string|string[]
	 */
	public var $genre;
	
	/**
	 * A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	 * see : https://schema.org/track
	 * @var \ItemList|\ItemList[]|\MusicRecording|\MusicRecording[]
	 */
	public var $track;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicGroup'
		);
		
		$serialized = so_json_serialize( $this->album );
		if ( ! empty( $serialized ) ) {
			$out['album'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->genre );
		if ( ! empty( $serialized ) ) {
			$out['genre'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->track );
		if ( ! empty( $serialized ) ) {
			$out['track'] = $serialized;
		}
		
		return $out;
	}
}
