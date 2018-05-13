<?php

class DeleteAction extends UpdateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DeleteAction';
	
	/**
	 * A sub property of object. The collection target of the action. Supersedes collection (see: https://schema.org/collection).
	 * see : https://schema.org/targetCollection
	 * @var \Thing|\Thing[]
	 */
	public var $target_collection;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DeleteAction'
		);
		
		$serialized = so_json_serialize( $this->target_collection );
		if ( ! empty( $serialized ) ) {
			$out['targetCollection'] = $serialized;
		}
		
		return $out;
	}
}
