<?php

class Episode extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Episode';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	 * see : https://schema.org/director
	 * @var \Person|\Person[]
	 */
	public var $director;
	
	/**
	 * Position of the episode within an ordered group of episodes.
	 * see : https://schema.org/episodeNumber
	 * @var integer|integer[]|string|string[]
	 */
	public var $episode_number;
	
	/**
	 * The composer of the soundtrack.
	 * see : https://schema.org/musicBy
	 * @var \MusicGroup|\MusicGroup[]|\Person|\Person[]
	 */
	public var $music_by;
	
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
	
	/**
	 * The production company or studio responsible for the item e.g. series, video game, episode etc.
	 * see : https://schema.org/productionCompany
	 * @var \Organization|\Organization[]
	 */
	public var $production_company;
	
	/**
	 * The trailer of a movie or tv/radio series, season, episode, etc.
	 * see : https://schema.org/trailer
	 * @var \VideoObject|\VideoObject[]
	 */
	public var $trailer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Episode'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->episode_number );
		if ( ! empty( $serialized ) ) {
			$out['episodeNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_season );
		if ( ! empty( $serialized ) ) {
			$out['partOfSeason'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_series );
		if ( ! empty( $serialized ) ) {
			$out['partOfSeries'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_company );
		if ( ! empty( $serialized ) ) {
			$out['productionCompany'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trailer );
		if ( ! empty( $serialized ) ) {
			$out['trailer'] = $serialized;
		}
		
		return $out;
	}
}
