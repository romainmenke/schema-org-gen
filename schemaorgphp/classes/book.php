<?php

class Book extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Book';
	
	/**
	 * Indicates whether the book is an abridged edition.
	 * see : http://bib.schema.org/abridged
	 * @var boolean|boolean[]
	 */
	public var $abridged;
	
	/**
	 * The edition of the book.
	 * see : https://schema.org/bookEdition
	 * @var string|string[]
	 */
	public var $book_edition;
	
	/**
	 * The format of the book.
	 * see : https://schema.org/bookFormat
	 * @var \BookFormatType|\BookFormatType[]
	 */
	public var $book_format;
	
	/**
	 * The illustrator of the book.
	 * see : https://schema.org/illustrator
	 * @var \Person|\Person[]
	 */
	public var $illustrator;
	
	/**
	 * The ISBN of the book.
	 * see : https://schema.org/isbn
	 * @var string|string[]
	 */
	public var $isbn;
	
	/**
	 * The number of pages in the book.
	 * see : https://schema.org/numberOfPages
	 * @var integer|integer[]
	 */
	public var $number_of_pages;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Book'
		);
		
		$serialized = so_json_serialize( $this->abridged );
		if ( ! empty( $serialized ) ) {
			$out['abridged'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->book_edition );
		if ( ! empty( $serialized ) ) {
			$out['bookEdition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->book_format );
		if ( ! empty( $serialized ) ) {
			$out['bookFormat'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->illustrator );
		if ( ! empty( $serialized ) ) {
			$out['illustrator'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->isbn );
		if ( ! empty( $serialized ) ) {
			$out['isbn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_pages );
		if ( ! empty( $serialized ) ) {
			$out['numberOfPages'] = $serialized;
		}
		
		return $out;
	}
}
