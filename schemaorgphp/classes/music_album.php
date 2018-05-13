<?php

class MusicAlbum extends MusicPlaylist implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicAlbum';
	
	/**
	 * Classification of the album by it&#39;s type of content: soundtrack, live album, studio album, etc.
	 * see : https://schema.org/albumProductionType
	 * @var \MusicAlbumProductionType|\MusicAlbumProductionType[]
	 */
	public var $album_production_type;
	
	/**
	 * A release of this album. Inverse property: releaseOf (see: https://schema.org/releaseOf).
	 * see : https://schema.org/albumRelease
	 * @var \MusicRelease|\MusicRelease[]
	 */
	public var $album_release;
	
	/**
	 * The kind of release which this album is: single, EP or album.
	 * see : https://schema.org/albumReleaseType
	 * @var \MusicAlbumReleaseType|\MusicAlbumReleaseType[]
	 */
	public var $album_release_type;
	
	/**
	 * The artist that performed this album or recording.
	 * see : https://schema.org/byArtist
	 * @var \MusicGroup|\MusicGroup[]
	 */
	public var $by_artist;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicAlbum'
		);
		
		$serialized = so_json_serialize( $this->album_production_type );
		if ( ! empty( $serialized ) ) {
			$out['albumProductionType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->album_release );
		if ( ! empty( $serialized ) ) {
			$out['albumRelease'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->album_release_type );
		if ( ! empty( $serialized ) ) {
			$out['albumReleaseType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->by_artist );
		if ( ! empty( $serialized ) ) {
			$out['byArtist'] = $serialized;
		}
		
		return $out;
	}
}
