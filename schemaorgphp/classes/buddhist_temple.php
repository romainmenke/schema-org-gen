<?php

class BuddhistTemple extends PlaceOfWorship implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BuddhistTemple';
	
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
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BuddhistTemple'
		);
		
		$serialized = so_json_serialize( $this->opening_hours );
		if ( ! empty( $serialized ) ) {
			$out['openingHours'] = $serialized;
		}
		
		return $out;
	}
}
