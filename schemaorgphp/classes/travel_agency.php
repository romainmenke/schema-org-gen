<?php

class TravelAgency extends LocalBusiness implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TravelAgency';
	
	/**
	 * The currency accepted (in ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217)).
	 * see : https://schema.org/currenciesAccepted
	 * @var string|string[]
	 */
	public var $currencies_accepted;
	
	/**
	 * The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.
	 * 
	 * 
	 * Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	 * Times are specified using 24:00 time. For example, 3pm is specified as 15:00. 
	 * Here is an example: &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;.
	 * If a business is open 7 days a week, then it can be specified as &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;.
	 * 
	 * 
	 * see : https://schema.org/openingHours
	 * @var string|string[]
	 */
	public var $opening_hours;
	
	/**
	 * Cash, credit card, etc.
	 * see : https://schema.org/paymentAccepted
	 * @var string|string[]
	 */
	public var $payment_accepted;
	
	/**
	 * The price range of the business, for example $$$.
	 * see : https://schema.org/priceRange
	 * @var string|string[]
	 */
	public var $price_range;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TravelAgency'
		);
		
		$serialized = so_json_serialize( $this->currencies_accepted );
		if ( ! empty( $serialized ) ) {
			$out['currenciesAccepted'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->opening_hours );
		if ( ! empty( $serialized ) ) {
			$out['openingHours'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payment_accepted );
		if ( ! empty( $serialized ) ) {
			$out['paymentAccepted'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_range );
		if ( ! empty( $serialized ) ) {
			$out['priceRange'] = $serialized;
		}
		
		return $out;
	}
}
