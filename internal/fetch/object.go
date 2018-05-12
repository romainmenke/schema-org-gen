package fetch

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/romainmenke/schema-org-gen/internal/parsers"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"golang.org/x/net/html"
)

func Object(ctx context.Context, o *typemap.ObjectSource, err error) error {
	if err != nil {
		return err
	}
	if o == nil {
		return nil
	}

	req, err := http.NewRequest("GET", "https://schema.org"+o.URL, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status code : %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return errors.New("nil resp body")
	}

	z := html.NewTokenizer(resp.Body)
	err = parsers.ObjectParser(z, o)
	if err != nil {
		return err
	}

	return nil
}
