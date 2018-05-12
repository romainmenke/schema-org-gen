package schemaorggo

import "encoding/json"

// ServiceChannel see : https://schema.org/ServiceChannel
type ServiceChannel struct {
	Intangible

	typeContext

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	AvailableLanguage interface{} `json:"availableLanguage,omitempty"` // types : Language Text

	// ProcessingTime see : https://schema.org/processingTime
	// Estimated processing time for the service using this channel.
	ProcessingTime *Duration `json:"processingTime,omitempty"` // types : Duration

	// ProvidesService see : https://schema.org/providesService
	// The service provided by this channel.
	ProvidesService *Service `json:"providesService,omitempty"` // types : Service

	// ServiceLocation see : https://schema.org/serviceLocation
	// The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
	ServiceLocation *Place `json:"serviceLocation,omitempty"` // types : Place

	// ServicePhone see : https://schema.org/servicePhone
	// The phone number to use to access the service.
	ServicePhone *ContactPoint `json:"servicePhone,omitempty"` // types : ContactPoint

	// ServicePostalAddress see : https://schema.org/servicePostalAddress
	// The address for accessing the service by mail.
	ServicePostalAddress *PostalAddress `json:"servicePostalAddress,omitempty"` // types : PostalAddress

	// ServiceSmsNumber see : https://schema.org/serviceSmsNumber
	// The number to access the service by text message.
	ServiceSmsNumber *ContactPoint `json:"serviceSmsNumber,omitempty"` // types : ContactPoint

	// ServiceUrl see : https://schema.org/serviceUrl
	// The website to access the service.
	ServiceUrl string `json:"serviceUrl,omitempty"` // types : URL

}

func (v ServiceChannel) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ServiceChannel"

	return json.Marshal(v)
}

func (v *ServiceChannel) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ServiceChannel"

	return json.Marshal(*v)
}
