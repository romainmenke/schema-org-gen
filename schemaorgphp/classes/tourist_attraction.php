<?php

class TouristAttraction extends Place implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TouristAttraction';
	
	/**
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $available_language;
	
	/**
	 * Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	 * see : https://schema.org/touristType
	 * @var \Audience|\Audience[]|string|string[]
	 */
	public var $tourist_type;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TouristAttraction'
		);
		
		$serialized = so_json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->tourist_type );
		if ( ! empty( $serialized ) ) {
			$out['touristType'] = $serialized;
		}
		
		return $out;
	}
}
