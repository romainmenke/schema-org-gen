<?php
namespace SchemaOrg;

// ParcelDelivery see : https://schema.org/ParcelDelivery
class ParcelDelivery implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'ParcelDelivery';

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
	 * &#39;carrier&#39; is an out-dated term indicating the &#39;provider&#39; for parcel delivery and flights.
	 * see : https://schema.org/carrier
	 *
	 * @var \Organization | \Organization[]
	 */
	public $carrier;

	/**
	 * Destination address.
	 * see : https://schema.org/deliveryAddress
	 *
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public $delivery_address;

	/**
	 * New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	 * see : https://schema.org/deliveryStatus
	 *
	 * @var \DeliveryEvent | \DeliveryEvent[]
	 */
	public $delivery_status;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The earliest date the package may arrive.
	 * see : https://schema.org/expectedArrivalFrom
	 *
	 * @var string | string[]
	 */
	public $expected_arrival_from;

	/**
	 * The latest date the package may arrive.
	 * see : https://schema.org/expectedArrivalUntil
	 *
	 * @var string | string[]
	 */
	public $expected_arrival_until;

	/**
	 * Method used for delivery or shipping.
	 * see : https://schema.org/hasDeliveryMethod
	 *
	 * @var \DeliveryMethod | \DeliveryMethod[]
	 */
	public $has_delivery_method;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Item(s) being shipped.
	 * see : https://schema.org/itemShipped
	 *
	 * @var \Product | \Product[]
	 */
	public $item_shipped;

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
	 * Shipper&#39;s address.
	 * see : https://schema.org/originAddress
	 *
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public $origin_address;

	/**
	 * The overall order the items in this delivery were included in.
	 * see : https://schema.org/partOfOrder
	 *
	 * @var \Order | \Order[]
	 */
	public $part_of_order;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	 * see : https://schema.org/provider
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $provider;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * Shipper tracking number.
	 * see : https://schema.org/trackingNumber
	 *
	 * @var string | string[]
	 */
	public $tracking_number;

	/**
	 * Tracking url for the parcel delivery.
	 * see : https://schema.org/trackingUrl
	 *
	 * @var string | string[]
	 */
	public $tracking_url;

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
			'@type'    => 'ParcelDelivery',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->carrier );
		if ( ! empty( $serialized ) ) {
			$out['carrier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->delivery_address );
		if ( ! empty( $serialized ) ) {
			$out['deliveryAddress'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->delivery_status );
		if ( ! empty( $serialized ) ) {
			$out['deliveryStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->expected_arrival_from );
		if ( ! empty( $serialized ) ) {
			$out['expectedArrivalFrom'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->expected_arrival_until );
		if ( ! empty( $serialized ) ) {
			$out['expectedArrivalUntil'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['hasDeliveryMethod'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->item_shipped );
		if ( ! empty( $serialized ) ) {
			$out['itemShipped'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->origin_address );
		if ( ! empty( $serialized ) ) {
			$out['originAddress'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->part_of_order );
		if ( ! empty( $serialized ) ) {
			$out['partOfOrder'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->tracking_number );
		if ( ! empty( $serialized ) ) {
			$out['trackingNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->tracking_url );
		if ( ! empty( $serialized ) ) {
			$out['trackingUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
