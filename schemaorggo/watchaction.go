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

func (v WatchAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ConsumeAction.IntoMap(intop)

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

func (v WatchAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "WatchAction"

	return data, nil
}

func (v WatchAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
