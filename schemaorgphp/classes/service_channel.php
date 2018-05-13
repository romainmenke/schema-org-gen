<?php

class ServiceChannel extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ServiceChannel';
	
	/**
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $available_language;
	
	/**
	 * Estimated processing time for the service using this channel.
	 * see : https://schema.org/processingTime
	 * @var \Duration|\Duration[]
	 */
	public var $processing_time;
	
	/**
	 * The service provided by this channel.
	 * see : https://schema.org/providesService
	 * @var \Service|\Service[]
	 */
	public var $provides_service;
	
	/**
	 * The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
	 * see : https://schema.org/serviceLocation
	 * @var \Place|\Place[]
	 */
	public var $service_location;
	
	/**
	 * The phone number to use to access the service.
	 * see : https://schema.org/servicePhone
	 * @var \ContactPoint|\ContactPoint[]
	 */
	public var $service_phone;
	
	/**
	 * The address for accessing the service by mail.
	 * see : https://schema.org/servicePostalAddress
	 * @var \PostalAddress|\PostalAddress[]
	 */
	public var $service_postal_address;
	
	/**
	 * The number to access the service by text message.
	 * see : https://schema.org/serviceSmsNumber
	 * @var \ContactPoint|\ContactPoint[]
	 */
	public var $service_sms_number;
	
	/**
	 * The website to access the service.
	 * see : https://schema.org/serviceUrl
	 * @var string|string[]
	 */
	public var $service_url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ServiceChannel'
		);
		
		$serialized = so_json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->processing_time );
		if ( ! empty( $serialized ) ) {
			$out['processingTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provides_service );
		if ( ! empty( $serialized ) ) {
			$out['providesService'] = $serialized;
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
		
		return $out;
	}
}
