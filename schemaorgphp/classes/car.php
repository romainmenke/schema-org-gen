<?php

namespace SchemaOrg;

// Car see : https://schema.org/Car
class Car implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Car';
	
	/**
	 * With properties from Product see : https://schema.org/Product
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	/**
	 * With properties from Vehicle see : https://schema.org/Vehicle
	 */
	
	
	/**
	 * The time needed to accelerate the vehicle from a given start velocity to a given target velocity.
	 * 
	 * Typical unit code(s): SEC for seconds
	 * 
	 * 
	 * Note: There are unfortunately no standard unit codes for seconds/0..100 km/h or seconds/0..60 mph. Simply use &quot;SEC&quot; for seconds and indicate the velocities in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue), or use valueReference (see: https://schema.org/valueReference) with a QuantitativeValue (see: https://schema.org/QuantitativeValue) of 0..60 mph or 0..100 km/h to specify the reference speeds.
	 * 
	 * 
	 * see : https://auto.schema.org/accelerationTime
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $acceleration_time;
	
	/**
	 * The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
	 * see : https://auto.schema.org/acrissCode
	 * @var string | string[]
	 */
	public var $acriss_code;
	
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
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience | \Audience[]
	 */
	public var $audience;
	
	/**
	 * An award won by or for this item.
	 * see : https://schema.org/award
	 * @var string | string[]
	 */
	public var $award;
	
	/**
	 * Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	 * see : https://auto.schema.org/bodyType
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public var $body_type;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public var $brand;
	
	/**
	 * The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
	 * 
	 * Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
	 * 
	 * Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * see : https://schema.org/cargoVolume
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $cargo_volume;
	
	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://pending.schema.org/category
	 * @var \PhysicalActivityCategory | \PhysicalActivityCategory[] | string | string[] | \Thing | \Thing[]
	 */
	public var $category;
	
	/**
	 * The color of the product.
	 * see : https://schema.org/color
	 * @var string | string[]
	 */
	public var $color;
	
	/**
	 * The date of the first registration of the vehicle with the respective public authorities.
	 * see : https://schema.org/dateVehicleFirstRegistered
	 * @var string | string[]
	 */
	public var $date_vehicle_first_registered;
	
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
	 * The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	 * see : https://schema.org/driveWheelConfiguration
	 * @var \DriveWheelConfigurationValue | \DriveWheelConfigurationValue[] | string | string[]
	 */
	public var $drive_wheel_configuration;
	
	/**
	 * The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put &quot;g/km&quot; into the unitText property of that value, since there is no UN/CEFACT Common Code for &quot;g/km&quot;.
	 * see : https://auto.schema.org/emissionsCO2
	 * @var float | float[]
	 */
	public var $emissionsc_o_2;
	
	/**
	 * The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.
	 * 
	 * Typical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
	 * see : https://auto.schema.org/fuelCapacity
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $fuel_capacity;
	
	/**
	 * The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).
	 * 
	 * 
	 * Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. L/100 km.
	 * Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
	 * Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel consumption to another value.
	 * 
	 * 
	 * see : https://schema.org/fuelConsumption
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $fuel_consumption;
	
	/**
	 * The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).
	 * 
	 * 
	 * Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. mpg or km/L.
	 * Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
	 * Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel economy to another value.
	 * 
	 * 
	 * see : https://schema.org/fuelEfficiency
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $fuel_efficiency;
	
	/**
	 * The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	 * see : https://schema.org/fuelType
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public var $fuel_type;
	
	/**
	 * The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	 * see : https://schema.org/gtin8
	 * @var string | string[]
	 */
	public var $gtin_8;
	
	/**
	 * The height of the item.
	 * see : https://schema.org/height
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $height;
	
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
	 * A pointer to another product (or multiple products) for which this product is a consumable.
	 * see : https://schema.org/isConsumableFor
	 * @var \Product | \Product[]
	 */
	public var $is_consumable_for;
	
	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public var $is_similar_to;
	
	/**
	 * A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	 * see : https://schema.org/itemCondition
	 * @var \OfferItemCondition | \OfferItemCondition[]
	 */
	public var $item_condition;
	
	/**
	 * A textual description of known damages, both repaired and unrepaired.
	 * see : https://schema.org/knownVehicleDamages
	 * @var string | string[]
	 */
	public var $known_vehicle_damages;
	
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
	 * The manufacturer of the product.
	 * see : https://schema.org/manufacturer
	 * @var \Organization | \Organization[]
	 */
	public var $manufacturer;
	
	/**
	 * A material that something is made from, e.g. leather, wool, cotton, paper.
	 * see : https://schema.org/material
	 * @var \Product | \Product[] | string | string[]
	 */
	public var $material;
	
	/**
	 * Indicates that the vehicle meets the respective emission standard.
	 * see : https://auto.schema.org/meetsEmissionStandard
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public var $meets_emission_standard;
	
	/**
	 * The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
	 * 
	 * Typical unit code(s): KMT for kilometers, SMI for statute miles
	 * see : https://schema.org/mileageFromOdometer
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $mileage_from_odometer;
	
	/**
	 * The release date of a vehicle model (often used to differentiate versions of the same make and model).
	 * see : https://auto.schema.org/modelDate
	 * @var string | string[]
	 */
	public var $model_date;
	
	/**
	 * The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	 * see : https://schema.org/mpn
	 * @var string | string[]
	 */
	public var $mpn;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * The number or type of airbags in the vehicle.
	 * see : https://schema.org/numberOfAirbags
	 * @var float | float[] | string | string[]
	 */
	public var $number_of_airbags;
	
	/**
	 * The number of axles.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfAxles
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $number_of_axles;
	
	/**
	 * The number of doors.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfDoors
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $number_of_doors;
	
	/**
	 * The total number of forward gears available for the transmission system of the vehicle.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfForwardGears
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $number_of_forward_gears;
	
	/**
	 * The number of owners of the vehicle, including the current one.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfPreviousOwners
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $number_of_previous_owners;
	
	/**
	 * The permitted weight of passengers and cargo, EXCLUDING the weight of the empty vehicle.
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * 
	 * 
	 * Note 1: Many databases specify the permitted TOTAL weight instead, which is the sum of weight (see: https://schema.org/weight) and payload (see: https://schema.org/payload)
	 * Note 2: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * Note 3: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	 * Note 4: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/payload
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $payload;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The product identifier, such as ISBN. For example: meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot;.
	 * see : https://schema.org/productID
	 * @var string | string[]
	 */
	public var $producti_d;
	
	/**
	 * The date of production of the item, e.g. vehicle.
	 * see : https://schema.org/productionDate
	 * @var string | string[]
	 */
	public var $production_date;
	
	/**
	 * The date the item e.g. vehicle was purchased by the current owner.
	 * see : https://schema.org/purchaseDate
	 * @var string | string[]
	 */
	public var $purchase_date;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review | \Review[]
	 */
	public var $review;
	
	/**
	 * The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * 
	 * 
	 * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference)
	 * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/roofLoad
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $roof_load;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.
	 * 
	 * Typical unit code(s): C62 for persons
	 * see : https://auto.schema.org/seatingCapacity
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $seating_capacity;
	
	/**
	 * The speed range of the vehicle. If the vehicle is powered by an engine, the upper limit of the speed range (indicated by maxValue (see: https://schema.org/maxValue) should be the maximum speed achievable under regular conditions.
	 * 
	 * Typical unit code(s): KMH for km/h, HM for mile per hour (0.447 04 m/s), KNT for knot
	 * 
	 * *Note 1: Use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate the range. Typically, the minimal value is zero.
	 * * Note 2: There are many different ways of measuring the speed range. You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	 * see : https://auto.schema.org/speed
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $speed;
	
	/**
	 * The position of the steering wheel or similar device (mostly for cars).
	 * see : https://schema.org/steeringPosition
	 * @var \SteeringPositionValue | \SteeringPositionValue[]
	 */
	public var $steering_position;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * The permitted vertical load (TWR) of a trailer attached to the vehicle. Also referred to as Tongue Load Rating (TLR) or Vertical Load Rating (VLR)
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * 
	 * 
	 * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	 * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/tongueWeight
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $tongue_weight;
	
	/**
	 * The permitted weight of a trailer attached to the vehicle.
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	 * * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * see : https://auto.schema.org/trailerWeight
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $trailer_weight;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * A short text indicating the configuration of the vehicle, e.g. &#39;5dr hatchback ST 2.5 MT 225 hp&#39; or &#39;limited edition&#39;.
	 * see : https://schema.org/vehicleConfiguration
	 * @var string | string[]
	 */
	public var $vehicle_configuration;
	
	/**
	 * Information about the engine or engines of the vehicle.
	 * see : https://schema.org/vehicleEngine
	 * @var \EngineSpecification | \EngineSpecification[]
	 */
	public var $vehicle_engine;
	
	/**
	 * The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	 * see : https://schema.org/vehicleIdentificationNumber
	 * @var string | string[]
	 */
	public var $vehicle_identification_number;
	
	/**
	 * The color or color combination of the interior of the vehicle.
	 * see : https://schema.org/vehicleInteriorColor
	 * @var string | string[]
	 */
	public var $vehicle_interior_color;
	
	/**
	 * The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	 * see : https://schema.org/vehicleInteriorType
	 * @var string | string[]
	 */
	public var $vehicle_interior_type;
	
	/**
	 * The release date of a vehicle model (often used to differentiate versions of the same make and model).
	 * see : https://schema.org/vehicleModelDate
	 * @var string | string[]
	 */
	public var $vehicle_model_date;
	
	/**
	 * The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
	 * 
	 * Typical unit code(s): C62 for persons.
	 * see : https://schema.org/vehicleSeatingCapacity
	 * @var float | float[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $vehicle_seating_capacity;
	
	/**
	 * Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	 * see : https://auto.schema.org/vehicleSpecialUsage
	 * @var \CarUsageType | \CarUsageType[] | string | string[]
	 */
	public var $vehicle_special_usage;
	
	/**
	 * The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	 * see : https://schema.org/vehicleTransmission
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public var $vehicle_transmission;
	
	/**
	 * The permitted total weight of the loaded vehicle, including passengers and cargo and the weight of the empty vehicle.
	 * 
	 * Typical unit code(s): KGM for kilogram, LBR for pound
	 * 
	 * 
	 * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	 * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	 * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/weightTotal
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $weight_total;
	
	/**
	 * The distance between the centers of the front and rear wheels.
	 * 
	 * Typical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
	 * see : https://auto.schema.org/wheelbase
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $wheelbase;
	
	/**
	 * The width of the item.
	 * see : https://schema.org/width
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $width;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Car'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->acceleration_time );
		if ( ! empty( $serialized ) ) {
			$out['accelerationTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->acriss_code );
		if ( ! empty( $serialized ) ) {
			$out['acrissCode'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->body_type );
		if ( ! empty( $serialized ) ) {
			$out['bodyType'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->emissionsc_o_2 );
		if ( ! empty( $serialized ) ) {
			$out['emissionsCO2'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fuel_capacity );
		if ( ! empty( $serialized ) ) {
			$out['fuelCapacity'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->is_consumable_for );
		if ( ! empty( $serialized ) ) {
			$out['isConsumableFor'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->meets_emission_standard );
		if ( ! empty( $serialized ) ) {
			$out['meetsEmissionStandard'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->mileage_from_odometer );
		if ( ! empty( $serialized ) ) {
			$out['mileageFromOdometer'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->model_date );
		if ( ! empty( $serialized ) ) {
			$out['modelDate'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->payload );
		if ( ! empty( $serialized ) ) {
			$out['payload'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->producti_d );
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
		
		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->roof_load );
		if ( ! empty( $serialized ) ) {
			$out['roofLoad'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->seating_capacity );
		if ( ! empty( $serialized ) ) {
			$out['seatingCapacity'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->speed );
		if ( ! empty( $serialized ) ) {
			$out['speed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->steering_position );
		if ( ! empty( $serialized ) ) {
			$out['steeringPosition'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->tongue_weight );
		if ( ! empty( $serialized ) ) {
			$out['tongueWeight'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->trailer_weight );
		if ( ! empty( $serialized ) ) {
			$out['trailerWeight'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->weight_total );
		if ( ! empty( $serialized ) ) {
			$out['weightTotal'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->wheelbase );
		if ( ! empty( $serialized ) ) {
			$out['wheelbase'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}
		
		return $out;
	}
}
