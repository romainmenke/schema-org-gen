<?php

class ScreeningEvent extends Event implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ScreeningEvent';
	
	/**
	 * Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	 * see : https://schema.org/subtitleLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $subtitle_language;
	
	/**
	 * The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	 * see : https://schema.org/videoFormat
	 * @var string|string[]
	 */
	public var $video_format;
	
	/**
	 * The movie presented during this event.
	 * see : https://schema.org/workPresented
	 * @var \Movie|\Movie[]
	 */
	public var $work_presented;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ScreeningEvent'
		);
		
		$serialized = so_json_serialize( $this->subtitle_language );
		if ( ! empty( $serialized ) ) {
			$out['subtitleLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->video_format );
		if ( ! empty( $serialized ) ) {
			$out['videoFormat'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->work_presented );
		if ( ! empty( $serialized ) ) {
			$out['workPresented'] = $serialized;
		}
		
		return $out;
	}
}
