<?php

class TechArticle extends Article implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TechArticle';
	
	/**
	 * Prerequisites needed to fulfill steps in article.
	 * see : https://schema.org/dependencies
	 * @var string|string[]
	 */
	public var $dependencies;
	
	/**
	 * Proficiency needed for this content; expected values: &#39;Beginner&#39;, &#39;Expert&#39;.
	 * see : https://schema.org/proficiencyLevel
	 * @var string|string[]
	 */
	public var $proficiency_level;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TechArticle'
		);
		
		$serialized = so_json_serialize( $this->dependencies );
		if ( ! empty( $serialized ) ) {
			$out['dependencies'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->proficiency_level );
		if ( ! empty( $serialized ) ) {
			$out['proficiencyLevel'] = $serialized;
		}
		
		return $out;
	}
}
