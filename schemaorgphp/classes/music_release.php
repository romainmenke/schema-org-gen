<?php

class MusicRelease extends MusicPlaylist implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicRelease';
	
	/**
	 * The catalog number for the release.
	 * see : https://schema.org/catalogNumber
	 * @var string|string[]
	 */
	public var $catalog_number;
	
	/**
	 * The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to &quot;Stefani Germanotta Band&quot;, but by Lady Gaga.
	 * see : https://schema.org/creditedTo
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $credited_to;
	
	/**
	 * The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://pending.schema.org/duration
	 * @var \Duration|\Duration[]
	 */
	public var $duration;
	
	/**
	 * Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc.).
	 * see : https://schema.org/musicReleaseFormat
	 * @var \MusicReleaseFormatType|\MusicReleaseFormatType[]
	 */
	public var $music_release_format;
	
	/**
	 * The label that issued the release.
	 * see : https://schema.org/recordLabel
	 * @var \Organization|\Organization[]
	 */
	public var $record_label;
	
	/**
	 * The album this is a release of. Inverse property: albumRelease (see: https://schema.org/albumRelease).
	 * see : https://schema.org/releaseOf
	 * @var \MusicAlbum|\MusicAlbum[]
	 */
	public var $release_of;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicRelease'
		);
		
		$serialized = so_json_serialize( $this->catalog_number );
		if ( ! empty( $serialized ) ) {
			$out['catalogNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->credited_to );
		if ( ! empty( $serialized ) ) {
			$out['creditedTo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duration );
		if ( ! empty( $serialized ) ) {
			$out['duration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_release_format );
		if ( ! empty( $serialized ) ) {
			$out['musicReleaseFormat'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->record_label );
		if ( ! empty( $serialized ) ) {
			$out['recordLabel'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->release_of );
		if ( ! empty( $serialized ) ) {
			$out['releaseOf'] = $serialized;
		}
		
		return $out;
	}
}
