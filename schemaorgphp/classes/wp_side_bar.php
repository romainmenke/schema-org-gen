<?php

class WPSideBar extends WebPageElement implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WPSideBar';
	
	/**
	 * A CSS selector, e.g. of a SpeakableSpecification (see: https://schema.org/SpeakableSpecification) or WebPageElement (see: https://schema.org/WebPageElement). In the latter case, multiple matches within a page can constitute a single conceptual &quot;Web page element&quot;.
	 * see : https://pending.schema.org/cssSelector
	 * @var \CssSelectorType|\CssSelectorType[]
	 */
	public var $css_selector;
	
	/**
	 * An XPath, e.g. of a SpeakableSpecification (see: https://schema.org/SpeakableSpecification) or WebPageElement (see: https://schema.org/WebPageElement). In the latter case, multiple matches within a page can constitute a single conceptual &quot;Web page element&quot;.
	 * see : https://pending.schema.org/xpath
	 * @var \XPathType|\XPathType[]
	 */
	public var $xpath;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WPSideBar'
		);
		
		$serialized = so_json_serialize( $this->css_selector );
		if ( ! empty( $serialized ) ) {
			$out['cssSelector'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->xpath );
		if ( ! empty( $serialized ) ) {
			$out['xpath'] = $serialized;
		}
		
		return $out;
	}
}
