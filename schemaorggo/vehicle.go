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
	AccelerationTime []*QuantitativeValue `json:"accelerationTime,omitempty"`

	// BodyType see : https://auto.schema.org/bodyType
	// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
	// types : QualitativeValue Text URL
	BodyType []interface{} `json:"bodyType,omitempty"`

	// CargoVolume see : https://schema.org/cargoVolume
	// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
	//
	// Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
	//
	// Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	// types : QuantitativeValue
	CargoVolume []*QuantitativeValue `json:"cargoVolume,omitempty"`

	// DateVehicleFirstRegistered see : https://schema.org/dateVehicleFirstRegistered
	// The date of the first registration of the vehicle with the respective public authorities.
	// types : Date
	DateVehicleFirstRegistered []Date `json:"dateVehicleFirstRegistered,omitempty"`

	// DriveWheelConfiguration see : https://schema.org/driveWheelConfiguration
	// The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	// types : DriveWheelConfigurationValue Text
	DriveWheelConfiguration []interface{} `json:"driveWheelConfiguration,omitempty"`

	// EmissionsCO2 see : https://auto.schema.org/emissionsCO2
	// The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put &quot;g/km&quot; into the unitText property of that value, since there is no UN/CEFACT Common Code for &quot;g/km&quot;.
	// types : Number
	EmissionsCO2 []float64 `json:"emissionsCO2,omitempty"`

	// FuelCapacity see : https://auto.schema.org/fuelCapacity
	// The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.
	//
	// Typical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
	// types : QuantitativeValue
	FuelCapacity []*QuantitativeValue `json:"fuelCapacity,omitempty"`

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
	FuelConsumption []*QuantitativeValue `json:"fuelConsumption,omitempty"`

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
	FuelEfficiency []*QuantitativeValue `json:"fuelEfficiency,omitempty"`

	// FuelType see : https://schema.org/fuelType
	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	// types : QualitativeValue Text URL
	FuelType []interface{} `json:"fuelType,omitempty"`

	// KnownVehicleDamages see : https://schema.org/knownVehicleDamages
	// A textual description of known damages, both repaired and unrepaired.
	// types : Text
	KnownVehicleDamages []string `json:"knownVehicleDamages,omitempty"`

	// MeetsEmissionStandard see : https://auto.schema.org/meetsEmissionStandard
	// Indicates that the vehicle meets the respective emission standard.
	// types : QualitativeValue Text URL
	MeetsEmissionStandard []interface{} `json:"meetsEmissionStandard,omitempty"`

	// MileageFromOdometer see : https://schema.org/mileageFromOdometer
	// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
	//
	// Typical unit code(s): KMT for kilometers, SMI for statute miles
	// types : QuantitativeValue
	MileageFromOdometer []*QuantitativeValue `json:"mileageFromOdometer,omitempty"`

	// ModelDate see : https://auto.schema.org/modelDate
	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	// types : Date
	ModelDate []Date `json:"modelDate,omitempty"`

	// NumberOfAirbags see : https://schema.org/numberOfAirbags
	// The number or type of airbags in the vehicle.
	// types : Number Text
	NumberOfAirbags []interface{} `json:"numberOfAirbags,omitempty"`

	// NumberOfAxles see : https://schema.org/numberOfAxles
	// The number of axles.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfAxles []interface{} `json:"numberOfAxles,omitempty"`

	// NumberOfDoors see : https://schema.org/numberOfDoors
	// The number of doors.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfDoors []interface{} `json:"numberOfDoors,omitempty"`

	// NumberOfForwardGears see : https://schema.org/numberOfForwardGears
	// The total number of forward gears available for the transmission system of the vehicle.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfForwardGears []interface{} `json:"numberOfForwardGears,omitempty"`

	// NumberOfPreviousOwners see : https://schema.org/numberOfPreviousOwners
	// The number of owners of the vehicle, including the current one.
	//
	// Typical unit code(s): C62
	// types : Number QuantitativeValue
	NumberOfPreviousOwners []interface{} `json:"numberOfPreviousOwners,omitempty"`

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
	Payload []*QuantitativeValue `json:"payload,omitempty"`

	// ProductionDate see : https://schema.org/productionDate
	// The date of production of the item, e.g. vehicle.
	// types : Date
	ProductionDate []Date `json:"productionDate,omitempty"`

	// PurchaseDate see : https://schema.org/purchaseDate
	// The date the item e.g. vehicle was purchased by the current owner.
	// types : Date
	PurchaseDate []Date `json:"purchaseDate,omitempty"`

	// SeatingCapacity see : https://auto.schema.org/seatingCapacity
	// The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.
	//
	// Typical unit code(s): C62 for persons
	// types : Number QuantitativeValue
	SeatingCapacity []interface{} `json:"seatingCapacity,omitempty"`

	// Speed see : https://auto.schema.org/speed
	// The speed range of the vehicle. If the vehicle is powered by an engine, the upper limit of the speed range (indicated by maxValue (see: https://schema.org/maxValue) should be the maximum speed achievable under regular conditions.
	//
	// Typical unit code(s): KMH for km/h, HM for mile per hour (0.447 04 m/s), KNT for knot
	//
	// *Note 1: Use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate the range. Typically, the minimal value is zero.
	// * Note 2: There are many different ways of measuring the speed range. You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	// types : QuantitativeValue
	Speed []*QuantitativeValue `json:"speed,omitempty"`

	// SteeringPosition see : https://schema.org/steeringPosition
	// The position of the steering wheel or similar device (mostly for cars).
	// types : SteeringPositionValue
	SteeringPosition []*SteeringPositionValue `json:"steeringPosition,omitempty"`

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
	TongueWeight []*QuantitativeValue `json:"tongueWeight,omitempty"`

	// TrailerWeight see : https://auto.schema.org/trailerWeight
	// The permitted weight of a trailer attached to the vehicle.
	//
	// Typical unit code(s): KGM for kilogram, LBR for pound
	// * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
	// * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
	// * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	// types : QuantitativeValue
	TrailerWeight []*QuantitativeValue `json:"trailerWeight,omitempty"`

	// VehicleConfiguration see : https://schema.org/vehicleConfiguration
	// A short text indicating the configuration of the vehicle, e.g. &#39;5dr hatchback ST 2.5 MT 225 hp&#39; or &#39;limited edition&#39;.
	// types : Text
	VehicleConfiguration []string `json:"vehicleConfiguration,omitempty"`

	// VehicleEngine see : https://schema.org/vehicleEngine
	// Information about the engine or engines of the vehicle.
	// types : EngineSpecification
	VehicleEngine []*EngineSpecification `json:"vehicleEngine,omitempty"`

	// VehicleIdentificationNumber see : https://schema.org/vehicleIdentificationNumber
	// The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
	// types : Text
	VehicleIdentificationNumber []string `json:"vehicleIdentificationNumber,omitempty"`

	// VehicleInteriorColor see : https://schema.org/vehicleInteriorColor
	// The color or color combination of the interior of the vehicle.
	// types : Text
	VehicleInteriorColor []string `json:"vehicleInteriorColor,omitempty"`

	// VehicleInteriorType see : https://schema.org/vehicleInteriorType
	// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
	// types : Text
	VehicleInteriorType []string `json:"vehicleInteriorType,omitempty"`

	// VehicleModelDate see : https://schema.org/vehicleModelDate
	// The release date of a vehicle model (often used to differentiate versions of the same make and model).
	// types : Date
	VehicleModelDate []Date `json:"vehicleModelDate,omitempty"`

	// VehicleSeatingCapacity see : https://schema.org/vehicleSeatingCapacity
	// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
	//
	// Typical unit code(s): C62 for persons.
	// types : Number QuantitativeValue
	VehicleSeatingCapacity []interface{} `json:"vehicleSeatingCapacity,omitempty"`

	// VehicleSpecialUsage see : https://auto.schema.org/vehicleSpecialUsage
	// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	// types : CarUsageType Text
	VehicleSpecialUsage []interface{} `json:"vehicleSpecialUsage,omitempty"`

	// VehicleTransmission see : https://schema.org/vehicleTransmission
	// The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	// types : QualitativeValue Text URL
	VehicleTransmission []interface{} `json:"vehicleTransmission,omitempty"`

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
	WeightTotal []*QuantitativeValue `json:"weightTotal,omitempty"`

	// Wheelbase see : https://auto.schema.org/wheelbase
	// The distance between the centers of the front and rear wheels.
	//
	// Typical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
	// types : QuantitativeValue
	Wheelbase []*QuantitativeValue `json:"wheelbase,omitempty"`
}

func (v Vehicle) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Product.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AccelerationTime
		if len(v.AccelerationTime) == 1 {
			value = v.AccelerationTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accelerationTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BodyType
		if len(v.BodyType) == 1 {
			value = v.BodyType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bodyType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CargoVolume
		if len(v.CargoVolume) == 1 {
			value = v.CargoVolume[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cargoVolume"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateVehicleFirstRegistered
		if len(v.DateVehicleFirstRegistered) == 1 {
			value = v.DateVehicleFirstRegistered[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateVehicleFirstRegistered"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DriveWheelConfiguration
		if len(v.DriveWheelConfiguration) == 1 {
			value = v.DriveWheelConfiguration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["driveWheelConfiguration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EmissionsCO2
		if len(v.EmissionsCO2) == 1 {
			value = v.EmissionsCO2[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["emissionsCO2"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FuelCapacity
		if len(v.FuelCapacity) == 1 {
			value = v.FuelCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fuelCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FuelConsumption
		if len(v.FuelConsumption) == 1 {
			value = v.FuelConsumption[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fuelConsumption"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FuelEfficiency
		if len(v.FuelEfficiency) == 1 {
			value = v.FuelEfficiency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fuelEfficiency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FuelType
		if len(v.FuelType) == 1 {
			value = v.FuelType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fuelType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.KnownVehicleDamages
		if len(v.KnownVehicleDamages) == 1 {
			value = v.KnownVehicleDamages[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["knownVehicleDamages"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MeetsEmissionStandard
		if len(v.MeetsEmissionStandard) == 1 {
			value = v.MeetsEmissionStandard[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["meetsEmissionStandard"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MileageFromOdometer
		if len(v.MileageFromOdometer) == 1 {
			value = v.MileageFromOdometer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mileageFromOdometer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ModelDate
		if len(v.ModelDate) == 1 {
			value = v.ModelDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["modelDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfAirbags
		if len(v.NumberOfAirbags) == 1 {
			value = v.NumberOfAirbags[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfAirbags"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfAxles
		if len(v.NumberOfAxles) == 1 {
			value = v.NumberOfAxles[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfAxles"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfDoors
		if len(v.NumberOfDoors) == 1 {
			value = v.NumberOfDoors[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfDoors"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfForwardGears
		if len(v.NumberOfForwardGears) == 1 {
			value = v.NumberOfForwardGears[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfForwardGears"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfPreviousOwners
		if len(v.NumberOfPreviousOwners) == 1 {
			value = v.NumberOfPreviousOwners[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfPreviousOwners"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Payload
		if len(v.Payload) == 1 {
			value = v.Payload[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["payload"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductionDate
		if len(v.ProductionDate) == 1 {
			value = v.ProductionDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productionDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PurchaseDate
		if len(v.PurchaseDate) == 1 {
			value = v.PurchaseDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["purchaseDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SeatingCapacity
		if len(v.SeatingCapacity) == 1 {
			value = v.SeatingCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seatingCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Speed
		if len(v.Speed) == 1 {
			value = v.Speed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["speed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SteeringPosition
		if len(v.SteeringPosition) == 1 {
			value = v.SteeringPosition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["steeringPosition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TongueWeight
		if len(v.TongueWeight) == 1 {
			value = v.TongueWeight[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["tongueWeight"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TrailerWeight
		if len(v.TrailerWeight) == 1 {
			value = v.TrailerWeight[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trailerWeight"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleConfiguration
		if len(v.VehicleConfiguration) == 1 {
			value = v.VehicleConfiguration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleConfiguration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleEngine
		if len(v.VehicleEngine) == 1 {
			value = v.VehicleEngine[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleEngine"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleIdentificationNumber
		if len(v.VehicleIdentificationNumber) == 1 {
			value = v.VehicleIdentificationNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleIdentificationNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleInteriorColor
		if len(v.VehicleInteriorColor) == 1 {
			value = v.VehicleInteriorColor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleInteriorColor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleInteriorType
		if len(v.VehicleInteriorType) == 1 {
			value = v.VehicleInteriorType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleInteriorType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleModelDate
		if len(v.VehicleModelDate) == 1 {
			value = v.VehicleModelDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleModelDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleSeatingCapacity
		if len(v.VehicleSeatingCapacity) == 1 {
			value = v.VehicleSeatingCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleSeatingCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleSpecialUsage
		if len(v.VehicleSpecialUsage) == 1 {
			value = v.VehicleSpecialUsage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleSpecialUsage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VehicleTransmission
		if len(v.VehicleTransmission) == 1 {
			value = v.VehicleTransmission[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vehicleTransmission"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WeightTotal
		if len(v.WeightTotal) == 1 {
			value = v.WeightTotal[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["weightTotal"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Wheelbase
		if len(v.Wheelbase) == 1 {
			value = v.Wheelbase[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["wheelbase"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Vehicle) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Vehicle"

	return data, nil
}

func (v Vehicle) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
