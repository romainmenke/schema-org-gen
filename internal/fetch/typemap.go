package fetch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/romainmenke/schema-org-gen/internal/parsers"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"golang.org/x/net/html"
)

func TypeMap(ctx context.Context) (*typemap.TypeMap, error) {
	req, err := http.NewRequest("GET", "https://schema.org/docs/full.html", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("unexpected status code : %d", resp.StatusCode)
	}
	if resp.Body == nil {
		return nil, errors.New("nil resp body")
	}

	z := html.NewTokenizer(resp.Body)

	tMap, err := parsers.IndexParser(z)
	if err != nil {
		return nil, err
	}

	log.Println("fetched the object list")

	return tMap, nil
}
