<?php
namespace SchemaOrg;

// Order see : https://schema.org/Order
class Order implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Order';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	 * see : https://schema.org/acceptedOffer
	 *
	 * @var \Offer | \Offer[]
	 */
	public $accepted_offer;

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
	 * The billing address for the order.
	 * see : https://schema.org/billingAddress
	 *
	 * @var \PostalAddress | \PostalAddress[]
	 */
	public $billing_address;

	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	 * see : https://schema.org/broker
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $broker;

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
	 * Any discount applied (to an Order).
	 * see : https://schema.org/discount
	 *
	 * @var float | float[] | string | string[]
	 */
	public $discount;

	/**
	 * Code used to redeem a discount.
	 * see : https://schema.org/discountCode
	 *
	 * @var string | string[]
	 */
	public $discount_code;

	/**
	 * The currency (in 3-letter ISO 4217 format) of the discount.
	 * see : https://schema.org/discountCurrency
	 *
	 * @var string | string[]
	 */
	public $discount_currency;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Was the offer accepted as a gift for someone other than the buyer.
	 * see : https://schema.org/isGift
	 *
	 * @var boolean | boolean[]
	 */
	public $is_gift;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * &#39;merchant&#39; is an out-dated term for &#39;seller&#39;.
	 * see : https://schema.org/merchant
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $merchant;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * Date order was placed.
	 * see : https://schema.org/orderDate
	 *
	 * @var string | string[]
	 */
	public $order_date;

	/**
	 * The delivery of the parcel related to this order or order item.
	 * see : https://schema.org/orderDelivery
	 *
	 * @var \ParcelDelivery | \ParcelDelivery[]
	 */
	public $order_delivery;

	/**
	 * The identifier of the transaction.
	 * see : https://schema.org/orderNumber
	 *
	 * @var string | string[]
	 */
	public $order_number;

	/**
	 * The current status of the order.
	 * see : https://schema.org/orderStatus
	 *
	 * @var \OrderStatus | \OrderStatus[]
	 */
	public $order_status;

	/**
	 * The item ordered.
	 * see : https://schema.org/orderedItem
	 *
	 * @var \Product | \Product[] | \OrderItem | \OrderItem[]
	 */
	public $ordered_item;

	/**
	 * The order is being paid as part of the referenced Invoice.
	 * see : https://schema.org/partOfInvoice
	 *
	 * @var \Invoice | \Invoice[]
	 */
	public $part_of_invoice;

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
	 * The URL for sending a payment.
	 * see : https://schema.org/paymentUrl
	 *
	 * @var string | string[]
	 */
	public $payment_url;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	 * see : https://schema.org/seller
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $seller;

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
			'@type'    => 'Order',
		);

		$serialized = \SchemaOrg\json_serialize( $this->accepted_offer );
		if ( ! empty( $serialized ) ) {
			$out['acceptedOffer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->billing_address );
		if ( ! empty( $serialized ) ) {
			$out['billingAddress'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->discount );
		if ( ! empty( $serialized ) ) {
			$out['discount'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->discount_code );
		if ( ! empty( $serialized ) ) {
			$out['discountCode'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->discount_currency );
		if ( ! empty( $serialized ) ) {
			$out['discountCurrency'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_gift );
		if ( ! empty( $serialized ) ) {
			$out['isGift'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->merchant );
		if ( ! empty( $serialized ) ) {
			$out['merchant'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->order_date );
		if ( ! empty( $serialized ) ) {
			$out['orderDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->order_delivery );
		if ( ! empty( $serialized ) ) {
			$out['orderDelivery'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->order_number );
		if ( ! empty( $serialized ) ) {
			$out['orderNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->order_status );
		if ( ! empty( $serialized ) ) {
			$out['orderStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->ordered_item );
		if ( ! empty( $serialized ) ) {
			$out['orderedItem'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->part_of_invoice );
		if ( ! empty( $serialized ) ) {
			$out['partOfInvoice'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->payment_url );
		if ( ! empty( $serialized ) ) {
			$out['paymentUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
