package schemaorggo

import "encoding/json"

// Vehicle see : https://schema.org/Vehicle
type Vehicle struct {
	Product

	typeContext

	// AccelerationTime see : https://auto.schema.org/accelerationTime
	// The time needed to accelerate the vehicle from a given start velocity to a given target velocity.
	//
	// Typical unit code(s): SEC for seconds
	//
	//
	// Note: There are unfortunately no standard unit codes for seconds/0..100 km/h or seconds/0..60 mph. Simply use &quot;SEC&quot; for seconds and indicate the velocities in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue), or use valueReference (see: https://schema.org/valueReference) with a QuantitativeValue (see: https://schema.org/QuantitativeValue) of 0..60 mph or 0..100 km/h to specify the reference speeds.
	//
	//
	// types : QuantitativeValue
	AccelerationTime *QuantitativeValue `json:"accelerationTime,omitempty"`

	// BodyType see : https://auto.schema.org/bodyType
	// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	// types : QualitativeValue Text URL
	BodyType interface{} `json:"bodyType,omitempty"`

	// CargoVolume see : https://schema.org/cargoVolume
	// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
	//
	// Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
	//
	// Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	// types : QuantitativeValue
	CargoVolume *QuantitativeValue `json:"cargoVolume,omitempty"`

	// DateVehicleFirstRegistered see : https://schema.org/dateVehicleFirstRegistered
	// The date of the first registration of the vehicle with the respective public authorities.
	// types : Date
	DateVehicleFirstRegistered Date `json:"dateVehicleFirstRegistered,omitempty"`

	// DriveWheelConfiguration see : https://schema.org/driveWheelConfiguration
	// The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	// types : DriveWheelConfigurationValue Text
	DriveWheelConfiguration interface{} `json:"driveWheelConfiguration,omitempty"`

	// EmissionsCO2 see : https://auto.schema.org/emissionsCO2
	// The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put &quot;g/km&quot; into the unitText property of that value, since there is no UN/CEFACT Common Code for &quot;g/km&quot;.
	// types : Number
	EmissionsCO2 float64 `json:"emissionsCO2,omitempty"`

	// FuelCapacity see : https://auto.schema.org/fuelCapacity
	// The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.
	//
	// Typical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
	// types : QuantitativeValue
	FuelCapacity *QuantitativeValue `json:"fuelCapacity,omitempty"`

	// FuelConsumption see : https://schema.org/fuelConsumption
	// The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).
	//
	//
	// Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. L/100 km.
	// Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
	// Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel consumption to another value.
	//
	//
	// types : QuantitativeValue
	FuelConsumption *QuantitativeValue `json:"fuelConsumption,omitempty"`

	// FuelEfficiency see : https://schema.org/fuelEfficiency
	// The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).
	//
	//
	// Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. mpg or km/L.
	// Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
	// Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel economy to another value.
	//
	//
	// types : QuantitativeValue
	FuelEfficiency *QuantitativeValue `json:"fuelEfficiency,omitempty"`

	// FuelType see : https://schema.org/fuelType
	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	// types : QualitativeValue Text URL
	FuelType interface{} `json:"fuelType,omitempty"`

	// KnownVehicleDamages see : https://schema.org/knownVehicleDamages
	// A textual description of known damages, both repaired and unrepaired.
	// types : Text
	KnownVehicleDamages string `json:"knownVehicleDamages,omitempty"`

	// MeetsEmissionStandard see : https://auto.schema.org/meetsEmissionStandard
	// Indicates that the vehicle meets the respective emission standard.
	// types : QualitativeValue Text URL
	MeetsEmissionStandard interface{} `json:"meetsEmissionStandard,omitempty"`

	// MileageFromOdometer see : https://schema.org/mileageFromOdometer
	// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
	//
	// Typical unit code(s): KMT for kilometers, SMI for statute miles
	// types : QuantitativeValue
	MileageFromOdometer *QuantitativeValue `json:"mileageFromOdometer,omitempty"`

	// ModelDate see : https://auto.schema.org/modelDate
	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	// types : Date
	ModelDate Date `json:"modelDate,omitempty"`

	// NumberOfAirbags see : https://schema.org/numberOfAirbags
	// The number or type of airbags in the vehicle.
	// types : Number Text
	NumberOfAirbags interface{} `json:"numberOfAirbags,omitempty"`

	// NumberOfAxles see : https://schema.org/numberOfAxles
	// The number of axles.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfAxles interface{} `json:"numberOfAxles,omitempty"`

	// NumberOfDoors see : https://schema.org/numberOfDoors
	// The number of doors.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfDoors interface{} `json:"numberOfDoors,omitempty"`

	// NumberOfForwardGears see : https://schema.org/numberOfForwardGears
	// The total number of forward gears available for the transmission system of the vehicle.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfForwardGears interface{} `json:"numberOfForwardGears,omitempty"`

	// NumberOfPreviousOwners see : https://schema.org/numberOfPreviousOwners
	// The number of owners of the vehicle, including the current one.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfPreviousOwners interface{} `json:"numberOfPreviousOwners,omitempty"`

	// Payload see : https://auto.schema.org/payload
	// The permitted weight of passengers and cargo, EXCLUDING the weight of the empty vehicle.
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	//
	//
	// Note 1: Many databases specify the permitted TOTAL weight instead, which is the sum of weight (see: https://schema.org/weight) and payload (see: https://schema.org/payload)
	// Note 2: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// Note 3: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	// Note 4: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	Payload *QuantitativeValue `json:"payload,omitempty"`

	// ProductionDate see : https://schema.org/productionDate
	// The date of production of the item, e.g. vehicle.
	// types : Date
	ProductionDate Date `json:"productionDate,omitempty"`

	// PurchaseDate see : https://schema.org/purchaseDate
	// The date the item e.g. vehicle was purchased by the current owner.
	// types : Date
	PurchaseDate Date `json:"purchaseDate,omitempty"`

	// SeatingCapacity see : https://auto.schema.org/seatingCapacity
	// The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.
	//
	// Typical unit code(s): C62 for persons
	// types : Number QuantitativeValue
	SeatingCapacity interface{} `json:"seatingCapacity,omitempty"`

	// Speed see : https://auto.schema.org/speed
	// The speed range of the vehicle. If the vehicle is powered by an engine, the upper limit of the speed range (indicated by maxValue (see: https://schema.org/maxValue) should be the maximum speed achievable under regular conditions.
	//
	// Typical unit code(s): KMH for km/h, HM for mile per hour (0.447 04 m/s), KNT for knot
	//
	// *Note 1: Use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate the range. Typically, the minimal value is zero.
	// * Note 2: There are many different ways of measuring the speed range. You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	// types : QuantitativeValue
	Speed *QuantitativeValue `json:"speed,omitempty"`

	// SteeringPosition see : https://schema.org/steeringPosition
	// The position of the steering wheel or similar device (mostly for cars).
	// types : SteeringPositionValue
	SteeringPosition *SteeringPositionValue `json:"steeringPosition,omitempty"`

	// TongueWeight see : https://auto.schema.org/tongueWeight
	// The permitted vertical load (TWR) of a trailer attached to the vehicle. Also referred to as Tongue Load Rating (TLR) or Vertical Load Rating (VLR)
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	//
	//
	// Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	// Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	TongueWeight *QuantitativeValue `json:"tongueWeight,omitempty"`

	// TrailerWeight see : https://auto.schema.org/trailerWeight
	// The permitted weight of a trailer attached to the vehicle.
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	// * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	// * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	// types : QuantitativeValue
	TrailerWeight *QuantitativeValue `json:"trailerWeight,omitempty"`

	// VehicleConfiguration see : https://schema.org/vehicleConfiguration
	// A short text indicating the configuration of the vehicle, e.g. &#39;5dr hatchback ST 2.5 MT 225 hp&#39; or &#39;limited edition&#39;.
	// types : Text
	VehicleConfiguration string `json:"vehicleConfiguration,omitempty"`

	// VehicleEngine see : https://schema.org/vehicleEngine
	// Information about the engine or engines of the vehicle.
	// types : EngineSpecification
	VehicleEngine *EngineSpecification `json:"vehicleEngine,omitempty"`

	// VehicleIdentificationNumber see : https://schema.org/vehicleIdentificationNumber
	// The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	// types : Text
	VehicleIdentificationNumber string `json:"vehicleIdentificationNumber,omitempty"`

	// VehicleInteriorColor see : https://schema.org/vehicleInteriorColor
	// The color or color combination of the interior of the vehicle.
	// types : Text
	VehicleInteriorColor string `json:"vehicleInteriorColor,omitempty"`

	// VehicleInteriorType see : https://schema.org/vehicleInteriorType
	// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	// types : Text
	VehicleInteriorType string `json:"vehicleInteriorType,omitempty"`

	// VehicleModelDate see : https://schema.org/vehicleModelDate
	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	// types : Date
	VehicleModelDate Date `json:"vehicleModelDate,omitempty"`

	// VehicleSeatingCapacity see : https://schema.org/vehicleSeatingCapacity
	// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
	//
	// Typical unit code(s): C62 for persons.
	// types : Number QuantitativeValue
	VehicleSeatingCapacity interface{} `json:"vehicleSeatingCapacity,omitempty"`

	// VehicleSpecialUsage see : https://auto.schema.org/vehicleSpecialUsage
	// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	// types : CarUsageType Text
	VehicleSpecialUsage interface{} `json:"vehicleSpecialUsage,omitempty"`

	// VehicleTransmission see : https://schema.org/vehicleTransmission
	// The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	// types : QualitativeValue Text URL
	VehicleTransmission interface{} `json:"vehicleTransmission,omitempty"`

	// WeightTotal see : https://auto.schema.org/weightTotal
	// The permitted total weight of the loaded vehicle, including passengers and cargo and the weight of the empty vehicle.
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	//
	//
	// Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	// Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	WeightTotal *QuantitativeValue `json:"weightTotal,omitempty"`

	// Wheelbase see : https://auto.schema.org/wheelbase
	// The distance between the centers of the front and rear wheels.
	//
	// Typical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
	// types : QuantitativeValue
	Wheelbase *QuantitativeValue `json:"wheelbase,omitempty"`
}

func (v Vehicle) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Vehicle"

	return json.Marshal(v)
}

func (v *Vehicle) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Vehicle"

	return json.Marshal(*v)
}
