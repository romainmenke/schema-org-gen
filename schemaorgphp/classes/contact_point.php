<?php

class ContactPoint extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ContactPoint';
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea|\AdministrativeArea[]|\GeoShape|\GeoShape[]|\Place|\Place[]|string|string[]
	 */
	public var $area_served;
	
	/**
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $available_language;
	
	/**
	 * An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
	 * see : https://schema.org/contactOption
	 * @var \ContactPointOption|\ContactPointOption[]
	 */
	public var $contact_option;
	
	/**
	 * A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
	 * see : https://schema.org/contactType
	 * @var string|string[]
	 */
	public var $contact_type;
	
	/**
	 * Email address.
	 * see : https://schema.org/email
	 * @var string|string[]
	 */
	public var $email;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string|string[]
	 */
	public var $fax_number;
	
	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 * @var \OpeningHoursSpecification|\OpeningHoursSpecification[]
	 */
	public var $hours_available;
	
	/**
	 * The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. &quot;iPhone&quot;) or a general category of products or services (e.g. &quot;smartphones&quot;).
	 * see : https://schema.org/productSupported
	 * @var \Product|\Product[]|string|string[]
	 */
	public var $product_supported;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string|string[]
	 */
	public var $telephone;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ContactPoint'
		);
		
		$serialized = so_json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contact_option );
		if ( ! empty( $serialized ) ) {
			$out['contactOption'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contact_type );
		if ( ! empty( $serialized ) ) {
			$out['contactType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->product_supported );
		if ( ! empty( $serialized ) ) {
			$out['productSupported'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		return $out;
	}
}
