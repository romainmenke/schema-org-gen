package schemaorg

// QuantitativeValue see : https://schema.org/QuantitativeValue
type QuantitativeValue struct {

StructuredValue// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty string `json:"additionalProperty"`

// MaxValue see : /maxValue
// The upper value of some characteristic or property.
MaxValue string `json:"maxValue"`

// MinValue see : /minValue
// The lower value of some characteristic or property.
MinValue string `json:"minValue"`

// UnitCode see : /unitCode
// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
UnitCode interface{} `json:"unitCode"` // types : Text URL

// UnitText see : /unitText
// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
// unitCode (see: https://schema.orgunitCode).
UnitText string `json:"unitText"`

// Value see : /value
// The value of the quantitative value or property value node.
// 
// 
// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is 'Number'.
// For PropertyValue (see: https://schema.org/PropertyValue), it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
// 
// 
Value interface{} `json:"value"` // types : Boolean Number StructuredValue Text

// ValueReference see : /valueReference
// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
ValueReference interface{} `json:"valueReference"` // types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue

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

// AdvanceBookingRequirement see : /advanceBookingRequirement
// The amount of time that is required between accepting the offer and the actual usage of the resource or service. 
AdvanceBookingRequirement interface{} `json:"advanceBookingRequirement"` // types : Demand Offer

// AnnualPercentageRate see : /annualPercentageRate
// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction. 
AnnualPercentageRate string `json:"annualPercentageRate"`

// CargoVolume see : /cargoVolume
// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.
// 
// Typical unit code(s): LTR for liters, FTQ for cubic foot/feet
// 
// Note: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges. 
CargoVolume string `json:"cargoVolume"`

// DeliveryLeadTime see : /deliveryLeadTime
// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup. 
DeliveryLeadTime interface{} `json:"deliveryLeadTime"` // types : Demand Offer

// Depth see : /depth
// The depth of the item. 
Depth interface{} `json:"depth"` // types : Product VisualArtwork

// DurationOfWarranty see : /durationOfWarranty
// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days. 
DurationOfWarranty string `json:"durationOfWarranty"`

// EligibleDuration see : /eligibleDuration
// The duration for which the given offer is valid. 
EligibleDuration interface{} `json:"eligibleDuration"` // types : Demand Offer

// EligibleQuantity see : /eligibleQuantity
// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity. 
EligibleQuantity interface{} `json:"eligibleQuantity"` // types : Demand Offer PriceSpecification

// FloorSize see : /floorSize
// The size of the accommodation, e.g. in square meter or squarefoot.
// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard 
FloorSize string `json:"floorSize"`

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

// Height see : /height
// The height of the item. 
Height interface{} `json:"height"` // types : MediaObject Person Product VisualArtwork

// InterestRate see : /interestRate
// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate. 
InterestRate string `json:"interestRate"`

// InventoryLevel see : /inventoryLevel
// The current approximate inventory level for the item or items. 
InventoryLevel interface{} `json:"inventoryLevel"` // types : Demand Offer SomeProducts

// LoanTerm see : /loanTerm
// The duration of the loan or credit agreement. 
LoanTerm string `json:"loanTerm"`

// MileageFromOdometer see : /mileageFromOdometer
// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.
// 
// Typical unit code(s): KMT for kilometers, SMI for statute miles 
MileageFromOdometer string `json:"mileageFromOdometer"`

// NumAdults see : /numAdults
// The number of adults staying in the unit. 
NumAdults string `json:"numAdults"`

// NumChildren see : /numChildren
// The number of children staying in the unit. 
NumChildren string `json:"numChildren"`

// NumberOfAxles see : /numberOfAxles
// The number of axles.
// 
// Typical unit code(s): C62 
NumberOfAxles string `json:"numberOfAxles"`

// NumberOfDoors see : /numberOfDoors
// The number of doors.
// 
// Typical unit code(s): C62 
NumberOfDoors string `json:"numberOfDoors"`

// NumberOfEmployees see : /numberOfEmployees
// The number of employees in an organization e.g. business. 
NumberOfEmployees interface{} `json:"numberOfEmployees"` // types : BusinessAudience Organization

// NumberOfForwardGears see : /numberOfForwardGears
// The total number of forward gears available for the transmission system of the vehicle.
// 
// Typical unit code(s): C62 
NumberOfForwardGears string `json:"numberOfForwardGears"`

// NumberOfPlayers see : /numberOfPlayers
// Indicate how many people can play this game (minimum, maximum, or range). 
NumberOfPlayers interface{} `json:"numberOfPlayers"` // types : Game VideoGameSeries

// NumberOfPreviousOwners see : /numberOfPreviousOwners
// The number of owners of the vehicle, including the current one.
// 
// Typical unit code(s): C62 
NumberOfPreviousOwners string `json:"numberOfPreviousOwners"`

// NumberOfRooms see : /numberOfRooms
// The number of rooms (excluding bathrooms and closets) of the acccommodation or lodging business.
// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue. 
NumberOfRooms interface{} `json:"numberOfRooms"` // types : Accommodation Apartment House SingleFamilyResidence Suite

// Occupancy see : /occupancy
// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
// Typical unit code(s): C62 for person 
Occupancy interface{} `json:"occupancy"` // types : Apartment HotelRoom SingleFamilyResidence Suite

// PartySize see : /partySize
// Number of people the reservation should accommodate. 
PartySize interface{} `json:"partySize"` // types : FoodEstablishmentReservation TaxiReservation

// RecipeYield see : /recipeYield
// The quantity produced by the recipe (for example, number of people served, number of servings, etc). 
RecipeYield string `json:"recipeYield"`

// ReferenceQuantity see : /referenceQuantity
// The reference quantity for which a certain price applies, e.g. 1 EUR per 4 kWh of electricity. This property is a replacement for unitOfMeasurement for the advanced cases where the price does not relate to a standard unit. 
ReferenceQuantity string `json:"referenceQuantity"`

// RequiredQuantity see : /requiredQuantity
// The required quantity of the item(s). 
RequiredQuantity string `json:"requiredQuantity"`

// ValueReference see : /valueReference
// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature. 
ValueReference interface{} `json:"valueReference"` // types : PropertyValue QualitativeValue QuantitativeValue

// VehicleSeatingCapacity see : /vehicleSeatingCapacity
// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.
// 
// Typical unit code(s): C62 for persons. 
VehicleSeatingCapacity string `json:"vehicleSeatingCapacity"`

// Weight see : /weight
// The weight of the product or person. 
Weight interface{} `json:"weight"` // types : Person Product

// Width see : /width
// The width of the item. 
Width interface{} `json:"width"` // types : MediaObject Product VisualArtwork

// YearlyRevenue see : /yearlyRevenue
// The size of the business in annual revenue. 
YearlyRevenue string `json:"yearlyRevenue"`

// YearsInOperation see : /yearsInOperation
// The age of the business. 
YearsInOperation string `json:"yearsInOperation"`

// Yield see : /yield
// The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles. 
Yield string `json:"yield"`

}

