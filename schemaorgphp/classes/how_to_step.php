<?php

class HowToStep extends ItemList implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToStep';
	
	/**
	 * For itemListElement values, you can use simple strings (e.g. &quot;Peter&quot;, &quot;Paul&quot;, &quot;Mary&quot;), existing entities, or use ListItem.
	 * 
	 * Text values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.
	 * 
	 * Note: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a &#39;position&#39; property in such cases.
	 * see : https://schema.org/itemListElement
	 * @var \ListItem|\ListItem[]|string|string[]|\Thing|\Thing[]
	 */
	public var $item_list_element;
	
	/**
	 * Type of ordering (e.g. Ascending, Descending, Unordered).
	 * see : https://schema.org/itemListOrder
	 * @var \ItemListOrderType|\ItemListOrderType[]|string|string[]
	 */
	public var $item_list_order;
	
	/**
	 * The number of items in an ItemList. Note that some descriptions might not fully describe all items in a list (e.g., multi-page pagination); in such cases, the numberOfItems would be for the entire list.
	 * see : https://schema.org/numberOfItems
	 * @var integer|integer[]
	 */
	public var $number_of_items;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToStep'
		);
		
		$serialized = so_json_serialize( $this->item_list_element );
		if ( ! empty( $serialized ) ) {
			$out['itemListElement'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->item_list_order );
		if ( ! empty( $serialized ) ) {
			$out['itemListOrder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_items );
		if ( ! empty( $serialized ) ) {
			$out['numberOfItems'] = $serialized;
		}
		
		return $out;
	}
}
