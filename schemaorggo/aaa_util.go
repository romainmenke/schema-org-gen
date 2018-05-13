package schemaorggo

// typeContext is used for fixed values
type typeContext struct {
	C string `json:"@context,omitempty"`
	T string `json:"@type"`
}
