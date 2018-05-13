<?php

class Report extends Article implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Report';
	
	/**
	 * The number or other unique designator assigned to a Report by the publishing organization.
	 * see : https://schema.org/reportNumber
	 * @var string|string[]
	 */
	public var $report_number;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Report'
		);
		
		$serialized = so_json_serialize( $this->report_number );
		if ( ! empty( $serialized ) ) {
			$out['reportNumber'] = $serialized;
		}
		
		return $out;
	}
}
