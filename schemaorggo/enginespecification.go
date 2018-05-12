package schemaorggo

import "encoding/json"

// EngineSpecification see : https://schema.org/EngineSpecification
type EngineSpecification struct {
	StructuredValue

	typeContext

	// EngineDisplacement see : http://auto.schema.org/engineDisplacement
	// The volume swept by all of the pistons inside the cylinders of an internal combustion engine in a single movement.
	//
	// Typical unit code(s): CMQ for cubic centimeter, LTR for liters, INQ for cubic inches
	// * Note 1: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	// * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	// types : QuantitativeValue
	EngineDisplacement *QuantitativeValue `json:"engineDisplacement,omitempty"`

	// EnginePower see : http://auto.schema.org/enginePower
	// The power of the vehicle&#39;s engine.
	//     Typical unit code(s): KWT for kilowatt, BHP for brake horsepower, N12 for metric horsepower (PS, with 1 PS = 735,49875 W)
	//
	//
	// Note 1: There are many different ways of measuring an engine&#39;s power. For an overview, see  http://en.wikipedia.org/wiki/Horsepower#Engine (see: https://schema.orghttp://en.wikipedia.org/wiki/Horsepower#Engine_power_test_codes)powertest_codes.
	// Note 2: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	// Note 3: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	EnginePower *QuantitativeValue `json:"enginePower,omitempty"`

	// EngineType see : http://auto.schema.org/engineType
	// The type of engine or engines powering the vehicle.
	// types : QualitativeValue Text URL
	EngineType interface{} `json:"engineType,omitempty"`

	// FuelType see : https://schema.org/fuelType
	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	// types : QualitativeValue Text URL
	FuelType interface{} `json:"fuelType,omitempty"`

	// Torque see : http://auto.schema.org/torque
	// The torque (turning force) of the vehicle&#39;s engine.
	//
	// Typical unit code(s): NU for newton metre (N m), F17 for pound-force per foot, or F48 for pound-force per inch
	//
	//
	// Note 1: You can link to information about how the given value has been determined (e.g. reference RPM) using the valueReference (see: https://schema.org/valueReference) property.
	// Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	//
	//
	// types : QuantitativeValue
	Torque *QuantitativeValue `json:"torque,omitempty"`
}

func (v EngineSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EngineSpecification"

	return json.Marshal(v)
}

func (v *EngineSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EngineSpecification"

	return json.Marshal(*v)
}
