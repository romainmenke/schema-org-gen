<?php

class VideoGame extends SoftwareApplication implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'VideoGame';
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * Cheat codes to the game.
	 * see : https://schema.org/cheatCode
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $cheat_code;
	
	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	 * see : https://schema.org/director
	 * @var \Person|\Person[]
	 */
	public var $director;
	
	/**
	 * The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	 * see : https://schema.org/gamePlatform
	 * @var string|string[]|\Thing|\Thing[]|string|string[]
	 */
	public var $game_platform;
	
	/**
	 * The server on which  it is possible to play the game. Inverse property: game (see: https://schema.org/game).
	 * see : https://schema.org/gameServer
	 * @var \GameServer|\GameServer[]
	 */
	public var $game_server;
	
	/**
	 * Links to tips, tactics, etc.
	 * see : https://schema.org/gameTip
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $game_tip;
	
	/**
	 * The composer of the soundtrack.
	 * see : https://schema.org/musicBy
	 * @var \MusicGroup|\MusicGroup[]|\Person|\Person[]
	 */
	public var $music_by;
	
	/**
	 * Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	 * see : https://schema.org/playMode
	 * @var \GamePlayMode|\GamePlayMode[]
	 */
	public var $play_mode;
	
	/**
	 * The trailer of a movie or tv/radio series, season, episode, etc.
	 * see : https://schema.org/trailer
	 * @var \VideoObject|\VideoObject[]
	 */
	public var $trailer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'VideoGame'
		);
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cheat_code );
		if ( ! empty( $serialized ) ) {
			$out['cheatCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_platform );
		if ( ! empty( $serialized ) ) {
			$out['gamePlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_server );
		if ( ! empty( $serialized ) ) {
			$out['gameServer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_tip );
		if ( ! empty( $serialized ) ) {
			$out['gameTip'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_by );
		if ( ! empty( $serialized ) ) {
			$out['musicBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->play_mode );
		if ( ! empty( $serialized ) ) {
			$out['playMode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trailer );
		if ( ! empty( $serialized ) ) {
			$out['trailer'] = $serialized;
		}
		
		return $out;
	}
}
