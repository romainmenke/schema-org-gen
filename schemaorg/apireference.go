package schemaorg

import "encoding/json"

// APIReference see : https://schema.org/APIReference
type APIReference struct {

typeContext

TechArticle

// AssemblyVersion see : /assemblyVersion
// Associated product/technology version. e.g., .NET Framework 4.5.
AssemblyVersion string `json:"assemblyVersion"`

// ExecutableLibraryName see : /executableLibraryName
// Library file name e.g., mscorlib.dll, system.web.dll. Supersedes assembly (see: https://schema.org/assembly).
ExecutableLibraryName string `json:"executableLibraryName"`

// ProgrammingModel see : /programmingModel
// Indicates whether API is managed or unmanaged.
ProgrammingModel string `json:"programmingModel"`

// TargetPlatform see : /targetPlatform
// Type of app development: phone, Metro style, desktop, XBox, etc.
TargetPlatform string `json:"targetPlatform"`

}

func (v *APIReference) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "APIReference"

	return json.Marshal(v)
}
