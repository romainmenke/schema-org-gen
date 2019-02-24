package schemaorggo

import "encoding/json"

// MedicalConditionStage see : https://schema.org/MedicalConditionStage
type MedicalConditionStage struct {

	// With properties from MedicalIntangible see : https://schema.org/MedicalIntangible
	//

	// With properties from MedicalEntity see : https://schema.org/MedicalEntity
	//

	// With properties from Thing see : https://schema.org/Thing
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

	// Code see : https://schema.org/code
	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	// types : MedicalCode
	Code []*MedicalCode `json:"code,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

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
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
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

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RecognizingAuthority see : https://schema.org/recognizingAuthority
	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	// types : Organization
	RecognizingAuthority []*Organization `json:"recognizingAuthority,omitempty"`

	// RelevantSpecialty see : https://schema.org/relevantSpecialty
	// If applicable, a medical specialty in which this entity is relevant.
	// types : MedicalSpecialty
	RelevantSpecialty []*MedicalSpecialty `json:"relevantSpecialty,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// StageAsNumber see : https://schema.org/stageAsNumber
	// The stage represented as a number, e.g. 3.
	// types : Number
	StageAsNumber []float64 `json:"stageAsNumber,omitempty"`

	// Study see : https://schema.org/study
	// A medical study or trial related to this entity.
	// types : MedicalStudy
	Study []*MedicalStudy `json:"study,omitempty"`

	// SubStageSuffix see : https://schema.org/subStageSuffix
	// The substage, e.g. &#39;a&#39; for Stage IIIa.
	// types : Text
	SubStageSuffix []string `json:"subStageSuffix,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v MedicalConditionStage) MarshalJSON() ([]byte, error) {
	type Alias MedicalConditionStage

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"MedicalConditionStage\","), b[1:]...), nil
}
