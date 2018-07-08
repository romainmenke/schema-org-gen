<?php

class DeliveryMethod extends Enumeration implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DeliveryMethod';
	
	/**
	 * Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	 * see : https://meta.schema.org/supersededBy
	 * @var \Class|\Class[]|\Enumeration|\Enumeration[]|\Property|\Property[]
	 */
	public var $superseded_by;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DeliveryMethod'
		);
		
		$serialized = so_json_serialize( $this->superseded_by );
		if ( ! empty( $serialized ) ) {
			$out['supersededBy'] = $serialized;
		}
		
		return $out;
	}
}
