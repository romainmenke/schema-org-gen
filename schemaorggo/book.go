package schemaorggo

import "encoding/json"

// Book see : https://schema.org/Book
type Book struct {
	CreativeWork

	typeContext

	// Abridged see : http://bib.schema.org/abridged
	// Indicates whether the book is an abridged edition.
	// types : Boolean
	Abridged []bool `json:"abridged,omitempty"`

	// BookEdition see : https://schema.org/bookEdition
	// The edition of the book.
	// types : Text
	BookEdition []string `json:"bookEdition,omitempty"`

	// BookFormat see : https://schema.org/bookFormat
	// The format of the book.
	// types : BookFormatType
	BookFormat []*BookFormatType `json:"bookFormat,omitempty"`

	// Illustrator see : https://schema.org/illustrator
	// The illustrator of the book.
	// types : Person
	Illustrator []*Person `json:"illustrator,omitempty"`

	// Isbn see : https://schema.org/isbn
	// The ISBN of the book.
	// types : Text
	Isbn []string `json:"isbn,omitempty"`

	// NumberOfPages see : https://schema.org/numberOfPages
	// The number of pages in the book.
	// types : Integer
	NumberOfPages []float64 `json:"numberOfPages,omitempty"`
}

func (v Book) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Abridged
		if len(v.Abridged) == 1 {
			value = v.Abridged[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["abridged"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BookEdition
		if len(v.BookEdition) == 1 {
			value = v.BookEdition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bookEdition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BookFormat
		if len(v.BookFormat) == 1 {
			value = v.BookFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bookFormat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Illustrator
		if len(v.Illustrator) == 1 {
			value = v.Illustrator[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["illustrator"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Isbn
		if len(v.Isbn) == 1 {
			value = v.Isbn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isbn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfPages
		if len(v.NumberOfPages) == 1 {
			value = v.NumberOfPages[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfPages"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Book) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Book"

	return data, nil
}

func (v Book) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
