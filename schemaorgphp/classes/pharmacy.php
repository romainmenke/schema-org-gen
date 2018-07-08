<?php

class Pharmacy extends MedicalOrganization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Pharmacy';
	
	/**
	 * Name or unique ID of network. (Networks are often reused across different insurance plans).
	 * see : https://pending.schema.org/healthPlanNetworkId
	 * @var string|string[]
	 */
	public var $health_plan_network_id;
	
	/**
	 * Whether the provider is accepting new patients.
	 * see : https://pending.schema.org/isAcceptingNewPatients
	 * @var boolean|boolean[]
	 */
	public var $is_accepting_new_patients;
	
	/**
	 * A medical specialty of the provider.
	 * see : https://health-lifesci.schema.org/medicalSpecialty
	 * @var \MedicalSpecialty|\MedicalSpecialty[]
	 */
	public var $medical_specialty;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Pharmacy'
		);
		
		$serialized = so_json_serialize( $this->health_plan_network_id );
		if ( ! empty( $serialized ) ) {
			$out['healthPlanNetworkId'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_accepting_new_patients );
		if ( ! empty( $serialized ) ) {
			$out['isAcceptingNewPatients'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->medical_specialty );
		if ( ! empty( $serialized ) ) {
			$out['medicalSpecialty'] = $serialized;
		}
		
		return $out;
	}
}
