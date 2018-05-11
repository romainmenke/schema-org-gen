package schemaorg

// Vehicle see : https://schema.org/Vehicle
type Vehicle struct {

Product// AccelerationTime see : https://auto.schema.org/accelerationTime
// The time needed to accelerate the vehicle from a given start velocity to a given target velocity.
// 
// Typical unit code(s): SEC for seconds
// 
// 
// Note: There are unfortunately no standard unit codes for seconds/0..100 km/h or seconds/0..60 mph. Simply use "SEC" for seconds and indicate the velocities in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue), or use valueReference (see: https://schema.org/valueReference) with a QuantitativeValue (see: https://schema.org/QuantitativeValue) of 0..60 mph or 0..100 km/h to specify the reference speeds.
// 
// 
AccelerationTime string `json:"accelerationTime"`

// BodyType see : https://auto.schema.org/bodyType
// Indicates the design and body style of the vehicle (e.g. station wagon, hatchback, etc.).
BodyType interface{} `json:"bodyType"` // types : QualitativeValue Text URL

// CargoVolume see : /cargoVolume
// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
// 
// Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
// 
// Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
CargoVolume string `json:"cargoVolume"`

// DateVehicleFirstRegistered see : /dateVehicleFirstRegistered
// The date of the first registration of the vehicle with the respective public authorities.
DateVehicleFirstRegistered string `json:"dateVehicleFirstRegistered"`

// DriveWheelConfiguration see : /driveWheelConfiguration
// The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle's engine via the drivetrain.
DriveWheelConfiguration interface{} `json:"driveWheelConfiguration"` // types : DriveWheelConfigurationValue Text

// EmissionsCO2 see : https://auto.schema.org/emissionsCO2
// The CO2 emissions in g/km. When used in combination with a QuantitativeValue, put "g/km" into the unitText property of that value, since there is no UN/CEFACT Common Code for "g/km".
EmissionsCO2 string `json:"emissionsCO2"`

// FuelCapacity see : https://auto.schema.org/fuelCapacity
// The capacity of the fuel tank or in the case of electric cars, the battery. If there are multiple components for storage, this should indicate the total of all storage of the same type.
// 
// Typical unit code(s): LTR for liters, GLL of US gallons, GLI for UK / imperial gallons, AMH for ampere-hours (for electrical vehicles).
FuelCapacity string `json:"fuelCapacity"`

// FuelConsumption see : /fuelConsumption
// The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).
// 
// 
// Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. L/100 km.
// Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
// Note 3: Often, the absolute value is useful only when related to driving speed ("at 80 km/h") or usage pattern ("city traffic"). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel consumption to another value.
// 
// 
FuelConsumption string `json:"fuelConsumption"`

// FuelEfficiency see : /fuelEfficiency
// The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).
// 
// 
// Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use unitText (see: https://schema.org/unitText) to indicate the unit of measurement, e.g. mpg or km/L.
// Note 2: There are two ways of indicating the fuel consumption, fuelConsumption (see: https://schema.org/fuelConsumption) (e.g. 8 liters per 100 km) and fuelEfficiency (see: https://schema.org/fuelEfficiency) (e.g. 30 miles per gallon). They are reciprocal.
// Note 3: Often, the absolute value is useful only when related to driving speed ("at 80 km/h") or usage pattern ("city traffic"). You can use valueReference (see: https://schema.org/valueReference) to link the value for the fuel economy to another value.
// 
// 
FuelEfficiency string `json:"fuelEfficiency"`

// FuelType see : /fuelType
// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
FuelType interface{} `json:"fuelType"` // types : QualitativeValue Text URL

// KnownVehicleDamages see : /knownVehicleDamages
// A textual description of known damages, both repaired and unrepaired.
KnownVehicleDamages string `json:"knownVehicleDamages"`

// MeetsEmissionStandard see : https://auto.schema.org/meetsEmissionStandard
// Indicates that the vehicle meets the respective emission standard.
MeetsEmissionStandard interface{} `json:"meetsEmissionStandard"` // types : QualitativeValue Text URL

// MileageFromOdometer see : /mileageFromOdometer
// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
// 
// Typical unit code(s): KMT for kilometers, SMI for statute miles
MileageFromOdometer string `json:"mileageFromOdometer"`

// ModelDate see : https://auto.schema.org/modelDate
// The release date of a vehicle model (often used to differentiate versions of the same make and model).
ModelDate string `json:"modelDate"`

// NumberOfAirbags see : /numberOfAirbags
// The number or type of airbags in the vehicle.
NumberOfAirbags interface{} `json:"numberOfAirbags"` // types : Number Text

// NumberOfAxles see : /numberOfAxles
// The number of axles.
// 
// Typical unit code(s): C62
NumberOfAxles interface{} `json:"numberOfAxles"` // types : Number QuantitativeValue

// NumberOfDoors see : /numberOfDoors
// The number of doors.
// 
// Typical unit code(s): C62
NumberOfDoors interface{} `json:"numberOfDoors"` // types : Number QuantitativeValue

// NumberOfForwardGears see : /numberOfForwardGears
// The total number of forward gears available for the transmission system of the vehicle.
// 
// Typical unit code(s): C62
NumberOfForwardGears interface{} `json:"numberOfForwardGears"` // types : Number QuantitativeValue

// NumberOfPreviousOwners see : /numberOfPreviousOwners
// The number of owners of the vehicle, including the current one.
// 
// Typical unit code(s): C62
NumberOfPreviousOwners interface{} `json:"numberOfPreviousOwners"` // types : Number QuantitativeValue

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
Payload string `json:"payload"`

// ProductionDate see : /productionDate
// The date of production of the item, e.g. vehicle.
ProductionDate string `json:"productionDate"`

// PurchaseDate see : /purchaseDate
// The date the item e.g. vehicle was purchased by the current owner.
PurchaseDate string `json:"purchaseDate"`

// SeatingCapacity see : https://auto.schema.org/seatingCapacity
// The number of persons that can be seated (e.g. in a vehicle), both in terms of the physical space available, and in terms of limitations set by law.
// 
// Typical unit code(s): C62 for persons
SeatingCapacity interface{} `json:"seatingCapacity"` // types : Number QuantitativeValue

// Speed see : https://auto.schema.org/speed
// The speed range of the vehicle. If the vehicle is powered by an engine, the upper limit of the speed range (indicated by maxValue (see: https://schema.org/maxValue) should be the maximum speed achievable under regular conditions.
// 
// Typical unit code(s): KMH for km/h, HM for mile per hour (0.447 04 m/s), KNT for knot
// 
// *Note 1: Use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate the range. Typically, the minimal value is zero.
// * Note 2: There are many different ways of measuring the speed range. You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
Speed string `json:"speed"`

// SteeringPosition see : /steeringPosition
// The position of the steering wheel or similar device (mostly for cars).
SteeringPosition string `json:"steeringPosition"`

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
TongueWeight string `json:"tongueWeight"`

// TrailerWeight see : https://auto.schema.org/trailerWeight
// The permitted weight of a trailer attached to the vehicle.
// 
// Typical unit code(s): KGM for kilogram, LBR for pound
// * Note 1: You can indicate additional information in the name (see: https://schema.org/name) of the QuantitativeValue (see: https://schema.org/QuantitativeValue) node.
// * Note 2: You may also link to a QualitativeValue (see: https://schema.org/QualitativeValue) node that provides additional information using valueReference (see: https://schema.org/valueReference).
// * Note 3: Note that you can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
TrailerWeight string `json:"trailerWeight"`

// VehicleConfiguration see : /vehicleConfiguration
// A short text indicating the configuration of the vehicle, e.g. '5dr hatchback ST 2.5 MT 225 hp' or 'limited edition'.
VehicleConfiguration string `json:"vehicleConfiguration"`

// VehicleEngine see : /vehicleEngine
// Information about the engine or engines of the vehicle.
VehicleEngine string `json:"vehicleEngine"`

// VehicleIdentificationNumber see : /vehicleIdentificationNumber
// The Vehicle Identification Number (VIN) is a unique serial number used by the automotive industry to identify individual motor vehicles.
VehicleIdentificationNumber string `json:"vehicleIdentificationNumber"`

// VehicleInteriorColor see : /vehicleInteriorColor
// The color or color combination of the interior of the vehicle.
VehicleInteriorColor string `json:"vehicleInteriorColor"`

// VehicleInteriorType see : /vehicleInteriorType
// The type or material of the interior of the vehicle (e.g. synthetic fabric, leather, wood, etc.). While most interior types are characterized by the material used, an interior type can also be based on vehicle usage or target audience.
VehicleInteriorType string `json:"vehicleInteriorType"`

// VehicleModelDate see : /vehicleModelDate
// The release date of a vehicle model (often used to differentiate versions of the same make and model).
VehicleModelDate string `json:"vehicleModelDate"`

// VehicleSeatingCapacity see : /vehicleSeatingCapacity
// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
// 
// Typical unit code(s): C62 for persons.
VehicleSeatingCapacity interface{} `json:"vehicleSeatingCapacity"` // types : Number QuantitativeValue

// VehicleSpecialUsage see : https://auto.schema.org/vehicleSpecialUsage
// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
VehicleSpecialUsage interface{} `json:"vehicleSpecialUsage"` // types : CarUsageType Text

// VehicleTransmission see : /vehicleTransmission
// The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) ("gearbox" for cars).
VehicleTransmission interface{} `json:"vehicleTransmission"` // types : QualitativeValue Text URL

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
WeightTotal string `json:"weightTotal"`

// Wheelbase see : https://auto.schema.org/wheelbase
// The distance between the centers of the front and rear wheels.
// 
// Typical unit code(s): CMT for centimeters, MTR for meters, INH for inches, FOT for foot/feet
Wheelbase string `json:"wheelbase"`

// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty string `json:"additionalProperty"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// Color see : /color
// The color of the product.
Color string `json:"color"`

// Depth see : /depth
// The depth of the item.
Depth interface{} `json:"depth"` // types : Distance QuantitativeValue

// Gtin12 see : /gtin12
// The GTIN-12 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-12.aspx) code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin12 string `json:"gtin12"`

// Gtin13 see : /gtin13
// The GTIN-13 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-13.aspx) code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin13 string `json:"gtin13"`

// Gtin14 see : /gtin14
// The GTIN-14 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-14.aspx) code of the product, or the product to which the offer refers. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin14 string `json:"gtin14"`

// Gtin8 see : /gtin8
// The GTIN-8 (see: https://schema.orghttp://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See GS1 GTIN Summary (see: https://schema.orghttp://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
Gtin8 string `json:"gtin8"`

// Height see : /height
// The height of the item.
Height interface{} `json:"height"` // types : Distance QuantitativeValue

// IsAccessoryOrSparePartFor see : /isAccessoryOrSparePartFor
// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
IsAccessoryOrSparePartFor string `json:"isAccessoryOrSparePartFor"`

// IsConsumableFor see : /isConsumableFor
// A pointer to another product (or multiple products) for which this product is a consumable.
IsConsumableFor string `json:"isConsumableFor"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products).
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products).
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// ItemCondition see : /itemCondition
// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
ItemCondition string `json:"itemCondition"`

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// Manufacturer see : /manufacturer
// The manufacturer of the product.
Manufacturer string `json:"manufacturer"`

// Material see : /material
// A material that something is made from, e.g. leather, wool, cotton, paper.
Material interface{} `json:"material"` // types : Product Text URL

// Model see : /model
// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
Model interface{} `json:"model"` // types : ProductModel Text

// Mpn see : /mpn
// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
Mpn string `json:"mpn"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers string `json:"offers"`

// ProductID see : /productID
// The product identifier, such as ISBN. For example: meta itemprop="productID" content="isbn:123-456-789".
ProductID string `json:"productID"`

// ProductionDate see : /productionDate
// The date of production of the item, e.g. vehicle.
ProductionDate string `json:"productionDate"`

// PurchaseDate see : /purchaseDate
// The date the item e.g. vehicle was purchased by the current owner.
PurchaseDate string `json:"purchaseDate"`

// ReleaseDate see : /releaseDate
// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
ReleaseDate string `json:"releaseDate"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Sku see : /sku
// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
Sku string `json:"sku"`

// Weight see : /weight
// The weight of the product or person.
Weight string `json:"weight"`

// Width see : /width
// The width of the item.
Width interface{} `json:"width"` // types : Distance QuantitativeValue

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// Aircraft see : /aircraft
// The kind of aircraft (e.g., "Boeing 747"). 
Aircraft string `json:"aircraft"`

}

