<?php
namespace SchemaOrg;

// Vehicle see : https://schema.org/Vehicle
class Vehicle implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Vehicle';

	/**
	 * With properties from Product see : https://schema.org/Product
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 *
	 * see : https://schema.org/additionalProperty
	 *
	 * @var \PropertyValue | \PropertyValue[]
	 */
	public $additional_property;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 *
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public $aggregate_rating;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * An intended audience, i.e. a group for whom something was created.
	 * see : https://schema.org/audience
	 *
	 * @var \Audience | \Audience[]
	 */
	public $audience;

	/**
	 * An award won by or for this item.
	 * see : https://schema.org/award
	 *
	 * @var string | string[]
	 */
	public $award;

	/**
	 * Awards won by or for this item.
	 * see : https://schema.org/awards
	 *
	 * @var string | string[]
	 */
	public $awards;

	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 *
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public $brand;

	/**
	 * The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.\n\nTypical unit code(s): LTR for liters, FTQ for cubic foot/feet\n\nNote: You can use [[minValue]] and [[maxValue]] to indicate ranges.
	 * see : https://schema.org/cargoVolume
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $cargo_volume;

	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://schema.org/category
	 *
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public $category;

	/**
	 * The color of the product.
	 * see : https://schema.org/color
	 *
	 * @var string | string[]
	 */
	public $color;

	/**
	 * The date of the first registration of the vehicle with the respective public authorities.
	 * see : https://schema.org/dateVehicleFirstRegistered
	 *
	 * @var string | string[]
	 */
	public $date_vehicle_first_registered;

	/**
	 * The depth of the item.
	 * see : https://schema.org/depth
	 *
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $depth;

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
	 * The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	 * see : https://schema.org/driveWheelConfiguration
	 *
	 * @var \DriveWheelConfigurationValue | \DriveWheelConfigurationValue[] | string | string[]
	 */
	public $drive_wheel_configuration;

	/**
	 * The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).\n\n* Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use [[unitText]] to indicate the unit of measurement, e.g. L/100 km.\n* Note 2: There are two ways of indicating the fuel consumption, [[fuelConsumption]] (e.g. 8 liters per 100 km) and [[fuelEfficiency]] (e.g. 30 miles per gallon). They are reciprocal.\n* Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use [[valueReference]] to link the value for the fuel consumption to another value.
	 * see : https://schema.org/fuelConsumption
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $fuel_consumption;

	/**
	 * The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).\n\n* Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use [[unitText]] to indicate the unit of measurement, e.g. mpg or km/L.\n* Note 2: There are two ways of indicating the fuel consumption, [[fuelConsumption]] (e.g. 8 liters per 100 km) and [[fuelEfficiency]] (e.g. 30 miles per gallon). They are reciprocal.\n* Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use [[valueReference]] to link the value for the fuel economy to another value.
	 * see : https://schema.org/fuelEfficiency
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $fuel_efficiency;

	/**
	 * The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	 * see : https://schema.org/fuelType
	 *
	 * @var string | string[] | \QualitativeValue | \QualitativeValue[]
	 */
	public $fuel_type;

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
	 * The height of the item.
	 * see : https://schema.org/height
	 *
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $height;

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
	 * A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	 * see : https://schema.org/isAccessoryOrSparePartFor
	 *
	 * @var \Product | \Product[]
	 */
	public $is_accessory_or_spare_part_for;

	/**
	 * A pointer to another product (or multiple products) for which this product is a consumable.
	 * see : https://schema.org/isConsumableFor
	 *
	 * @var \Product | \Product[]
	 */
	public $is_consumable_for;

	/**
	 * A pointer to another, somehow related product (or multiple products).
	 * see : https://schema.org/isRelatedTo
	 *
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public $is_related_to;

	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 *
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public $is_similar_to;

	/**
	 * A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	 * see : https://schema.org/itemCondition
	 *
	 * @var \OfferItemCondition | \OfferItemCondition[]
	 */
	public $item_condition;

	/**
	 * A textual description of known damages, both repaired and unrepaired.
	 * see : https://schema.org/knownVehicleDamages
	 *
	 * @var string | string[]
	 */
	public $known_vehicle_damages;

	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 *
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $logo;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The manufacturer of the product.
	 * see : https://schema.org/manufacturer
	 *
	 * @var \Organization | \Organization[]
	 */
	public $manufacturer;

	/**
	 * A material that something is made from, e.g. leather, wool, cotton, paper.
	 * see : https://schema.org/material
	 *
	 * @var string | string[] | \Product | \Product[]
	 */
	public $material;

	/**
	 * The total distance travelled by the particular vehicle since its initial production, as read from its odometer.\n\nTypical unit code(s): KMT for kilometers, SMI for statute miles
	 * see : https://schema.org/mileageFromOdometer
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $mileage_from_odometer;

	/**
	 * The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	 * see : https://schema.org/model
	 *
	 * @var \ProductModel | \ProductModel[] | string | string[]
	 */
	public $model;

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
	 * The number or type of airbags in the vehicle.
	 * see : https://schema.org/numberOfAirbags
	 *
	 * @var float | float[] | string | string[]
	 */
	public $number_of_airbags;

	/**
	 * The number of axles.\n\nTypical unit code(s): C62
	 * see : https://schema.org/numberOfAxles
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[] | float | float[]
	 */
	public $number_of_axles;

	/**
	 * The number of doors.\n\nTypical unit code(s): C62
	 * see : https://schema.org/numberOfDoors
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[] | float | float[]
	 */
	public $number_of_doors;

	/**
	 * The total number of forward gears available for the transmission system of the vehicle.\n\nTypical unit code(s): C62
	 * see : https://schema.org/numberOfForwardGears
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[] | float | float[]
	 */
	public $number_of_forward_gears;

	/**
	 * The number of owners of the vehicle, including the current one.\n\nTypical unit code(s): C62
	 * see : https://schema.org/numberOfPreviousOwners
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[] | float | float[]
	 */
	public $number_of_previous_owners;

	/**
	 * An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 *
	 * @var \Offer | \Offer[]
	 */
	public $offers;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The product identifier, such as ISBN. For example: ``` meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot; ```.
	 * see : https://schema.org/productID
	 *
	 * @var string | string[]
	 */
	public $product_id;

	/**
	 * The date of production of the item, e.g. vehicle.
	 * see : https://schema.org/productionDate
	 *
	 * @var string | string[]
	 */
	public $production_date;

	/**
	 * The date the item e.g. vehicle was purchased by the current owner.
	 * see : https://schema.org/purchaseDate
	 *
	 * @var string | string[]
	 */
	public $purchase_date;

	/**
	 * The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	 * see : https://schema.org/releaseDate
	 *
	 * @var string | string[]
	 */
	public $release_date;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * Review of the item.
	 * see : https://schema.org/reviews
	 *
	 * @var \Review | \Review[]
	 */
	public $reviews;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	 * see : https://schema.org/sku
	 *
	 * @var string | string[]
	 */
	public $sku;

	/**
	 * A slogan or motto associated with the item.
	 * see : https://schema.org/slogan
	 *
	 * @var string | string[]
	 */
	public $slogan;

	/**
	 * The position of the steering wheel or similar device (mostly for cars).
	 * see : https://schema.org/steeringPosition
	 *
	 * @var \SteeringPositionValue | \SteeringPositionValue[]
	 */
	public $steering_position;

	/**
	 * A CreativeWork or Event about this Thing..
	 * see : https://schema.org/subjectOf
	 *
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * A short text indicating the configuration of the vehicle, e.g. &#39;5dr hatchback ST 2.5 MT 225 hp&#39; or &#39;limited edition&#39;.
	 * see : https://schema.org/vehicleConfiguration
	 *
	 * @var string | string[]
	 */
	public $vehicle_configuration;

	/**
	 * Information about the engine or engines of the vehicle.
	 * see : https://schema.org/vehicleEngine
	 *
	 * @var \EngineSpecification | \EngineSpecification[]
	 */
	public $vehicle_engine;

	/**
	 * The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	 * see : https://schema.org/vehicleIdentificationNumber
	 *
	 * @var string | string[]
	 */
	public $vehicle_identification_number;

	/**
	 * The color or color combination of the interior of the vehicle.
	 * see : https://schema.org/vehicleInteriorColor
	 *
	 * @var string | string[]
	 */
	public $vehicle_interior_color;

	/**
	 * The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	 * see : https://schema.org/vehicleInteriorType
	 *
	 * @var string | string[]
	 */
	public $vehicle_interior_type;

	/**
	 * The release date of a vehicle model (often used to differentiate versions of the same make and model).
	 * see : https://schema.org/vehicleModelDate
	 *
	 * @var string | string[]
	 */
	public $vehicle_model_date;

	/**
	 * The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.\n\nTypical unit code(s): C62 for persons.
	 * see : https://schema.org/vehicleSeatingCapacity
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[] | float | float[]
	 */
	public $vehicle_seating_capacity;

	/**
	 * Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	 * see : https://schema.org/vehicleSpecialUsage
	 *
	 * @var string | string[]
	 */
	public $vehicle_special_usage;

	/**
	 * The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	 * see : https://schema.org/vehicleTransmission
	 *
	 * @var string | string[] | \QualitativeValue | \QualitativeValue[]
	 */
	public $vehicle_transmission;

	/**
	 * The weight of the product or person.
	 * see : https://schema.org/weight
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $weight;

	/**
	 * The width of the item.
	 * see : https://schema.org/width
	 *
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $width;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Vehicle',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}

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

		$serialized = \SchemaOrg\json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->awards );
		if ( ! empty( $serialized ) ) {
			$out['awards'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->cargo_volume );
		if ( ! empty( $serialized ) ) {
			$out['cargoVolume'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->color );
		if ( ! empty( $serialized ) ) {
			$out['color'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_vehicle_first_registered );
		if ( ! empty( $serialized ) ) {
			$out['dateVehicleFirstRegistered'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->depth );
		if ( ! empty( $serialized ) ) {
			$out['depth'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->drive_wheel_configuration );
		if ( ! empty( $serialized ) ) {
			$out['driveWheelConfiguration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fuel_consumption );
		if ( ! empty( $serialized ) ) {
			$out['fuelConsumption'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fuel_efficiency );
		if ( ! empty( $serialized ) ) {
			$out['fuelEfficiency'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fuel_type );
		if ( ! empty( $serialized ) ) {
			$out['fuelType'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_accessory_or_spare_part_for );
		if ( ! empty( $serialized ) ) {
			$out['isAccessoryOrSparePartFor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_consumable_for );
		if ( ! empty( $serialized ) ) {
			$out['isConsumableFor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_related_to );
		if ( ! empty( $serialized ) ) {
			$out['isRelatedTo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_similar_to );
		if ( ! empty( $serialized ) ) {
			$out['isSimilarTo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->item_condition );
		if ( ! empty( $serialized ) ) {
			$out['itemCondition'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->known_vehicle_damages );
		if ( ! empty( $serialized ) ) {
			$out['knownVehicleDamages'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->manufacturer );
		if ( ! empty( $serialized ) ) {
			$out['manufacturer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->material );
		if ( ! empty( $serialized ) ) {
			$out['material'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mileage_from_odometer );
		if ( ! empty( $serialized ) ) {
			$out['mileageFromOdometer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->model );
		if ( ! empty( $serialized ) ) {
			$out['model'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mpn );
		if ( ! empty( $serialized ) ) {
			$out['mpn'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_airbags );
		if ( ! empty( $serialized ) ) {
			$out['numberOfAirbags'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_axles );
		if ( ! empty( $serialized ) ) {
			$out['numberOfAxles'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_doors );
		if ( ! empty( $serialized ) ) {
			$out['numberOfDoors'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_forward_gears );
		if ( ! empty( $serialized ) ) {
			$out['numberOfForwardGears'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_previous_owners );
		if ( ! empty( $serialized ) ) {
			$out['numberOfPreviousOwners'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->product_id );
		if ( ! empty( $serialized ) ) {
			$out['productID'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->production_date );
		if ( ! empty( $serialized ) ) {
			$out['productionDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->purchase_date );
		if ( ! empty( $serialized ) ) {
			$out['purchaseDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->release_date );
		if ( ! empty( $serialized ) ) {
			$out['releaseDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->reviews );
		if ( ! empty( $serialized ) ) {
			$out['reviews'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sku );
		if ( ! empty( $serialized ) ) {
			$out['sku'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->slogan );
		if ( ! empty( $serialized ) ) {
			$out['slogan'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->steering_position );
		if ( ! empty( $serialized ) ) {
			$out['steeringPosition'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_configuration );
		if ( ! empty( $serialized ) ) {
			$out['vehicleConfiguration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_engine );
		if ( ! empty( $serialized ) ) {
			$out['vehicleEngine'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_identification_number );
		if ( ! empty( $serialized ) ) {
			$out['vehicleIdentificationNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_interior_color );
		if ( ! empty( $serialized ) ) {
			$out['vehicleInteriorColor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_interior_type );
		if ( ! empty( $serialized ) ) {
			$out['vehicleInteriorType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_model_date );
		if ( ! empty( $serialized ) ) {
			$out['vehicleModelDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_seating_capacity );
		if ( ! empty( $serialized ) ) {
			$out['vehicleSeatingCapacity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_special_usage );
		if ( ! empty( $serialized ) ) {
			$out['vehicleSpecialUsage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vehicle_transmission );
		if ( ! empty( $serialized ) ) {
			$out['vehicleTransmission'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->weight );
		if ( ! empty( $serialized ) ) {
			$out['weight'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}

		return $out;
	}
}
