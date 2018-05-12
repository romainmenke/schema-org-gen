package schemaorggo

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {
	StructuredValue

	typeContext

	// Calories see : https://schema.org/calories
	// The number of calories.
	Calories *Energy `json:"calories,omitempty"` // types : Energy

	// CarbohydrateContent see : https://schema.org/carbohydrateContent
	// The number of grams of carbohydrates.
	CarbohydrateContent *Mass `json:"carbohydrateContent,omitempty"` // types : Mass

	// CholesterolContent see : https://schema.org/cholesterolContent
	// The number of milligrams of cholesterol.
	CholesterolContent *Mass `json:"cholesterolContent,omitempty"` // types : Mass

	// FatContent see : https://schema.org/fatContent
	// The number of grams of fat.
	FatContent *Mass `json:"fatContent,omitempty"` // types : Mass

	// FiberContent see : https://schema.org/fiberContent
	// The number of grams of fiber.
	FiberContent *Mass `json:"fiberContent,omitempty"` // types : Mass

	// ProteinContent see : https://schema.org/proteinContent
	// The number of grams of protein.
	ProteinContent *Mass `json:"proteinContent,omitempty"` // types : Mass

	// SaturatedFatContent see : https://schema.org/saturatedFatContent
	// The number of grams of saturated fat.
	SaturatedFatContent *Mass `json:"saturatedFatContent,omitempty"` // types : Mass

	// ServingSize see : https://schema.org/servingSize
	// The serving size, in terms of the number of volume or mass.
	ServingSize string `json:"servingSize,omitempty"` // types : Text

	// SodiumContent see : https://schema.org/sodiumContent
	// The number of milligrams of sodium.
	SodiumContent *Mass `json:"sodiumContent,omitempty"` // types : Mass

	// SugarContent see : https://schema.org/sugarContent
	// The number of grams of sugar.
	SugarContent *Mass `json:"sugarContent,omitempty"` // types : Mass

	// TransFatContent see : https://schema.org/transFatContent
	// The number of grams of trans fat.
	TransFatContent *Mass `json:"transFatContent,omitempty"` // types : Mass

	// UnsaturatedFatContent see : https://schema.org/unsaturatedFatContent
	// The number of grams of unsaturated fat.
	UnsaturatedFatContent *Mass `json:"unsaturatedFatContent,omitempty"` // types : Mass

}

func (v NutritionInformation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NutritionInformation"

	return json.Marshal(v)
}

func (v *NutritionInformation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "NutritionInformation"

	return json.Marshal(*v)
}
