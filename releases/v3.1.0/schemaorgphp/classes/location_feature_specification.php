<?php
namespace SchemaOrg;

// LocationFeatureSpecification see : https://schema.org/LocationFeatureSpecification
class LocationFeatureSpecification implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'LocationFeatureSpecification';

	/**
	 * With properties from PropertyValue see : https://schema.org/PropertyValue
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
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 *
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $hours_available;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The upper value of some characteristic or property.
	 * see : https://schema.org/maxValue
	 *
	 * @var float | float[]
	 */
	public $max_value;

	/**
	 * The lower value of some characteristic or property.
	 * see : https://schema.org/minValue
	 *
	 * @var float | float[]
	 */
	public $min_value;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
	 * (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)
	 * a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	 * Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
	 * see : https://schema.org/propertyID
	 *
	 * @var string | string[]
	 */
	public $property_id;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	 * see : https://schema.org/unitCode
	 *
	 * @var string | string[]
	 */
	public $unit_code;

	/**
	 * A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	 * &lt;a href=&#39;unitCode&#39;&gt;unitCode&lt;/a&gt;.
	 * see : https://schema.org/unitText
	 *
	 * @var string | string[]
	 */
	public $unit_text;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 *
	 * @var string | string[]
	 */
	public $valid_from;

	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 *
	 * @var string | string[]
	 */
	public $valid_through;

	/**
	 * The value of the quantitative value or property value node.\n\n* For [[QuantitativeValue]] and [[MonetaryAmount]], the recommended type for values is &#39;Number&#39;.\n* For [[PropertyValue]], it can be &#39;Text;&#39;, &#39;Number&#39;, &#39;Boolean&#39;, or &#39;StructuredValue&#39;.
	 * see : https://schema.org/value
	 *
	 * @var float | float[] | string | string[] | boolean | boolean[] | \StructuredValue | \StructuredValue[]
	 */
	public $value;

	/**
	 * A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	 * see : https://schema.org/valueReference
	 *
	 * @var \Enumeration | \Enumeration[] | \StructuredValue | \StructuredValue[] | \PropertyValue | \PropertyValue[] | \QualitativeValue | \QualitativeValue[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $value_reference;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'LocationFeatureSpecification',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->max_value );
		if ( ! empty( $serialized ) ) {
			$out['maxValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->min_value );
		if ( ! empty( $serialized ) ) {
			$out['minValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->property_id );
		if ( ! empty( $serialized ) ) {
			$out['propertyID'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->unit_code );
		if ( ! empty( $serialized ) ) {
			$out['unitCode'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->unit_text );
		if ( ! empty( $serialized ) ) {
			$out['unitText'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value );
		if ( ! empty( $serialized ) ) {
			$out['value'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_reference );
		if ( ! empty( $serialized ) ) {
			$out['valueReference'] = $serialized;
		}

		return $out;
	}
}
