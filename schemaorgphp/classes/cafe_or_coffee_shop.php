<?php

class CafeOrCoffeeShop extends FoodEstablishment implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CafeOrCoffeeShop';
	
	/**
	 * Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings Yes or No.
	 * see : https://schema.org/acceptsReservations
	 * @var boolean|boolean[]|string|string[]|string|string[]
	 */
	public var $accepts_reservations;
	
	/**
	 * Either the actual menu as a structured representation, as text, or a URL of the menu. Supersedes menu (see: https://schema.org/menu).
	 * see : https://schema.org/hasMenu
	 * @var \Menu|\Menu[]|string|string[]|string|string[]
	 */
	public var $has_menu;
	
	/**
	 * The cuisine of the restaurant.
	 * see : https://schema.org/servesCuisine
	 * @var string|string[]
	 */
	public var $serves_cuisine;
	
	/**
	 * An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	 * see : https://schema.org/starRating
	 * @var \Rating|\Rating[]
	 */
	public var $star_rating;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CafeOrCoffeeShop'
		);
		
		$serialized = so_json_serialize( $this->accepts_reservations );
		if ( ! empty( $serialized ) ) {
			$out['acceptsReservations'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_menu );
		if ( ! empty( $serialized ) ) {
			$out['hasMenu'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->serves_cuisine );
		if ( ! empty( $serialized ) ) {
			$out['servesCuisine'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->star_rating );
		if ( ! empty( $serialized ) ) {
			$out['starRating'] = $serialized;
		}
		
		return $out;
	}
}
