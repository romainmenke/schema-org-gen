<?php

// ServiceChannel see : https://schema.org/ServiceChannel
class ServiceChannel implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ServiceChannel';
	
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
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language | \Language[] | string | string[]
	 */
	public var $available_language;
	
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
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * Estimated processing time for the service using this channel.
	 * see : https://schema.org/processingTime
	 * @var \Duration | \Duration[]
	 */
	public var $processing_time;
	
	/**
	 * The service provided by this channel.
	 * see : https://schema.org/providesService
	 * @var \Service | \Service[]
	 */
	public var $provides_service;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
	 * see : https://schema.org/serviceLocation
	 * @var \Place | \Place[]
	 */
	public var $service_location;
	
	/**
	 * The phone number to use to access the service.
	 * see : https://schema.org/servicePhone
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public var $service_phone;
	
	/**
	 * The address for accessing the service by mail.
	 * see : https://schema.org/servicePostalAddress
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public var $service_postal_address;
	
	/**
	 * The number to access the service by text message.
	 * see : https://schema.org/serviceSmsNumber
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public var $service_sms_number;
	
	/**
	 * The website to access the service.
	 * see : https://schema.org/serviceUrl
	 * @var string | string[]
	 */
	public var $service_url;
	
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
			'@type' => 'ServiceChannel'
		);
		
		$serialized = so_json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->processing_time );
		if ( ! empty( $serialized ) ) {
			$out['processingTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provides_service );
		if ( ! empty( $serialized ) ) {
			$out['providesService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->service_location );
		if ( ! empty( $serialized ) ) {
			$out['serviceLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->service_phone );
		if ( ! empty( $serialized ) ) {
			$out['servicePhone'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->service_postal_address );
		if ( ! empty( $serialized ) ) {
			$out['servicePostalAddress'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->service_sms_number );
		if ( ! empty( $serialized ) ) {
			$out['serviceSmsNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->service_url );
		if ( ! empty( $serialized ) ) {
			$out['serviceUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
