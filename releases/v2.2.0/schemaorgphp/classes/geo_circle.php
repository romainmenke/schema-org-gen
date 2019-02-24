<?php
namespace SchemaOrg;

// GeoCircle see : https://schema.org/GeoCircle
class GeoCircle implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'GeoCircle';

	/**
	 * With properties from GeoShape see : https://schema.org/GeoShape
	 */

	/**
	 * With properties from StructuredValue see : https://schema.org/StructuredValue
	 */

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 *
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $address;

	/**
	 * The country. For example, USA. You can also provide the two-letter &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_3166-1&#39;&gt;ISO 3166-1 alpha-2 country code&lt;/a&gt;.
	 * see : https://schema.org/addressCountry
	 *
	 * @var string | string[] | \Country | \Country[]
	 */
	public $address_country;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
	 * see : https://schema.org/box
	 *
	 * @var string | string[]
	 */
	public $box;

	/**
	 * A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
	 * see : https://schema.org/circle
	 *
	 * @var string | string[]
	 */
	public $circle;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The elevation of a location (&lt;a href=&quot;https://en.wikipedia.org/wiki/World_Geodetic_System&quot;&gt;WGS 84&lt;/a&gt;).
	 * see : https://schema.org/elevation
	 *
	 * @var float | float[] | string | string[]
	 */
	public $elevation;

	/**
	 * Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	 * see : https://schema.org/geoMidpoint
	 *
	 * @var \GeoCoordinates | \GeoCoordinates[]
	 */
	public $geo_midpoint;

	/**
	 * Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
	 * see : https://schema.org/geoRadius
	 *
	 * @var string | string[] | float | float[] | \Distance | \Distance[]
	 */
	public $geo_radius;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
	 * see : https://schema.org/line
	 *
	 * @var string | string[]
	 */
	public $line;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
	 * see : https://schema.org/polygon
	 *
	 * @var string | string[]
	 */
	public $polygon;

	/**
	 * The postal code. For example, 94043.
	 * see : https://schema.org/postalCode
	 *
	 * @var string | string[]
	 */
	public $postal_code;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'GeoCircle',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->address_country );
		if ( ! empty( $serialized ) ) {
			$out['addressCountry'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->box );
		if ( ! empty( $serialized ) ) {
			$out['box'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->circle );
		if ( ! empty( $serialized ) ) {
			$out['circle'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->elevation );
		if ( ! empty( $serialized ) ) {
			$out['elevation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->geo_midpoint );
		if ( ! empty( $serialized ) ) {
			$out['geoMidpoint'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->geo_radius );
		if ( ! empty( $serialized ) ) {
			$out['geoRadius'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->line );
		if ( ! empty( $serialized ) ) {
			$out['line'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->polygon );
		if ( ! empty( $serialized ) ) {
			$out['polygon'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->postal_code );
		if ( ! empty( $serialized ) ) {
			$out['postalCode'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
