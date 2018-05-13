<?php

class DataFeed extends Dataset implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DataFeed';
	
	/**
	 * An item within in a data feed. Data feeds may have many elements.
	 * see : https://schema.org/dataFeedElement
	 * @var \DataFeedItem|\DataFeedItem[]|string|string[]|\Thing|\Thing[]
	 */
	public var $data_feed_element;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DataFeed'
		);
		
		$serialized = so_json_serialize( $this->data_feed_element );
		if ( ! empty( $serialized ) ) {
			$out['dataFeedElement'] = $serialized;
		}
		
		return $out;
	}
}
