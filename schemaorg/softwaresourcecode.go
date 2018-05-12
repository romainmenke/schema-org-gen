package schemaorg

import "encoding/json"

// SoftwareSourceCode see : https://schema.org/SoftwareSourceCode
type SoftwareSourceCode struct {

CreativeWork

typeContext

// CodeRepository see : https://schema.org/codeRepository
// Link to the repository where the un-compiled, human readable code and related code is located (SVN, github, CodePlex).
CodeRepository string `json:"codeRepository"`

// CodeSampleType see : https://schema.org/codeSampleType
// What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template. Supersedes sampleType (see: https://schema.org/sampleType).
CodeSampleType string `json:"codeSampleType"`

// ProgrammingLanguage see : https://schema.org/programmingLanguage
// The computer programming language.
ProgrammingLanguage interface{} `json:"programmingLanguage"` // types : ComputerLanguage Text

// RuntimePlatform see : https://schema.org/runtimePlatform
// Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0). Supersedes runtime (see: https://schema.org/runtime).
RuntimePlatform string `json:"runtimePlatform"`

// TargetProduct see : https://schema.org/targetProduct
// Target Operating System / Product to which the code applies.  If applies to several versions, just the product name can be used.
TargetProduct *SoftwareApplication `json:"targetProduct"`

}

func (v *SoftwareSourceCode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SoftwareSourceCode"

	return json.Marshal(v)
}
