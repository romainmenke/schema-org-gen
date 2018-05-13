package schemaorggo

import "encoding/json"

// TrainTrip see : https://schema.org/TrainTrip
type TrainTrip struct {
	Intangible

	typeContext

	// ArrivalPlatform see : https://schema.org/arrivalPlatform
	// The platform where the train arrives.
	// types : Text
	ArrivalPlatform []string `json:"arrivalPlatform,omitempty"`

	// ArrivalStation see : https://schema.org/arrivalStation
	// The station where the train trip ends.
	// types : TrainStation
	ArrivalStation []*TrainStation `json:"arrivalStation,omitempty"`

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime []DateTime `json:"arrivalTime,omitempty"`

	// DeparturePlatform see : https://schema.org/departurePlatform
	// The platform from which the train departs.
	// types : Text
	DeparturePlatform []string `json:"departurePlatform,omitempty"`

	// DepartureStation see : https://schema.org/departureStation
	// The station from which the train departs.
	// types : TrainStation
	DepartureStation []*TrainStation `json:"departureStation,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	// types : DateTime
	DepartureTime []DateTime `json:"departureTime,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// TrainName see : https://schema.org/trainName
	// The name of the train (e.g. The Orient Express).
	// types : Text
	TrainName []string `json:"trainName,omitempty"`

	// TrainNumber see : https://schema.org/trainNumber
	// The unique identifier for the train.
	// types : Text
	TrainNumber []string `json:"trainNumber,omitempty"`
}

func (v TrainTrip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ArrivalPlatform
		if len(v.ArrivalPlatform) == 1 {
			value = v.ArrivalPlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalPlatform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalStation
		if len(v.ArrivalStation) == 1 {
			value = v.ArrivalStation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalStation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArrivalTime
		if len(v.ArrivalTime) == 1 {
			value = v.ArrivalTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DeparturePlatform
		if len(v.DeparturePlatform) == 1 {
			value = v.DeparturePlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departurePlatform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureStation
		if len(v.DepartureStation) == 1 {
			value = v.DepartureStation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureStation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTime
		if len(v.DepartureTime) == 1 {
			value = v.DepartureTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Provider
		if len(v.Provider) == 1 {
			value = v.Provider[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["provider"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TrainName
		if len(v.TrainName) == 1 {
			value = v.TrainName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trainName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TrainNumber
		if len(v.TrainNumber) == 1 {
			value = v.TrainNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trainNumber"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TrainTrip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TrainTrip"

	return data, nil
}

func (v TrainTrip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
