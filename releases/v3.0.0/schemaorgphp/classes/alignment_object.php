<?php
namespace SchemaOrg;

// AlignmentObject see : https://schema.org/AlignmentObject
class AlignmentObject implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'AlignmentObject';

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
	 * A category of alignment between the learning resource and the framework node. Recommended values include: &#39;assesses&#39;, &#39;teaches&#39;, &#39;requires&#39;, &#39;textComplexity&#39;, &#39;readingLevel&#39;, &#39;educationalSubject&#39;, and &#39;educationLevel&#39;.
	 * see : https://schema.org/alignmentType
	 *
	 * @var string | string[]
	 */
	public $alignment_type;

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
	 * The framework to which the resource being described is aligned.
	 * see : https://schema.org/educationalFramework
	 *
	 * @var string | string[]
	 */
	public $educational_framework;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
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
	 * The description of a node in an established educational framework.
	 * see : https://schema.org/targetDescription
	 *
	 * @var string | string[]
	 */
	public $target_description;

	/**
	 * The name of a node in an established educational framework.
	 * see : https://schema.org/targetName
	 *
	 * @var string | string[]
	 */
	public $target_name;

	/**
	 * The URL of a node in an established educational framework.
	 * see : https://schema.org/targetUrl
	 *
	 * @var string | string[]
	 */
	public $target_url;

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
			'@type'    => 'AlignmentObject',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alignment_type );
		if ( ! empty( $serialized ) ) {
			$out['alignmentType'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->educational_framework );
		if ( ! empty( $serialized ) ) {
			$out['educationalFramework'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->target_description );
		if ( ! empty( $serialized ) ) {
			$out['targetDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->target_name );
		if ( ! empty( $serialized ) ) {
			$out['targetName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->target_url );
		if ( ! empty( $serialized ) ) {
			$out['targetUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
