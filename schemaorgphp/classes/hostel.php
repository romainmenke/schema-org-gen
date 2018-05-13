<?php

class Hostel extends LodgingBusiness implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Hostel';
	
	/**
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 * @var \LocationFeatureSpecification|\LocationFeatureSpecification[]
	 */
	public var $amenity_feature;
	
	/**
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience|\Audience[]
	 */
	public var $audience;
	
	/**
	 * A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	 * see : https://schema.org/availableLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $available_language;
	
	/**
	 * The earliest someone may check into a lodging establishment.
	 * see : https://schema.org/checkinTime
	 * @var string|string[]
	 */
	public var $checkin_time;
	
	/**
	 * The latest someone may check out of a lodging establishment.
	 * see : https://schema.org/checkoutTime
	 * @var string|string[]
	 */
	public var $checkout_time;
	
	/**
	 * Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	 * see : https://schema.org/petsAllowed
	 * @var boolean|boolean[]|string|string[]
	 */
	public var $pets_allowed;
	
	/**
	 * An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	 * see : https://schema.org/starRating
	 * @var \Rating|\Rating[]
	 */
	public var $star_rating;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Hostel'
		);
		
		$serialized = so_json_serialize( $this->amenity_feature );
		if ( ! empty( $serialized ) ) {
			$out['amenityFeature'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_language );
		if ( ! empty( $serialized ) ) {
			$out['availableLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->checkin_time );
		if ( ! empty( $serialized ) ) {
			$out['checkinTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->checkout_time );
		if ( ! empty( $serialized ) ) {
			$out['checkoutTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->pets_allowed );
		if ( ! empty( $serialized ) ) {
			$out['petsAllowed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->star_rating );
		if ( ! empty( $serialized ) ) {
			$out['starRating'] = $serialized;
		}
		
		return $out;
	}
}
