<?php

class Order extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Order';
	
	/**
	 * The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	 * see : https://schema.org/acceptedOffer
	 * @var \Offer|\Offer[]
	 */
	public var $accepted_offer;
	
	/**
	 * The billing address for the order.
	 * see : https://schema.org/billingAddress
	 * @var \PostalAddress|\PostalAddress[]
	 */
	public var $billing_address;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $broker;
	
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
	 * Any discount applied (to an Order).
	 * see : https://schema.org/discount
	 * @var float|float[]|string|string[]
	 */
	public var $discount;
	
	/**
	 * Code used to redeem a discount.
	 * see : https://schema.org/discountCode
	 * @var string|string[]
	 */
	public var $discount_code;
	
	/**
	 * The currency of the discount.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://schema.org/discountCurrency
	 * @var string|string[]
	 */
	public var $discount_currency;
	
	/**
	 * Was the offer accepted as a gift for someone other than the buyer.
	 * see : https://schema.org/isGift
	 * @var boolean|boolean[]
	 */
	public var $is_gift;
	
	/**
	 * Date order was placed.
	 * see : https://schema.org/orderDate
	 * @var string|string[]
	 */
	public var $order_date;
	
	/**
	 * The delivery of the parcel related to this order or order item.
	 * see : https://schema.org/orderDelivery
	 * @var \ParcelDelivery|\ParcelDelivery[]
	 */
	public var $order_delivery;
	
	/**
	 * The identifier of the transaction.
	 * see : https://schema.org/orderNumber
	 * @var string|string[]
	 */
	public var $order_number;
	
	/**
	 * The current status of the order.
	 * see : https://schema.org/orderStatus
	 * @var \OrderStatus|\OrderStatus[]
	 */
	public var $order_status;
	
	/**
	 * The item ordered.
	 * see : https://schema.org/orderedItem
	 * @var \OrderItem|\OrderItem[]|\Product|\Product[]
	 */
	public var $ordered_item;
	
	/**
	 * The order is being paid as part of the referenced Invoice.
	 * see : https://schema.org/partOfInvoice
	 * @var \Invoice|\Invoice[]
	 */
	public var $part_of_invoice;
	
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
	 * The URL for sending a payment.
	 * see : https://schema.org/paymentUrl
	 * @var string|string[]
	 */
	public var $payment_url;
	
	/**
	 * An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	 * see : https://schema.org/seller
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $seller;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Order'
		);
		
		$serialized = so_json_serialize( $this->accepted_offer );
		if ( ! empty( $serialized ) ) {
			$out['acceptedOffer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->billing_address );
		if ( ! empty( $serialized ) ) {
			$out['billingAddress'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->confirmation_number );
		if ( ! empty( $serialized ) ) {
			$out['confirmationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->customer );
		if ( ! empty( $serialized ) ) {
			$out['customer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->discount );
		if ( ! empty( $serialized ) ) {
			$out['discount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->discount_code );
		if ( ! empty( $serialized ) ) {
			$out['discountCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->discount_currency );
		if ( ! empty( $serialized ) ) {
			$out['discountCurrency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_gift );
		if ( ! empty( $serialized ) ) {
			$out['isGift'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_date );
		if ( ! empty( $serialized ) ) {
			$out['orderDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_delivery );
		if ( ! empty( $serialized ) ) {
			$out['orderDelivery'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_number );
		if ( ! empty( $serialized ) ) {
			$out['orderNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->order_status );
		if ( ! empty( $serialized ) ) {
			$out['orderStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ordered_item );
		if ( ! empty( $serialized ) ) {
			$out['orderedItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->part_of_invoice );
		if ( ! empty( $serialized ) ) {
			$out['partOfInvoice'] = $serialized;
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
		
		$serialized = so_json_serialize( $this->payment_url );
		if ( ! empty( $serialized ) ) {
			$out['paymentUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}
		
		return $out;
	}
}
