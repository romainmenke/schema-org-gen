<?php

class MenuItem extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MenuItem';
	
	/**
	 * Additional menu item(s) such as a side dish of salad or side order of fries that can be added to this menu item. Additionally it can be a menu section containing allowed add-on menu items for this menu item.
	 * see : http://pending.schema.org/menuAddOn
	 * @var \MenuItem|\MenuItem[]|\MenuSection|\MenuSection[]
	 */
	public var $menu_add_on;
	
	/**
	 * Nutrition information about the recipe or menu item.
	 * see : https://schema.org/nutrition
	 * @var \NutritionInformation|\NutritionInformation[]
	 */
	public var $nutrition;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer|\Offer[]
	 */
	public var $offers;
	
	/**
	 * Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	 * see : https://schema.org/suitableForDiet
	 * @var \RestrictedDiet|\RestrictedDiet[]
	 */
	public var $suitable_for_diet;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MenuItem'
		);
		
		$serialized = so_json_serialize( $this->menu_add_on );
		if ( ! empty( $serialized ) ) {
			$out['menuAddOn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->nutrition );
		if ( ! empty( $serialized ) ) {
			$out['nutrition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suitable_for_diet );
		if ( ! empty( $serialized ) ) {
			$out['suitableForDiet'] = $serialized;
		}
		
		return $out;
	}
}
