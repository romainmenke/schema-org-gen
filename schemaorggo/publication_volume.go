package schemaorggo

import "encoding/json"

// PublicationVolume see : https://schema.org/PublicationVolume
type PublicationVolume struct {
	CreativeWork

	typeContext

	// PageEnd see : https://schema.org/pageEnd
	// The page on which the work ends; for example &quot;138&quot; or &quot;xvi&quot;.
	// types : Integer Text
	PageEnd []interface{} `json:"pageEnd,omitempty"`

	// PageStart see : https://schema.org/pageStart
	// The page on which the work starts; for example &quot;135&quot; or &quot;xiii&quot;.
	// types : Integer Text
	PageStart []interface{} `json:"pageStart,omitempty"`

	// Pagination see : https://schema.org/pagination
	// Any description of pages that is not separated into pageStart and pageEnd; for example, &quot;1-6, 9, 55&quot; or &quot;10-12, 46-49&quot;.
	// types : Text
	Pagination []string `json:"pagination,omitempty"`

	// VolumeNumber see : https://schema.org/volumeNumber
	// Identifies the volume of publication or multi-part work; for example, &quot;iii&quot; or &quot;2&quot;.
	// types : Integer Text
	VolumeNumber []interface{} `json:"volumeNumber,omitempty"`
}

func (v PublicationVolume) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.PageEnd
		if len(v.PageEnd) == 1 {
			value = v.PageEnd[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pageEnd"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PageStart
		if len(v.PageStart) == 1 {
			value = v.PageStart[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pageStart"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Pagination
		if len(v.Pagination) == 1 {
			value = v.Pagination[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["pagination"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.VolumeNumber
		if len(v.VolumeNumber) == 1 {
			value = v.VolumeNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["volumeNumber"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v PublicationVolume) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PublicationVolume"

	return data, nil
}

func (v PublicationVolume) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
