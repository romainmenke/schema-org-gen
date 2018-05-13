<?php

class GameServer extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GameServer';
	
	/**
	 * Video game which is played on this server. Inverse property: gameServer (see: https://schema.org/gameServer).
	 * see : https://schema.org/game
	 * @var \VideoGame|\VideoGame[]
	 */
	public var $game;
	
	/**
	 * Number of players on the server.
	 * see : https://schema.org/playersOnline
	 * @var integer|integer[]
	 */
	public var $players_online;
	
	/**
	 * Status of a game server.
	 * see : https://schema.org/serverStatus
	 * @var \GameServerStatus|\GameServerStatus[]
	 */
	public var $server_status;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GameServer'
		);
		
		$serialized = so_json_serialize( $this->game );
		if ( ! empty( $serialized ) ) {
			$out['game'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->players_online );
		if ( ! empty( $serialized ) ) {
			$out['playersOnline'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->server_status );
		if ( ! empty( $serialized ) ) {
			$out['serverStatus'] = $serialized;
		}
		
		return $out;
	}
}
