<?php

class Clip extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Clip';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * Position of the clip within an ordered group of clips.
	 * see : https://schema.org/clipNumber
	 * @var integer|integer[]|string|string[]
	 */
	public var $clip_number;
	
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
	 * The episode to which this clip belongs.
	 * see : https://schema.org/partOfEpisode
	 * @var \Episode|\Episode[]
	 */
	public var $part_of_episode;
	
	/**
	 * The season to which this episode belongs.
	 * see : https://schema.org/partOfSeason
	 * @var \CreativeWorkSeason|\CreativeWorkSeason[]
	 */
	public var $part_of_season;
	
	/**
	 * The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	 * see : https://schema.org/partOfSeries
	 * @var \CreativeWorkSeries|\CreativeWorkSeries[]
	 */
	public var $part_of_series;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Clip'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->clip_number );
		if ( ! empty( $serialized ) ) {
			$out['clipNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_episode );
		if ( ! empty( $serialized ) ) {
			$out['partOfEpisode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_season );
		if ( ! empty( $serialized ) ) {
			$out['partOfSeason'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_series );
		if ( ! empty( $serialized ) ) {
			$out['partOfSeries'] = $serialized;
		}
		
		return $out;
	}
}
