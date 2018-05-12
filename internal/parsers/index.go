package parsers

import (
	"io"
	"time"

	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"golang.org/x/net/html"
)

func IndexParser(tokenizer *html.Tokenizer) (*typemap.TypeMap, error) {

	inTree := false

	tm := &typemap.TypeMap{
		UpdatedAt: time.Now(),
	}
	var currentObjectSource *typemap.ObjectSource
	var parentObjectSource *typemap.ObjectSource

TOKENIZER_LOOP:
	for {
		tt := tokenizer.Next()
		if tokenizer.Err() != nil {
			break TOKENIZER_LOOP
		}

		switch {
		case tt == html.ErrorToken:
			break TOKENIZER_LOOP
		case tt == html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == "div" {
				for _, attr := range t.Attr {
					if attr.Key == "id" && attr.Val == "thing_tree" {
						inTree = true
					}
				}
			}

			if inTree {

				if t.Data == "ul" {
					if currentObjectSource != nil {
						parentObjectSource = currentObjectSource
					}
				}

				if t.Data == "a" {
					currentObjectSource = &typemap.ObjectSource{}

					for _, attr := range t.Attr {
						if attr.Key == "href" {
							currentObjectSource.URL = attr.Val
						}
					}

					if parentObjectSource != nil {
						currentObjectSource.Parent = parentObjectSource
						parentObjectSource.Children = append(parentObjectSource.Children, currentObjectSource)
					}
				}

			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "div" && inTree {
				inTree = false
				break TOKENIZER_LOOP
			}

			if inTree {

				if t.Data == "ul" {

					if parentObjectSource == nil {
						tm.ObjectSources = append(tm.ObjectSources, currentObjectSource)
					}

					if currentObjectSource != nil {
						currentObjectSource = currentObjectSource.Parent
					}

					if parentObjectSource != nil {
						parentObjectSource = parentObjectSource.Parent
					}

				}

			}

		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	return tm, nil
}
