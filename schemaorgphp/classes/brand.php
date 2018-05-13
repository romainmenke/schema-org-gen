<?php

class Brand extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Brand';
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Brand'
		);
		
		return $out;
	}
}
