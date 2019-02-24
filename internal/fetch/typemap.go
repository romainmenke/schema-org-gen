package fetch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/parsers"
	"golang.org/x/net/html"
)

func TypemapRDF(ctx context.Context, version string) (ast.Typemap, error) {
	var (
		req *http.Request
		err error
	)

	if version == "master" {
		req, err = http.NewRequest("GET", "https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/schema.rdfa", nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest("GET", "https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/"+version+"/schema.rdfa", nil)
		if err != nil {
			return nil, err
		}
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

	tMap, err := parsers.RDFAParser(z)
	if err != nil {
		return nil, err
	}

	log.Println("fetched the object list")

	return tMap, nil
}
