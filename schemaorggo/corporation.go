package schemaorggo

import "encoding/json"

// Corporation see : https://schema.org/Corporation
type Corporation struct {
	Organization

	typeContext

	// TickerSymbol see : https://schema.org/tickerSymbol
	// The exchange traded instrument associated with a Corporation object. The tickerSymbol is expressed as an exchange and an instrument name separated by a space character. For the exchange component of the tickerSymbol attribute, we reccommend using the controlled vocaulary of Market Identifier Codes (MIC) specified in ISO15022.
	// types : Text
	TickerSymbol string `json:"tickerSymbol,omitempty"`
}

func (v Corporation) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Corporation"

	return json.Marshal(v)
}

func (v *Corporation) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Corporation"

	return json.Marshal(*v)
}
