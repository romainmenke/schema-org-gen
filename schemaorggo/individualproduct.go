package schemaorggo

import "encoding/json"

// IndividualProduct see : https://schema.org/IndividualProduct
type IndividualProduct struct {
	Product

	typeContext

	// SerialNumber see : https://schema.org/serialNumber
	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	SerialNumber string `json:"serialNumber,omitempty"` // types : Text

}

func (v IndividualProduct) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "IndividualProduct"

	return json.Marshal(v)
}

func (v *IndividualProduct) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "IndividualProduct"

	return json.Marshal(*v)
}
