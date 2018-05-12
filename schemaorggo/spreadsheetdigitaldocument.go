package schemaorggo

import "encoding/json"

// SpreadsheetDigitalDocument see : https://schema.org/SpreadsheetDigitalDocument
type SpreadsheetDigitalDocument struct {
	DigitalDocument

	typeContext

	// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to "public".
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"` // types : DigitalDocumentPermission

}

func (v SpreadsheetDigitalDocument) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SpreadsheetDigitalDocument"

	return json.Marshal(v)
}

func (v *SpreadsheetDigitalDocument) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SpreadsheetDigitalDocument"

	return json.Marshal(*v)
}
