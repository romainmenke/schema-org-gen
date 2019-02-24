package parsers

import (
	"io"
	"log"
	"time"

	"github.com/romainmenke/schema-org-gen/internal/ast"
	"golang.org/x/net/html"
)

func RDFAParser(tokenizer *html.Tokenizer) (ast.Typemap, error) {

	start := time.Now()

	var partials []rdfaPartial

TOKENIZER_LOOP:
	for {
		partial, err := rdfaPartialParser(tokenizer)
		if err == io.EOF {
			break TOKENIZER_LOOP
		}
		if err != nil {
			return nil, err
		}
		if partial.label == "" {
			// log.Println(" --- > empty : ", partial.comment, partial.typeOf, partial.url)
			continue TOKENIZER_LOOP
		}

		partials = append(partials, partial)
	}
	err := tokenizer.Err()
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	var classes []rdfaPartial
	var properties []rdfaPartial

	for _, partial := range partials {
		if partial.typeOf == "rdfs:Class" {
			classes = append(classes, partial)
			continue
		}

		if partial.typeOf == "rdf:Property" {
			properties = append(properties, partial)
			continue
		}
	}

	fields := make(map[string][]ast.Field)

	for _, property := range properties {
		fieldTypes := []ast.FieldType{}
		for _, rangeIncludes := range property.rangeIncludes {
			fieldTypeURL := ""
			for _, class := range classes {
				if class.label == rangeIncludes {
					fieldTypeURL = class.url
				}
			}

			fieldTypes = append(fieldTypes, ast.FieldType{
				Type: rangeIncludes,
				URL:  fieldTypeURL,
			})
		}

		field := ast.Field{
			Comment: property.comment,
			URL:     property.url,
			Name:    property.label,
			Types:   fieldTypes,
		}

		for _, domainIncludes := range property.domainIncludes {
			fields[domainIncludes] = append(fields[domainIncludes], field)
		}
	}

	objects := []ast.Object{}
	for _, class := range classes {
		object := ast.Object{
			Comment: class.comment,
			URL:     class.url,
			Name:    class.label,
		}

		if objectFields, ok := fields[class.label]; ok {
			for _, field := range objectFields {
				object = ast.AddField(object, field)
			}
		}

		for _, subClass := range class.subClassOf {
			object = ast.AddSuperClass(object, subClass)
		}

		objects = append(objects, object)
	}

	tm := ast.Typemap(objects)

	for i, object := range tm {
		object.SubClassOf = ast.SuperClasses(tm, object)
		for _, super := range object.SubClassOf {

			for _, field := range super.Fields {
				object = ast.AddField(object, field)
				tm[i] = object
			}
		}
	}

	log.Println(tm.Len())
	log.Println(time.Since(start))

	return tm, nil
}

type rdfaPartial struct {
	typeOf         string
	url            string
	label          string
	comment        string
	subClassOf     []string
	domainIncludes []string
	rangeIncludes  []string
}

func rdfaPartialParser(tokenizer *html.Tokenizer) (rdfaPartial, error) {

	var partial = rdfaPartial{}
	var innerHTMLSetter func(string)

TOKENIZER_LOOP:
	for {
		tt := tokenizer.Next()
		if tokenizer.Err() != nil {
			break TOKENIZER_LOOP
		}

		switch {
		case tt == html.TextToken:
			t := tokenizer.Token()

			if innerHTMLSetter != nil {
				innerHTMLSetter(t.Data)
				innerHTMLSetter = nil
				continue TOKENIZER_LOOP
			}
		case tt == html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == "div" {
				for _, attr := range t.Attr {
					if attr.Key == "typeof" {
						partial.typeOf = attr.Val
					}

					if attr.Key == "resource" {
						partial.url = attr.Val
					}
				}

				continue TOKENIZER_LOOP
			}

			if t.Data == "span" {
				for _, attr := range t.Attr {
					if attr.Key == "property" && attr.Val == "rdfs:label" {
						innerHTMLSetter = func(str string) {
							partial.label = str
						}
					}

					if attr.Key == "property" && attr.Val == "rdfs:comment" {
						innerHTMLSetter = func(str string) {
							partial.comment = str
						}
					}
				}

				continue TOKENIZER_LOOP
			}

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "property" && attr.Val == "rdfs:subClassOf" {
						innerHTMLSetter = func(str string) {
							partial.subClassOf = append(partial.subClassOf, str)
						}
					}

					if attr.Key == "property" && attr.Val == "http://schema.org/domainIncludes" {
						innerHTMLSetter = func(str string) {
							partial.domainIncludes = append(partial.domainIncludes, str)
						}
					}

					if attr.Key == "property" && attr.Val == "http://schema.org/rangeIncludes" {
						innerHTMLSetter = func(str string) {
							partial.rangeIncludes = append(partial.rangeIncludes, str)
						}
					}
				}

				continue TOKENIZER_LOOP
			}

		case tt == html.EndTagToken:
			t := tokenizer.Token()

			if t.Data == "div" {
				break TOKENIZER_LOOP
			}
		}
	}
	err := tokenizer.Err()
	if err != nil {
		return partial, err
	}

	return partial, nil
}
