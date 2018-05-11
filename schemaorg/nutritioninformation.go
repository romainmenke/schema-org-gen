package schemaorg

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {

typeContext

StructuredValue

// Calories see : /calories
// The number of calories.
Calories *Energy `json:"calories"`

// CarbohydrateContent see : /carbohydrateContent
// The number of grams of carbohydrates.
CarbohydrateContent *Mass `json:"carbohydrateContent"`

// CholesterolContent see : /cholesterolContent
// The number of milligrams of cholesterol.
CholesterolContent *Mass `json:"cholesterolContent"`

// FatContent see : /fatContent
// The number of grams of fat.
FatContent *Mass `json:"fatContent"`

// FiberContent see : /fiberContent
// The number of grams of fiber.
FiberContent *Mass `json:"fiberContent"`

// ProteinContent see : /proteinContent
// The number of grams of protein.
ProteinContent *Mass `json:"proteinContent"`

// SaturatedFatContent see : /saturatedFatContent
// The number of grams of saturated fat.
SaturatedFatContent *Mass `json:"saturatedFatContent"`

// ServingSize see : /servingSize
// The serving size, in terms of the number of volume or mass.
ServingSize string `json:"servingSize"`

// SodiumContent see : /sodiumContent
// The number of milligrams of sodium.
SodiumContent *Mass `json:"sodiumContent"`

// SugarContent see : /sugarContent
// The number of grams of sugar.
SugarContent *Mass `json:"sugarContent"`

// TransFatContent see : /transFatContent
// The number of grams of trans fat.
TransFatContent *Mass `json:"transFatContent"`

// UnsaturatedFatContent see : /unsaturatedFatContent
// The number of grams of unsaturated fat.
UnsaturatedFatContent *Mass `json:"unsaturatedFatContent"`

}

func (v *NutritionInformation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NutritionInformation"

	return json.Marshal(v)
}
