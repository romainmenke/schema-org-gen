<?php

class Hospital extends CivicStructure implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Hospital';
	
	/**
	 * A medical service available from this provider.
	 * see : http://health-lifesci.schema.org/availableService
	 * @var \MedicalProcedure|\MedicalProcedure[]|\MedicalTest|\MedicalTest[]|\MedicalTherapy|\MedicalTherapy[]
	 */
	public var $available_service;
	
	/**
	 * A medical specialty of the provider.
	 * see : http://health-lifesci.schema.org/medicalSpecialty
	 * @var \MedicalSpecialty|\MedicalSpecialty[]
	 */
	public var $medical_specialty;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Hospital'
		);
		
		$serialized = so_json_serialize( $this->available_service );
		if ( ! empty( $serialized ) ) {
			$out['availableService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->medical_specialty );
		if ( ! empty( $serialized ) ) {
			$out['medicalSpecialty'] = $serialized;
		}
		
		return $out;
	}
}
