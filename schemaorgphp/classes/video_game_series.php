<?php

class VideoGameSeries extends CreativeWorkSeries implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'VideoGameSeries';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	 * see : https://schema.org/characterAttribute
	 * @var \Thing|\Thing[]
	 */
	public var $character_attribute;
	
	/**
	 * Cheat codes to the game.
	 * see : https://schema.org/cheatCode
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $cheat_code;
	
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
	 * An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	 * see : https://schema.org/gameItem
	 * @var \Thing|\Thing[]
	 */
	public var $game_item;
	
	/**
	 * Real or fictional location of the game (or part of game).
	 * see : https://schema.org/gameLocation
	 * @var \Place|\Place[]|\PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $game_location;
	
	/**
	 * The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	 * see : https://schema.org/gamePlatform
	 * @var string|string[]|\Thing|\Thing[]|string|string[]
	 */
	public var $game_platform;
	
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
	 * Indicate how many people can play this game (minimum, maximum, or range).
	 * see : https://schema.org/numberOfPlayers
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_players;
	
	/**
	 * The number of seasons in this series.
	 * see : https://schema.org/numberOfSeasons
	 * @var integer|integer[]
	 */
	public var $number_of_seasons;
	
	/**
	 * Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	 * see : https://schema.org/playMode
	 * @var \GamePlayMode|\GamePlayMode[]
	 */
	public var $play_mode;
	
	/**
	 * The production company or studio responsible for the item e.g. series, video game, episode etc.
	 * see : https://schema.org/productionCompany
	 * @var \Organization|\Organization[]
	 */
	public var $production_company;
	
	/**
	 * The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	 * see : https://schema.org/quest
	 * @var \Thing|\Thing[]
	 */
	public var $quest;
	
	/**
	 * The trailer of a movie or tv/radio series, season, episode, etc.
	 * see : https://schema.org/trailer
	 * @var \VideoObject|\VideoObject[]
	 */
	public var $trailer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'VideoGameSeries'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->character_attribute );
		if ( ! empty( $serialized ) ) {
			$out['characterAttribute'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cheat_code );
		if ( ! empty( $serialized ) ) {
			$out['cheatCode'] = $serialized;
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
		
		$serialized = so_json_serialize( $this->game_item );
		if ( ! empty( $serialized ) ) {
			$out['gameItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_location );
		if ( ! empty( $serialized ) ) {
			$out['gameLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_platform );
		if ( ! empty( $serialized ) ) {
			$out['gamePlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_episodes );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEpisodes'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_players );
		if ( ! empty( $serialized ) ) {
			$out['numberOfPlayers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_seasons );
		if ( ! empty( $serialized ) ) {
			$out['numberOfSeasons'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->play_mode );
		if ( ! empty( $serialized ) ) {
			$out['playMode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_company );
		if ( ! empty( $serialized ) ) {
			$out['productionCompany'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->quest );
		if ( ! empty( $serialized ) ) {
			$out['quest'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trailer );
		if ( ! empty( $serialized ) ) {
			$out['trailer'] = $serialized;
		}
		
		return $out;
	}
}
