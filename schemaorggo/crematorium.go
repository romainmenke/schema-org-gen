package schemaorggo

import "encoding/json"

// Crematorium see : https://schema.org/Crematorium
type Crematorium struct {
	CivicStructure

	typeContext

	// OpeningHours see : https://schema.org/openingHours
	// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.
	//
	//
	// Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	// Times are specified using 24:00 time. For example, 3pm is specified as 15:00.
	// Here is an example: &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;.
	// If a business is open 7 days a week, then it can be specified as &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;.
	//
	//
	// types : Text
	OpeningHours []string `json:"openingHours,omitempty"`
}

func (v Crematorium) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CivicStructure.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.OpeningHours
		if len(v.OpeningHours) == 1 {
			value = v.OpeningHours[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["openingHours"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Crematorium) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Crematorium"

	return data, nil
}

func (v Crematorium) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
