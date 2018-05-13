<?php

class CreativeWorkSeries extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CreativeWorkSeries';
	
	/**
	 * The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/endDate
	 * @var string|string[]|string|string[]
	 */
	public var $end_date;
	
	/**
	 * The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	 * see : https://schema.org/issn
	 * @var string|string[]
	 */
	public var $issn;
	
	/**
	 * The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/startDate
	 * @var string|string[]|string|string[]
	 */
	public var $start_date;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CreativeWorkSeries'
		);
		
		$serialized = so_json_serialize( $this->end_date );
		if ( ! empty( $serialized ) ) {
			$out['endDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->issn );
		if ( ! empty( $serialized ) ) {
			$out['issn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_date );
		if ( ! empty( $serialized ) ) {
			$out['startDate'] = $serialized;
		}
		
		return $out;
	}
}
