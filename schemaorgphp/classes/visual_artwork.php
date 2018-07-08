<?php

class VisualArtwork extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'VisualArtwork';
	
	/**
	 * The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, &#39;artEdition&#39; refers to the total number of copies (in this example &quot;20&quot;).
	 * see : https://schema.org/artEdition
	 * @var integer|integer[]|string|string[]
	 */
	public var $art_edition;
	
	/**
	 * The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
	 * see : https://schema.org/artMedium
	 * @var string|string[]|string|string[]
	 */
	public var $art_medium;
	
	/**
	 * e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
	 * see : https://schema.org/artform
	 * @var string|string[]|string|string[]
	 */
	public var $artform;
	
	/**
	 * The primary artist for a work
	 *     in a medium other than pencils or digital line art--for example, if the
	 *     primary artwork is done in watercolors or digital paints.
	 * see : https://bib.schema.org/artist
	 * @var \Person|\Person[]
	 */
	public var $artist;
	
	/**
	 * The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc. Supersedes surface (see: https://schema.org/surface).
	 * see : https://schema.org/artworkSurface
	 * @var string|string[]|string|string[]
	 */
	public var $artwork_surface;
	
	/**
	 * The individual who adds color to inked drawings.
	 * see : https://bib.schema.org/colorist
	 * @var \Person|\Person[]
	 */
	public var $colorist;
	
	/**
	 * The depth of the item.
	 * see : https://schema.org/depth
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $depth;
	
	/**
	 * The height of the item.
	 * see : https://schema.org/height
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $height;
	
	/**
	 * The individual who traces over the pencil drawings in ink after pencils are complete.
	 * see : https://bib.schema.org/inker
	 * @var \Person|\Person[]
	 */
	public var $inker;
	
	/**
	 * The individual who adds lettering, including speech balloons and sound effects, to artwork.
	 * see : https://bib.schema.org/letterer
	 * @var \Person|\Person[]
	 */
	public var $letterer;
	
	/**
	 * The individual who draws the primary narrative artwork.
	 * see : https://bib.schema.org/penciler
	 * @var \Person|\Person[]
	 */
	public var $penciler;
	
	/**
	 * The width of the item.
	 * see : https://schema.org/width
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $width;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'VisualArtwork'
		);
		
		$serialized = so_json_serialize( $this->art_edition );
		if ( ! empty( $serialized ) ) {
			$out['artEdition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->art_medium );
		if ( ! empty( $serialized ) ) {
			$out['artMedium'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->artform );
		if ( ! empty( $serialized ) ) {
			$out['artform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->artist );
		if ( ! empty( $serialized ) ) {
			$out['artist'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->artwork_surface );
		if ( ! empty( $serialized ) ) {
			$out['artworkSurface'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->colorist );
		if ( ! empty( $serialized ) ) {
			$out['colorist'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->depth );
		if ( ! empty( $serialized ) ) {
			$out['depth'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->inker );
		if ( ! empty( $serialized ) ) {
			$out['inker'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->letterer );
		if ( ! empty( $serialized ) ) {
			$out['letterer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->penciler );
		if ( ! empty( $serialized ) ) {
			$out['penciler'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}
		
		return $out;
	}
}
