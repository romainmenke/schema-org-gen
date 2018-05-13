<?php

class WebSite extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WebSite';
	
	/**
	 * The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	 * see : https://schema.org/issn
	 * @var string|string[]
	 */
	public var $issn;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WebSite'
		);
		
		$serialized = so_json_serialize( $this->issn );
		if ( ! empty( $serialized ) ) {
			$out['issn'] = $serialized;
		}
		
		return $out;
	}
}
