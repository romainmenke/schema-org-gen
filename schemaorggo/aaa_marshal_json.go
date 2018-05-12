package schemaorggo

type Marshaler interface {
	MarshalJSONWithTypeContext() ([]byte, error)
}

// Alternative to "json.Marshal", this ensures "@type" and "@context" are set correctly.
func MarshalJSON(v Marshaler) ([]byte, error) {
	return v.MarshalJSONWithTypeContext()
}
