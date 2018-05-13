<?php

class NewsArticle extends Article implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'NewsArticle';
	
	/**
	 * A dateline (see: https://schema.orghttps://en.wikipedia.org/wiki/Dateline) is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.
	 * see : https://schema.org/dateline
	 * @var string|string[]
	 */
	public var $dateline;
	
	/**
	 * The number of the column in which the NewsArticle appears in the print edition.
	 * see : https://schema.org/printColumn
	 * @var string|string[]
	 */
	public var $print_column;
	
	/**
	 * The edition of the print product in which the NewsArticle appears.
	 * see : https://schema.org/printEdition
	 * @var string|string[]
	 */
	public var $print_edition;
	
	/**
	 * If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
	 * see : https://schema.org/printPage
	 * @var string|string[]
	 */
	public var $print_page;
	
	/**
	 * If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
	 * see : https://schema.org/printSection
	 * @var string|string[]
	 */
	public var $print_section;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'NewsArticle'
		);
		
		$serialized = so_json_serialize( $this->dateline );
		if ( ! empty( $serialized ) ) {
			$out['dateline'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->print_column );
		if ( ! empty( $serialized ) ) {
			$out['printColumn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->print_edition );
		if ( ! empty( $serialized ) ) {
			$out['printEdition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->print_page );
		if ( ! empty( $serialized ) ) {
			$out['printPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->print_section );
		if ( ! empty( $serialized ) ) {
			$out['printSection'] = $serialized;
		}
		
		return $out;
	}
}
