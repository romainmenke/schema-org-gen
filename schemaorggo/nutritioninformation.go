package schemaorggo

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {
	StructuredValue

	typeContext

	// Calories see : https://schema.org/calories
	// The number of calories.
	// types : Energy
	Calories *Energy `json:"calories,omitempty"`

	// CarbohydrateContent see : https://schema.org/carbohydrateContent
	// The number of grams of carbohydrates.
	// types : Mass
	CarbohydrateContent *Mass `json:"carbohydrateContent,omitempty"`

	// CholesterolContent see : https://schema.org/cholesterolContent
	// The number of milligrams of cholesterol.
	// types : Mass
	CholesterolContent *Mass `json:"cholesterolContent,omitempty"`

	// FatContent see : https://schema.org/fatContent
	// The number of grams of fat.
	// types : Mass
	FatContent *Mass `json:"fatContent,omitempty"`

	// FiberContent see : https://schema.org/fiberContent
	// The number of grams of fiber.
	// types : Mass
	FiberContent *Mass `json:"fiberContent,omitempty"`

	// ProteinContent see : https://schema.org/proteinContent
	// The number of grams of protein.
	// types : Mass
	ProteinContent *Mass `json:"proteinContent,omitempty"`

	// SaturatedFatContent see : https://schema.org/saturatedFatContent
	// The number of grams of saturated fat.
	// types : Mass
	SaturatedFatContent *Mass `json:"saturatedFatContent,omitempty"`

	// ServingSize see : https://schema.org/servingSize
	// The serving size, in terms of the number of volume or mass.
	// types : Text
	ServingSize string `json:"servingSize,omitempty"`

	// SodiumContent see : https://schema.org/sodiumContent
	// The number of milligrams of sodium.
	// types : Mass
	SodiumContent *Mass `json:"sodiumContent,omitempty"`

	// SugarContent see : https://schema.org/sugarContent
	// The number of grams of sugar.
	// types : Mass
	SugarContent *Mass `json:"sugarContent,omitempty"`

	// TransFatContent see : https://schema.org/transFatContent
	// The number of grams of trans fat.
	// types : Mass
	TransFatContent *Mass `json:"transFatContent,omitempty"`

	// UnsaturatedFatContent see : https://schema.org/unsaturatedFatContent
	// The number of grams of unsaturated fat.
	// types : Mass
	UnsaturatedFatContent *Mass `json:"unsaturatedFatContent,omitempty"`
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
