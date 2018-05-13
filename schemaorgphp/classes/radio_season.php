<?php

class RadioSeason extends CreativeWorkSeason implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'RadioSeason';
	
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
	 * The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/endDate
	 * @var string|string[]|string|string[]
	 */
	public var $end_date;
	
	/**
	 * An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	 * see : https://schema.org/episode
	 * @var \Episode|\Episode[]
	 */
	public var $episode;
	
	/**
	 * The number of episodes in this season or series.
	 * see : https://schema.org/numberOfEpisodes
	 * @var integer|integer[]
	 */
	public var $number_of_episodes;
	
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
	 * Position of the season within an ordered group of seasons.
	 * see : https://schema.org/seasonNumber
	 * @var integer|integer[]|string|string[]
	 */
	public var $season_number;
	
	/**
	 * The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/startDate
	 * @var string|string[]|string|string[]
	 */
	public var $start_date;
	
	/**
	 * The trailer of a movie or tv/radio series, season, episode, etc.
	 * see : https://schema.org/trailer
	 * @var \VideoObject|\VideoObject[]
	 */
	public var $trailer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'RadioSeason'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->end_date );
		if ( ! empty( $serialized ) ) {
			$out['endDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->episode );
		if ( ! empty( $serialized ) ) {
			$out['episode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_episodes );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEpisodes'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_series );
		if ( ! empty( $serialized ) ) {
			$out['partOfSeries'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_company );
		if ( ! empty( $serialized ) ) {
			$out['productionCompany'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->season_number );
		if ( ! empty( $serialized ) ) {
			$out['seasonNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_date );
		if ( ! empty( $serialized ) ) {
			$out['startDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trailer );
		if ( ! empty( $serialized ) ) {
			$out['trailer'] = $serialized;
		}
		
		return $out;
	}
}
