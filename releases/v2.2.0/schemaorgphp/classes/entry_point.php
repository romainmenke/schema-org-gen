<?php
namespace SchemaOrg;

// EntryPoint see : https://schema.org/EntryPoint
class EntryPoint implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'EntryPoint';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An application that can complete the request.
	 * see : https://schema.org/actionApplication
	 *
	 * @var \SoftwareApplication | \SoftwareApplication[]
	 */
	public $action_application;

	/**
	 * The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	 * see : https://schema.org/actionPlatform
	 *
	 * @var \
	   | \
	  [] | \
	 | \
	[]
	 */
	public $action_platform;

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
	 * An application that can complete the request.
	 * see : https://schema.org/application
	 *
	 * @var \SoftwareApplication | \SoftwareApplication[]
	 */
	public $application;

	/**
	 * The supported content type(s) for an EntryPoint response.
	 * see : https://schema.org/contentType
	 *
	 * @var string | string[]
	 */
	public $content_type;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The supported encoding type(s) for an EntryPoint request.
	 * see : https://schema.org/encodingType
	 *
	 * @var string | string[]
	 */
	public $encoding_type;

	/**
	 * An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	 * see : https://schema.org/httpMethod
	 *
	 * @var string | string[]
	 */
	public $http_method;

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

	/**
	 * A url template (RFC6570) that will be used to construct the target of the execution of the action.
	 * see : https://schema.org/urlTemplate
	 *
	 * @var string | string[]
	 */
	public $url_template;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'EntryPoint',
		);

		$serialized = \SchemaOrg\json_serialize( $this->action_application );
		if ( ! empty( $serialized ) ) {
			$out['actionApplication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->action_platform );
		if ( ! empty( $serialized ) ) {
			$out['actionPlatform'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->application );
		if ( ! empty( $serialized ) ) {
			$out['application'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->content_type );
		if ( ! empty( $serialized ) ) {
			$out['contentType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->encoding_type );
		if ( ! empty( $serialized ) ) {
			$out['encodingType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->http_method );
		if ( ! empty( $serialized ) ) {
			$out['httpMethod'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url_template );
		if ( ! empty( $serialized ) ) {
			$out['urlTemplate'] = $serialized;
		}

		return $out;
	}
}
