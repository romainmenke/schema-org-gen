<?php
namespace SchemaOrg;

// URL see : https://schema.org/URL
class URL implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'URL';

	/**
	 * With properties from Text see : https://schema.org
	 */


	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'URL',
		);

		return $out;
	}
}
