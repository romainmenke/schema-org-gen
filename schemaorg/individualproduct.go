package schemaorg

import "encoding/json"

// IndividualProduct see : https://schema.org/IndividualProduct
type IndividualProduct struct {

typeContext

Product

// SerialNumber see : /serialNumber
// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
SerialNumber string `json:"serialNumber"`

}

func (v *IndividualProduct) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "IndividualProduct"

	return json.Marshal(v)
}
