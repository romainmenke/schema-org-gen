<?php

class HowToTip extends ListItem implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToTip';
	
	/**
	 * An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	 * see : https://schema.org/item
	 * @var \Thing|\Thing[]
	 */
	public var $item;
	
	/**
	 * A link to the ListItem that follows the current one.
	 * see : https://schema.org/nextItem
	 * @var \ListItem|\ListItem[]
	 */
	public var $next_item;
	
	/**
	 * The position of an item in a series or sequence of items.
	 * see : https://schema.org/position
	 * @var integer|integer[]|string|string[]
	 */
	public var $position;
	
	/**
	 * A link to the ListItem that preceeds the current one.
	 * see : https://schema.org/previousItem
	 * @var \ListItem|\ListItem[]
	 */
	public var $previous_item;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToTip'
		);
		
		$serialized = so_json_serialize( $this->item );
		if ( ! empty( $serialized ) ) {
			$out['item'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->next_item );
		if ( ! empty( $serialized ) ) {
			$out['nextItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->position );
		if ( ! empty( $serialized ) ) {
			$out['position'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->previous_item );
		if ( ! empty( $serialized ) ) {
			$out['previousItem'] = $serialized;
		}
		
		return $out;
	}
}
