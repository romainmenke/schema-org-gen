<?php

class Vehicle extends Product implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Vehicle';
	
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $acceleration_time;
	
	/**
	 * Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	 * see : https://auto.schema.org/bodyType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
	 */
	public var $body_type;
	
	/**
	 * The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
	 * 
	 * Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
	 * 
	 * Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * see : https://schema.org/cargoVolume
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $cargo_volume;
	
	/**
	 * The date of the first registration of the vehicle with the respective public authorities.
	 * see : https://schema.org/dateVehicleFirstRegistered
	 * @var string|string[]
	 */
	public var $date_vehicle_first_registered;
	
	/**
	 * The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	 * see : https://schema.org/driveWheelConfiguration
	 * @var \DriveWheelConfigurationValue|\DriveWheelConfigurationValue[]|string|string[]
	 */
	public var $drive_wheel_configuration;
	
	/**
	 * The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put &quot;g/km&quot; into the unitText property of that value, since there is no UN/CEFACT Common Code for &quot;g/km&quot;.
	 * see : https://auto.schema.org/emissionsCO2
	 * @var float|float[]
	 */
	public var $emissionsc_o_2;
	
	/**
	 * The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.
	 * 
	 * Typical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
	 * see : https://auto.schema.org/fuelCapacity
	 * @var \QuantitativeValue|\QuantitativeValue[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $fuel_efficiency;
	
	/**
	 * The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	 * see : https://schema.org/fuelType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
	 */
	public var $fuel_type;
	
	/**
	 * A textual description of known damages, both repaired and unrepaired.
	 * see : https://schema.org/knownVehicleDamages
	 * @var string|string[]
	 */
	public var $known_vehicle_damages;
	
	/**
	 * Indicates that the vehicle meets the respective emission standard.
	 * see : https://auto.schema.org/meetsEmissionStandard
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
	 */
	public var $meets_emission_standard;
	
	/**
	 * The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
	 * 
	 * Typical unit code(s): KMT for kilometers, SMI for statute miles
	 * see : https://schema.org/mileageFromOdometer
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $mileage_from_odometer;
	
	/**
	 * The release date of a vehicle model (often used to differentiate versions of the same make and model).
	 * see : https://auto.schema.org/modelDate
	 * @var string|string[]
	 */
	public var $model_date;
	
	/**
	 * The number or type of airbags in the vehicle.
	 * see : https://schema.org/numberOfAirbags
	 * @var float|float[]|string|string[]
	 */
	public var $number_of_airbags;
	
	/**
	 * The number of axles.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfAxles
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_axles;
	
	/**
	 * The number of doors.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfDoors
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_doors;
	
	/**
	 * The total number of forward gears available for the transmission system of the vehicle.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfForwardGears
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_forward_gears;
	
	/**
	 * The number of owners of the vehicle, including the current one.
	 * 
	 * Typical unit code(s): C62
	 * see : https://schema.org/numberOfPreviousOwners
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $payload;
	
	/**
	 * The date of production of the item, e.g. vehicle.
	 * see : https://schema.org/productionDate
	 * @var string|string[]
	 */
	public var $production_date;
	
	/**
	 * The date the item e.g. vehicle was purchased by the current owner.
	 * see : https://schema.org/purchaseDate
	 * @var string|string[]
	 */
	public var $purchase_date;
	
	/**
	 * The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.
	 * 
	 * Typical unit code(s): C62 for persons
	 * see : https://auto.schema.org/seatingCapacity
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $speed;
	
	/**
	 * The position of the steering wheel or similar device (mostly for cars).
	 * see : https://schema.org/steeringPosition
	 * @var \SteeringPositionValue|\SteeringPositionValue[]
	 */
	public var $steering_position;
	
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $trailer_weight;
	
	/**
	 * A short text indicating the configuration of the vehicle, e.g. &#39;5dr hatchback ST 2.5 MT 225 hp&#39; or &#39;limited edition&#39;.
	 * see : https://schema.org/vehicleConfiguration
	 * @var string|string[]
	 */
	public var $vehicle_configuration;
	
	/**
	 * Information about the engine or engines of the vehicle.
	 * see : https://schema.org/vehicleEngine
	 * @var \EngineSpecification|\EngineSpecification[]
	 */
	public var $vehicle_engine;
	
	/**
	 * The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	 * see : https://schema.org/vehicleIdentificationNumber
	 * @var string|string[]
	 */
	public var $vehicle_identification_number;
	
	/**
	 * The color or color combination of the interior of the vehicle.
	 * see : https://schema.org/vehicleInteriorColor
	 * @var string|string[]
	 */
	public var $vehicle_interior_color;
	
	/**
	 * The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	 * see : https://schema.org/vehicleInteriorType
	 * @var string|string[]
	 */
	public var $vehicle_interior_type;
	
	/**
	 * The release date of a vehicle model (often used to differentiate versions of the same make and model).
	 * see : https://schema.org/vehicleModelDate
	 * @var string|string[]
	 */
	public var $vehicle_model_date;
	
	/**
	 * The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
	 * 
	 * Typical unit code(s): C62 for persons.
	 * see : https://schema.org/vehicleSeatingCapacity
	 * @var float|float[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $vehicle_seating_capacity;
	
	/**
	 * Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	 * see : https://auto.schema.org/vehicleSpecialUsage
	 * @var \CarUsageType|\CarUsageType[]|string|string[]
	 */
	public var $vehicle_special_usage;
	
	/**
	 * The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	 * see : https://schema.org/vehicleTransmission
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]|string|string[]
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
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $weight_total;
	
	/**
	 * The distance between the centers of the front and rear wheels.
	 * 
	 * Typical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
	 * see : https://auto.schema.org/wheelbase
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $wheelbase;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Vehicle'
		);
		
		$serialized = so_json_serialize( $this->acceleration_time );
		if ( ! empty( $serialized ) ) {
			$out['accelerationTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->body_type );
		if ( ! empty( $serialized ) ) {
			$out['bodyType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cargo_volume );
		if ( ! empty( $serialized ) ) {
			$out['cargoVolume'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_vehicle_first_registered );
		if ( ! empty( $serialized ) ) {
			$out['dateVehicleFirstRegistered'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->drive_wheel_configuration );
		if ( ! empty( $serialized ) ) {
			$out['driveWheelConfiguration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->emissionsc_o_2 );
		if ( ! empty( $serialized ) ) {
			$out['emissionsCO2'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fuel_capacity );
		if ( ! empty( $serialized ) ) {
			$out['fuelCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fuel_consumption );
		if ( ! empty( $serialized ) ) {
			$out['fuelConsumption'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fuel_efficiency );
		if ( ! empty( $serialized ) ) {
			$out['fuelEfficiency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fuel_type );
		if ( ! empty( $serialized ) ) {
			$out['fuelType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->known_vehicle_damages );
		if ( ! empty( $serialized ) ) {
			$out['knownVehicleDamages'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->meets_emission_standard );
		if ( ! empty( $serialized ) ) {
			$out['meetsEmissionStandard'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->mileage_from_odometer );
		if ( ! empty( $serialized ) ) {
			$out['mileageFromOdometer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->model_date );
		if ( ! empty( $serialized ) ) {
			$out['modelDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_airbags );
		if ( ! empty( $serialized ) ) {
			$out['numberOfAirbags'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_axles );
		if ( ! empty( $serialized ) ) {
			$out['numberOfAxles'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_doors );
		if ( ! empty( $serialized ) ) {
			$out['numberOfDoors'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_forward_gears );
		if ( ! empty( $serialized ) ) {
			$out['numberOfForwardGears'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_previous_owners );
		if ( ! empty( $serialized ) ) {
			$out['numberOfPreviousOwners'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->payload );
		if ( ! empty( $serialized ) ) {
			$out['payload'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_date );
		if ( ! empty( $serialized ) ) {
			$out['productionDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->purchase_date );
		if ( ! empty( $serialized ) ) {
			$out['purchaseDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seating_capacity );
		if ( ! empty( $serialized ) ) {
			$out['seatingCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->speed );
		if ( ! empty( $serialized ) ) {
			$out['speed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->steering_position );
		if ( ! empty( $serialized ) ) {
			$out['steeringPosition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->tongue_weight );
		if ( ! empty( $serialized ) ) {
			$out['tongueWeight'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trailer_weight );
		if ( ! empty( $serialized ) ) {
			$out['trailerWeight'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_configuration );
		if ( ! empty( $serialized ) ) {
			$out['vehicleConfiguration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_engine );
		if ( ! empty( $serialized ) ) {
			$out['vehicleEngine'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_identification_number );
		if ( ! empty( $serialized ) ) {
			$out['vehicleIdentificationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_interior_color );
		if ( ! empty( $serialized ) ) {
			$out['vehicleInteriorColor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_interior_type );
		if ( ! empty( $serialized ) ) {
			$out['vehicleInteriorType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_model_date );
		if ( ! empty( $serialized ) ) {
			$out['vehicleModelDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_seating_capacity );
		if ( ! empty( $serialized ) ) {
			$out['vehicleSeatingCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_special_usage );
		if ( ! empty( $serialized ) ) {
			$out['vehicleSpecialUsage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vehicle_transmission );
		if ( ! empty( $serialized ) ) {
			$out['vehicleTransmission'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->weight_total );
		if ( ! empty( $serialized ) ) {
			$out['weightTotal'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->wheelbase );
		if ( ! empty( $serialized ) ) {
			$out['wheelbase'] = $serialized;
		}
		
		return $out;
	}
}
