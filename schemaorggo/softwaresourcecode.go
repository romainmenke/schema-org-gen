package schemaorggo

import "encoding/json"

// SoftwareSourceCode see : https://schema.org/SoftwareSourceCode
type SoftwareSourceCode struct {
	CreativeWork

	typeContext

	// CodeRepository see : https://schema.org/codeRepository
	// Link to the repository where the un-compiled, human readable code and related code is located (SVN, github, CodePlex).
	// types : URL
	CodeRepository string `json:"codeRepository,omitempty"`

	// CodeSampleType see : https://schema.org/codeSampleType
	// What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template. Supersedes sampleType (see: https://schema.org/sampleType).
	// types : Text
	CodeSampleType string `json:"codeSampleType,omitempty"`

	// ProgrammingLanguage see : https://schema.org/programmingLanguage
	// The computer programming language.
	// types : ComputerLanguage Text
	ProgrammingLanguage interface{} `json:"programmingLanguage,omitempty"`

	// RuntimePlatform see : https://schema.org/runtimePlatform
	// Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0). Supersedes runtime (see: https://schema.org/runtime).
	// types : Text
	RuntimePlatform string `json:"runtimePlatform,omitempty"`

	// TargetProduct see : https://schema.org/targetProduct
	// Target Operating System / Product to which the code applies.  If applies to several versions, just the product name can be used.
	// types : SoftwareApplication
	TargetProduct *SoftwareApplication `json:"targetProduct,omitempty"`
}

func (v SoftwareSourceCode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SoftwareSourceCode"

	return json.Marshal(v)
}

func (v *SoftwareSourceCode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SoftwareSourceCode"

	return json.Marshal(*v)
}
