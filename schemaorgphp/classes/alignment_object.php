<?php

class AlignmentObject extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AlignmentObject';
	
	/**
	 * A category of alignment between the learning resource and the framework node. Recommended values include: &#39;assesses&#39;, &#39;teaches&#39;, &#39;requires&#39;, &#39;textComplexity&#39;, &#39;readingLevel&#39;, &#39;educationalSubject&#39;, and &#39;educationalLevel&#39;.
	 * see : https://schema.org/alignmentType
	 * @var string|string[]
	 */
	public var $alignment_type;
	
	/**
	 * The framework to which the resource being described is aligned.
	 * see : https://schema.org/educationalFramework
	 * @var string|string[]
	 */
	public var $educational_framework;
	
	/**
	 * The description of a node in an established educational framework.
	 * see : https://schema.org/targetDescription
	 * @var string|string[]
	 */
	public var $target_description;
	
	/**
	 * The name of a node in an established educational framework.
	 * see : https://schema.org/targetName
	 * @var string|string[]
	 */
	public var $target_name;
	
	/**
	 * The URL of a node in an established educational framework.
	 * see : https://schema.org/targetUrl
	 * @var string|string[]
	 */
	public var $target_url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AlignmentObject'
		);
		
		$serialized = so_json_serialize( $this->alignment_type );
		if ( ! empty( $serialized ) ) {
			$out['alignmentType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->educational_framework );
		if ( ! empty( $serialized ) ) {
			$out['educationalFramework'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target_description );
		if ( ! empty( $serialized ) ) {
			$out['targetDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target_name );
		if ( ! empty( $serialized ) ) {
			$out['targetName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target_url );
		if ( ! empty( $serialized ) ) {
			$out['targetUrl'] = $serialized;
		}
		
		return $out;
	}
}
