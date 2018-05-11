package schemaorg

// Thing see : https://schema.org/Thing
type Thing struct {

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

// About see : /about
// The subject matter of the content. 
About interface{} `json:"about"` // types : CommunicateAction CreativeWork Event

// ActionOption see : /actionOption
// A sub property of object. The options subject to this action.  Supersedes option (see: https://schema.org/option).
ActionOption string `json:"actionOption"`

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy. 
Category interface{} `json:"category"` // types : Class Invoice Offer PhysicalActivity Product Property Service

// CharacterAttribute see : /characterAttribute
// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage). 
CharacterAttribute interface{} `json:"characterAttribute"` // types : Game VideoGameSeries

// DataFeedElement see : /dataFeedElement
// An item within in a data feed. Data feeds may have many elements. 
DataFeedElement string `json:"dataFeedElement"`

// DefaultValue see : /defaultValue
// The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it's an ID reference to one of the current values. 
DefaultValue string `json:"defaultValue"`

// Error see : /error
// For failed actions, more information on the cause of the failure. 
Error string `json:"error"`

// GameItem see : /gameItem
// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character. 
GameItem interface{} `json:"gameItem"` // types : Game VideoGameSeries

// GamePlatform see : /gamePlatform
// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms). 
GamePlatform interface{} `json:"gamePlatform"` // types : VideoGame VideoGameSeries

// Instrument see : /instrument
// The object that helped the agent perform the action. e.g. John wrote a book with a pen. 
Instrument string `json:"instrument"`

// Item see : /item
// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’. 
Item interface{} `json:"item"` // types : DataFeedItem ListItem

// ItemListElement see : /itemListElement
// For itemListElement values, you can use simple strings (e.g. "Peter", "Paul", "Mary"), existing entities, or use ListItem.
// 
// Text values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.
// 
// Note: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a 'position' property in such cases. 
ItemListElement string `json:"itemListElement"`

// ItemReviewed see : /itemReviewed
// The item that is being reviewed/rated. 
ItemReviewed interface{} `json:"itemReviewed"` // types : AggregateRating Review

// MainEntity see : /mainEntity
// Indicates the primary entity described in some page or other CreativeWork.  inverse property: mainEntityOfPage (see: https://schema.org/mainEntityOfPage).
MainEntity string `json:"mainEntity"`

// Mentions see : /mentions
// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept. 
Mentions string `json:"mentions"`

// Object see : /object
// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn't). e.g. John read a book. 
Object string `json:"object"`

// Quest see : /quest
// The task that a player-controlled character, or group of characters may complete in order to gain a reward. 
Quest interface{} `json:"quest"` // types : Game VideoGameSeries

// Replacee see : /replacee
// A sub property of object. The object that is being replaced. 
Replacee string `json:"replacee"`

// Replacer see : /replacer
// A sub property of object. The object that replaces. 
Replacer string `json:"replacer"`

// RequiredCollateral see : /requiredCollateral
// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.) 
RequiredCollateral string `json:"requiredCollateral"`

// ReservationFor see : /reservationFor
// The thing -- flight, event, restaurant,etc. being reserved. 
ReservationFor string `json:"reservationFor"`

// Result see : /result
// The result produced in the action. e.g. John wrote a book. 
Result string `json:"result"`

// ServiceOutput see : /serviceOutput
// The tangible thing generated by the service, e.g. a passport, permit, etc.  Supersedes produces (see: https://schema.org/produces).
ServiceOutput string `json:"serviceOutput"`

// TargetCollection see : /targetCollection
// A sub property of object. The collection target of the action.  Supersedes collection (see: https://schema.org/collection).
TargetCollection string `json:"targetCollection"`

}

