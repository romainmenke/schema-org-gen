<?php

class Dataset extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Dataset';
	
	/**
	 * A downloadable form of this dataset, at a specific location, in a specific format.
	 * see : https://schema.org/distribution
	 * @var \DataDownload|\DataDownload[]
	 */
	public var $distribution;
	
	/**
	 * A data catalog which contains this dataset. Supersedes catalog (see: https://schema.org/catalog), includedDataCatalog (see: https://schema.org/includedDataCatalog). Inverse property: dataset (see: https://schema.org/dataset).
	 * see : https://schema.org/includedInDataCatalog
	 * @var \DataCatalog|\DataCatalog[]
	 */
	public var $included_in_data_catalog;
	
	/**
	 * The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	 * see : https://schema.org/issn
	 * @var string|string[]
	 */
	public var $issn;
	
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
	 * The variableMeasured property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue.
	 * see : https://pending.schema.org/variableMeasured
	 * @var \PropertyValue|\PropertyValue[]|string|string[]
	 */
	public var $variable_measured;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Dataset'
		);
		
		$serialized = so_json_serialize( $this->distribution );
		if ( ! empty( $serialized ) ) {
			$out['distribution'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->included_in_data_catalog );
		if ( ! empty( $serialized ) ) {
			$out['includedInDataCatalog'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->issn );
		if ( ! empty( $serialized ) ) {
			$out['issn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->measurement_technique );
		if ( ! empty( $serialized ) ) {
			$out['measurementTechnique'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->variable_measured );
		if ( ! empty( $serialized ) ) {
			$out['variableMeasured'] = $serialized;
		}
		
		return $out;
	}
}
