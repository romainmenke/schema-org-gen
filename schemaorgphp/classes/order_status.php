<?php

class OrderStatus extends Enumeration implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OrderStatus';
	
	/**
	 * Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	 * see : http://meta.schema.org/supersededBy
	 * @var \Class|\Class[]|\Enumeration|\Enumeration[]|\Property|\Property[]
	 */
	public var $superseded_by;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OrderStatus'
		);
		
		$serialized = so_json_serialize( $this->superseded_by );
		if ( ! empty( $serialized ) ) {
			$out['supersededBy'] = $serialized;
		}
		
		return $out;
	}
}
