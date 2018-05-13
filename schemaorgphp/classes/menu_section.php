<?php

class MenuSection extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MenuSection';
	
	/**
	 * A food or drink item contained in a menu or menu section.
	 * see : https://schema.org/hasMenuItem
	 * @var \MenuItem|\MenuItem[]
	 */
	public var $has_menu_item;
	
	/**
	 * A subgrouping of the menu (by dishes, course, serving time period, etc.).
	 * see : https://schema.org/hasMenuSection
	 * @var \MenuSection|\MenuSection[]
	 */
	public var $has_menu_section;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MenuSection'
		);
		
		$serialized = so_json_serialize( $this->has_menu_item );
		if ( ! empty( $serialized ) ) {
			$out['hasMenuItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_menu_section );
		if ( ! empty( $serialized ) ) {
			$out['hasMenuSection'] = $serialized;
		}
		
		return $out;
	}
}
