<?php

class PropertyValue extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PropertyValue';
	
	/**
	 * The upper value of some characteristic or property.
	 * see : https://schema.org/maxValue
	 * @var float|float[]
	 */
	public var $max_value;
	
	/**
	 * A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	 * corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	 * 
	 * For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: &quot;mass spectrometry&quot; or &quot;nmr spectroscopy&quot; or &quot;colorimetry&quot; or &quot;immunofluorescence&quot;.
	 * 
	 * If the variableMeasured (see: https://schema.org/variableMeasured) is &quot;depression rating&quot;, the measurementTechnique (see: https://schema.org/measurementTechnique) could be &quot;Zung Scale&quot; or &quot;HAM-D&quot; or &quot;Beck Depression Inventory&quot;.
	 * 
	 * If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	 * see : https://pending.schema.org/measurementTechnique
	 * @var string|string[]|string|string[]
	 */
	public var $measurement_technique;
	
	/**
	 * The lower value of some characteristic or property.
	 * see : https://schema.org/minValue
	 * @var float|float[]
	 */
	public var $min_value;
	
	/**
	 * A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
	 * (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)
	 * a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	 * Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
	 * see : https://schema.org/propertyID
	 * @var string|string[]|string|string[]
	 */
	public var $propertyi_d;
	
	/**
	 * The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	 * see : https://schema.org/unitCode
	 * @var string|string[]|string|string[]
	 */
	public var $unit_code;
	
	/**
	 * A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	 * unitCode (see: https://schema.orgunitCode).
	 * see : https://schema.org/unitText
	 * @var string|string[]
	 */
	public var $unit_text;
	
	/**
	 * The value of the quantitative value or property value node.
	 * 
	 * 
	 * For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is &#39;Number&#39;.
	 * For PropertyValue (see: https://schema.org/PropertyValue), it can be &#39;Text;&#39;, &#39;Number&#39;, &#39;Boolean&#39;, or &#39;StructuredValue&#39;.
	 * 
	 * 
	 * see : https://schema.org/value
	 * @var boolean|boolean[]|float|float[]|\StructuredValue|\StructuredValue[]|string|string[]
	 */
	public var $value;
	
	/**
	 * A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	 * see : https://schema.org/valueReference
	 * @var \Enumeration|\Enumeration[]|\PropertyValue|\PropertyValue[]|\QualitativeValue|\QualitativeValue[]|\QuantitativeValue|\QuantitativeValue[]|\StructuredValue|\StructuredValue[]
	 */
	public var $value_reference;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PropertyValue'
		);
		
		$serialized = so_json_serialize( $this->max_value );
		if ( ! empty( $serialized ) ) {
			$out['maxValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->measurement_technique );
		if ( ! empty( $serialized ) ) {
			$out['measurementTechnique'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->min_value );
		if ( ! empty( $serialized ) ) {
			$out['minValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->propertyi_d );
		if ( ! empty( $serialized ) ) {
			$out['propertyID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_code );
		if ( ! empty( $serialized ) ) {
			$out['unitCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unit_text );
		if ( ! empty( $serialized ) ) {
			$out['unitText'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value );
		if ( ! empty( $serialized ) ) {
			$out['value'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_reference );
		if ( ! empty( $serialized ) ) {
			$out['valueReference'] = $serialized;
		}
		
		return $out;
	}
}
