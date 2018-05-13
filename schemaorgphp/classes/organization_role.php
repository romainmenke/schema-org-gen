<?php

class OrganizationRole extends Role implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OrganizationRole';
	
	/**
	 * A number associated with a role in an organization, for example, the number on an athlete&#39;s jersey.
	 * see : https://schema.org/numberedPosition
	 * @var float|float[]
	 */
	public var $numbered_position;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OrganizationRole'
		);
		
		$serialized = so_json_serialize( $this->numbered_position );
		if ( ! empty( $serialized ) ) {
			$out['numberedPosition'] = $serialized;
		}
		
		return $out;
	}
}
