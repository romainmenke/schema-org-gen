<?php

class SportsEvent extends Event implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SportsEvent';
	
	/**
	 * The away team in a sports event.
	 * see : https://schema.org/awayTeam
	 * @var \Person|\Person[]|\SportsTeam|\SportsTeam[]
	 */
	public var $away_team;
	
	/**
	 * A competitor in a sports event.
	 * see : https://schema.org/competitor
	 * @var \Person|\Person[]|\SportsTeam|\SportsTeam[]
	 */
	public var $competitor;
	
	/**
	 * The home team in a sports event.
	 * see : https://schema.org/homeTeam
	 * @var \Person|\Person[]|\SportsTeam|\SportsTeam[]
	 */
	public var $home_team;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SportsEvent'
		);
		
		$serialized = so_json_serialize( $this->away_team );
		if ( ! empty( $serialized ) ) {
			$out['awayTeam'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->competitor );
		if ( ! empty( $serialized ) ) {
			$out['competitor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->home_team );
		if ( ! empty( $serialized ) ) {
			$out['homeTeam'] = $serialized;
		}
		
		return $out;
	}
}
