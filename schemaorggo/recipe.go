package schemaorggo

import "encoding/json"

// Recipe see : https://schema.org/Recipe
type Recipe struct {
	HowTo

	typeContext

	// CookTime see : https://schema.org/cookTime
	// The time it takes to actually cook the dish, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	CookTime []*Duration `json:"cookTime,omitempty"`

	// CookingMethod see : https://schema.org/cookingMethod
	// The method of cooking, such as Frying, Steaming, ...
	// types : Text
	CookingMethod []string `json:"cookingMethod,omitempty"`

	// Nutrition see : https://schema.org/nutrition
	// Nutrition information about the recipe or menu item.
	// types : NutritionInformation
	Nutrition []*NutritionInformation `json:"nutrition,omitempty"`

	// RecipeCategory see : https://schema.org/recipeCategory
	// The category of the recipe—for example, appetizer, entree, etc.
	// types : Text
	RecipeCategory []string `json:"recipeCategory,omitempty"`

	// RecipeCuisine see : https://schema.org/recipeCuisine
	// The cuisine of the recipe (for example, French or Ethiopian).
	// types : Text
	RecipeCuisine []string `json:"recipeCuisine,omitempty"`

	// RecipeIngredient see : https://schema.org/recipeIngredient
	// A single ingredient used in the recipe, e.g. sugar, flour or garlic. Supersedes ingredients (see: https://schema.org/ingredients).
	// types : Text
	RecipeIngredient []string `json:"recipeIngredient,omitempty"`

	// RecipeInstructions see : https://schema.org/recipeInstructions
	// A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	// types : CreativeWork ItemList Text
	RecipeInstructions []interface{} `json:"recipeInstructions,omitempty"`

	// RecipeYield see : https://schema.org/recipeYield
	// The quantity produced by the recipe (for example, number of people served, number of servings, etc).
	// types : QuantitativeValue Text
	RecipeYield []interface{} `json:"recipeYield,omitempty"`

	// SuitableForDiet see : https://schema.org/suitableForDiet
	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	// types : RestrictedDiet
	SuitableForDiet []*RestrictedDiet `json:"suitableForDiet,omitempty"`
}

func (v Recipe) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.HowTo.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.CookTime
		if len(v.CookTime) == 1 {
			value = v.CookTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cookTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CookingMethod
		if len(v.CookingMethod) == 1 {
			value = v.CookingMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cookingMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Nutrition
		if len(v.Nutrition) == 1 {
			value = v.Nutrition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["nutrition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecipeCategory
		if len(v.RecipeCategory) == 1 {
			value = v.RecipeCategory[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipeCategory"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecipeCuisine
		if len(v.RecipeCuisine) == 1 {
			value = v.RecipeCuisine[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipeCuisine"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecipeIngredient
		if len(v.RecipeIngredient) == 1 {
			value = v.RecipeIngredient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipeIngredient"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecipeInstructions
		if len(v.RecipeInstructions) == 1 {
			value = v.RecipeInstructions[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipeInstructions"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecipeYield
		if len(v.RecipeYield) == 1 {
			value = v.RecipeYield[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipeYield"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuitableForDiet
		if len(v.SuitableForDiet) == 1 {
			value = v.SuitableForDiet[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suitableForDiet"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Recipe) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Recipe"

	return data, nil
}

func (v Recipe) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
