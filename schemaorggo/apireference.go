package schemaorggo

import "encoding/json"

// APIReference see : https://schema.org/APIReference
type APIReference struct {
	TechArticle

	typeContext

	// AssemblyVersion see : https://schema.org/assemblyVersion
	// Associated product/technology version. e.g., .NET Framework 4.5.
	// types : Text
	AssemblyVersion []string `json:"assemblyVersion,omitempty"`

	// ExecutableLibraryName see : https://schema.org/executableLibraryName
	// Library file name e.g., mscorlib.dll, system.web.dll. Supersedes assembly (see: https://schema.org/assembly).
	// types : Text
	ExecutableLibraryName []string `json:"executableLibraryName,omitempty"`

	// ProgrammingModel see : https://schema.org/programmingModel
	// Indicates whether API is managed or unmanaged.
	// types : Text
	ProgrammingModel []string `json:"programmingModel,omitempty"`

	// TargetPlatform see : https://schema.org/targetPlatform
	// Type of app development: phone, Metro style, desktop, XBox, etc.
	// types : Text
	TargetPlatform []string `json:"targetPlatform,omitempty"`
}

func (v APIReference) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TechArticle.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AssemblyVersion
		if len(v.AssemblyVersion) == 1 {
			value = v.AssemblyVersion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["assemblyVersion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExecutableLibraryName
		if len(v.ExecutableLibraryName) == 1 {
			value = v.ExecutableLibraryName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["executableLibraryName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProgrammingModel
		if len(v.ProgrammingModel) == 1 {
			value = v.ProgrammingModel[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["programmingModel"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TargetPlatform
		if len(v.TargetPlatform) == 1 {
			value = v.TargetPlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["targetPlatform"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v APIReference) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "APIReference"

	return data, nil
}

func (v APIReference) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
