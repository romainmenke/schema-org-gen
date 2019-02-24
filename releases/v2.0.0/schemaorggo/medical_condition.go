package schemaorggo

import "encoding/json"

// MedicalCondition see : https://schema.org/MedicalCondition
type MedicalCondition struct {

	// With properties from MedicalEntity see : https://schema.org/MedicalEntity
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AssociatedAnatomy see : https://schema.org/associatedAnatomy
	// The anatomy of the underlying organ system or structures associated with this entity.
	// types : AnatomicalStructure AnatomicalSystem SuperficialAnatomy
	AssociatedAnatomy []interface{} `json:"associatedAnatomy,omitempty"`

	// Cause see : https://schema.org/cause
	// An underlying cause. More specifically, one of the causative agent(s) that are most directly responsible for the pathophysiologic process that eventually results in the occurrence.
	// types : MedicalCause
	Cause []*MedicalCause `json:"cause,omitempty"`

	// Code see : https://schema.org/code
	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	// types : MedicalCode
	Code []*MedicalCode `json:"code,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DifferentialDiagnosis see : https://schema.org/differentialDiagnosis
	// One of a set of differential diagnoses for the condition. Specifically, a closely-related or competing diagnosis typically considered later in the cognitive process whereby this medical condition is distinguished from others most likely responsible for a similar collection of signs and symptoms to reach the most parsimonious diagnosis or diagnoses in a patient.
	// types : DDxElement
	DifferentialDiagnosis []*DDxElement `json:"differentialDiagnosis,omitempty"`

	// Epidemiology see : https://schema.org/epidemiology
	// The characteristics of associated patients, such as age, gender, race etc.
	// types : Text
	Epidemiology []string `json:"epidemiology,omitempty"`

	// ExpectedPrognosis see : https://schema.org/expectedPrognosis
	// The likely outcome in either the short term or long term of the medical condition.
	// types : Text
	ExpectedPrognosis []string `json:"expectedPrognosis,omitempty"`

	// Guideline see : https://schema.org/guideline
	// A medical guideline related to this entity.
	// types : MedicalGuideline
	Guideline []*MedicalGuideline `json:"guideline,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

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

	// MedicineSystem see : https://schema.org/medicineSystem
	// The system of medicine that includes this MedicalEntity, for example &#39;evidence-based&#39;, &#39;homeopathic&#39;, &#39;chiropractic&#39;, etc.
	// types : MedicineSystem
	MedicineSystem []*MedicineSystem `json:"medicineSystem,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NaturalProgression see : https://schema.org/naturalProgression
	// The expected progression of the condition if it is not treated and allowed to progress naturally.
	// types : Text
	NaturalProgression []string `json:"naturalProgression,omitempty"`

	// Pathophysiology see : https://schema.org/pathophysiology
	// Changes in the normal mechanical, physical, and biochemical functions that are associated with this activity or condition.
	// types : Text
	Pathophysiology []string `json:"pathophysiology,omitempty"`

	// PossibleComplication see : https://schema.org/possibleComplication
	// A possible unexpected and unfavorable evolution of a medical condition. Complications may include worsening of the signs or symptoms of the disease, extension of the condition to other organ systems, etc.
	// types : Text
	PossibleComplication []string `json:"possibleComplication,omitempty"`

	// PossibleTreatment see : https://schema.org/possibleTreatment
	// A possible treatment to address this condition, sign or symptom.
	// types : MedicalTherapy
	PossibleTreatment []*MedicalTherapy `json:"possibleTreatment,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PrimaryPrevention see : https://schema.org/primaryPrevention
	// A preventative therapy used to prevent an initial occurrence of the medical condition, such as vaccination.
	// types : MedicalTherapy
	PrimaryPrevention []*MedicalTherapy `json:"primaryPrevention,omitempty"`

	// RecognizingAuthority see : https://schema.org/recognizingAuthority
	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	// types : Organization
	RecognizingAuthority []*Organization `json:"recognizingAuthority,omitempty"`

	// RelevantSpecialty see : https://schema.org/relevantSpecialty
	// If applicable, a medical specialty in which this entity is relevant.
	// types : MedicalSpecialty
	RelevantSpecialty []*MedicalSpecialty `json:"relevantSpecialty,omitempty"`

	// RiskFactor see : https://schema.org/riskFactor
	// A modifiable or non-modifiable factor that increases the risk of a patient contracting this condition, e.g. age,  coexisting condition.
	// types : MedicalRiskFactor
	RiskFactor []*MedicalRiskFactor `json:"riskFactor,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SecondaryPrevention see : https://schema.org/secondaryPrevention
	// A preventative therapy used to prevent reoccurrence of the medical condition after an initial episode of the condition.
	// types : MedicalTherapy
	SecondaryPrevention []*MedicalTherapy `json:"secondaryPrevention,omitempty"`

	// SignOrSymptom see : https://schema.org/signOrSymptom
	// A sign or symptom of this condition. Signs are objective or physically observable manifestations of the medical condition while symptoms are the subjective experience of the medical condition.
	// types : MedicalSignOrSymptom
	SignOrSymptom []*MedicalSignOrSymptom `json:"signOrSymptom,omitempty"`

	// Stage see : https://schema.org/stage
	// The stage of the condition, if applicable.
	// types : MedicalConditionStage
	Stage []*MedicalConditionStage `json:"stage,omitempty"`

	// Study see : https://schema.org/study
	// A medical study or trial related to this entity.
	// types : MedicalStudy
	Study []*MedicalStudy `json:"study,omitempty"`

	// Subtype see : https://schema.org/subtype
	// A more specific type of the condition, where applicable, for example &#39;Type 1 Diabetes&#39;, &#39;Type 2 Diabetes&#39;, or &#39;Gestational Diabetes&#39; for Diabetes.
	// types : Text
	Subtype []string `json:"subtype,omitempty"`

	// TypicalTest see : https://schema.org/typicalTest
	// A medical test typically performed given this condition.
	// types : MedicalTest
	TypicalTest []*MedicalTest `json:"typicalTest,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v MedicalCondition) MarshalJSON() ([]byte, error) {
	type Alias MedicalCondition

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"MedicalCondition\","), b[1:]...), nil
}
