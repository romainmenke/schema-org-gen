package fetch

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/romainmenke/schema-org-gen/internal/parsers"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"golang.org/x/net/html"
	"golang.org/x/time/rate"
)

var ratelimiter = rate.NewLimiter(10, 1)

func Object(ctx context.Context, o *typemap.ObjectSource, err error) error {
	if err != nil {
		return err
	}
	if o == nil {
		return nil
	}
	if strings.HasPrefix(o.URL, "#") {
		return nil
	}

	req, err := http.NewRequest("GET", "https://schema.org"+o.URL, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s", "https://schema.org"+o.URL))
	}

	req = req.WithContext(ctx)

	ratelimiter.Wait(ctx)

	start := time.Now()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s", "https://schema.org"+o.URL))
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode/100 != 2 {
		return errors.Wrap(fmt.Errorf("unexpected status code : %d", resp.StatusCode), fmt.Sprintf("%s", "https://schema.org"+o.URL))
	}
	if resp.Body == nil {
		return errors.Wrap(errors.New("nil resp body"), fmt.Sprintf("%s", "https://schema.org"+o.URL))
	}

	z := html.NewTokenizer(resp.Body)
	err = parsers.ObjectParser(z, o)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s", "https://schema.org"+o.URL))
	}

	log.Println(fmt.Sprintf("%-29s %-39s %10v", o.Object.Name, "https://schema.org"+o.URL, time.Since(start)))

	return nil
}
