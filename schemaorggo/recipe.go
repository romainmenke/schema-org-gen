package schemaorggo

import "encoding/json"

// Recipe see : https://schema.org/Recipe
type Recipe struct {
	HowTo

	typeContext

	// CookTime see : https://schema.org/cookTime
	// The time it takes to actually cook the dish, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	CookTime *Duration `json:"cookTime,omitempty"`

	// CookingMethod see : https://schema.org/cookingMethod
	// The method of cooking, such as Frying, Steaming, ...
	// types : Text
	CookingMethod string `json:"cookingMethod,omitempty"`

	// Nutrition see : https://schema.org/nutrition
	// Nutrition information about the recipe or menu item.
	// types : NutritionInformation
	Nutrition *NutritionInformation `json:"nutrition,omitempty"`

	// RecipeCategory see : https://schema.org/recipeCategory
	// The category of the recipe—for example, appetizer, entree, etc.
	// types : Text
	RecipeCategory string `json:"recipeCategory,omitempty"`

	// RecipeCuisine see : https://schema.org/recipeCuisine
	// The cuisine of the recipe (for example, French or Ethiopian).
	// types : Text
	RecipeCuisine string `json:"recipeCuisine,omitempty"`

	// RecipeIngredient see : https://schema.org/recipeIngredient
	// A single ingredient used in the recipe, e.g. sugar, flour or garlic. Supersedes ingredients (see: https://schema.org/ingredients).
	// types : Text
	RecipeIngredient string `json:"recipeIngredient,omitempty"`

	// RecipeInstructions see : https://schema.org/recipeInstructions
	// A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	// types : CreativeWork ItemList Text
	RecipeInstructions interface{} `json:"recipeInstructions,omitempty"`

	// RecipeYield see : https://schema.org/recipeYield
	// The quantity produced by the recipe (for example, number of people served, number of servings, etc).
	// types : QuantitativeValue Text
	RecipeYield interface{} `json:"recipeYield,omitempty"`

	// SuitableForDiet see : https://schema.org/suitableForDiet
	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	// types : RestrictedDiet
	SuitableForDiet *RestrictedDiet `json:"suitableForDiet,omitempty"`
}

func (v Recipe) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Recipe"

	return json.Marshal(v)
}

func (v *Recipe) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Recipe"

	return json.Marshal(*v)
}
