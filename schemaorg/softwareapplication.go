package schemaorg

import "encoding/json"

// SoftwareApplication see : https://schema.org/SoftwareApplication
type SoftwareApplication struct {

CreativeWork

typeContext

// ApplicationCategory see : https://schema.org/applicationCategory
// Type of software application, e.g. 'Game, Multimedia'.
ApplicationCategory interface{} `json:"applicationCategory"` // types : Text URL

// ApplicationSubCategory see : https://schema.org/applicationSubCategory
// Subcategory of the application, e.g. 'Arcade Game'.
ApplicationSubCategory interface{} `json:"applicationSubCategory"` // types : Text URL

// ApplicationSuite see : https://schema.org/applicationSuite
// The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
ApplicationSuite string `json:"applicationSuite"`

// AvailableOnDevice see : https://schema.org/availableOnDevice
// Device required to run the application. Used in cases where a specific make/model is required to run the application. Supersedes device (see: https://schema.org/device).
AvailableOnDevice string `json:"availableOnDevice"`

// CountriesNotSupported see : https://schema.org/countriesNotSupported
// Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
CountriesNotSupported string `json:"countriesNotSupported"`

// CountriesSupported see : https://schema.org/countriesSupported
// Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
CountriesSupported string `json:"countriesSupported"`

// DownloadUrl see : https://schema.org/downloadUrl
// If the file can be downloaded, URL to download the binary.
DownloadUrl string `json:"downloadUrl"`

// FeatureList see : https://schema.org/featureList
// Features or modules provided by this application (and possibly required by other applications).
FeatureList interface{} `json:"featureList"` // types : Text URL

// FileSize see : https://schema.org/fileSize
// Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
FileSize string `json:"fileSize"`

// InstallUrl see : https://schema.org/installUrl
// URL at which the app may be installed, if different from the URL of the item.
InstallUrl string `json:"installUrl"`

// MemoryRequirements see : https://schema.org/memoryRequirements
// Minimum memory requirements.
MemoryRequirements interface{} `json:"memoryRequirements"` // types : Text URL

// OperatingSystem see : https://schema.org/operatingSystem
// Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
OperatingSystem string `json:"operatingSystem"`

// Permissions see : https://schema.org/permissions
// Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
Permissions string `json:"permissions"`

// ProcessorRequirements see : https://schema.org/processorRequirements
// Processor architecture required to run the application (e.g. IA64).
ProcessorRequirements string `json:"processorRequirements"`

// ReleaseNotes see : https://schema.org/releaseNotes
// Description of what changed in this version.
ReleaseNotes interface{} `json:"releaseNotes"` // types : Text URL

// Screenshot see : https://schema.org/screenshot
// A link to a screenshot image of the app.
Screenshot interface{} `json:"screenshot"` // types : ImageObject URL

// SoftwareAddOn see : https://schema.org/softwareAddOn
// Additional content for a software application.
SoftwareAddOn *SoftwareApplication `json:"softwareAddOn"`

// SoftwareHelp see : https://schema.org/softwareHelp
// Software application help.
SoftwareHelp *CreativeWork `json:"softwareHelp"`

// SoftwareRequirements see : https://schema.org/softwareRequirements
// Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime). Supersedes requirements (see: https://schema.org/requirements).
SoftwareRequirements interface{} `json:"softwareRequirements"` // types : Text URL

// SoftwareVersion see : https://schema.org/softwareVersion
// Version of the software instance.
SoftwareVersion string `json:"softwareVersion"`

// StorageRequirements see : https://schema.org/storageRequirements
// Storage requirements (free space required).
StorageRequirements interface{} `json:"storageRequirements"` // types : Text URL

// SupportingData see : https://schema.org/supportingData
// Supporting data for a SoftwareApplication.
SupportingData *DataFeed `json:"supportingData"`

}

func (v *SoftwareApplication) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SoftwareApplication"

	return json.Marshal(v)
}
