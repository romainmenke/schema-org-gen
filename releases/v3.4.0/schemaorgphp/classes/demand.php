<?php
namespace SchemaOrg;

// Demand see : https://schema.org/Demand
class Demand implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Demand';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * The payment method(s) accepted by seller for this offer.
	 * see : https://schema.org/acceptedPaymentMethod
	 *
	 * @var \LoanOrCredit | \LoanOrCredit[] | \PaymentMethod | \PaymentMethod[]
	 */
	public $accepted_payment_method;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * The amount of time that is required between accepting the offer and the actual usage of the resource or service.
	 * see : https://schema.org/advanceBookingRequirement
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $advance_booking_requirement;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The geographic area where a service or offered item is provided.
	 * see : https://schema.org/areaServed
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | string | string[]
	 */
	public $area_served;

	/**
	 * The availability of this item&amp;#x2014;for example In stock, Out of stock, Pre-order, etc.
	 * see : https://schema.org/availability
	 *
	 * @var \ItemAvailability | \ItemAvailability[]
	 */
	public $availability;

	/**
	 * The end of the availability of the product or service included in the offer.
	 * see : https://schema.org/availabilityEnds
	 *
	 * @var string | string[]
	 */
	public $availability_ends;

	/**
	 * The beginning of the availability of the product or service included in the offer.
	 * see : https://schema.org/availabilityStarts
	 *
	 * @var string | string[]
	 */
	public $availability_starts;

	/**
	 * The place(s) from which the offer can be obtained (e.g. store locations).
	 * see : https://schema.org/availableAtOrFrom
	 *
	 * @var \Place | \Place[]
	 */
	public $available_at_or_from;

	/**
	 * The delivery method(s) available for this offer.
	 * see : https://schema.org/availableDeliveryMethod
	 *
	 * @var \DeliveryMethod | \DeliveryMethod[]
	 */
	public $available_delivery_method;

	/**
	 * The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	 * see : https://schema.org/businessFunction
	 *
	 * @var \BusinessFunction | \BusinessFunction[]
	 */
	public $business_function;

	/**
	 * The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
	 * see : https://schema.org/deliveryLeadTime
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $delivery_lead_time;

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
	 * The type(s) of customers for which the given offer is valid.
	 * see : https://schema.org/eligibleCustomerType
	 *
	 * @var \BusinessEntityType | \BusinessEntityType[]
	 */
	public $eligible_customer_type;

	/**
	 * The duration for which the given offer is valid.
	 * see : https://schema.org/eligibleDuration
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $eligible_duration;

	/**
	 * The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	 * see : https://schema.org/eligibleQuantity
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $eligible_quantity;

	/**
	 * The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.\n\nSee also [[ineligibleRegion]].
	 *
	 * see : https://schema.org/eligibleRegion
	 *
	 * @var \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public $eligible_region;

	/**
	 * The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	 * see : https://schema.org/eligibleTransactionVolume
	 *
	 * @var \PriceSpecification | \PriceSpecification[]
	 */
	public $eligible_transaction_volume;

	/**
	 * The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin12
	 *
	 * @var string | string[]
	 */
	public $gtin_12;

	/**
	 * The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin13
	 *
	 * @var string | string[]
	 */
	public $gtin_13;

	/**
	 * The GTIN-14 code of the product, or the product to which the offer refers. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin14
	 *
	 * @var string | string[]
	 */
	public $gtin_14;

	/**
	 * The [GTIN-8](http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin8
	 *
	 * @var string | string[]
	 */
	public $gtin_8;

	/**
	 * The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	 *
	 * see : https://schema.org/identifier
	 *
	 * @var string | string[] | \PropertyValue | \PropertyValue[]
	 */
	public $identifier;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * This links to a node or nodes indicating the exact quantity of the products included in the offer.
	 * see : https://schema.org/includesObject
	 *
	 * @var \TypeAndQuantityNode | \TypeAndQuantityNode[]
	 */
	public $includes_object;

	/**
	 * The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.\n\nSee also [[eligibleRegion]].
	 *
	 * see : https://schema.org/ineligibleRegion
	 *
	 * @var \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public $ineligible_region;

	/**
	 * The current approximate inventory level for the item or items.
	 * see : https://schema.org/inventoryLevel
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $inventory_level;

	/**
	 * A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	 * see : https://schema.org/itemCondition
	 *
	 * @var \OfferItemCondition | \OfferItemCondition[]
	 */
	public $item_condition;

	/**
	 * The item being offered.
	 * see : https://schema.org/itemOffered
	 *
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public $item_offered;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	 * see : https://schema.org/mpn
	 *
	 * @var string | string[]
	 */
	public $mpn;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * One or more detailed price specifications, indicating the unit price and delivery or payment charges.
	 * see : https://schema.org/priceSpecification
	 *
	 * @var \PriceSpecification | \PriceSpecification[]
	 */
	public $price_specification;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
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
	 * The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	 * see : https://schema.org/serialNumber
	 *
	 * @var string | string[]
	 */
	public $serial_number;

	/**
	 * The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	 * see : https://schema.org/sku
	 *
	 * @var string | string[]
	 */
	public $sku;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 *
	 * @var string | string[]
	 */
	public $valid_from;

	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 *
	 * @var string | string[]
	 */
	public $valid_through;

	/**
	 * The warranty promise(s) included in the offer.
	 * see : https://schema.org/warranty
	 *
	 * @var \WarrantyPromise | \WarrantyPromise[]
	 */
	public $warranty;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Demand',
		);

		$serialized = \SchemaOrg\json_serialize( $this->accepted_payment_method );
		if ( ! empty( $serialized ) ) {
			$out['acceptedPaymentMethod'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->advance_booking_requirement );
		if ( ! empty( $serialized ) ) {
			$out['advanceBookingRequirement'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->availability );
		if ( ! empty( $serialized ) ) {
			$out['availability'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->availability_ends );
		if ( ! empty( $serialized ) ) {
			$out['availabilityEnds'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->availability_starts );
		if ( ! empty( $serialized ) ) {
			$out['availabilityStarts'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->available_at_or_from );
		if ( ! empty( $serialized ) ) {
			$out['availableAtOrFrom'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->available_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['availableDeliveryMethod'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->business_function );
		if ( ! empty( $serialized ) ) {
			$out['businessFunction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->delivery_lead_time );
		if ( ! empty( $serialized ) ) {
			$out['deliveryLeadTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_customer_type );
		if ( ! empty( $serialized ) ) {
			$out['eligibleCustomerType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_duration );
		if ( ! empty( $serialized ) ) {
			$out['eligibleDuration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_quantity );
		if ( ! empty( $serialized ) ) {
			$out['eligibleQuantity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_region );
		if ( ! empty( $serialized ) ) {
			$out['eligibleRegion'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->eligible_transaction_volume );
		if ( ! empty( $serialized ) ) {
			$out['eligibleTransactionVolume'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->gtin_12 );
		if ( ! empty( $serialized ) ) {
			$out['gtin12'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->gtin_13 );
		if ( ! empty( $serialized ) ) {
			$out['gtin13'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->gtin_14 );
		if ( ! empty( $serialized ) ) {
			$out['gtin14'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->gtin_8 );
		if ( ! empty( $serialized ) ) {
			$out['gtin8'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->includes_object );
		if ( ! empty( $serialized ) ) {
			$out['includesObject'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->ineligible_region );
		if ( ! empty( $serialized ) ) {
			$out['ineligibleRegion'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->inventory_level );
		if ( ! empty( $serialized ) ) {
			$out['inventoryLevel'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->item_condition );
		if ( ! empty( $serialized ) ) {
			$out['itemCondition'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->item_offered );
		if ( ! empty( $serialized ) ) {
			$out['itemOffered'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mpn );
		if ( ! empty( $serialized ) ) {
			$out['mpn'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->price_specification );
		if ( ! empty( $serialized ) ) {
			$out['priceSpecification'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->serial_number );
		if ( ! empty( $serialized ) ) {
			$out['serialNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sku );
		if ( ! empty( $serialized ) ) {
			$out['sku'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->warranty );
		if ( ! empty( $serialized ) ) {
			$out['warranty'] = $serialized;
		}

		return $out;
	}
}
