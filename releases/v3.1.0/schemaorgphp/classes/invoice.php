<?php
namespace SchemaOrg;

// Invoice see : https://schema.org/Invoice
class Invoice implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Invoice';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * The identifier for the account the payment will be applied to.
	 * see : https://schema.org/accountId
	 *
	 * @var string | string[]
	 */
	public $account_id;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The time interval used to compute the invoice.
	 * see : https://schema.org/billingPeriod
	 *
	 * @var \Duration | \Duration[]
	 */
	public $billing_period;

	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	 * see : https://schema.org/broker
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $broker;

	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://schema.org/category
	 *
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public $category;

	/**
	 * A number that confirms the given order or payment has been received.
	 * see : https://schema.org/confirmationNumber
	 *
	 * @var string | string[]
	 */
	public $confirmation_number;

	/**
	 * Party placing the order or paying the invoice.
	 * see : https://schema.org/customer
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $customer;

	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The minimum payment required at this time.
	 * see : https://schema.org/minimumPaymentDue
	 *
	 * @var \PriceSpecification | \PriceSpecification[] | \MonetaryAmount | \MonetaryAmount[]
	 */
	public $minimum_payment_due;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The date that payment is due.
	 * see : https://schema.org/paymentDue
	 *
	 * @var string | string[]
	 */
	public $payment_due;

	/**
	 * The date that payment is due.
	 * see : https://schema.org/paymentDueDate
	 *
	 * @var string | string[]
	 */
	public $payment_due_date;

	/**
	 * The name of the credit card or other method of payment for the order.
	 * see : https://schema.org/paymentMethod
	 *
	 * @var \PaymentMethod | \PaymentMethod[]
	 */
	public $payment_method;

	/**
	 * An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	 * see : https://schema.org/paymentMethodId
	 *
	 * @var string | string[]
	 */
	public $payment_method_id;

	/**
	 * The status of payment; whether the invoice has been paid or not.
	 * see : https://schema.org/paymentStatus
	 *
	 * @var \PaymentStatusType | \PaymentStatusType[] | string | string[]
	 */
	public $payment_status;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	 * see : https://schema.org/provider
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $provider;

	/**
	 * The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice.
	 * see : https://schema.org/referencesOrder
	 *
	 * @var \Order | \Order[]
	 */
	public $references_order;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The date the invoice is scheduled to be paid.
	 * see : https://schema.org/scheduledPaymentDate
	 *
	 * @var string | string[]
	 */
	public $scheduled_payment_date;

	/**
	 * The total amount due.
	 * see : https://schema.org/totalPaymentDue
	 *
	 * @var \PriceSpecification | \PriceSpecification[] | \MonetaryAmount | \MonetaryAmount[]
	 */
	public $total_payment_due;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Invoice',
		);

		$serialized = \SchemaOrg\json_serialize( $this->account_id );
		if ( ! empty( $serialized ) ) {
			$out['accountId'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->billing_period );
		if ( ! empty( $serialized ) ) {
			$out['billingPeriod'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->confirmation_number );
		if ( ! empty( $serialized ) ) {
			$out['confirmationNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->customer );
		if ( ! empty( $serialized ) ) {
			$out['customer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->minimum_payment_due );
		if ( ! empty( $serialized ) ) {
			$out['minimumPaymentDue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->payment_due );
		if ( ! empty( $serialized ) ) {
			$out['paymentDue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->payment_due_date );
		if ( ! empty( $serialized ) ) {
			$out['paymentDueDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->payment_method );
		if ( ! empty( $serialized ) ) {
			$out['paymentMethod'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->payment_method_id );
		if ( ! empty( $serialized ) ) {
			$out['paymentMethodId'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->payment_status );
		if ( ! empty( $serialized ) ) {
			$out['paymentStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->references_order );
		if ( ! empty( $serialized ) ) {
			$out['referencesOrder'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->scheduled_payment_date );
		if ( ! empty( $serialized ) ) {
			$out['scheduledPaymentDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->total_payment_due );
		if ( ! empty( $serialized ) ) {
			$out['totalPaymentDue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
