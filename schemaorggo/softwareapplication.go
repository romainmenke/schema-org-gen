package schemaorggo

import "encoding/json"

// SoftwareApplication see : https://schema.org/SoftwareApplication
type SoftwareApplication struct {
	CreativeWork

	typeContext

	// ApplicationCategory see : https://schema.org/applicationCategory
	// Type of software application, e.g. 'Game, Multimedia'.
	ApplicationCategory string `json:"applicationCategory,omitempty"` // types : Text URL

	// ApplicationSubCategory see : https://schema.org/applicationSubCategory
	// Subcategory of the application, e.g. 'Arcade Game'.
	ApplicationSubCategory string `json:"applicationSubCategory,omitempty"` // types : Text URL

	// ApplicationSuite see : https://schema.org/applicationSuite
	// The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
	ApplicationSuite string `json:"applicationSuite,omitempty"` // types : Text

	// AvailableOnDevice see : https://schema.org/availableOnDevice
	// Device required to run the application. Used in cases where a specific make/model is required to run the application. Supersedes device (see: https://schema.org/device).
	AvailableOnDevice string `json:"availableOnDevice,omitempty"` // types : Text

	// CountriesNotSupported see : https://schema.org/countriesNotSupported
	// Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	CountriesNotSupported string `json:"countriesNotSupported,omitempty"` // types : Text

	// CountriesSupported see : https://schema.org/countriesSupported
	// Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	CountriesSupported string `json:"countriesSupported,omitempty"` // types : Text

	// DownloadUrl see : https://schema.org/downloadUrl
	// If the file can be downloaded, URL to download the binary.
	DownloadUrl string `json:"downloadUrl,omitempty"` // types : URL

	// FeatureList see : https://schema.org/featureList
	// Features or modules provided by this application (and possibly required by other applications).
	FeatureList string `json:"featureList,omitempty"` // types : Text URL

	// FileSize see : https://schema.org/fileSize
	// Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
	FileSize string `json:"fileSize,omitempty"` // types : Text

	// InstallUrl see : https://schema.org/installUrl
	// URL at which the app may be installed, if different from the URL of the item.
	InstallUrl string `json:"installUrl,omitempty"` // types : URL

	// MemoryRequirements see : https://schema.org/memoryRequirements
	// Minimum memory requirements.
	MemoryRequirements string `json:"memoryRequirements,omitempty"` // types : Text URL

	// OperatingSystem see : https://schema.org/operatingSystem
	// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
	OperatingSystem string `json:"operatingSystem,omitempty"` // types : Text

	// Permissions see : https://schema.org/permissions
	// Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
	Permissions string `json:"permissions,omitempty"` // types : Text

	// ProcessorRequirements see : https://schema.org/processorRequirements
	// Processor architecture required to run the application (e.g. IA64).
	ProcessorRequirements string `json:"processorRequirements,omitempty"` // types : Text

	// ReleaseNotes see : https://schema.org/releaseNotes
	// Description of what changed in this version.
	ReleaseNotes string `json:"releaseNotes,omitempty"` // types : Text URL

	// Screenshot see : https://schema.org/screenshot
	// A link to a screenshot image of the app.
	Screenshot interface{} `json:"screenshot,omitempty"` // types : ImageObject URL

	// SoftwareAddOn see : https://schema.org/softwareAddOn
	// Additional content for a software application.
	SoftwareAddOn *SoftwareApplication `json:"softwareAddOn,omitempty"` // types : SoftwareApplication

	// SoftwareHelp see : https://schema.org/softwareHelp
	// Software application help.
	SoftwareHelp *CreativeWork `json:"softwareHelp,omitempty"` // types : CreativeWork

	// SoftwareRequirements see : https://schema.org/softwareRequirements
	// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime). Supersedes requirements (see: https://schema.org/requirements).
	SoftwareRequirements string `json:"softwareRequirements,omitempty"` // types : Text URL

	// SoftwareVersion see : https://schema.org/softwareVersion
	// Version of the software instance.
	SoftwareVersion string `json:"softwareVersion,omitempty"` // types : Text

	// StorageRequirements see : https://schema.org/storageRequirements
	// Storage requirements (free space required).
	StorageRequirements string `json:"storageRequirements,omitempty"` // types : Text URL

	// SupportingData see : https://schema.org/supportingData
	// Supporting data for a SoftwareApplication.
	SupportingData *DataFeed `json:"supportingData,omitempty"` // types : DataFeed

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
