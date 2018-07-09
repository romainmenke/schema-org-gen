<?php

namespace SchemaOrg;

// PostalAddress see : https://schema.org/PostalAddress
class PostalAddress implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PostalAddress';
	
	/**
	 * With properties from ContactPoint see : https://schema.org/ContactPoint
	 */
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from StructuredValue see : https://schema.org/StructuredValue
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
	 * The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	 * see : https://schema.org/addressCountry
	 * @var \Country | \Country[] | string | string[]
	 */
	public var $address_country;
	
	/**
	 * The locality. For example, Mountain View.
	 * see : https://schema.org/addressLocality
	 * @var string | string[]
	 */
	public var $address_locality;
	
	/**
	 * The region. For example, CA.
	 * see : https://schema.org/addressRegion
	 * @var string | string[]
	 */
	public var $address_region;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public var $area_served;
	
	/**
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language | \Language[] | string | string[]
	 */
	public var $available_language;
	
	/**
	 * An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
	 * see : https://schema.org/contactOption
	 * @var \ContactPointOption | \ContactPointOption[]
	 */
	public var $contact_option;
	
	/**
	 * A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
	 * see : https://schema.org/contactType
	 * @var string | string[]
	 */
	public var $contact_type;
	
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
	 * Email address.
	 * see : https://schema.org/email
	 * @var string | string[]
	 */
	public var $email;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string | string[]
	 */
	public var $fax_number;
	
	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public var $hours_available;
	
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
	 * The post office box number for PO box addresses.
	 * see : https://schema.org/postOfficeBoxNumber
	 * @var string | string[]
	 */
	public var $post_office_box_number;
	
	/**
	 * The postal code. For example, 94043.
	 * see : https://schema.org/postalCode
	 * @var string | string[]
	 */
	public var $postal_code;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. &quot;iPhone&quot;) or a general category of products or services (e.g. &quot;smartphones&quot;).
	 * see : https://schema.org/productSupported
	 * @var \Product | \Product[] | string | string[]
	 */
	public var $product_supported;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The street address. For example, 1600 Amphitheatre Pkwy.
	 * see : https://schema.org/streetAddress
	 * @var string | string[]
	 */
	public var $street_address;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string | string[]
	 */
	public var $telephone;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PostalAddress'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->address_country );
		if ( ! empty( $serialized ) ) {
			$out['addressCountry'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->address_locality );
		if ( ! empty( $serialized ) ) {
			$out['addressLocality'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->address_region );
		if ( ! empty( $serialized ) ) {
			$out['addressRegion'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contact_option );
		if ( ! empty( $serialized ) ) {
			$out['contactOption'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contact_type );
		if ( ! empty( $serialized ) ) {
			$out['contactType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->post_office_box_number );
		if ( ! empty( $serialized ) ) {
			$out['postOfficeBoxNumber'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->postal_code );
		if ( ! empty( $serialized ) ) {
			$out['postalCode'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->product_supported );
		if ( ! empty( $serialized ) ) {
			$out['productSupported'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->street_address );
		if ( ! empty( $serialized ) ) {
			$out['streetAddress'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
