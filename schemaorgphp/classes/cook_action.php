<?php

class CookAction extends CreateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CookAction';
	
	/**
	 * A sub property of location. The specific food establishment where the action occurred.
	 * see : https://schema.org/foodEstablishment
	 * @var \FoodEstablishment|\FoodEstablishment[]|\Place|\Place[]
	 */
	public var $food_establishment;
	
	/**
	 * A sub property of location. The specific food event where the action occurred.
	 * see : https://schema.org/foodEvent
	 * @var \FoodEvent|\FoodEvent[]
	 */
	public var $food_event;
	
	/**
	 * A sub property of instrument. The recipe/instructions used to perform the action.
	 * see : https://schema.org/recipe
	 * @var \Recipe|\Recipe[]
	 */
	public var $recipe;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CookAction'
		);
		
		$serialized = so_json_serialize( $this->food_establishment );
		if ( ! empty( $serialized ) ) {
			$out['foodEstablishment'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->food_event );
		if ( ! empty( $serialized ) ) {
			$out['foodEvent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe );
		if ( ! empty( $serialized ) ) {
			$out['recipe'] = $serialized;
		}
		
		return $out;
	}
}
