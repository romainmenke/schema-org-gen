package schemaorggo

import "encoding/json"

// NutritionInformation see : https://schema.org/NutritionInformation
type NutritionInformation struct {
	StructuredValue

	typeContext

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

	// FatContent see : https://schema.org/fatContent
	// The number of grams of fat.
	// types : Mass
	FatContent []*Mass `json:"fatContent,omitempty"`

	// FiberContent see : https://schema.org/fiberContent
	// The number of grams of fiber.
	// types : Mass
	FiberContent []*Mass `json:"fiberContent,omitempty"`

	// ProteinContent see : https://schema.org/proteinContent
	// The number of grams of protein.
	// types : Mass
	ProteinContent []*Mass `json:"proteinContent,omitempty"`

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
}

func (v NutritionInformation) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Calories
		if len(v.Calories) == 1 {
			value = v.Calories[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["calories"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CarbohydrateContent
		if len(v.CarbohydrateContent) == 1 {
			value = v.CarbohydrateContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["carbohydrateContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CholesterolContent
		if len(v.CholesterolContent) == 1 {
			value = v.CholesterolContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cholesterolContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FatContent
		if len(v.FatContent) == 1 {
			value = v.FatContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fatContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FiberContent
		if len(v.FiberContent) == 1 {
			value = v.FiberContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fiberContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProteinContent
		if len(v.ProteinContent) == 1 {
			value = v.ProteinContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["proteinContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SaturatedFatContent
		if len(v.SaturatedFatContent) == 1 {
			value = v.SaturatedFatContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["saturatedFatContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServingSize
		if len(v.ServingSize) == 1 {
			value = v.ServingSize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["servingSize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SodiumContent
		if len(v.SodiumContent) == 1 {
			value = v.SodiumContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sodiumContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SugarContent
		if len(v.SugarContent) == 1 {
			value = v.SugarContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sugarContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TransFatContent
		if len(v.TransFatContent) == 1 {
			value = v.TransFatContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["transFatContent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnsaturatedFatContent
		if len(v.UnsaturatedFatContent) == 1 {
			value = v.UnsaturatedFatContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unsaturatedFatContent"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v NutritionInformation) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "NutritionInformation"

	return data, nil
}

func (v NutritionInformation) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
