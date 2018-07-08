<?php

class QuoteAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'QuoteAction';
	
	/**
	 * The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
	 * 
	 * Usage guidelines:
	 * 
	 * 
	 * Use the priceCurrency (see: https://schema.org/priceCurrency) property (with standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;) instead of including ambiguous symbols (see: https://schema.orghttp://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as &#39;$&#39; in the value.
	 * Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.
	 * Note that both RDFa (see: https://schema.orghttp://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values alongside more human-friendly formatting.
	 * Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	 * 
	 * 
	 * see : https://schema.org/price
	 * @var float|float[]|string|string[]
	 */
	public var $price;
	
	/**
	 * One or more detailed price specifications, indicating the unit price and delivery or payment charges.
	 * see : https://schema.org/priceSpecification
	 * @var \PriceSpecification|\PriceSpecification[]
	 */
	public var $price_specification;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'QuoteAction'
		);
		
		$serialized = so_json_serialize( $this->price );
		if ( ! empty( $serialized ) ) {
			$out['price'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_specification );
		if ( ! empty( $serialized ) ) {
			$out['priceSpecification'] = $serialized;
		}
		
		return $out;
	}
}
