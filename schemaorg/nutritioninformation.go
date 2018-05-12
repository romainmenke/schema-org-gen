package schemaorg

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {

StructuredValue

typeContext

// Calories see : https://schema.org/calories
// The number of calories.
Calories *Energy `json:"calories"`

// CarbohydrateContent see : https://schema.org/carbohydrateContent
// The number of grams of carbohydrates.
CarbohydrateContent *Mass `json:"carbohydrateContent"`

// CholesterolContent see : https://schema.org/cholesterolContent
// The number of milligrams of cholesterol.
CholesterolContent *Mass `json:"cholesterolContent"`

// FatContent see : https://schema.org/fatContent
// The number of grams of fat.
FatContent *Mass `json:"fatContent"`

// FiberContent see : https://schema.org/fiberContent
// The number of grams of fiber.
FiberContent *Mass `json:"fiberContent"`

// ProteinContent see : https://schema.org/proteinContent
// The number of grams of protein.
ProteinContent *Mass `json:"proteinContent"`

// SaturatedFatContent see : https://schema.org/saturatedFatContent
// The number of grams of saturated fat.
SaturatedFatContent *Mass `json:"saturatedFatContent"`

// ServingSize see : https://schema.org/servingSize
// The serving size, in terms of the number of volume or mass.
ServingSize string `json:"servingSize"`

// SodiumContent see : https://schema.org/sodiumContent
// The number of milligrams of sodium.
SodiumContent *Mass `json:"sodiumContent"`

// SugarContent see : https://schema.org/sugarContent
// The number of grams of sugar.
SugarContent *Mass `json:"sugarContent"`

// TransFatContent see : https://schema.org/transFatContent
// The number of grams of trans fat.
TransFatContent *Mass `json:"transFatContent"`

// UnsaturatedFatContent see : https://schema.org/unsaturatedFatContent
// The number of grams of unsaturated fat.
UnsaturatedFatContent *Mass `json:"unsaturatedFatContent"`

}

func (v *NutritionInformation) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NutritionInformation"

	return json.Marshal(v)
}
