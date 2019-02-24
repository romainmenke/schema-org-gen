package schemaorggo

import "encoding/json"

// Car see : https://schema.org/Car
type Car struct {

	// With properties from Vehicle see : https://schema.org/Vehicle
	//

	// With properties from Product see : https://schema.org/Product
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// types : PropertyValue
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created.
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Awards see : https://schema.org/awards
	// Awards won by or for this item.
	// types : Text
	Awards []string `json:"awards,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// CargoVolume see : https://schema.org/cargoVolume
	// The available volume for cargo or luggage. For automobiles, this is usually the trunk volume.\n\nTypical unit code(s): LTR for liters, FTQ for cubic foot/feet\n\nNote: You can use [[minValue]] and [[maxValue]] to indicate ranges.
	// types : QuantitativeValue
	CargoVolume []*QuantitativeValue `json:"cargoVolume,omitempty"`

	// Category see : https://schema.org/category
	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	// types : Text Thing
	Category []interface{} `json:"category,omitempty"`

	// Color see : https://schema.org/color
	// The color of the product.
	// types : Text
	Color []string `json:"color,omitempty"`

	// DateVehicleFirstRegistered see : https://schema.org/dateVehicleFirstRegistered
	// The date of the first registration of the vehicle with the respective public authorities.
	// types : Date
	DateVehicleFirstRegistered []Date `json:"dateVehicleFirstRegistered,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	// types : Distance QuantitativeValue
	Depth []interface{} `json:"depth,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// DriveWheelConfiguration see : https://schema.org/driveWheelConfiguration
	// The drive wheel configuration, i.e. which roadwheels will receive torque from the vehicle&#39;s engine via the drivetrain.
	// types : DriveWheelConfigurationValue Text
	DriveWheelConfiguration []interface{} `json:"driveWheelConfiguration,omitempty"`

	// FuelConsumption see : https://schema.org/fuelConsumption
	// The amount of fuel consumed for traveling a particular distance or temporal duration with the given vehicle (e.g. liters per 100 km).\n\n* Note 1: There are unfortunately no standard unit codes for liters per 100 km.  Use [[unitText]] to indicate the unit of measurement, e.g. L/100 km.\n* Note 2: There are two ways of indicating the fuel consumption, [[fuelConsumption]] (e.g. 8 liters per 100 km) and [[fuelEfficiency]] (e.g. 30 miles per gallon). They are reciprocal.\n* Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use [[valueReference]] to link the value for the fuel consumption to another value.
	// types : QuantitativeValue
	FuelConsumption []*QuantitativeValue `json:"fuelConsumption,omitempty"`

	// FuelEfficiency see : https://schema.org/fuelEfficiency
	// The distance traveled per unit of fuel used; most commonly miles per gallon (mpg) or kilometers per liter (km/L).\n\n* Note 1: There are unfortunately no standard unit codes for miles per gallon or kilometers per liter. Use [[unitText]] to indicate the unit of measurement, e.g. mpg or km/L.\n* Note 2: There are two ways of indicating the fuel consumption, [[fuelConsumption]] (e.g. 8 liters per 100 km) and [[fuelEfficiency]] (e.g. 30 miles per gallon). They are reciprocal.\n* Note 3: Often, the absolute value is useful only when related to driving speed (&quot;at 80 km/h&quot;) or usage pattern (&quot;city traffic&quot;). You can use [[valueReference]] to link the value for the fuel economy to another value.
	// types : QuantitativeValue
	FuelEfficiency []*QuantitativeValue `json:"fuelEfficiency,omitempty"`

	// FuelType see : https://schema.org/fuelType
	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	// types : Text QualitativeValue URL
	FuelType []interface{} `json:"fuelType,omitempty"`

	// Gtin12 see : https://schema.org/gtin12
	// The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin12 []string `json:"gtin12,omitempty"`

	// Gtin13 see : https://schema.org/gtin13
	// The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin13 []string `json:"gtin13,omitempty"`

	// Gtin14 see : https://schema.org/gtin14
	// The GTIN-14 code of the product, or the product to which the offer refers. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin14 []string `json:"gtin14,omitempty"`

	// Gtin8 see : https://schema.org/gtin8
	// The [GTIN-8](http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx) code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See [GS1 GTIN Summary](http://www.gs1.org/barcodes/technical/idkeys/gtin) for more details.
	// types : Text
	Gtin8 []string `json:"gtin8,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IsAccessoryOrSparePartFor see : https://schema.org/isAccessoryOrSparePartFor
	// A pointer to another product (or multiple products) for which this product is an accessory or spare part.
	// types : Product
	IsAccessoryOrSparePartFor []*Product `json:"isAccessoryOrSparePartFor,omitempty"`

	// IsConsumableFor see : https://schema.org/isConsumableFor
	// A pointer to another product (or multiple products) for which this product is a consumable.
	// types : Product
	IsConsumableFor []*Product `json:"isConsumableFor,omitempty"`

	// IsRelatedTo see : https://schema.org/isRelatedTo
	// A pointer to another, somehow related product (or multiple products).
	// types : Product Service
	IsRelatedTo []interface{} `json:"isRelatedTo,omitempty"`

	// IsSimilarTo see : https://schema.org/isSimilarTo
	// A pointer to another, functionally similar product (or multiple products).
	// types : Product Service
	IsSimilarTo []interface{} `json:"isSimilarTo,omitempty"`

	// ItemCondition see : https://schema.org/itemCondition
	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	// types : OfferItemCondition
	ItemCondition []*OfferItemCondition `json:"itemCondition,omitempty"`

	// KnownVehicleDamages see : https://schema.org/knownVehicleDamages
	// A textual description of known damages, both repaired and unrepaired.
	// types : Text
	KnownVehicleDamages []string `json:"knownVehicleDamages,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo []interface{} `json:"logo,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Manufacturer see : https://schema.org/manufacturer
	// The manufacturer of the product.
	// types : Organization
	Manufacturer []*Organization `json:"manufacturer,omitempty"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	// types : Text URL Product
	Material []interface{} `json:"material,omitempty"`

	// MileageFromOdometer see : https://schema.org/mileageFromOdometer
	// The total distance travelled by the particular vehicle since its initial production, as read from its odometer.\n\nTypical unit code(s): KMT for kilometers, SMI for statute miles
	// types : QuantitativeValue
	MileageFromOdometer []*QuantitativeValue `json:"mileageFromOdometer,omitempty"`

	// Model see : https://schema.org/model
	// The model of the product. Use with the URL of a ProductModel or a textual representation of the model identifier. The URL of the ProductModel can be from an external source. It is recommended to additionally provide strong product identifiers via the gtin8/gtin13/gtin14 and mpn properties.
	// types : ProductModel Text
	Model []interface{} `json:"model,omitempty"`

	// Mpn see : https://schema.org/mpn
	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	// types : Text
	Mpn []string `json:"mpn,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NumberOfAirbags see : https://schema.org/numberOfAirbags
	// The number or type of airbags in the vehicle.
	// types : Number Text
	NumberOfAirbags []interface{} `json:"numberOfAirbags,omitempty"`

	// NumberOfAxles see : https://schema.org/numberOfAxles
	// The number of axles.\n\nTypical unit code(s): C62
	// types : QuantitativeValue Number
	NumberOfAxles []interface{} `json:"numberOfAxles,omitempty"`

	// NumberOfDoors see : https://schema.org/numberOfDoors
	// The number of doors.\n\nTypical unit code(s): C62
	// types : QuantitativeValue Number
	NumberOfDoors []interface{} `json:"numberOfDoors,omitempty"`

	// NumberOfForwardGears see : https://schema.org/numberOfForwardGears
	// The total number of forward gears available for the transmission system of the vehicle.\n\nTypical unit code(s): C62
	// types : QuantitativeValue Number
	NumberOfForwardGears []interface{} `json:"numberOfForwardGears,omitempty"`

	// NumberOfPreviousOwners see : https://schema.org/numberOfPreviousOwners
	// The number of owners of the vehicle, including the current one.\n\nTypical unit code(s): C62
	// types : QuantitativeValue Number
	NumberOfPreviousOwners []interface{} `json:"numberOfPreviousOwners,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// ProductID see : https://schema.org/productID
	// The product identifier, such as ISBN. For example: ``` meta itemprop=&quot;productID&quot; content=&quot;isbn:123-456-789&quot; ```.
	// types : Text
	ProductID []string `json:"productID,omitempty"`

	// ProductionDate see : https://schema.org/productionDate
	// The date of production of the item, e.g. vehicle.
	// types : Date
	ProductionDate []Date `json:"productionDate,omitempty"`

	// PurchaseDate see : https://schema.org/purchaseDate
	// The date the item e.g. vehicle was purchased by the current owner.
	// types : Date
	PurchaseDate []Date `json:"purchaseDate,omitempty"`

	// ReleaseDate see : https://schema.org/releaseDate
	// The release date of a product or product model. This can be used to distinguish the exact variant of a product.
	// types : Date
	ReleaseDate []Date `json:"releaseDate,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Reviews see : https://schema.org/reviews
	// Review of the item.
	// types : Review
	Reviews []*Review `json:"reviews,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Sku see : https://schema.org/sku
	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	// types : Text
	Sku []string `json:"sku,omitempty"`

	// Slogan see : https://schema.org/slogan
	// A slogan or motto associated with the item.
	// types : Text
	Slogan []string `json:"slogan,omitempty"`

	// SteeringPosition see : https://schema.org/steeringPosition
	// The position of the steering wheel or similar device (mostly for cars).
	// types : SteeringPositionValue
	SteeringPosition []*SteeringPositionValue `json:"steeringPosition,omitempty"`

	// SubjectOf see : https://schema.org/subjectOf
	// A CreativeWork or Event about this Thing..
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

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
	// The number of passengers that can be seated in the vehicle, both in terms of the physical space available, and in terms of limitations set by law.\n\nTypical unit code(s): C62 for persons.
	// types : QuantitativeValue Number
	VehicleSeatingCapacity []interface{} `json:"vehicleSeatingCapacity,omitempty"`

	// VehicleSpecialUsage see : https://schema.org/vehicleSpecialUsage
	// Indicates whether the vehicle has been used for special purposes, like commercial rental, driving school, or as a taxi. The legislation in many countries requires this information to be revealed when offering a car for sale.
	// types : Text
	VehicleSpecialUsage []string `json:"vehicleSpecialUsage,omitempty"`

	// VehicleTransmission see : https://schema.org/vehicleTransmission
	// The type of component used for transmitting the power from a rotating power source to the wheels or other relevant component(s) (&quot;gearbox&quot; for cars).
	// types : Text QualitativeValue URL
	VehicleTransmission []interface{} `json:"vehicleTransmission,omitempty"`

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	// types : QuantitativeValue
	Weight []*QuantitativeValue `json:"weight,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width []interface{} `json:"width,omitempty"`
}

func (v Car) MarshalJSON() ([]byte, error) {
	type Alias Car

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Car\","), b[1:]...), nil
}
