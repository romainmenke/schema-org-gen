<?php

class SearchAction extends Action implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SearchAction';
	
	/**
	 * A sub property of instrument. The query used on this action.
	 * see : https://schema.org/query
	 * @var string|string[]
	 */
	public var $query;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SearchAction'
		);
		
		$serialized = so_json_serialize( $this->query );
		if ( ! empty( $serialized ) ) {
			$out['query'] = $serialized;
		}
		
		return $out;
	}
}
