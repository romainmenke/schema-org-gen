<?php

class HowToSection extends ItemList implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToSection';
	
	/**
	 * The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	 * see : https://schema.org/steps
	 * @var \CreativeWork|\CreativeWork[]|\ItemList|\ItemList[]|string|string[]
	 */
	public var $steps;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToSection'
		);
		
		$serialized = so_json_serialize( $this->steps );
		if ( ! empty( $serialized ) ) {
			$out['steps'] = $serialized;
		}
		
		return $out;
	}
}
