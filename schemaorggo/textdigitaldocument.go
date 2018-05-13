package schemaorggo

import "encoding/json"

// TextDigitalDocument see : https://schema.org/TextDigitalDocument
type TextDigitalDocument struct {
	DigitalDocument

	typeContext

	// HasDigitalDocumentPermission see : https://schema.org/hasDigitalDocumentPermission
	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to &quot;public&quot;.
	// types : DigitalDocumentPermission
	HasDigitalDocumentPermission []*DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"`
}

func (v TextDigitalDocument) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.DigitalDocument.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.HasDigitalDocumentPermission
		if len(v.HasDigitalDocumentPermission) == 1 {
			value = v.HasDigitalDocumentPermission[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasDigitalDocumentPermission"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TextDigitalDocument) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TextDigitalDocument"

	return data, nil
}

func (v TextDigitalDocument) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
