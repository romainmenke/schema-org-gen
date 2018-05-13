<?php

class RadioSeries extends CreativeWorkSeries implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'RadioSeries';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	 * see : https://schema.org/containsSeason
	 * @var \CreativeWorkSeason|\CreativeWorkSeason[]
	 */
	public var $contains_season;
	
	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	 * see : https://schema.org/director
	 * @var \Person|\Person[]
	 */
	public var $director;
	
	/**
	 * An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	 * see : https://schema.org/episode
	 * @var \Episode|\Episode[]
	 */
	public var $episode;
	
	/**
	 * The composer of the soundtrack.
	 * see : https://schema.org/musicBy
	 * @var \MusicGroup|\MusicGroup[]|\Person|\Person[]
	 */
	public var $music_by;
	
	/**
	 * The number of episodes in this season or series.
	 * see : https://schema.org/numberOfEpisodes
	 * @var integer|integer[]
	 */
	public var $number_of_episodes;
	
	/**
	 * The number of seasons in this series.
	 * see : https://schema.org/numberOfSeasons
	 * @var integer|integer[]
	 */
	public var $number_of_seasons;
	
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
			'@type' => 'RadioSeries'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contains_season );
		if ( ! empty( $serialized ) ) {
			$out['containsSeason'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->episode );
		if ( ! empty( $serialized ) ) {
			$out['episode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_episodes );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEpisodes'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_seasons );
		if ( ! empty( $serialized ) ) {
			$out['numberOfSeasons'] = $serialized;
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
