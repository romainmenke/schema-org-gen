<?php

class Corporation extends Organization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Corporation';
	
	/**
	 * The exchange traded instrument associated with a Corporation object. The tickerSymbol is expressed as an exchange and an instrument name separated by a space character. For the exchange component of the tickerSymbol attribute, we reccommend using the controlled vocaulary of Market Identifier Codes (MIC) specified in ISO15022.
	 * see : https://schema.org/tickerSymbol
	 * @var string|string[]
	 */
	public var $ticker_symbol;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Corporation'
		);
		
		$serialized = so_json_serialize( $this->ticker_symbol );
		if ( ! empty( $serialized ) ) {
			$out['tickerSymbol'] = $serialized;
		}
		
		return $out;
	}
}
