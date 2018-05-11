package schemaorg

import "encoding/json"

// SoftwareApplication see : https://schema.org/SoftwareApplication
type SoftwareApplication struct {

typeContext

CreativeWork

// ApplicationCategory see : /applicationCategory
// Type of software application, e.g. 'Game, Multimedia'.
ApplicationCategory interface{} `json:"applicationCategory"` // types : Text URL

// ApplicationSubCategory see : /applicationSubCategory
// Subcategory of the application, e.g. 'Arcade Game'.
ApplicationSubCategory interface{} `json:"applicationSubCategory"` // types : Text URL

// ApplicationSuite see : /applicationSuite
// The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
ApplicationSuite string `json:"applicationSuite"`

// AvailableOnDevice see : /availableOnDevice
// Device required to run the application. Used in cases where a specific make/model is required to run the application. Supersedes device (see: https://schema.org/device).
AvailableOnDevice string `json:"availableOnDevice"`

// CountriesNotSupported see : /countriesNotSupported
// Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
CountriesNotSupported string `json:"countriesNotSupported"`

// CountriesSupported see : /countriesSupported
// Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
CountriesSupported string `json:"countriesSupported"`

// DownloadUrl see : /downloadUrl
// If the file can be downloaded, URL to download the binary.
DownloadUrl string `json:"downloadUrl"`

// FeatureList see : /featureList
// Features or modules provided by this application (and possibly required by other applications).
FeatureList interface{} `json:"featureList"` // types : Text URL

// FileSize see : /fileSize
// Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
FileSize string `json:"fileSize"`

// InstallUrl see : /installUrl
// URL at which the app may be installed, if different from the URL of the item.
InstallUrl string `json:"installUrl"`

// MemoryRequirements see : /memoryRequirements
// Minimum memory requirements.
MemoryRequirements interface{} `json:"memoryRequirements"` // types : Text URL

// OperatingSystem see : /operatingSystem
// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
OperatingSystem string `json:"operatingSystem"`

// Permissions see : /permissions
// Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
Permissions string `json:"permissions"`

// ProcessorRequirements see : /processorRequirements
// Processor architecture required to run the application (e.g. IA64).
ProcessorRequirements string `json:"processorRequirements"`

// ReleaseNotes see : /releaseNotes
// Description of what changed in this version.
ReleaseNotes interface{} `json:"releaseNotes"` // types : Text URL

// Screenshot see : /screenshot
// A link to a screenshot image of the app.
Screenshot interface{} `json:"screenshot"` // types : ImageObject URL

// SoftwareAddOn see : /softwareAddOn
// Additional content for a software application.
SoftwareAddOn *SoftwareApplication `json:"softwareAddOn"`

// SoftwareHelp see : /softwareHelp
// Software application help.
SoftwareHelp *CreativeWork `json:"softwareHelp"`

// SoftwareRequirements see : /softwareRequirements
// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime). Supersedes requirements (see: https://schema.org/requirements).
SoftwareRequirements interface{} `json:"softwareRequirements"` // types : Text URL

// SoftwareVersion see : /softwareVersion
// Version of the software instance.
SoftwareVersion string `json:"softwareVersion"`

// StorageRequirements see : /storageRequirements
// Storage requirements (free space required).
StorageRequirements interface{} `json:"storageRequirements"` // types : Text URL

// SupportingData see : /supportingData
// Supporting data for a SoftwareApplication.
SupportingData *DataFeed `json:"supportingData"`

}

func (v *SoftwareApplication) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SoftwareApplication"

	return json.Marshal(v)
}
