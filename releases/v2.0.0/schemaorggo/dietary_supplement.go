package schemaorggo

import "encoding/json"

// DietarySupplement see : https://schema.org/DietarySupplement
type DietarySupplement struct {

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

	// AdverseOutcome see : https://schema.org/adverseOutcome
	// A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	// types : MedicalEntity
	AdverseOutcome []*MedicalEntity `json:"adverseOutcome,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Background see : https://schema.org/background
	// Descriptive information establishing a historical perspective on the supplement. May include the rationale for the name, the population where the supplement first came to prominence, etc.
	// types : Text
	Background []string `json:"background,omitempty"`

	// Code see : https://schema.org/code
	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	// types : MedicalCode
	Code []*MedicalCode `json:"code,omitempty"`

	// Contraindication see : https://schema.org/contraindication
	// A contraindication for this therapy.
	// types : MedicalContraindication
	Contraindication []*MedicalContraindication `json:"contraindication,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DosageForm see : https://schema.org/dosageForm
	// A dosage form in which this drug/supplement is available, e.g. &#39;tablet&#39;, &#39;suspension&#39;, &#39;injection&#39;.
	// types : Text
	DosageForm []string `json:"dosageForm,omitempty"`

	// DuplicateTherapy see : https://schema.org/duplicateTherapy
	// A therapy that duplicates or overlaps this one.
	// types : MedicalTherapy
	DuplicateTherapy []*MedicalTherapy `json:"duplicateTherapy,omitempty"`

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

	// IsProprietary see : https://schema.org/isProprietary
	// True if this item&#39;s name is a proprietary/brand name (vs. generic name).
	// types : Boolean
	IsProprietary []bool `json:"isProprietary,omitempty"`

	// LegalStatus see : https://schema.org/legalStatus
	// The drug or supplement&#39;s legal status, including any controlled substance schedules that apply.
	// types : DrugLegalStatus
	LegalStatus []*DrugLegalStatus `json:"legalStatus,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Manufacturer see : https://schema.org/manufacturer
	// The manufacturer of the product.
	// types : Organization
	Manufacturer []*Organization `json:"manufacturer,omitempty"`

	// MaximumIntake see : https://schema.org/maximumIntake
	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	// types : MaximumDoseSchedule
	MaximumIntake []*MaximumDoseSchedule `json:"maximumIntake,omitempty"`

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

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RecognizingAuthority see : https://schema.org/recognizingAuthority
	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	// types : Organization
	RecognizingAuthority []*Organization `json:"recognizingAuthority,omitempty"`

	// RecommendedIntake see : https://schema.org/recommendedIntake
	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	// types : RecommendedDoseSchedule
	RecommendedIntake []*RecommendedDoseSchedule `json:"recommendedIntake,omitempty"`

	// RelevantSpecialty see : https://schema.org/relevantSpecialty
	// If applicable, a medical specialty in which this entity is relevant.
	// types : MedicalSpecialty
	RelevantSpecialty []*MedicalSpecialty `json:"relevantSpecialty,omitempty"`

	// SafetyConsideration see : https://schema.org/safetyConsideration
	// Any potential safety concern associated with the supplement. May include interactions with other drugs and foods, pregnancy, breastfeeding, known adverse reactions, and documented efficacy of the supplement.
	// types : Text
	SafetyConsideration []string `json:"safetyConsideration,omitempty"`

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

	// TargetPopulation see : https://schema.org/targetPopulation
	// Characteristics of the population for which this is intended, or which typically uses it, e.g. &#39;adults&#39;.
	// types : Text
	TargetPopulation []string `json:"targetPopulation,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v DietarySupplement) MarshalJSON() ([]byte, error) {
	type Alias DietarySupplement

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"DietarySupplement\","), b[1:]...), nil
}
