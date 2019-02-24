package schemaorggo

import "encoding/json"

// ServiceChannel see : https://schema.org/ServiceChannel
type ServiceChannel struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with the item. Please use one of the language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47). See also [[inLanguage]]
	// types : Language Text
	AvailableLanguage []interface{} `json:"availableLanguage,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// ProcessingTime see : https://schema.org/processingTime
	// Estimated processing time for the service using this channel.
	// types : Duration
	ProcessingTime []*Duration `json:"processingTime,omitempty"`

	// ProvidesService see : https://schema.org/providesService
	// The service provided by this channel.
	// types : Service
	ProvidesService []*Service `json:"providesService,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// ServiceLocation see : https://schema.org/serviceLocation
	// The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
	// types : Place
	ServiceLocation []*Place `json:"serviceLocation,omitempty"`

	// ServicePhone see : https://schema.org/servicePhone
	// The phone number to use to access the service.
	// types : ContactPoint
	ServicePhone []*ContactPoint `json:"servicePhone,omitempty"`

	// ServicePostalAddress see : https://schema.org/servicePostalAddress
	// The address for accessing the service by mail.
	// types : PostalAddress
	ServicePostalAddress []*PostalAddress `json:"servicePostalAddress,omitempty"`

	// ServiceSmsNumber see : https://schema.org/serviceSmsNumber
	// The number to access the service by text message.
	// types : ContactPoint
	ServiceSmsNumber []*ContactPoint `json:"serviceSmsNumber,omitempty"`

	// ServiceUrl see : https://schema.org/serviceUrl
	// The website to access the service.
	// types : URL
	ServiceUrl []string `json:"serviceUrl,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v ServiceChannel) MarshalJSON() ([]byte, error) {
	type Alias ServiceChannel

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"ServiceChannel\","), b[1:]...), nil
}
