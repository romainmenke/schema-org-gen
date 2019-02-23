<?php
namespace SchemaOrg;

// Integer see : https://schema.org/Integer
class Integer implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Integer';

	/**
	 * With properties from Number see : https://schema.org
	 */


	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Integer',
		);

		return $out;
	}
}
