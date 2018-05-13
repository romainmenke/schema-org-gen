package schemaorggo

import "encoding/json"

// AlignmentObject see : https://schema.org/AlignmentObject
type AlignmentObject struct {
	Intangible

	typeContext

	// AlignmentType see : https://schema.org/alignmentType
	// A category of alignment between the learning resource and the framework node. Recommended values include: &#39;assesses&#39;, &#39;teaches&#39;, &#39;requires&#39;, &#39;textComplexity&#39;, &#39;readingLevel&#39;, &#39;educationalSubject&#39;, and &#39;educationalLevel&#39;.
	// types : Text
	AlignmentType []string `json:"alignmentType,omitempty"`

	// EducationalFramework see : https://schema.org/educationalFramework
	// The framework to which the resource being described is aligned.
	// types : Text
	EducationalFramework []string `json:"educationalFramework,omitempty"`

	// TargetDescription see : https://schema.org/targetDescription
	// The description of a node in an established educational framework.
	// types : Text
	TargetDescription []string `json:"targetDescription,omitempty"`

	// TargetName see : https://schema.org/targetName
	// The name of a node in an established educational framework.
	// types : Text
	TargetName []string `json:"targetName,omitempty"`

	// TargetUrl see : https://schema.org/targetUrl
	// The URL of a node in an established educational framework.
	// types : URL
	TargetUrl []string `json:"targetUrl,omitempty"`
}

func (v AlignmentObject) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AlignmentType
		if len(v.AlignmentType) == 1 {
			value = v.AlignmentType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alignmentType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EducationalFramework
		if len(v.EducationalFramework) == 1 {
			value = v.EducationalFramework[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationalFramework"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TargetDescription
		if len(v.TargetDescription) == 1 {
			value = v.TargetDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TargetName
		if len(v.TargetName) == 1 {
			value = v.TargetName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TargetUrl
		if len(v.TargetUrl) == 1 {
			value = v.TargetUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetUrl"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v AlignmentObject) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AlignmentObject"

	return data, nil
}

func (v AlignmentObject) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
