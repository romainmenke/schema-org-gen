<?php

class GenderType extends Enumeration implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GenderType';
	
	/**
	 * Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	 * see : https://meta.schema.org/supersededBy
	 * @var \Class|\Class[]|\Enumeration|\Enumeration[]|\Property|\Property[]
	 */
	public var $superseded_by;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GenderType'
		);
		
		$serialized = so_json_serialize( $this->superseded_by );
		if ( ! empty( $serialized ) ) {
			$out['supersededBy'] = $serialized;
		}
		
		return $out;
	}
}
