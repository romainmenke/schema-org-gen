package schemaorg


import (
	"fmt"
	"time"
)
// Data Types see : https://schema.org/DataType

// typeContext is used for fixed values
type typeContext struct {
	C string `json:"@context"`
	T string `json:"@type"`
}
// Date https://en.wikipedia.org/wiki/ISO_8601
type Date time.Time

func (v Date)MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("\"%s\"", time.Time(v).Format("2006-01-02"))
	return []byte(jsonValue), nil
}
// DateTime https://en.wikipedia.org/wiki/ISO_8601
type DateTime time.Time

// Time https://schema.org/Time
type Time string

