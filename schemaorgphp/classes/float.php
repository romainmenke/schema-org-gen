<?php
namespace SchemaOrg;

// Float see : https://schema.org/Float
class Float implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Float';

	/**
	 * With properties from Number see : https://schema.org
	 */


	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Float',
		);

		return $out;
	}
}
