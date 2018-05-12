package parsers

import (
	"errors"
	"fmt"
	"io"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"golang.org/x/net/html"
)

func fieldNameParser(tokenizer *html.Tokenizer) (string, string, error) {

	label := ""
	url := ""

	propNameLink := false

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

			if propNameLink {

				label = t.Data
				break TOKENIZER_LOOP
			}

		case tt == html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						url = attr.Val
					}
				}

				propNameLink = true
			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "a" {
				propNameLink = false
			}
		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = errors.New("unexpected EOF in sub parser")
	}
	if err != nil {
		return "", "", err
	}

	return label, url, nil

}

func fieldTypeParser(tokenizer *html.Tokenizer) ([]ast.FieldType, error) {

	fieldTypes := []ast.FieldType{}
	f := ast.FieldType{}

	fieldTypeLink := false

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

			if fieldTypeLink {

				f.Type = t.Data
			}

		case tt == html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						f.URL = attr.Val
					}
				}

				fieldTypeLink = true
			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "a" {
				fieldTypeLink = false
				fieldTypes = append(fieldTypes, f)
				f = ast.FieldType{}
			}

			if t.Data == "td" {
				break TOKENIZER_LOOP
			}
		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = errors.New("unexpected EOF in sub parser")
	}
	if err != nil {
		return nil, err
	}

	return fieldTypes, nil
}

func fieldCommentParser(tokenizer *html.Tokenizer) (string, error) {

	comment := ""
	link := ""

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

			comment = comment + t.Data

			if link != "" {
				comment = comment + fmt.Sprintf(" (see: %s)", link)
				link = ""
			}

		case tt == html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						link = "https://schema.org" + attr.Val
					}
				}
			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "td" {
				break TOKENIZER_LOOP
			}
		}
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = errors.New("unexpected EOF in sub parser")
	}
	if err != nil {
		return "", err
	}

	return comment, nil
}
