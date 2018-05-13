<?php

class TVEpisode extends Episode implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TVEpisode';
	
	/**
	 * The country of the principal offices of the production company or individual responsible for the movie or program.
	 * see : https://schema.org/countryOfOrigin
	 * @var \Country|\Country[]
	 */
	public var $country_of_origin;
	
	/**
	 * Languages in which subtitles/captions are available, in IETF BCP 47 standard format (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	 * see : https://schema.org/subtitleLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $subtitle_language;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TVEpisode'
		);
		
		$serialized = so_json_serialize( $this->country_of_origin );
		if ( ! empty( $serialized ) ) {
			$out['countryOfOrigin'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->subtitle_language );
		if ( ! empty( $serialized ) ) {
			$out['subtitleLanguage'] = $serialized;
		}
		
		return $out;
	}
}
