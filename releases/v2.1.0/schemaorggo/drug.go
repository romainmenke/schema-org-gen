package schemaorggo

import "encoding/json"

// Drug see : https://schema.org/Drug
type Drug struct {

	// With properties from MedicalTherapy see : https://schema.org/MedicalTherapy
	//

	// With properties from MedicalEntity see : https://schema.org/MedicalEntity
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// ActiveIngredient see : https://schema.org/activeIngredient
	// An active ingredient, typically chemical compounds and/or biologic substances.
	// types : Text
	ActiveIngredient []string `json:"activeIngredient,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AdministrationRoute see : https://schema.org/administrationRoute
	// A route by which this drug may be administered, e.g. &#39;oral&#39;.
	// types : Text
	AdministrationRoute []string `json:"administrationRoute,omitempty"`

	// AdverseOutcome see : https://schema.org/adverseOutcome
	// A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	// types : MedicalEntity
	AdverseOutcome []*MedicalEntity `json:"adverseOutcome,omitempty"`

	// AlcoholWarning see : https://schema.org/alcoholWarning
	// Any precaution, guidance, contraindication, etc. related to consumption of alcohol while taking this drug.
	// types : Text
	AlcoholWarning []string `json:"alcoholWarning,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AvailableStrength see : https://schema.org/availableStrength
	// An available dosage strength for the drug.
	// types : DrugStrength
	AvailableStrength []*DrugStrength `json:"availableStrength,omitempty"`

	// BreastfeedingWarning see : https://schema.org/breastfeedingWarning
	// Any precaution, guidance, contraindication, etc. related to this drug&#39;s use by breastfeeding mothers.
	// types : Text
	BreastfeedingWarning []string `json:"breastfeedingWarning,omitempty"`

	// ClinicalPharmacology see : https://schema.org/clinicalPharmacology
	// Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
	// types : Text
	ClinicalPharmacology []string `json:"clinicalPharmacology,omitempty"`

	// Code see : https://schema.org/code
	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	// types : MedicalCode
	Code []*MedicalCode `json:"code,omitempty"`

	// Contraindication see : https://schema.org/contraindication
	// A contraindication for this therapy.
	// types : MedicalContraindication
	Contraindication []*MedicalContraindication `json:"contraindication,omitempty"`

	// Cost see : https://schema.org/cost
	// Cost per unit of the drug, as reported by the source being tagged.
	// types : DrugCost
	Cost []*DrugCost `json:"cost,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DosageForm see : https://schema.org/dosageForm
	// A dosage form in which this drug/supplement is available, e.g. &#39;tablet&#39;, &#39;suspension&#39;, &#39;injection&#39;.
	// types : Text
	DosageForm []string `json:"dosageForm,omitempty"`

	// DoseSchedule see : https://schema.org/doseSchedule
	// A dosing schedule for the drug for a given population, either observed, recommended, or maximum dose based on the type used.
	// types : DoseSchedule
	DoseSchedule []*DoseSchedule `json:"doseSchedule,omitempty"`

	// DrugClass see : https://schema.org/drugClass
	// The class of drug this belongs to (e.g., statins).
	// types : DrugClass
	DrugClass []*DrugClass `json:"drugClass,omitempty"`

	// DuplicateTherapy see : https://schema.org/duplicateTherapy
	// A therapy that duplicates or overlaps this one.
	// types : MedicalTherapy
	DuplicateTherapy []*MedicalTherapy `json:"duplicateTherapy,omitempty"`

	// FoodWarning see : https://schema.org/foodWarning
	// Any precaution, guidance, contraindication, etc. related to consumption of specific foods while taking this drug.
	// types : Text
	FoodWarning []string `json:"foodWarning,omitempty"`

	// Guideline see : https://schema.org/guideline
	// A medical guideline related to this entity.
	// types : MedicalGuideline
	Guideline []*MedicalGuideline `json:"guideline,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// Indication see : https://schema.org/indication
	// A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	// types : MedicalIndication
	Indication []*MedicalIndication `json:"indication,omitempty"`

	// InteractingDrug see : https://schema.org/interactingDrug
	// Another drug that is known to interact with this drug in a way that impacts the effect of this drug or causes a risk to the patient. Note: disease interactions are typically captured as contraindications.
	// types : Drug
	InteractingDrug []*Drug `json:"interactingDrug,omitempty"`

	// IsAvailableGenerically see : https://schema.org/isAvailableGenerically
	// True if the drug is available in a generic form (regardless of name).
	// types : Boolean
	IsAvailableGenerically []bool `json:"isAvailableGenerically,omitempty"`

	// IsProprietary see : https://schema.org/isProprietary
	// True if this item&#39;s name is a proprietary/brand name (vs. generic name).
	// types : Boolean
	IsProprietary []bool `json:"isProprietary,omitempty"`

	// LabelDetails see : https://schema.org/labelDetails
	// Link to the drug&#39;s label details.
	// types : URL
	LabelDetails []string `json:"labelDetails,omitempty"`

	// LegalStatus see : https://schema.org/legalStatus
	// The drug or supplement&#39;s legal status, including any controlled substance schedules that apply.
	// types : DrugLegalStatus
	LegalStatus []*DrugLegalStatus `json:"legalStatus,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Manufacturer see : https://schema.org/manufacturer
	// The manufacturer of the product.
	// types : Organization
	Manufacturer []*Organization `json:"manufacturer,omitempty"`

	// MechanismOfAction see : https://schema.org/mechanismOfAction
	// The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
	// types : Text
	MechanismOfAction []string `json:"mechanismOfAction,omitempty"`

	// MedicineSystem see : https://schema.org/medicineSystem
	// The system of medicine that includes this MedicalEntity, for example &#39;evidence-based&#39;, &#39;homeopathic&#39;, &#39;chiropractic&#39;, etc.
	// types : MedicineSystem
	MedicineSystem []*MedicineSystem `json:"medicineSystem,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NonProprietaryName see : https://schema.org/nonProprietaryName
	// The generic name of this drug or supplement.
	// types : Text
	NonProprietaryName []string `json:"nonProprietaryName,omitempty"`

	// Overdosage see : https://schema.org/overdosage
	// Any information related to overdose on a drug, including signs or symptoms, treatments, contact information for emergency response.
	// types : Text
	Overdosage []string `json:"overdosage,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PregnancyCategory see : https://schema.org/pregnancyCategory
	// Pregnancy category of this drug.
	// types : DrugPregnancyCategory
	PregnancyCategory []*DrugPregnancyCategory `json:"pregnancyCategory,omitempty"`

	// PregnancyWarning see : https://schema.org/pregnancyWarning
	// Any precaution, guidance, contraindication, etc. related to this drug&#39;s use during pregnancy.
	// types : Text
	PregnancyWarning []string `json:"pregnancyWarning,omitempty"`

	// PrescribingInfo see : https://schema.org/prescribingInfo
	// Link to prescribing information for the drug.
	// types : URL
	PrescribingInfo []string `json:"prescribingInfo,omitempty"`

	// PrescriptionStatus see : https://schema.org/prescriptionStatus
	// Indicates whether this drug is available by prescription or over-the-counter.
	// types : DrugPrescriptionStatus
	PrescriptionStatus []*DrugPrescriptionStatus `json:"prescriptionStatus,omitempty"`

	// RecognizingAuthority see : https://schema.org/recognizingAuthority
	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	// types : Organization
	RecognizingAuthority []*Organization `json:"recognizingAuthority,omitempty"`

	// RelatedDrug see : https://schema.org/relatedDrug
	// Any other drug related to this one, for example commonly-prescribed alternatives.
	// types : Drug
	RelatedDrug []*Drug `json:"relatedDrug,omitempty"`

	// RelevantSpecialty see : https://schema.org/relevantSpecialty
	// If applicable, a medical specialty in which this entity is relevant.
	// types : MedicalSpecialty
	RelevantSpecialty []*MedicalSpecialty `json:"relevantSpecialty,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SeriousAdverseOutcome see : https://schema.org/seriousAdverseOutcome
	// A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	// types : MedicalEntity
	SeriousAdverseOutcome []*MedicalEntity `json:"seriousAdverseOutcome,omitempty"`

	// Study see : https://schema.org/study
	// A medical study or trial related to this entity.
	// types : MedicalStudy
	Study []*MedicalStudy `json:"study,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// Warning see : https://schema.org/warning
	// Any FDA or other warnings about the drug (text or URL).
	// types : Text URL
	Warning []string `json:"warning,omitempty"`
}

func (v Drug) MarshalJSON() ([]byte, error) {
	type Alias Drug

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Drug\","), b[1:]...), nil
}
