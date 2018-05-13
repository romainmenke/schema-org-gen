<?php

class ProductModel extends Product implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ProductModel';
	
	/**
	 * A pointer to a base product from which this product is a variant. It is safe to infer that the variant inherits all product features from the base model, unless defined locally. This is not transitive.
	 * see : https://schema.org/isVariantOf
	 * @var \ProductModel|\ProductModel[]
	 */
	public var $is_variant_of;
	
	/**
	 * A pointer from a previous, often discontinued variant of the product to its newer variant.
	 * see : https://schema.org/predecessorOf
	 * @var \ProductModel|\ProductModel[]
	 */
	public var $predecessor_of;
	
	/**
	 * A pointer from a newer variant of a product  to its previous, often discontinued predecessor.
	 * see : https://schema.org/successorOf
	 * @var \ProductModel|\ProductModel[]
	 */
	public var $successor_of;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ProductModel'
		);
		
		$serialized = so_json_serialize( $this->is_variant_of );
		if ( ! empty( $serialized ) ) {
			$out['isVariantOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->predecessor_of );
		if ( ! empty( $serialized ) ) {
			$out['predecessorOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->successor_of );
		if ( ! empty( $serialized ) ) {
			$out['successorOf'] = $serialized;
		}
		
		return $out;
	}
}
