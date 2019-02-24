package schemaorggo

import "encoding/json"

// AggregateRating see : https://schema.org/AggregateRating
type AggregateRating struct {

	// With properties from Rating see : https://schema.org/Rating
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Author see : https://schema.org/author
	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	// types : Organization Person
	Author []interface{} `json:"author,omitempty"`

	// BestRating see : https://schema.org/bestRating
	// The highest value allowed in this rating system. If bestRating is omitted, 5 is assumed.
	// types : Number Text
	BestRating []interface{} `json:"bestRating,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// ItemReviewed see : https://schema.org/itemReviewed
	// The item that is being reviewed/rated.
	// types : Thing
	ItemReviewed []*Thing `json:"itemReviewed,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RatingCount see : https://schema.org/ratingCount
	// The count of total number of ratings.
	// types : Integer
	RatingCount []float64 `json:"ratingCount,omitempty"`

	// RatingValue see : https://schema.org/ratingValue
	// The rating for the content.
	// types : Text Number
	RatingValue []interface{} `json:"ratingValue,omitempty"`

	// ReviewAspect see : https://schema.org/reviewAspect
	// This Review or Rating is relevant to this part or facet of the itemReviewed.
	// types : Text
	ReviewAspect []string `json:"reviewAspect,omitempty"`

	// ReviewCount see : https://schema.org/reviewCount
	// The count of total number of reviews.
	// types : Integer
	ReviewCount []float64 `json:"reviewCount,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SubjectOf see : https://schema.org/subjectOf
	// A CreativeWork or Event about this Thing..
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// WorstRating see : https://schema.org/worstRating
	// The lowest value allowed in this rating system. If worstRating is omitted, 1 is assumed.
	// types : Number Text
	WorstRating []interface{} `json:"worstRating,omitempty"`
}

func (v AggregateRating) MarshalJSON() ([]byte, error) {
	type Alias AggregateRating

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"AggregateRating\","), b[1:]...), nil
}
