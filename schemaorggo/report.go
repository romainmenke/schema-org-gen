package schemaorggo

import "encoding/json"

// Report see : https://schema.org/Report
type Report struct {
	Article

	typeContext

	// ReportNumber see : https://schema.org/reportNumber
	// The number or other unique designator assigned to a Report by the publishing organization.
	// types : Text
	ReportNumber string `json:"reportNumber,omitempty"`
}

func (v Report) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Report"

	return json.Marshal(v)
}

func (v *Report) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Report"

	return json.Marshal(*v)
}
