package schemaorggo

import "encoding/json"

// WatchAction see : https://schema.org/WatchAction
type WatchAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf []*Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v WatchAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ConsumeAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ExpectsAcceptanceOf
		if len(v.ExpectsAcceptanceOf) == 1 {
			value = v.ExpectsAcceptanceOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["expectsAcceptanceOf"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v WatchAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WatchAction"

	return data, nil
}

func (v WatchAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
