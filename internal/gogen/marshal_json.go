package gogen

import (
	"context"
	"fmt"
)

func writeMarshalJSON(ctx context.Context, dir string, packageName string) error {
	f, err := newObjectFile(dir, "aaa_marshal_json")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	f.WriteString(`
type Marshaler interface {
	MarshalJSONWithTypeContext() ([]byte, error)
}

// Alternative to "json.Marshal", this ensures "@type" and "@context" are set correctly.
func MarshalJSON(v Marshaler) ([]byte, error) {
	return v.MarshalJSONWithTypeContext()
}
`)

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
