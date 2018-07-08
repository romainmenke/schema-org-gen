<?php

class Invoice extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Invoice';
	
	/**
	 * The identifier for the account the payment will be applied to.
	 * see : https://schema.org/accountId
	 * @var string|string[]
	 */
	public var $account_id;
	
	/**
	 * The time interval used to compute the invoice.
	 * see : https://schema.org/billingPeriod
	 * @var \Duration|\Duration[]
	 */
	public var $billing_period;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $broker;
	
	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://pending.schema.org/category
	 * @var \PhysicalActivityCategory|\PhysicalActivityCategory[]|string|string[]|\Thing|\Thing[]
	 */
	public var $category;
	
	/**
	 * A number that confirms the given order or payment has been received.
	 * see : https://schema.org/confirmationNumber
	 * @var string|string[]
	 */
	public var $confirmation_number;
	
	/**
	 * Party placing the order or paying the invoice.
	 * see : https://schema.org/customer
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $customer;
	
	/**
	 * The minimum payment required at this time.
	 * see : https://schema.org/minimumPaymentDue
	 * @var \MonetaryAmount|\MonetaryAmount[]|\PriceSpecification|\PriceSpecification[]
	 */
	public var $minimum_payment_due;
	
	/**
	 * The date that payment is due. Supersedes paymentDue (see: https://schema.org/paymentDue).
	 * see : https://schema.org/paymentDueDate
	 * @var string|string[]
	 */
	public var $payment_due_date;
	
	/**
	 * The name of the credit card or other method of payment for the order.
	 * see : https://schema.org/paymentMethod
	 * @var \PaymentMethod|\PaymentMethod[]
	 */
	public var $payment_method;
	
	/**
	 * An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	 * see : https://schema.org/paymentMethodId
	 * @var string|string[]
	 */
	public var $payment_method_id;
	
	/**
	 * The status of payment; whether the invoice has been paid or not.
	 * see : https://schema.org/paymentStatus
	 * @var \PaymentStatusType|\PaymentStatusType[]|string|string[]
	 */
	public var $payment_status;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	/**
	 * The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice.
	 * see : https://schema.org/referencesOrder
	 * @var \Order|\Order[]
	 */
	public var $references_order;
	
	/**
	 * The date the invoice is scheduled to be paid.
	 * see : https://schema.org/scheduledPaymentDate
	 * @var string|string[]
	 */
	public var $scheduled_payment_date;
	
	/**
	 * The total amount due.
	 * see : https://schema.org/totalPaymentDue
	 * @var \MonetaryAmount|\MonetaryAmount[]|\PriceSpecification|\PriceSpecification[]
	 */
	public var $total_payment_due;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Invoice'
		);
		
		$serialized = so_json_serialize( $this->account_id );
		if ( ! empty( $serialized ) ) {
			$out['accountId'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->billing_period );
		if ( ! empty( $serialized ) ) {
			$out['billingPeriod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->confirmation_number );
		if ( ! empty( $serialized ) ) {
			$out['confirmationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->customer );
		if ( ! empty( $serialized ) ) {
			$out['customer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->minimum_payment_due );
		if ( ! empty( $serialized ) ) {
			$out['minimumPaymentDue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payment_due_date );
		if ( ! empty( $serialized ) ) {
			$out['paymentDueDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payment_method );
		if ( ! empty( $serialized ) ) {
			$out['paymentMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payment_method_id );
		if ( ! empty( $serialized ) ) {
			$out['paymentMethodId'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payment_status );
		if ( ! empty( $serialized ) ) {
			$out['paymentStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->references_order );
		if ( ! empty( $serialized ) ) {
			$out['referencesOrder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->scheduled_payment_date );
		if ( ! empty( $serialized ) ) {
			$out['scheduledPaymentDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->total_payment_due );
		if ( ! empty( $serialized ) ) {
			$out['totalPaymentDue'] = $serialized;
		}
		
		return $out;
	}
}
