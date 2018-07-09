package parsers

import (
	"errors"
	"io"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"github.com/romainmenke/schema-org-gen/internal/typemap"
	"golang.org/x/net/html"
)

func ObjectParser(tokenizer *html.Tokenizer, o *typemap.ObjectSource) error {
	o.Object = &ast.Object{}
	inDefinition := false
	inProperties := false
	inField := false

	parsedTitle := false

	capturedObject := false

	o.Object.URL = "https://schema.org" + o.URL

	currentField := ast.Field{}

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

			if t.Data == "h1" && !parsedTitle {
				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "page-title" {
						title, err := objectTitleParser(tokenizer)
						if err != nil {
							return err
						}

						parsedTitle = true
						o.Object.Name = title
					}
				}
			}

			if t.Data == "table" {
				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "definition-table" {
						inDefinition = true
					}
				}

				if capturedObject {
					break TOKENIZER_LOOP
				}
			}

			if inDefinition {
				if t.Data == "tbody" {
					for _, attr := range t.Attr {
						if attr.Key == "class" && attr.Val == "supertype" {
							inProperties = true
						}
					}

					if capturedObject {
						break TOKENIZER_LOOP
					}
				}
			}

			if inProperties {

				if t.Data == "tr" {
					inField = true
					capturedObject = true
				}
			}

			if inField {

				if t.Data == "th" {
					for _, attr := range t.Attr {
						if attr.Key == "class" && attr.Val == "prop-nam" {
							label, url, err := fieldNameParser(tokenizer)
							if err != nil {
								return err
							}

							currentField.Name = label
							currentField.URL = url
						}
					}
				}
				if t.Data == "td" {
					for _, attr := range t.Attr {
						if attr.Key == "class" && attr.Val == "prop-ect" {
							fieldTypes, err := fieldTypeParser(tokenizer)
							if err != nil {
								return err
							}

							currentField.Types = fieldTypes
						}
					}
				}
				if t.Data == "td" {
					for _, attr := range t.Attr {
						if attr.Key == "class" && attr.Val == "prop-desc" {
							comment, err := fieldCommentParser(tokenizer)
							if err != nil {
								return err
							}

							currentField.Comment = comment
						}
					}
				}

			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "table" {
				inDefinition = false
			}

			if t.Data == "tbody" {
				inProperties = false
				break TOKENIZER_LOOP
			}

			if inDefinition && inProperties && inField {

				if t.Data == "tr" {
					inField = false
					o.Object.AddField(currentField)
					currentField = ast.Field{}
				}
			}

		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return err
	}

	return nil
}

func objectTitleParser(tokenizer *html.Tokenizer) (string, error) {

TOKENIZER_LOOP:
	for {
		tt := tokenizer.Next()
		if tokenizer.Err() != nil {
			break TOKENIZER_LOOP
		}

		switch {
		case tt == html.ErrorToken:
			break TOKENIZER_LOOP
		case tt == html.TextToken:
			t := tokenizer.Token()

			return t.Data, nil
		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = errors.New("unexpected EOF in sub parser")
	}
	if err != nil {
		return "", err
	}

	return "", nil
}
