<?php

class MusicPlaylist extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicPlaylist';
	
	/**
	 * The number of tracks in this album or playlist.
	 * see : https://schema.org/numTracks
	 * @var integer|integer[]
	 */
	public var $num_tracks;
	
	/**
	 * A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording. Supersedes tracks (see: https://schema.org/tracks).
	 * see : https://schema.org/track
	 * @var \ItemList|\ItemList[]|\MusicRecording|\MusicRecording[]
	 */
	public var $track;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicPlaylist'
		);
		
		$serialized = so_json_serialize( $this->num_tracks );
		if ( ! empty( $serialized ) ) {
			$out['numTracks'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->track );
		if ( ! empty( $serialized ) ) {
			$out['track'] = $serialized;
		}
		
		return $out;
	}
}
