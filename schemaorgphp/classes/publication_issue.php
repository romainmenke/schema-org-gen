<?php

class PublicationIssue extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PublicationIssue';
	
	/**
	 * Identifies the issue of publication; for example, &quot;iii&quot; or &quot;2&quot;.
	 * see : https://schema.org/issueNumber
	 * @var integer|integer[]|string|string[]
	 */
	public var $issue_number;
	
	/**
	 * The page on which the work ends; for example &quot;138&quot; or &quot;xvi&quot;.
	 * see : https://schema.org/pageEnd
	 * @var integer|integer[]|string|string[]
	 */
	public var $page_end;
	
	/**
	 * The page on which the work starts; for example &quot;135&quot; or &quot;xiii&quot;.
	 * see : https://schema.org/pageStart
	 * @var integer|integer[]|string|string[]
	 */
	public var $page_start;
	
	/**
	 * Any description of pages that is not separated into pageStart and pageEnd; for example, &quot;1-6, 9, 55&quot; or &quot;10-12, 46-49&quot;.
	 * see : https://schema.org/pagination
	 * @var string|string[]
	 */
	public var $pagination;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PublicationIssue'
		);
		
		$serialized = so_json_serialize( $this->issue_number );
		if ( ! empty( $serialized ) ) {
			$out['issueNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->page_end );
		if ( ! empty( $serialized ) ) {
			$out['pageEnd'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->page_start );
		if ( ! empty( $serialized ) ) {
			$out['pageStart'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->pagination );
		if ( ! empty( $serialized ) ) {
			$out['pagination'] = $serialized;
		}
		
		return $out;
	}
}
