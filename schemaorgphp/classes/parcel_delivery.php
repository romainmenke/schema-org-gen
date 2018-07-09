<?php

namespace SchemaOrg;

// ParcelDelivery see : https://schema.org/ParcelDelivery
class ParcelDelivery implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ParcelDelivery';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * Destination address.
	 * see : https://schema.org/deliveryAddress
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public var $delivery_address;
	
	/**
	 * New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	 * see : https://schema.org/deliveryStatus
	 * @var \DeliveryEvent | \DeliveryEvent[]
	 */
	public var $delivery_status;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public var $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public var $disambiguating_description;
	
	/**
	 * The earliest date the package may arrive.
	 * see : https://schema.org/expectedArrivalFrom
	 * @var string | string[]
	 */
	public var $expected_arrival_from;
	
	/**
	 * The latest date the package may arrive.
	 * see : https://schema.org/expectedArrivalUntil
	 * @var string | string[]
	 */
	public var $expected_arrival_until;
	
	/**
	 * Method used for delivery or shipping.
	 * see : https://schema.org/hasDeliveryMethod
	 * @var \DeliveryMethod | \DeliveryMethod[]
	 */
	public var $has_delivery_method;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public var $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $image;
	
	/**
	 * Item(s) being shipped.
	 * see : https://schema.org/itemShipped
	 * @var \Product | \Product[]
	 */
	public var $item_shipped;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * Shipper&#39;s address.
	 * see : https://schema.org/originAddress
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public var $origin_address;
	
	/**
	 * The overall order the items in this delivery were included in.
	 * see : https://schema.org/partOfOrder
	 * @var \Order | \Order[]
	 */
	public var $part_of_order;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $provider;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * Shipper tracking number.
	 * see : https://schema.org/trackingNumber
	 * @var string | string[]
	 */
	public var $tracking_number;
	
	/**
	 * Tracking url for the parcel delivery.
	 * see : https://schema.org/trackingUrl
	 * @var string | string[]
	 */
	public var $tracking_url;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ParcelDelivery'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
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
