<?php

class Recipe extends HowTo implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Recipe';
	
	/**
	 * The time it takes to actually cook the dish, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/cookTime
	 * @var \Duration|\Duration[]
	 */
	public var $cook_time;
	
	/**
	 * The method of cooking, such as Frying, Steaming, ...
	 * see : https://schema.org/cookingMethod
	 * @var string|string[]
	 */
	public var $cooking_method;
	
	/**
	 * Nutrition information about the recipe or menu item.
	 * see : https://schema.org/nutrition
	 * @var \NutritionInformation|\NutritionInformation[]
	 */
	public var $nutrition;
	
	/**
	 * The category of the recipe—for example, appetizer, entree, etc.
	 * see : https://schema.org/recipeCategory
	 * @var string|string[]
	 */
	public var $recipe_category;
	
	/**
	 * The cuisine of the recipe (for example, French or Ethiopian).
	 * see : https://schema.org/recipeCuisine
	 * @var string|string[]
	 */
	public var $recipe_cuisine;
	
	/**
	 * A single ingredient used in the recipe, e.g. sugar, flour or garlic. Supersedes ingredients (see: https://schema.org/ingredients).
	 * see : https://schema.org/recipeIngredient
	 * @var string|string[]
	 */
	public var $recipe_ingredient;
	
	/**
	 * A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	 * see : https://schema.org/recipeInstructions
	 * @var \CreativeWork|\CreativeWork[]|\ItemList|\ItemList[]|string|string[]
	 */
	public var $recipe_instructions;
	
	/**
	 * The quantity produced by the recipe (for example, number of people served, number of servings, etc).
	 * see : https://schema.org/recipeYield
	 * @var \QuantitativeValue|\QuantitativeValue[]|string|string[]
	 */
	public var $recipe_yield;
	
	/**
	 * Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	 * see : https://schema.org/suitableForDiet
	 * @var \RestrictedDiet|\RestrictedDiet[]
	 */
	public var $suitable_for_diet;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Recipe'
		);
		
		$serialized = so_json_serialize( $this->cook_time );
		if ( ! empty( $serialized ) ) {
			$out['cookTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cooking_method );
		if ( ! empty( $serialized ) ) {
			$out['cookingMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->nutrition );
		if ( ! empty( $serialized ) ) {
			$out['nutrition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe_category );
		if ( ! empty( $serialized ) ) {
			$out['recipeCategory'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe_cuisine );
		if ( ! empty( $serialized ) ) {
			$out['recipeCuisine'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe_ingredient );
		if ( ! empty( $serialized ) ) {
			$out['recipeIngredient'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe_instructions );
		if ( ! empty( $serialized ) ) {
			$out['recipeInstructions'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipe_yield );
		if ( ! empty( $serialized ) ) {
			$out['recipeYield'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suitable_for_diet );
		if ( ! empty( $serialized ) ) {
			$out['suitableForDiet'] = $serialized;
		}
		
		return $out;
	}
}
