<?php
namespace SchemaOrg;

// DataType see : https://schema.org/DataType
class DataType implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'DataType';

	/**
	 * With properties from Rdfs:Class see : https://schema.org
	 */


	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'DataType',
		);

		return $out;
	}
}
