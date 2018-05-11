package schemaorg

import "encoding/json"

// ServiceChannel see : https://schema.org/ServiceChannel
type ServiceChannel struct {

typeContext

Intangible

// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// ProcessingTime see : /processingTime
// Estimated processing time for the service using this channel.
ProcessingTime *Duration `json:"processingTime"`

// ProvidesService see : /providesService
// The service provided by this channel.
ProvidesService *Service `json:"providesService"`

// ServiceLocation see : /serviceLocation
// The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
ServiceLocation *Place `json:"serviceLocation"`

// ServicePhone see : /servicePhone
// The phone number to use to access the service.
ServicePhone *ContactPoint `json:"servicePhone"`

// ServicePostalAddress see : /servicePostalAddress
// The address for accessing the service by mail.
ServicePostalAddress *PostalAddress `json:"servicePostalAddress"`

// ServiceSmsNumber see : /serviceSmsNumber
// The number to access the service by text message.
ServiceSmsNumber *ContactPoint `json:"serviceSmsNumber"`

// ServiceUrl see : /serviceUrl
// The website to access the service.
ServiceUrl string `json:"serviceUrl"`

}

func (v *ServiceChannel) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ServiceChannel"

	return json.Marshal(v)
}
