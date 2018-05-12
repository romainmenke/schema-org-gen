package schemaorggo

import "encoding/json"

// Recipe see : https://schema.org/Recipe
type Recipe struct {
	HowTo

	typeContext

	// CookTime see : https://schema.org/cookTime
	// The time it takes to actually cook the dish, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	CookTime *Duration `json:"cookTime,omitempty"` // types : Duration

	// CookingMethod see : https://schema.org/cookingMethod
	// The method of cooking, such as Frying, Steaming, ...
	CookingMethod string `json:"cookingMethod,omitempty"` // types : Text

	// Nutrition see : https://schema.org/nutrition
	// Nutrition information about the recipe or menu item.
	Nutrition *NutritionInformation `json:"nutrition,omitempty"` // types : NutritionInformation

	// RecipeCategory see : https://schema.org/recipeCategory
	// The category of the recipe—for example, appetizer, entree, etc.
	RecipeCategory string `json:"recipeCategory,omitempty"` // types : Text

	// RecipeCuisine see : https://schema.org/recipeCuisine
	// The cuisine of the recipe (for example, French or Ethiopian).
	RecipeCuisine string `json:"recipeCuisine,omitempty"` // types : Text

	// RecipeIngredient see : https://schema.org/recipeIngredient
	// A single ingredient used in the recipe, e.g. sugar, flour or garlic. Supersedes ingredients (see: https://schema.org/ingredients).
	RecipeIngredient string `json:"recipeIngredient,omitempty"` // types : Text

	// RecipeInstructions see : https://schema.org/recipeInstructions
	// A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	RecipeInstructions interface{} `json:"recipeInstructions,omitempty"` // types : CreativeWork ItemList Text

	// RecipeYield see : https://schema.org/recipeYield
	// The quantity produced by the recipe (for example, number of people served, number of servings, etc).
	RecipeYield interface{} `json:"recipeYield,omitempty"` // types : QuantitativeValue Text

	// SuitableForDiet see : https://schema.org/suitableForDiet
	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	SuitableForDiet *RestrictedDiet `json:"suitableForDiet,omitempty"` // types : RestrictedDiet

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
