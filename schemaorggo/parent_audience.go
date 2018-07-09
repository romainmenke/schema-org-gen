package schemaorggo

import "encoding/json"

// ParentAudience see : https://schema.org/ParentAudience
type ParentAudience struct {
	typeContext

	// With properties from Audience see : https://schema.org/Audience
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from PeopleAudience see : https://schema.org/PeopleAudience
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

	// AudienceType see : https://schema.org/audienceType
	// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	// types : Text
	AudienceType []string `json:"audienceType,omitempty"`

	// ChildMaxAge see : https://schema.org/childMaxAge
	// Maximal age of the child.
	// types : Number
	ChildMaxAge []float64 `json:"childMaxAge,omitempty"`

	// ChildMinAge see : https://schema.org/childMinAge
	// Minimal age of the child.
	// types : Number
	ChildMinAge []float64 `json:"childMinAge,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// GeographicArea see : https://schema.org/geographicArea
	// The geographic area associated with the audience.
	// types : AdministrativeArea
	GeographicArea []*AdministrativeArea `json:"geographicArea,omitempty"`

	// HealthCondition see : https://health-lifesci.schema.org/healthCondition
	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	// types : MedicalCondition
	HealthCondition []interface{} `json:"healthCondition,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RequiredGender see : https://schema.org/requiredGender
	// Audiences defined by a person&#39;s gender.
	// types : Text
	RequiredGender []string `json:"requiredGender,omitempty"`

	// RequiredMaxAge see : https://schema.org/requiredMaxAge
	// Audiences defined by a person&#39;s maximum age.
	// types : Integer
	RequiredMaxAge []float64 `json:"requiredMaxAge,omitempty"`

	// RequiredMinAge see : https://schema.org/requiredMinAge
	// Audiences defined by a person&#39;s minimum age.
	// types : Integer
	RequiredMinAge []float64 `json:"requiredMinAge,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// SuggestedGender see : https://schema.org/suggestedGender
	// The gender of the person or audience.
	// types : Text
	SuggestedGender []string `json:"suggestedGender,omitempty"`

	// SuggestedMaxAge see : https://schema.org/suggestedMaxAge
	// Maximal age recommended for viewing content.
	// types : Number
	SuggestedMaxAge []float64 `json:"suggestedMaxAge,omitempty"`

	// SuggestedMinAge see : https://schema.org/suggestedMinAge
	// Minimal age recommended for viewing content.
	// types : Number
	SuggestedMinAge []float64 `json:"suggestedMinAge,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v ParentAudience) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AdditionalType
		if len(v.AdditionalType) == 1 {
			value = v.AdditionalType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlternateName
		if len(v.AlternateName) == 1 {
			value = v.AlternateName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alternateName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AudienceType
		if len(v.AudienceType) == 1 {
			value = v.AudienceType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audienceType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ChildMaxAge
		if len(v.ChildMaxAge) == 1 {
			value = v.ChildMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["childMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ChildMinAge
		if len(v.ChildMinAge) == 1 {
			value = v.ChildMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["childMinAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Description
		if len(v.Description) == 1 {
			value = v.Description[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["description"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DisambiguatingDescription
		if len(v.DisambiguatingDescription) == 1 {
			value = v.DisambiguatingDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["disambiguatingDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeographicArea
		if len(v.GeographicArea) == 1 {
			value = v.GeographicArea[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geographicArea"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HealthCondition
		if len(v.HealthCondition) == 1 {
			value = v.HealthCondition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["healthCondition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Identifier
		if len(v.Identifier) == 1 {
			value = v.Identifier[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["identifier"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Image
		if len(v.Image) == 1 {
			value = v.Image[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["image"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainEntityOfPage
		if len(v.MainEntityOfPage) == 1 {
			value = v.MainEntityOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainEntityOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Name
		if len(v.Name) == 1 {
			value = v.Name[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["name"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PotentialAction
		if len(v.PotentialAction) == 1 {
			value = v.PotentialAction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["potentialAction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredGender
		if len(v.RequiredGender) == 1 {
			value = v.RequiredGender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredGender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredMaxAge
		if len(v.RequiredMaxAge) == 1 {
			value = v.RequiredMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiredMinAge
		if len(v.RequiredMinAge) == 1 {
			value = v.RequiredMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiredMinAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SameAs
		if len(v.SameAs) == 1 {
			value = v.SameAs[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sameAs"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubjectOf
		if len(v.SubjectOf) == 1 {
			value = v.SubjectOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subjectOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedGender
		if len(v.SuggestedGender) == 1 {
			value = v.SuggestedGender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedGender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedMaxAge
		if len(v.SuggestedMaxAge) == 1 {
			value = v.SuggestedMaxAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedMaxAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedMinAge
		if len(v.SuggestedMinAge) == 1 {
			value = v.SuggestedMinAge[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedMinAge"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Url
		if len(v.Url) == 1 {
			value = v.Url[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["url"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ParentAudience) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ParentAudience"

	return data, nil
}

func (v ParentAudience) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
