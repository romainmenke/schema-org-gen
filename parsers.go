package main

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func objectParser(tokenizer *html.Tokenizer, o *ObjectSource) error {
	o.Object = &Object{}
	inDefinition := false
	inProperties := false
	inField := false

	parsedTitle := false

	capturedObject := false

	o.Object.URL = "https://schema.org" + o.URL
	if o.Parent != nil && o.Parent.Object != nil {
		o.Object.ParentObject = o.Parent.Object
	}

	currentField := Field{}

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
					o.Object.Fields = append(o.Object.Fields, currentField)
					currentField = Field{}
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

func fieldTypeParser(tokenizer *html.Tokenizer) ([]FieldType, error) {

	fieldTypes := []FieldType{}
	f := FieldType{}

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
				f = FieldType{}
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
