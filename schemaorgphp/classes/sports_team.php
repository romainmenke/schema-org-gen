<?php

class SportsTeam extends SportsOrganization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SportsTeam';
	
	/**
	 * A person that acts as performing member of a sports team; a player as opposed to a coach.
	 * see : https://schema.org/athlete
	 * @var \Person|\Person[]
	 */
	public var $athlete;
	
	/**
	 * A person that acts in a coaching role for a sports team.
	 * see : https://schema.org/coach
	 * @var \Person|\Person[]
	 */
	public var $coach;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SportsTeam'
		);
		
		$serialized = so_json_serialize( $this->athlete );
		if ( ! empty( $serialized ) ) {
			$out['athlete'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->coach );
		if ( ! empty( $serialized ) ) {
			$out['coach'] = $serialized;
		}
		
		return $out;
	}
}
