package schemaorggo

import (
	"fmt"
	"time"
)

// Data Types see : https://schema.org/DataType

// Date https://en.wikipedia.org/wiki/ISO_8601
type Date time.Time

func (v Date) MarshalJSON() ([]byte, error) {
	if time.Time(v).IsZero() {
		return []byte("null"), nil
	}

	jsonValue := fmt.Sprintf("\"%s\"", time.Time(v).Format("2006-01-02"))
	return []byte(jsonValue), nil
}

// DateTime https://en.wikipedia.org/wiki/ISO_8601
type DateTime time.Time

// Time https://schema.org/Time
type Time string
