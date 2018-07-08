<?php

class Rating extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Rating';
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Rating'
		);
		
		return $out;
	}
}
