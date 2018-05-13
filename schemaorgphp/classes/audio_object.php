<?php

class AudioObject extends MediaObject implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AudioObject';
	
	/**
	 * If this MediaObject is an AudioObject or VideoObject, the transcript of that object.
	 * see : https://schema.org/transcript
	 * @var string|string[]
	 */
	public var $transcript;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AudioObject'
		);
		
		$serialized = so_json_serialize( $this->transcript );
		if ( ! empty( $serialized ) ) {
			$out['transcript'] = $serialized;
		}
		
		return $out;
	}
}
