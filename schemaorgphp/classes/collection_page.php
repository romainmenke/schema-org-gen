<?php

class CollectionPage extends WebPage implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CollectionPage';
	
	/**
	 * A set of links that can help a user understand and navigate a website hierarchy.
	 * see : https://schema.org/breadcrumb
	 * @var \BreadcrumbList|\BreadcrumbList[]|string|string[]
	 */
	public var $breadcrumb;
	
	/**
	 * Date on which the content on this web page was last reviewed for accuracy and/or completeness.
	 * see : https://schema.org/lastReviewed
	 * @var string|string[]
	 */
	public var $last_reviewed;
	
	/**
	 * Indicates if this web page element is the main subject of the page. Supersedes aspect (see: https://schema.orghttp://health-lifesci.schema.org/aspect).
	 * see : https://schema.org/mainContentOfPage
	 * @var \WebPageElement|\WebPageElement[]
	 */
	public var $main_content_of_page;
	
	/**
	 * Indicates the main image on the page.
	 * see : https://schema.org/primaryImageOfPage
	 * @var \ImageObject|\ImageObject[]
	 */
	public var $primary_image_of_page;
	
	/**
	 * A link related to this web page, for example to other related web pages.
	 * see : https://schema.org/relatedLink
	 * @var string|string[]
	 */
	public var $related_link;
	
	/**
	 * People or organizations that have reviewed the content on this web page for accuracy and/or completeness.
	 * see : https://schema.org/reviewedBy
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $reviewed_by;
	
	/**
	 * One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most. Supersedes significantLinks (see: https://schema.org/significantLinks).
	 * see : https://schema.org/significantLink
	 * @var string|string[]
	 */
	public var $significant_link;
	
	/**
	 * Indicates sections of a Web page that are particularly &#39;speakable&#39; in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the &#39;speakable&#39; property serves to indicate the parts most likely to be generally useful for speech.
	 * 
	 * The speakable property can be repeated an arbitrary number of times, with three kinds of possible &#39;content-locator&#39; values:
	 * 
	 * 1.) id-value URL references - uses id-value of an element in the page being annotated. The simplest use of speakable has (potentially relative) URL values, referencing identified sections of the document concerned.
	 * 
	 * 2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the cssSelector (see: https://schema.org/cssSelector) property.
	 * 
	 * 3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the xpath (see: https://schema.org/xpath) property.
	 * 
	 * For more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this
	 * we define a supporting type, SpeakableSpecification (see: https://schema.org/SpeakableSpecification)  which is defined to be a possible value of the speakable property.
	 * see : http://pending.schema.org/speakable
	 * @var \SpeakableSpecification|\SpeakableSpecification[]|string|string[]
	 */
	public var $speakable;
	
	/**
	 * One of the domain specialities to which this web page&#39;s content applies.
	 * see : https://schema.org/specialty
	 * @var \Specialty|\Specialty[]
	 */
	public var $specialty;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CollectionPage'
		);
		
		$serialized = so_json_serialize( $this->breadcrumb );
		if ( ! empty( $serialized ) ) {
			$out['breadcrumb'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->last_reviewed );
		if ( ! empty( $serialized ) ) {
			$out['lastReviewed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->main_content_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainContentOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->primary_image_of_page );
		if ( ! empty( $serialized ) ) {
			$out['primaryImageOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->related_link );
		if ( ! empty( $serialized ) ) {
			$out['relatedLink'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reviewed_by );
		if ( ! empty( $serialized ) ) {
			$out['reviewedBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->significant_link );
		if ( ! empty( $serialized ) ) {
			$out['significantLink'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->speakable );
		if ( ! empty( $serialized ) ) {
			$out['speakable'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->specialty );
		if ( ! empty( $serialized ) ) {
			$out['specialty'] = $serialized;
		}
		
		return $out;
	}
}
