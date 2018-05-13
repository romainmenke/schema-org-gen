<?php

class NutritionInformation extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'NutritionInformation';
	
	/**
	 * The number of calories.
	 * see : https://schema.org/calories
	 * @var \Energy|\Energy[]
	 */
	public var $calories;
	
	/**
	 * The number of grams of carbohydrates.
	 * see : https://schema.org/carbohydrateContent
	 * @var \Mass|\Mass[]
	 */
	public var $carbohydrate_content;
	
	/**
	 * The number of milligrams of cholesterol.
	 * see : https://schema.org/cholesterolContent
	 * @var \Mass|\Mass[]
	 */
	public var $cholesterol_content;
	
	/**
	 * The number of grams of fat.
	 * see : https://schema.org/fatContent
	 * @var \Mass|\Mass[]
	 */
	public var $fat_content;
	
	/**
	 * The number of grams of fiber.
	 * see : https://schema.org/fiberContent
	 * @var \Mass|\Mass[]
	 */
	public var $fiber_content;
	
	/**
	 * The number of grams of protein.
	 * see : https://schema.org/proteinContent
	 * @var \Mass|\Mass[]
	 */
	public var $protein_content;
	
	/**
	 * The number of grams of saturated fat.
	 * see : https://schema.org/saturatedFatContent
	 * @var \Mass|\Mass[]
	 */
	public var $saturated_fat_content;
	
	/**
	 * The serving size, in terms of the number of volume or mass.
	 * see : https://schema.org/servingSize
	 * @var string|string[]
	 */
	public var $serving_size;
	
	/**
	 * The number of milligrams of sodium.
	 * see : https://schema.org/sodiumContent
	 * @var \Mass|\Mass[]
	 */
	public var $sodium_content;
	
	/**
	 * The number of grams of sugar.
	 * see : https://schema.org/sugarContent
	 * @var \Mass|\Mass[]
	 */
	public var $sugar_content;
	
	/**
	 * The number of grams of trans fat.
	 * see : https://schema.org/transFatContent
	 * @var \Mass|\Mass[]
	 */
	public var $trans_fat_content;
	
	/**
	 * The number of grams of unsaturated fat.
	 * see : https://schema.org/unsaturatedFatContent
	 * @var \Mass|\Mass[]
	 */
	public var $unsaturated_fat_content;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'NutritionInformation'
		);
		
		$serialized = so_json_serialize( $this->calories );
		if ( ! empty( $serialized ) ) {
			$out['calories'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->carbohydrate_content );
		if ( ! empty( $serialized ) ) {
			$out['carbohydrateContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cholesterol_content );
		if ( ! empty( $serialized ) ) {
			$out['cholesterolContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fat_content );
		if ( ! empty( $serialized ) ) {
			$out['fatContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fiber_content );
		if ( ! empty( $serialized ) ) {
			$out['fiberContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->protein_content );
		if ( ! empty( $serialized ) ) {
			$out['proteinContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->saturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['saturatedFatContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->serving_size );
		if ( ! empty( $serialized ) ) {
			$out['servingSize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sodium_content );
		if ( ! empty( $serialized ) ) {
			$out['sodiumContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sugar_content );
		if ( ! empty( $serialized ) ) {
			$out['sugarContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->trans_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['transFatContent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unsaturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['unsaturatedFatContent'] = $serialized;
		}
		
		return $out;
	}
}
