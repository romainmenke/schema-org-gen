package schemaorggo

import "encoding/json"

// OpeningHoursSpecification see : https://schema.org/OpeningHoursSpecification
type OpeningHoursSpecification struct {
	StructuredValue

	typeContext

	// Closes see : https://schema.org/closes
	// The closing hour of the place or service on the given day(s) of the week.
	// types : Time
	Closes []Time `json:"closes,omitempty"`

	// DayOfWeek see : https://schema.org/dayOfWeek
	// The day of the week for which these opening hours are valid.
	// types : DayOfWeek
	DayOfWeek []*DayOfWeek `json:"dayOfWeek,omitempty"`

	// Opens see : https://schema.org/opens
	// The opening hour of the place or service on the given day(s) of the week.
	// types : Time
	Opens []Time `json:"opens,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`
}

func (v OpeningHoursSpecification) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Closes
		if len(v.Closes) == 1 {
			value = v.Closes[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["closes"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DayOfWeek
		if len(v.DayOfWeek) == 1 {
			value = v.DayOfWeek[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dayOfWeek"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Opens
		if len(v.Opens) == 1 {
			value = v.Opens[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["opens"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidFrom
		if len(v.ValidFrom) == 1 {
			value = v.ValidFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidThrough
		if len(v.ValidThrough) == 1 {
			value = v.ValidThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validThrough"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v OpeningHoursSpecification) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "OpeningHoursSpecification"

	return data, nil
}

func (v OpeningHoursSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
