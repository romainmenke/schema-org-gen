<?php
namespace SchemaOrg;

// NutritionInformation see : https://schema.org/NutritionInformation
class NutritionInformation implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'NutritionInformation';

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
	 * The number of calories.
	 * see : https://schema.org/calories
	 *
	 * @var \Energy | \Energy[]
	 */
	public $calories;

	/**
	 * The number of grams of carbohydrates.
	 * see : https://schema.org/carbohydrateContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $carbohydrate_content;

	/**
	 * The number of milligrams of cholesterol.
	 * see : https://schema.org/cholesterolContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $cholesterol_content;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The number of grams of fat.
	 * see : https://schema.org/fatContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $fat_content;

	/**
	 * The number of grams of fiber.
	 * see : https://schema.org/fiberContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $fiber_content;

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
	 * The number of grams of protein.
	 * see : https://schema.org/proteinContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $protein_content;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The number of grams of saturated fat.
	 * see : https://schema.org/saturatedFatContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $saturated_fat_content;

	/**
	 * The serving size, in terms of the number of volume or mass.
	 * see : https://schema.org/servingSize
	 *
	 * @var string | string[]
	 */
	public $serving_size;

	/**
	 * The number of milligrams of sodium.
	 * see : https://schema.org/sodiumContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $sodium_content;

	/**
	 * The number of grams of sugar.
	 * see : https://schema.org/sugarContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $sugar_content;

	/**
	 * The number of grams of trans fat.
	 * see : https://schema.org/transFatContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $trans_fat_content;

	/**
	 * The number of grams of unsaturated fat.
	 * see : https://schema.org/unsaturatedFatContent
	 *
	 * @var \Mass | \Mass[]
	 */
	public $unsaturated_fat_content;

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
			'@type'    => 'NutritionInformation',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->calories );
		if ( ! empty( $serialized ) ) {
			$out['calories'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->carbohydrate_content );
		if ( ! empty( $serialized ) ) {
			$out['carbohydrateContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->cholesterol_content );
		if ( ! empty( $serialized ) ) {
			$out['cholesterolContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fat_content );
		if ( ! empty( $serialized ) ) {
			$out['fatContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fiber_content );
		if ( ! empty( $serialized ) ) {
			$out['fiberContent'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->protein_content );
		if ( ! empty( $serialized ) ) {
			$out['proteinContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->saturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['saturatedFatContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->serving_size );
		if ( ! empty( $serialized ) ) {
			$out['servingSize'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sodium_content );
		if ( ! empty( $serialized ) ) {
			$out['sodiumContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sugar_content );
		if ( ! empty( $serialized ) ) {
			$out['sugarContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->trans_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['transFatContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->unsaturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['unsaturatedFatContent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
