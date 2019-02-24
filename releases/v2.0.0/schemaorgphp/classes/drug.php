<?php
namespace SchemaOrg;

// Drug see : https://schema.org/Drug
class Drug implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Drug';

	/**
	 * With properties from MedicalTherapy see : https://schema.org/MedicalTherapy
	 */

	/**
	 * With properties from MedicalEntity see : https://schema.org/MedicalEntity
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An active ingredient, typically chemical compounds and/or biologic substances.
	 * see : https://schema.org/activeIngredient
	 *
	 * @var string | string[]
	 */
	public $active_ingredient;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * A route by which this drug may be administered, e.g. &#39;oral&#39;.
	 * see : https://schema.org/administrationRoute
	 *
	 * @var string | string[]
	 */
	public $administration_route;

	/**
	 * A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	 * see : https://schema.org/adverseOutcome
	 *
	 * @var \MedicalEntity | \MedicalEntity[]
	 */
	public $adverse_outcome;

	/**
	 * Any precaution, guidance, contraindication, etc. related to consumption of alcohol while taking this drug.
	 * see : https://schema.org/alcoholWarning
	 *
	 * @var string | string[]
	 */
	public $alcohol_warning;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * An available dosage strength for the drug.
	 * see : https://schema.org/availableStrength
	 *
	 * @var \DrugStrength | \DrugStrength[]
	 */
	public $available_strength;

	/**
	 * Any precaution, guidance, contraindication, etc. related to this drug&#39;s use by breastfeeding mothers.
	 * see : https://schema.org/breastfeedingWarning
	 *
	 * @var string | string[]
	 */
	public $breastfeeding_warning;

	/**
	 * Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
	 * see : https://schema.org/clinicalPharmacology
	 *
	 * @var string | string[]
	 */
	public $clinical_pharmacology;

	/**
	 * A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	 * see : https://schema.org/code
	 *
	 * @var \MedicalCode | \MedicalCode[]
	 */
	public $code;

	/**
	 * A contraindication for this therapy.
	 * see : https://schema.org/contraindication
	 *
	 * @var \MedicalContraindication | \MedicalContraindication[]
	 */
	public $contraindication;

	/**
	 * Cost per unit of the drug, as reported by the source being tagged.
	 * see : https://schema.org/cost
	 *
	 * @var \DrugCost | \DrugCost[]
	 */
	public $cost;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A dosage form in which this drug/supplement is available, e.g. &#39;tablet&#39;, &#39;suspension&#39;, &#39;injection&#39;.
	 * see : https://schema.org/dosageForm
	 *
	 * @var string | string[]
	 */
	public $dosage_form;

	/**
	 * A dosing schedule for the drug for a given population, either observed, recommended, or maximum dose based on the type used.
	 * see : https://schema.org/doseSchedule
	 *
	 * @var \DoseSchedule | \DoseSchedule[]
	 */
	public $dose_schedule;

	/**
	 * The class of drug this belongs to (e.g., statins).
	 * see : https://schema.org/drugClass
	 *
	 * @var \DrugClass | \DrugClass[]
	 */
	public $drug_class;

	/**
	 * A therapy that duplicates or overlaps this one.
	 * see : https://schema.org/duplicateTherapy
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $duplicate_therapy;

	/**
	 * Any precaution, guidance, contraindication, etc. related to consumption of specific foods while taking this drug.
	 * see : https://schema.org/foodWarning
	 *
	 * @var string | string[]
	 */
	public $food_warning;

	/**
	 * A medical guideline related to this entity.
	 * see : https://schema.org/guideline
	 *
	 * @var \MedicalGuideline | \MedicalGuideline[]
	 */
	public $guideline;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	 * see : https://schema.org/indication
	 *
	 * @var \MedicalIndication | \MedicalIndication[]
	 */
	public $indication;

	/**
	 * Another drug that is known to interact with this drug in a way that impacts the effect of this drug or causes a risk to the patient. Note: disease interactions are typically captured as contraindications.
	 * see : https://schema.org/interactingDrug
	 *
	 * @var \Drug | \Drug[]
	 */
	public $interacting_drug;

	/**
	 * True if the drug is available in a generic form (regardless of name).
	 * see : https://schema.org/isAvailableGenerically
	 *
	 * @var boolean | boolean[]
	 */
	public $is_available_generically;

	/**
	 * True if this item&#39;s name is a proprietary/brand name (vs. generic name).
	 * see : https://schema.org/isProprietary
	 *
	 * @var boolean | boolean[]
	 */
	public $is_proprietary;

	/**
	 * Link to the drug&#39;s label details.
	 * see : https://schema.org/labelDetails
	 *
	 * @var string | string[]
	 */
	public $label_details;

	/**
	 * The drug or supplement&#39;s legal status, including any controlled substance schedules that apply.
	 * see : https://schema.org/legalStatus
	 *
	 * @var \DrugLegalStatus | \DrugLegalStatus[]
	 */
	public $legal_status;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	 *       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	 *       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	 *       between the page and the primary entity.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       Related properties include sameAs, about, and url.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	 *       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	 *       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	 *       serves more to clarify which of several entities is the main one for that page.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	 *       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	 *       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	 *       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	 *       describes some other entity. For example, one web page may display a news article about a particular person.
	 *       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	 *       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The manufacturer of the product.
	 * see : https://schema.org/manufacturer
	 *
	 * @var \Organization | \Organization[]
	 */
	public $manufacturer;

	/**
	 * The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
	 * see : https://schema.org/mechanismOfAction
	 *
	 * @var string | string[]
	 */
	public $mechanism_of_action;

	/**
	 * The system of medicine that includes this MedicalEntity, for example &#39;evidence-based&#39;, &#39;homeopathic&#39;, &#39;chiropractic&#39;, etc.
	 * see : https://schema.org/medicineSystem
	 *
	 * @var \MedicineSystem | \MedicineSystem[]
	 */
	public $medicine_system;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The generic name of this drug or supplement.
	 * see : https://schema.org/nonProprietaryName
	 *
	 * @var string | string[]
	 */
	public $non_proprietary_name;

	/**
	 * Any information related to overdose on a drug, including signs or symptoms, treatments, contact information for emergency response.
	 * see : https://schema.org/overdosage
	 *
	 * @var string | string[]
	 */
	public $overdosage;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * Pregnancy category of this drug.
	 * see : https://schema.org/pregnancyCategory
	 *
	 * @var \DrugPregnancyCategory | \DrugPregnancyCategory[]
	 */
	public $pregnancy_category;

	/**
	 * Any precaution, guidance, contraindication, etc. related to this drug&#39;s use during pregnancy.
	 * see : https://schema.org/pregnancyWarning
	 *
	 * @var string | string[]
	 */
	public $pregnancy_warning;

	/**
	 * Link to prescribing information for the drug.
	 * see : https://schema.org/prescribingInfo
	 *
	 * @var string | string[]
	 */
	public $prescribing_info;

	/**
	 * Indicates whether this drug is available by prescription or over-the-counter.
	 * see : https://schema.org/prescriptionStatus
	 *
	 * @var \DrugPrescriptionStatus | \DrugPrescriptionStatus[]
	 */
	public $prescription_status;

	/**
	 * If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	 * see : https://schema.org/recognizingAuthority
	 *
	 * @var \Organization | \Organization[]
	 */
	public $recognizing_authority;

	/**
	 * Any other drug related to this one, for example commonly-prescribed alternatives.
	 * see : https://schema.org/relatedDrug
	 *
	 * @var \Drug | \Drug[]
	 */
	public $related_drug;

	/**
	 * If applicable, a medical specialty in which this entity is relevant.
	 * see : https://schema.org/relevantSpecialty
	 *
	 * @var \MedicalSpecialty | \MedicalSpecialty[]
	 */
	public $relevant_specialty;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	 * see : https://schema.org/seriousAdverseOutcome
	 *
	 * @var \MedicalEntity | \MedicalEntity[]
	 */
	public $serious_adverse_outcome;

	/**
	 * A medical study or trial related to this entity.
	 * see : https://schema.org/study
	 *
	 * @var \MedicalStudy | \MedicalStudy[]
	 */
	public $study;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * Any FDA or other warnings about the drug (text or URL).
	 * see : https://schema.org/warning
	 *
	 * @var string | string[]
	 */
	public $warning;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Drug',
		);

		$serialized = \SchemaOrg\json_serialize( $this->active_ingredient );
		if ( ! empty( $serialized ) ) {
			$out['activeIngredient'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->administration_route );
		if ( ! empty( $serialized ) ) {
			$out['administrationRoute'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->adverse_outcome );
		if ( ! empty( $serialized ) ) {
			$out['adverseOutcome'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alcohol_warning );
		if ( ! empty( $serialized ) ) {
			$out['alcoholWarning'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->available_strength );
		if ( ! empty( $serialized ) ) {
			$out['availableStrength'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->breastfeeding_warning );
		if ( ! empty( $serialized ) ) {
			$out['breastfeedingWarning'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->clinical_pharmacology );
		if ( ! empty( $serialized ) ) {
			$out['clinicalPharmacology'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->code );
		if ( ! empty( $serialized ) ) {
			$out['code'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contraindication );
		if ( ! empty( $serialized ) ) {
			$out['contraindication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->cost );
		if ( ! empty( $serialized ) ) {
			$out['cost'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dosage_form );
		if ( ! empty( $serialized ) ) {
			$out['dosageForm'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dose_schedule );
		if ( ! empty( $serialized ) ) {
			$out['doseSchedule'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->drug_class );
		if ( ! empty( $serialized ) ) {
			$out['drugClass'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->duplicate_therapy );
		if ( ! empty( $serialized ) ) {
			$out['duplicateTherapy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->food_warning );
		if ( ! empty( $serialized ) ) {
			$out['foodWarning'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->guideline );
		if ( ! empty( $serialized ) ) {
			$out['guideline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->indication );
		if ( ! empty( $serialized ) ) {
			$out['indication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->interacting_drug );
		if ( ! empty( $serialized ) ) {
			$out['interactingDrug'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_available_generically );
		if ( ! empty( $serialized ) ) {
			$out['isAvailableGenerically'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_proprietary );
		if ( ! empty( $serialized ) ) {
			$out['isProprietary'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->label_details );
		if ( ! empty( $serialized ) ) {
			$out['labelDetails'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->legal_status );
		if ( ! empty( $serialized ) ) {
			$out['legalStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->manufacturer );
		if ( ! empty( $serialized ) ) {
			$out['manufacturer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mechanism_of_action );
		if ( ! empty( $serialized ) ) {
			$out['mechanismOfAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->medicine_system );
		if ( ! empty( $serialized ) ) {
			$out['medicineSystem'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->non_proprietary_name );
		if ( ! empty( $serialized ) ) {
			$out['nonProprietaryName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->overdosage );
		if ( ! empty( $serialized ) ) {
			$out['overdosage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->pregnancy_category );
		if ( ! empty( $serialized ) ) {
			$out['pregnancyCategory'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->pregnancy_warning );
		if ( ! empty( $serialized ) ) {
			$out['pregnancyWarning'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->prescribing_info );
		if ( ! empty( $serialized ) ) {
			$out['prescribingInfo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->prescription_status );
		if ( ! empty( $serialized ) ) {
			$out['prescriptionStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recognizing_authority );
		if ( ! empty( $serialized ) ) {
			$out['recognizingAuthority'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->related_drug );
		if ( ! empty( $serialized ) ) {
			$out['relatedDrug'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->relevant_specialty );
		if ( ! empty( $serialized ) ) {
			$out['relevantSpecialty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->serious_adverse_outcome );
		if ( ! empty( $serialized ) ) {
			$out['seriousAdverseOutcome'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->study );
		if ( ! empty( $serialized ) ) {
			$out['study'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->warning );
		if ( ! empty( $serialized ) ) {
			$out['warning'] = $serialized;
		}

		return $out;
	}
}
