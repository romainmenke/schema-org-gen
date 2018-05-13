package schemaorggo

import "encoding/json"

// SoftwareSourceCode see : https://schema.org/SoftwareSourceCode
type SoftwareSourceCode struct {
	CreativeWork

	typeContext

	// CodeRepository see : https://schema.org/codeRepository
	// Link to the repository where the un-compiled, human readable code and related code is located (SVN, github, CodePlex).
	// types : URL
	CodeRepository []string `json:"codeRepository,omitempty"`

	// CodeSampleType see : https://schema.org/codeSampleType
	// What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template. Supersedes sampleType (see: https://schema.org/sampleType).
	// types : Text
	CodeSampleType []string `json:"codeSampleType,omitempty"`

	// ProgrammingLanguage see : https://schema.org/programmingLanguage
	// The computer programming language.
	// types : ComputerLanguage Text
	ProgrammingLanguage []interface{} `json:"programmingLanguage,omitempty"`

	// RuntimePlatform see : https://schema.org/runtimePlatform
	// Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0). Supersedes runtime (see: https://schema.org/runtime).
	// types : Text
	RuntimePlatform []string `json:"runtimePlatform,omitempty"`

	// TargetProduct see : https://schema.org/targetProduct
	// Target Operating System / Product to which the code applies.  If applies to several versions, just the product name can be used.
	// types : SoftwareApplication
	TargetProduct []*SoftwareApplication `json:"targetProduct,omitempty"`
}

func (v SoftwareSourceCode) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CodeRepository
		if len(v.CodeRepository) == 1 {
			value = v.CodeRepository[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["codeRepository"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CodeSampleType
		if len(v.CodeSampleType) == 1 {
			value = v.CodeSampleType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["codeSampleType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProgrammingLanguage
		if len(v.ProgrammingLanguage) == 1 {
			value = v.ProgrammingLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["programmingLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RuntimePlatform
		if len(v.RuntimePlatform) == 1 {
			value = v.RuntimePlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["runtimePlatform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TargetProduct
		if len(v.TargetProduct) == 1 {
			value = v.TargetProduct[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetProduct"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v SoftwareSourceCode) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "SoftwareSourceCode"

	return data, nil
}

func (v SoftwareSourceCode) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
