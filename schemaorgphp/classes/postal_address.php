<?php

class PostalAddress extends ContactPoint implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PostalAddress';
	
	/**
	 * The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	 * see : https://schema.org/addressCountry
	 * @var \Country|\Country[]|string|string[]
	 */
	public var $address_country;
	
	/**
	 * The locality. For example, Mountain View.
	 * see : https://schema.org/addressLocality
	 * @var string|string[]
	 */
	public var $address_locality;
	
	/**
	 * The region. For example, CA.
	 * see : https://schema.org/addressRegion
	 * @var string|string[]
	 */
	public var $address_region;
	
	/**
	 * The post office box number for PO box addresses.
	 * see : https://schema.org/postOfficeBoxNumber
	 * @var string|string[]
	 */
	public var $post_office_box_number;
	
	/**
	 * The postal code. For example, 94043.
	 * see : https://schema.org/postalCode
	 * @var string|string[]
	 */
	public var $postal_code;
	
	/**
	 * The street address. For example, 1600 Amphitheatre Pkwy.
	 * see : https://schema.org/streetAddress
	 * @var string|string[]
	 */
	public var $street_address;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PostalAddress'
		);
		
		$serialized = so_json_serialize( $this->address_country );
		if ( ! empty( $serialized ) ) {
			$out['addressCountry'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address_locality );
		if ( ! empty( $serialized ) ) {
			$out['addressLocality'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address_region );
		if ( ! empty( $serialized ) ) {
			$out['addressRegion'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->post_office_box_number );
		if ( ! empty( $serialized ) ) {
			$out['postOfficeBoxNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->postal_code );
		if ( ! empty( $serialized ) ) {
			$out['postalCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->street_address );
		if ( ! empty( $serialized ) ) {
			$out['streetAddress'] = $serialized;
		}
		
		return $out;
	}
}
