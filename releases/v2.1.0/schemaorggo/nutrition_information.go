package schemaorggo

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {

	// With properties from StructuredValue see : https://schema.org/StructuredValue
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

	// Calories see : https://schema.org/calories
	// The number of calories.
	// types : Energy
	Calories []*Energy `json:"calories,omitempty"`

	// CarbohydrateContent see : https://schema.org/carbohydrateContent
	// The number of grams of carbohydrates.
	// types : Mass
	CarbohydrateContent []*Mass `json:"carbohydrateContent,omitempty"`

	// CholesterolContent see : https://schema.org/cholesterolContent
	// The number of milligrams of cholesterol.
	// types : Mass
	CholesterolContent []*Mass `json:"cholesterolContent,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// FatContent see : https://schema.org/fatContent
	// The number of grams of fat.
	// types : Mass
	FatContent []*Mass `json:"fatContent,omitempty"`

	// FiberContent see : https://schema.org/fiberContent
	// The number of grams of fiber.
	// types : Mass
	FiberContent []*Mass `json:"fiberContent,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
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

	// ProteinContent see : https://schema.org/proteinContent
	// The number of grams of protein.
	// types : Mass
	ProteinContent []*Mass `json:"proteinContent,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SaturatedFatContent see : https://schema.org/saturatedFatContent
	// The number of grams of saturated fat.
	// types : Mass
	SaturatedFatContent []*Mass `json:"saturatedFatContent,omitempty"`

	// ServingSize see : https://schema.org/servingSize
	// The serving size, in terms of the number of volume or mass.
	// types : Text
	ServingSize []string `json:"servingSize,omitempty"`

	// SodiumContent see : https://schema.org/sodiumContent
	// The number of milligrams of sodium.
	// types : Mass
	SodiumContent []*Mass `json:"sodiumContent,omitempty"`

	// SugarContent see : https://schema.org/sugarContent
	// The number of grams of sugar.
	// types : Mass
	SugarContent []*Mass `json:"sugarContent,omitempty"`

	// TransFatContent see : https://schema.org/transFatContent
	// The number of grams of trans fat.
	// types : Mass
	TransFatContent []*Mass `json:"transFatContent,omitempty"`

	// UnsaturatedFatContent see : https://schema.org/unsaturatedFatContent
	// The number of grams of unsaturated fat.
	// types : Mass
	UnsaturatedFatContent []*Mass `json:"unsaturatedFatContent,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v NutritionInformation) MarshalJSON() ([]byte, error) {
	type Alias NutritionInformation

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"NutritionInformation\","), b[1:]...), nil
}
