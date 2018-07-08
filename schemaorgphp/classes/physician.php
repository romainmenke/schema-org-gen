<?php

class Physician extends MedicalOrganization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Physician';
	
	/**
	 * A medical service available from this provider.
	 * see : https://health-lifesci.schema.org/availableService
	 * @var \MedicalProcedure|\MedicalProcedure[]|\MedicalTest|\MedicalTest[]|\MedicalTherapy|\MedicalTherapy[]
	 */
	public var $available_service;
	
	/**
	 * A hospital with which the physician or office is affiliated.
	 * see : https://health-lifesci.schema.org/hospitalAffiliation
	 * @var \Hospital|\Hospital[]
	 */
	public var $hospital_affiliation;
	
	/**
	 * A medical specialty of the provider.
	 * see : https://health-lifesci.schema.org/medicalSpecialty
	 * @var \MedicalSpecialty|\MedicalSpecialty[]
	 */
	public var $medical_specialty;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Physician'
		);
		
		$serialized = so_json_serialize( $this->available_service );
		if ( ! empty( $serialized ) ) {
			$out['availableService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->hospital_affiliation );
		if ( ! empty( $serialized ) ) {
			$out['hospitalAffiliation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->medical_specialty );
		if ( ! empty( $serialized ) ) {
			$out['medicalSpecialty'] = $serialized;
		}
		
		return $out;
	}
}
