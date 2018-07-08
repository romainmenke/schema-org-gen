<?php

class PriceSpecification extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PriceSpecification';
	
	/**
	 * The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	 * see : https://schema.org/eligibleQuantity
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $eligible_quantity;
	
	/**
	 * The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	 * see : https://schema.org/eligibleTransactionVolume
	 * @var \PriceSpecification|\PriceSpecification[]
	 */
	public var $eligible_transaction_volume;
	
	/**
	 * The highest price if the price is a range.
	 * see : https://schema.org/maxPrice
	 * @var float|float[]
	 */
	public var $max_price;
	
	/**
	 * The lowest price if the price is a range.
	 * see : https://schema.org/minPrice
	 * @var float|float[]
	 */
	public var $min_price;
	
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
	 * The currency of the price, or a price component when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://schema.org/priceCurrency
	 * @var string|string[]
	 */
	public var $price_currency;
	
	/**
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 * @var string|string[]
	 */
	public var $valid_from;
	
	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 * @var string|string[]
	 */
	public var $valid_through;
	
	/**
	 * Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
	 * see : https://schema.org/valueAddedTaxIncluded
	 * @var boolean|boolean[]
	 */
	public var $value_added_tax_included;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PriceSpecification'
		);
		
		$serialized = so_json_serialize( $this->eligible_quantity );
		if ( ! empty( $serialized ) ) {
			$out['eligibleQuantity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->eligible_transaction_volume );
		if ( ! empty( $serialized ) ) {
			$out['eligibleTransactionVolume'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->max_price );
		if ( ! empty( $serialized ) ) {
			$out['maxPrice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->min_price );
		if ( ! empty( $serialized ) ) {
			$out['minPrice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price );
		if ( ! empty( $serialized ) ) {
			$out['price'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_added_tax_included );
		if ( ! empty( $serialized ) ) {
			$out['valueAddedTaxIncluded'] = $serialized;
		}
		
		return $out;
	}
}
