package schemaorggo

import "encoding/json"

// ServiceChannel see : https://schema.org/ServiceChannel
type ServiceChannel struct {
	Intangible

	typeContext

	// AvailableLanguage see : https://schema.org/availableLanguage
	// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
	// types : Language Text
	AvailableLanguage []interface{} `json:"availableLanguage,omitempty"`

	// ProcessingTime see : https://schema.org/processingTime
	// Estimated processing time for the service using this channel.
	// types : Duration
	ProcessingTime []*Duration `json:"processingTime,omitempty"`

	// ProvidesService see : https://schema.org/providesService
	// The service provided by this channel.
	// types : Service
	ProvidesService []*Service `json:"providesService,omitempty"`

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
}

func (v ServiceChannel) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AvailableLanguage
		if len(v.AvailableLanguage) == 1 {
			value = v.AvailableLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["availableLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProcessingTime
		if len(v.ProcessingTime) == 1 {
			value = v.ProcessingTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["processingTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProvidesService
		if len(v.ProvidesService) == 1 {
			value = v.ProvidesService[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["providesService"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServiceLocation
		if len(v.ServiceLocation) == 1 {
			value = v.ServiceLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serviceLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServicePhone
		if len(v.ServicePhone) == 1 {
			value = v.ServicePhone[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["servicePhone"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServicePostalAddress
		if len(v.ServicePostalAddress) == 1 {
			value = v.ServicePostalAddress[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["servicePostalAddress"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServiceSmsNumber
		if len(v.ServiceSmsNumber) == 1 {
			value = v.ServiceSmsNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serviceSmsNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServiceUrl
		if len(v.ServiceUrl) == 1 {
			value = v.ServiceUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serviceUrl"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ServiceChannel) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ServiceChannel"

	return data, nil
}

func (v ServiceChannel) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
