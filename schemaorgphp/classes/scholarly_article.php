<?php

class ScholarlyArticle extends Article implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ScholarlyArticle';
	
	/**
	 * The actual body of the article.
	 * see : https://schema.org/articleBody
	 * @var string|string[]
	 */
	public var $article_body;
	
	/**
	 * Articles may belong to one or more &#39;sections&#39; in a magazine or newspaper, such as Sports, Lifestyle, etc.
	 * see : https://schema.org/articleSection
	 * @var string|string[]
	 */
	public var $article_section;
	
	/**
	 * For an Article (see: https://schema.org/Article), typically a NewsArticle (see: https://schema.org/NewsArticle), the backstory property provides a textual summary giving a brief explanation of why and how an article was created. In a journalistic setting this could include information about reporting process, methods, interviews, data sources, etc.
	 * see : https://pending.schema.org/backstory
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $backstory;
	
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
	 * see : https://pending.schema.org/speakable
	 * @var \SpeakableSpecification|\SpeakableSpecification[]|string|string[]
	 */
	public var $speakable;
	
	/**
	 * The number of words in the text of the Article.
	 * see : https://schema.org/wordCount
	 * @var integer|integer[]
	 */
	public var $word_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ScholarlyArticle'
		);
		
		$serialized = so_json_serialize( $this->article_body );
		if ( ! empty( $serialized ) ) {
			$out['articleBody'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->article_section );
		if ( ! empty( $serialized ) ) {
			$out['articleSection'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->backstory );
		if ( ! empty( $serialized ) ) {
			$out['backstory'] = $serialized;
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
		
		$serialized = so_json_serialize( $this->speakable );
		if ( ! empty( $serialized ) ) {
			$out['speakable'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->word_count );
		if ( ! empty( $serialized ) ) {
			$out['wordCount'] = $serialized;
		}
		
		return $out;
	}
}
