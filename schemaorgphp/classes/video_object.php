<?php

class VideoObject extends MediaObject implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'VideoObject';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * The caption for this object.
	 * see : https://schema.org/caption
	 * @var string|string[]
	 */
	public var $caption;
	
	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	 * see : https://schema.org/director
	 * @var \Person|\Person[]
	 */
	public var $director;
	
	/**
	 * The composer of the soundtrack.
	 * see : https://schema.org/musicBy
	 * @var \MusicGroup|\MusicGroup[]|\Person|\Person[]
	 */
	public var $music_by;
	
	/**
	 * Thumbnail image for an image or video.
	 * see : https://schema.org/thumbnail
	 * @var \ImageObject|\ImageObject[]
	 */
	public var $thumbnail;
	
	/**
	 * If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	 * see : https://schema.org/transcript
	 * @var string|string[]
	 */
	public var $transcript;
	
	/**
	 * The frame size of the video.
	 * see : https://schema.org/videoFrameSize
	 * @var string|string[]
	 */
	public var $video_frame_size;
	
	/**
	 * The quality of the video.
	 * see : https://schema.org/videoQuality
	 * @var string|string[]
	 */
	public var $video_quality;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'VideoObject'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->caption );
		if ( ! empty( $serialized ) ) {
			$out['caption'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->thumbnail );
		if ( ! empty( $serialized ) ) {
			$out['thumbnail'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->transcript );
		if ( ! empty( $serialized ) ) {
			$out['transcript'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->video_frame_size );
		if ( ! empty( $serialized ) ) {
			$out['videoFrameSize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->video_quality );
		if ( ! empty( $serialized ) ) {
			$out['videoQuality'] = $serialized;
		}
		
		return $out;
	}
}
