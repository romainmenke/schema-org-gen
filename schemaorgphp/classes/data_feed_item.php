<?php

class DataFeedItem extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DataFeedItem';
	
	/**
	 * The date on which the CreativeWork was created or the item was added to a DataFeed.
	 * see : https://schema.org/dateCreated
	 * @var string|string[]|string|string[]
	 */
	public var $date_created;
	
	/**
	 * The datetime the item was removed from the DataFeed.
	 * see : https://schema.org/dateDeleted
	 * @var string|string[]
	 */
	public var $date_deleted;
	
	/**
	 * The date on which the CreativeWork was most recently modified or when the item&#39;s entry was modified within a DataFeed.
	 * see : https://schema.org/dateModified
	 * @var string|string[]|string|string[]
	 */
	public var $date_modified;
	
	/**
	 * An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	 * see : https://schema.org/item
	 * @var \Thing|\Thing[]
	 */
	public var $item;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DataFeedItem'
		);
		
		$serialized = so_json_serialize( $this->date_created );
		if ( ! empty( $serialized ) ) {
			$out['dateCreated'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_deleted );
		if ( ! empty( $serialized ) ) {
			$out['dateDeleted'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_modified );
		if ( ! empty( $serialized ) ) {
			$out['dateModified'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->item );
		if ( ! empty( $serialized ) ) {
			$out['item'] = $serialized;
		}
		
		return $out;
	}
}
