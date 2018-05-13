<?php

class Barcode extends ImageObject implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Barcode';
	
	/**
	 * The caption for this object.
	 * see : https://schema.org/caption
	 * @var string|string[]
	 */
	public var $caption;
	
	/**
	 * exif data for this object.
	 * see : https://schema.org/exifData
	 * @var \PropertyValue|\PropertyValue[]|string|string[]
	 */
	public var $exif_data;
	
	/**
	 * Indicates whether this image is representative of the content of the page.
	 * see : https://schema.org/representativeOfPage
	 * @var boolean|boolean[]
	 */
	public var $representative_of_page;
	
	/**
	 * Thumbnail image for an image or video.
	 * see : https://schema.org/thumbnail
	 * @var \ImageObject|\ImageObject[]
	 */
	public var $thumbnail;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Barcode'
		);
		
		$serialized = so_json_serialize( $this->caption );
		if ( ! empty( $serialized ) ) {
			$out['caption'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->exif_data );
		if ( ! empty( $serialized ) ) {
			$out['exifData'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->representative_of_page );
		if ( ! empty( $serialized ) ) {
			$out['representativeOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->thumbnail );
		if ( ! empty( $serialized ) ) {
			$out['thumbnail'] = $serialized;
		}
		
		return $out;
	}
}
