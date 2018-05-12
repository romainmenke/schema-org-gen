package schemaorggo

import "encoding/json"

// SoftwareApplication see : https://schema.org/SoftwareApplication
type SoftwareApplication struct {
	CreativeWork

	typeContext

	// ApplicationCategory see : https://schema.org/applicationCategory
	// Type of software application, e.g. &#39;Game, Multimedia&#39;.
	// types : Text URL
	ApplicationCategory string `json:"applicationCategory,omitempty"`

	// ApplicationSubCategory see : https://schema.org/applicationSubCategory
	// Subcategory of the application, e.g. &#39;Arcade Game&#39;.
	// types : Text URL
	ApplicationSubCategory string `json:"applicationSubCategory,omitempty"`

	// ApplicationSuite see : https://schema.org/applicationSuite
	// The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
	// types : Text
	ApplicationSuite string `json:"applicationSuite,omitempty"`

	// AvailableOnDevice see : https://schema.org/availableOnDevice
	// Device required to run the application. Used in cases where a specific make/model is required to run the application. Supersedes device (see: https://schema.org/device).
	// types : Text
	AvailableOnDevice string `json:"availableOnDevice,omitempty"`

	// CountriesNotSupported see : https://schema.org/countriesNotSupported
	// Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	// types : Text
	CountriesNotSupported string `json:"countriesNotSupported,omitempty"`

	// CountriesSupported see : https://schema.org/countriesSupported
	// Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	// types : Text
	CountriesSupported string `json:"countriesSupported,omitempty"`

	// DownloadUrl see : https://schema.org/downloadUrl
	// If the file can be downloaded, URL to download the binary.
	// types : URL
	DownloadUrl string `json:"downloadUrl,omitempty"`

	// FeatureList see : https://schema.org/featureList
	// Features or modules provided by this application (and possibly required by other applications).
	// types : Text URL
	FeatureList string `json:"featureList,omitempty"`

	// FileSize see : https://schema.org/fileSize
	// Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
	// types : Text
	FileSize string `json:"fileSize,omitempty"`

	// InstallUrl see : https://schema.org/installUrl
	// URL at which the app may be installed, if different from the URL of the item.
	// types : URL
	InstallUrl string `json:"installUrl,omitempty"`

	// MemoryRequirements see : https://schema.org/memoryRequirements
	// Minimum memory requirements.
	// types : Text URL
	MemoryRequirements string `json:"memoryRequirements,omitempty"`

	// OperatingSystem see : https://schema.org/operatingSystem
	// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
	// types : Text
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// Permissions see : https://schema.org/permissions
	// Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
	// types : Text
	Permissions string `json:"permissions,omitempty"`

	// ProcessorRequirements see : https://schema.org/processorRequirements
	// Processor architecture required to run the application (e.g. IA64).
	// types : Text
	ProcessorRequirements string `json:"processorRequirements,omitempty"`

	// ReleaseNotes see : https://schema.org/releaseNotes
	// Description of what changed in this version.
	// types : Text URL
	ReleaseNotes string `json:"releaseNotes,omitempty"`

	// Screenshot see : https://schema.org/screenshot
	// A link to a screenshot image of the app.
	// types : ImageObject URL
	Screenshot interface{} `json:"screenshot,omitempty"`

	// SoftwareAddOn see : https://schema.org/softwareAddOn
	// Additional content for a software application.
	// types : SoftwareApplication
	SoftwareAddOn *SoftwareApplication `json:"softwareAddOn,omitempty"`

	// SoftwareHelp see : https://schema.org/softwareHelp
	// Software application help.
	// types : CreativeWork
	SoftwareHelp *CreativeWork `json:"softwareHelp,omitempty"`

	// SoftwareRequirements see : https://schema.org/softwareRequirements
	// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime). Supersedes requirements (see: https://schema.org/requirements).
	// types : Text URL
	SoftwareRequirements string `json:"softwareRequirements,omitempty"`

	// SoftwareVersion see : https://schema.org/softwareVersion
	// Version of the software instance.
	// types : Text
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// StorageRequirements see : https://schema.org/storageRequirements
	// Storage requirements (free space required).
	// types : Text URL
	StorageRequirements string `json:"storageRequirements,omitempty"`

	// SupportingData see : https://schema.org/supportingData
	// Supporting data for a SoftwareApplication.
	// types : DataFeed
	SupportingData *DataFeed `json:"supportingData,omitempty"`
}

func (v SoftwareApplication) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SoftwareApplication"

	return json.Marshal(v)
}

func (v *SoftwareApplication) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SoftwareApplication"

	return json.Marshal(*v)
}
