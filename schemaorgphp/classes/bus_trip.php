<?php

namespace SchemaOrg;

// BusTrip see : https://schema.org/BusTrip
class BusTrip implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BusTrip';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	/**
	 * With properties from Trip see : https://schema.org/Trip
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
	 * The stop or station from which the bus arrives.
	 * see : https://schema.org/arrivalBusStop
	 * @var \BusStation | \BusStation[] | \BusStop | \BusStop[]
	 */
	public var $arrival_bus_stop;
	
	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 * @var string | string[]
	 */
	public var $arrival_time;
	
	/**
	 * The name of the bus (e.g. Bolt Express).
	 * see : https://schema.org/busName
	 * @var string | string[]
	 */
	public var $bus_name;
	
	/**
	 * The unique identifier for the bus.
	 * see : https://schema.org/busNumber
	 * @var string | string[]
	 */
	public var $bus_number;
	
	/**
	 * The stop or station from which the bus departs.
	 * see : https://schema.org/departureBusStop
	 * @var \BusStation | \BusStation[] | \BusStop | \BusStop[]
	 */
	public var $departure_bus_stop;
	
	/**
	 * The expected departure time.
	 * see : https://schema.org/departureTime
	 * @var string | string[]
	 */
	public var $departure_time;
	
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
	 * Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense). Inverse property: isPartOf (see: https://schema.org/isPartOf).
	 * see : https://schema.org/hasPart
	 * @var \CreativeWork | \CreativeWork[] | \Trip | \Trip[]
	 */
	public var $has_part;
	
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
	 * Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	 * see : https://schema.org/isPartOf
	 * @var \CreativeWork | \CreativeWork[] | \Trip | \Trip[]
	 */
	public var $is_part_of;
	
	/**
	 * Destination(s) ( Place (see: https://schema.org/Place) ) that make up a trip. For a trip where destination order is important use ItemList (see: https://schema.org/ItemList) to specify that order (see examples).
	 * see : https://pending.schema.org/itinerary
	 * @var \ItemList | \ItemList[] | \Place | \Place[]
	 */
	public var $itinerary;
	
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
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer | \Offer[]
	 */
	public var $offers;
	
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
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BusTrip'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->arrival_bus_stop );
		if ( ! empty( $serialized ) ) {
			$out['arrivalBusStop'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->bus_name );
		if ( ! empty( $serialized ) ) {
			$out['busName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->bus_number );
		if ( ! empty( $serialized ) ) {
			$out['busNumber'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->departure_bus_stop );
		if ( ! empty( $serialized ) ) {
			$out['departureBusStop'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->has_part );
		if ( ! empty( $serialized ) ) {
			$out['hasPart'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->is_part_of );
		if ( ! empty( $serialized ) ) {
			$out['isPartOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->itinerary );
		if ( ! empty( $serialized ) ) {
			$out['itinerary'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
