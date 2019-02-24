<?php
namespace SchemaOrg;

// UnitPriceSpecification see : https://schema.org/UnitPriceSpecification
class UnitPriceSpecification implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'UnitPriceSpecification';

	/**
	 * With properties from PriceSpecification see : https://schema.org/PriceSpecification
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
	 * This property specifies the minimal quantity and rounding increment that will be the basis for the billing. The unit of measurement is specified by the unitCode property.
	 * see : https://schema.org/billingIncrement
	 *
	 * @var float | float[]
	 */
	public $billing_increment;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	 * see : https://schema.org/eligibleQuantity
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $eligible_quantity;

	/**
	 * The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	 * see : https://schema.org/eligibleTransactionVolume
	 *
	 * @var \PriceSpecification | \PriceSpecification[]
	 */
	public $eligible_transaction_volume;

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
	 * The highest price if the price is a range.
	 * see : https://schema.org/maxPrice
	 *
	 * @var float | float[]
	 */
	public $max_price;

	/**
	 * The lowest price if the price is a range.
	 * see : https://schema.org/minPrice
	 *
	 * @var float | float[]
	 */
	public $min_price;

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
	 * The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
	 * &lt;br /&gt;
	 * &lt;br /&gt;
	 *       Usage guidelines:
	 * &lt;br /&gt;
	 * &lt;ul&gt;
	 * &lt;li&gt;Use the &lt;a href=&quot;/priceCurrency&quot;&gt;priceCurrency&lt;/a&gt; property (with &lt;a href=&quot;http://en.wikipedia.org/wiki/ISO_4217#Active_codes&quot;&gt;ISO 4217 codes&lt;/a&gt; e.g. &quot;USD&quot;) instead of
	 *       including &lt;a href=&quot;http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign&quot;&gt;ambiguous symbols&lt;/a&gt; such as &#39;$&#39; in the value.
	 * &lt;/li&gt;
	 * &lt;li&gt;
	 *       Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.
	 * &lt;/li&gt;
	 * &lt;li&gt;
	 *       Note that both &lt;a href=&quot;http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute&quot;&gt;RDFa&lt;/a&gt; and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values
	 *       alongside more human-friendly formatting.
	 * &lt;/li&gt;
	 * &lt;li&gt;
	 *       Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	 * &lt;/li&gt;
	 * &lt;/ul&gt;
	 *
	 * see : https://schema.org/price
	 *
	 * @var float | float[] | string | string[]
	 */
	public $price;

	/**
	 * The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification and its subtypes.
	 * see : https://schema.org/priceCurrency
	 *
	 * @var string | string[]
	 */
	public $price_currency;

	/**
	 * A short text or acronym indicating multiple price specifications for the same offer, e.g. SRP for the suggested retail price or INVOICE for the invoice price, mostly used in the car industry.
	 * see : https://schema.org/priceType
	 *
	 * @var string | string[]
	 */
	public $price_type;

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
	 * The end of the validity of offer, price specification, or opening hours data.
	 * see : https://schema.org/validThrough
	 *
	 * @var string | string[]
	 */
	public $valid_through;

	/**
	 * Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
	 * see : https://schema.org/valueAddedTaxIncluded
	 *
	 * @var boolean | boolean[]
	 */
	public $value_added_tax_included;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'UnitPriceSpecification',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->billing_increment );
		if ( ! empty( $serialized ) ) {
			$out['billingIncrement'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_quantity );
		if ( ! empty( $serialized ) ) {
			$out['eligibleQuantity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_transaction_volume );
		if ( ! empty( $serialized ) ) {
			$out['eligibleTransactionVolume'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->max_price );
		if ( ! empty( $serialized ) ) {
			$out['maxPrice'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->min_price );
		if ( ! empty( $serialized ) ) {
			$out['minPrice'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->price );
		if ( ! empty( $serialized ) ) {
			$out['price'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->price_type );
		if ( ! empty( $serialized ) ) {
			$out['priceType'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->value_added_tax_included );
		if ( ! empty( $serialized ) ) {
			$out['valueAddedTaxIncluded'] = $serialized;
		}

		return $out;
	}
}
