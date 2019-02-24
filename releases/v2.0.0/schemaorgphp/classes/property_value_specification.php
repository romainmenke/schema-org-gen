<?php
namespace SchemaOrg;

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
class PropertyValueSpecification implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'PropertyValueSpecification';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
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
	 * The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	 * see : https://schema.org/defaultValue
	 *
	 * @var \Thing | \Thing[] | string | string[]
	 */
	public $default_value;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	 *       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	 *       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	 *       between the page and the primary entity.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       Related properties include sameAs, about, and url.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	 *       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	 *       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	 *       serves more to clarify which of several entities is the main one for that page.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	 *       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	 *       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	 *       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	 *       describes some other entity. For example, one web page may display a news article about a particular person.
	 *       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	 *       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	 *
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
	 * Whether multiple values are allowed for the property.  Default is false.
	 * see : https://schema.org/multipleValues
	 *
	 * @var boolean | boolean[]
	 */
	public $multiple_values;

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
	 * Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	 * see : https://schema.org/readonlyValue
	 *
	 * @var boolean | boolean[]
	 */
	public $readonly_value;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	 * see : https://schema.org/stepValue
	 *
	 * @var float | float[]
	 */
	public $step_value;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * Specifies the allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMaxLength
	 *
	 * @var float | float[]
	 */
	public $value_max_length;

	/**
	 * Specifies the minimum allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMinLength
	 *
	 * @var float | float[]
	 */
	public $value_min_length;

	/**
	 * Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML&#39;s input@name.
	 * see : https://schema.org/valueName
	 *
	 * @var string | string[]
	 */
	public $value_name;

	/**
	 * Specifies a regular expression for testing literal values according to the HTML spec.
	 * see : https://schema.org/valuePattern
	 *
	 * @var string | string[]
	 */
	public $value_pattern;

	/**
	 * Whether the property must be filled in to complete the action.  Default is false.
	 * see : https://schema.org/valueRequired
	 *
	 * @var boolean | boolean[]
	 */
	public $value_required;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'PropertyValueSpecification',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->default_value );
		if ( ! empty( $serialized ) ) {
			$out['defaultValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->multiple_values );
		if ( ! empty( $serialized ) ) {
			$out['multipleValues'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->readonly_value );
		if ( ! empty( $serialized ) ) {
			$out['readonlyValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->step_value );
		if ( ! empty( $serialized ) ) {
			$out['stepValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_max_length );
		if ( ! empty( $serialized ) ) {
			$out['valueMaxLength'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_min_length );
		if ( ! empty( $serialized ) ) {
			$out['valueMinLength'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_name );
		if ( ! empty( $serialized ) ) {
			$out['valueName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_pattern );
		if ( ! empty( $serialized ) ) {
			$out['valuePattern'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->value_required );
		if ( ! empty( $serialized ) ) {
			$out['valueRequired'] = $serialized;
		}

		return $out;
	}
}
