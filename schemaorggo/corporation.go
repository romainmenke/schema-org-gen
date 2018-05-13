package schemaorggo

import "encoding/json"

// Corporation see : https://schema.org/Corporation
type Corporation struct {
	Organization

	typeContext

	// TickerSymbol see : https://schema.org/tickerSymbol
	// The exchange traded instrument associated with a Corporation object. The tickerSymbol is expressed as an exchange and an instrument name separated by a space character. For the exchange component of the tickerSymbol attribute, we reccommend using the controlled vocaulary of Market Identifier Codes (MIC) specified in ISO15022.
	// types : Text
	TickerSymbol []string `json:"tickerSymbol,omitempty"`
}

func (v Corporation) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Organization.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.TickerSymbol
		if len(v.TickerSymbol) == 1 {
			value = v.TickerSymbol[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["tickerSymbol"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Corporation) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Corporation"

	return data, nil
}

func (v Corporation) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
