<?php

class Game extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Game';
	
	/**
	 * A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	 * see : https://schema.org/characterAttribute
	 * @var \Thing|\Thing[]
	 */
	public var $character_attribute;
	
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
	 * Indicate how many people can play this game (minimum, maximum, or range).
	 * see : https://schema.org/numberOfPlayers
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_players;
	
	/**
	 * The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	 * see : https://schema.org/quest
	 * @var \Thing|\Thing[]
	 */
	public var $quest;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Game'
		);
		
		$serialized = so_json_serialize( $this->character_attribute );
		if ( ! empty( $serialized ) ) {
			$out['characterAttribute'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_item );
		if ( ! empty( $serialized ) ) {
			$out['gameItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->game_location );
		if ( ! empty( $serialized ) ) {
			$out['gameLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_players );
		if ( ! empty( $serialized ) ) {
			$out['numberOfPlayers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->quest );
		if ( ! empty( $serialized ) ) {
			$out['quest'] = $serialized;
		}
		
		return $out;
	}
}
