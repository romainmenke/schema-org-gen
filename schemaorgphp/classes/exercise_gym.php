<?php

class ExerciseGym extends SportsActivityLocation implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ExerciseGym';
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ExerciseGym'
		);
		
		return $out;
	}
}
