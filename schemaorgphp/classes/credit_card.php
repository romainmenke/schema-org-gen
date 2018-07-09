<?php

namespace SchemaOrg;

// CreditCard see : https://schema.org/CreditCard
class CreditCard implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CreditCard';
	
	/**
	 * With properties from Enumeration see : https://schema.org/Enumeration
	 */
	
	/**
	 * With properties from FinancialProduct see : https://schema.org/FinancialProduct
	 */
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from LoanOrCredit see : https://schema.org/LoanOrCredit
	 */
	
	/**
	 * With properties from PaymentCard see : https://schema.org/PaymentCard
	 */
	
	/**
	 * With properties from PaymentMethod see : https://schema.org/PaymentMethod
	 */
	
	/**
	 * With properties from Service see : https://schema.org/Service
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * The amount of money.
	 * see : https://schema.org/amount
	 * @var \MonetaryAmount | \MonetaryAmount[] | float | float[]
	 */
	public var $amount;
	
	/**
	 * The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
	 * see : https://schema.org/annualPercentageRate
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $annual_percentage_rate;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public var $area_served;
	
	/**
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience | \Audience[]
	 */
	public var $audience;
	
	/**
	 * A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
	 * see : https://schema.org/availableChannel
	 * @var \ServiceChannel | \ServiceChannel[]
	 */
	public var $available_channel;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string | string[]
	 */
	public var $award;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public var $brand;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $broker;
	
	/**
	 * A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
	 * see : https://pending.schema.org/cashBack
	 * @var boolean | boolean[] | float | float[]
	 */
	public var $cash_back;
	
	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://pending.schema.org/category
	 * @var \PhysicalActivityCategory | \PhysicalActivityCategory[] | string | string[] | \Thing | \Thing[]
	 */
	public var $category;
	
	/**
	 * A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
	 * see : https://pending.schema.org/contactlessPayment
	 * @var boolean | boolean[]
	 */
	public var $contactless_payment;
	
	/**
	 * The currency in which the monetary amount is expressed.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://pending.schema.org/currency
	 * @var string | string[]
	 */
	public var $currency;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public var $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public var $disambiguating_description;
	
	/**
	 * Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	 * see : https://schema.org/feesAndCommissionsSpecification
	 * @var string | string[]
	 */
	public var $fees_and_commissions_specification;
	
	/**
	 * A floor limit is the amount of money above which credit card transactions must be authorized.
	 * see : https://pending.schema.org/floorLimit
	 * @var \MonetaryAmount | \MonetaryAmount[]
	 */
	public var $floor_limit;
	
	/**
	 * The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
	 * see : https://pending.schema.org/gracePeriod
	 * @var \Duration | \Duration[]
	 */
	public var $grace_period;
	
	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public var $has_offer_catalog;
	
	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public var $hours_available;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public var $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $image;
	
	/**
	 * The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
	 * see : https://schema.org/interestRate
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $interest_rate;
	
	/**
	 * A pointer to another, somehow related product (or multiple products).
	 * see : https://schema.org/isRelatedTo
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public var $is_related_to;
	
	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public var $is_similar_to;
	
	/**
	 * A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
	 * see : https://pending.schema.org/loanRepaymentForm
	 * @var \RepaymentSpecification | \RepaymentSpecification[]
	 */
	public var $loan_repayment_form;
	
	/**
	 * The duration of the loan or credit agreement.
	 * see : https://schema.org/loanTerm
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $loan_term;
	
	/**
	 * The type of a loan or credit.
	 * see : https://pending.schema.org/loanType
	 * @var string | string[]
	 */
	public var $loan_type;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $logo;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
	 * see : https://pending.schema.org/monthlyMinimumRepaymentAmount
	 * @var \MonetaryAmount | \MonetaryAmount[] | float | float[]
	 */
	public var $monthly_minimum_repayment_amount;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer | \Offer[]
	 */
	public var $offers;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $provider;
	
	/**
	 * Indicates the mobility of a provided service (e.g. &#39;static&#39;, &#39;dynamic&#39;).
	 * see : https://schema.org/providerMobility
	 * @var string | string[]
	 */
	public var $provider_mobility;
	
	/**
	 * The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
	 * see : https://pending.schema.org/recourseLoan
	 * @var boolean | boolean[]
	 */
	public var $recourse_loan;
	
	/**
	 * Whether the terms for payment of interest can be renegotiated during the life of the loan.
	 * see : https://pending.schema.org/renegotiableLoan
	 * @var boolean | boolean[]
	 */
	public var $renegotiable_loan;
	
	/**
	 * Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
	 * see : https://schema.org/requiredCollateral
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public var $required_collateral;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review | \Review[]
	 */
	public var $review;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The tangible thing generated by the service, e.g. a passport, permit, etc. Supersedes produces (see: https://schema.org/produces).
	 * see : https://schema.org/serviceOutput
	 * @var \Thing | \Thing[]
	 */
	public var $service_output;
	
	/**
	 * The type of service being offered, e.g. veterans&#39; benefits, emergency relief, etc.
	 * see : https://schema.org/serviceType
	 * @var string | string[]
	 */
	public var $service_type;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	 * see : https://meta.schema.org/supersededBy
	 * @var \Class | \Class[] | \Enumeration | \Enumeration[] | \Property | \Property[]
	 */
	public var $superseded_by;
	
	/**
	 * Human-readable terms of service documentation.
	 * see : https://pending.schema.org/termsOfService
	 * @var string | string[]
	 */
	public var $terms_of_service;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CreditCard'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->amount );
		if ( ! empty( $serialized ) ) {
			$out['amount'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->annual_percentage_rate );
		if ( ! empty( $serialized ) ) {
			$out['annualPercentageRate'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->available_channel );
		if ( ! empty( $serialized ) ) {
			$out['availableChannel'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->cash_back );
		if ( ! empty( $serialized ) ) {
			$out['cashBack'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contactless_payment );
		if ( ! empty( $serialized ) ) {
			$out['contactlessPayment'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->currency );
		if ( ! empty( $serialized ) ) {
			$out['currency'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fees_and_commissions_specification );
		if ( ! empty( $serialized ) ) {
			$out['feesAndCommissionsSpecification'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->floor_limit );
		if ( ! empty( $serialized ) ) {
			$out['floorLimit'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->grace_period );
		if ( ! empty( $serialized ) ) {
			$out['gracePeriod'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->interest_rate );
		if ( ! empty( $serialized ) ) {
			$out['interestRate'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->is_related_to );
		if ( ! empty( $serialized ) ) {
			$out['isRelatedTo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->is_similar_to );
		if ( ! empty( $serialized ) ) {
			$out['isSimilarTo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->loan_repayment_form );
		if ( ! empty( $serialized ) ) {
			$out['loanRepaymentForm'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->loan_term );
		if ( ! empty( $serialized ) ) {
			$out['loanTerm'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->loan_type );
		if ( ! empty( $serialized ) ) {
			$out['loanType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->monthly_minimum_repayment_amount );
		if ( ! empty( $serialized ) ) {
			$out['monthlyMinimumRepaymentAmount'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->provider_mobility );
		if ( ! empty( $serialized ) ) {
			$out['providerMobility'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->recourse_loan );
		if ( ! empty( $serialized ) ) {
			$out['recourseLoan'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->renegotiable_loan );
		if ( ! empty( $serialized ) ) {
			$out['renegotiableLoan'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->required_collateral );
		if ( ! empty( $serialized ) ) {
			$out['requiredCollateral'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->service_output );
		if ( ! empty( $serialized ) ) {
			$out['serviceOutput'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->service_type );
		if ( ! empty( $serialized ) ) {
			$out['serviceType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->superseded_by );
		if ( ! empty( $serialized ) ) {
			$out['supersededBy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->terms_of_service );
		if ( ! empty( $serialized ) ) {
			$out['termsOfService'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
