package schemaorggo

import "encoding/json"

// Book see : https://schema.org/Book
type Book struct {
	CreativeWork

	typeContext

	// Abridged see : http://bib.schema.org/abridged
	// Indicates whether the book is an abridged edition.
	// types : Boolean
	Abridged bool `json:"abridged,omitempty"`

	// BookEdition see : https://schema.org/bookEdition
	// The edition of the book.
	// types : Text
	BookEdition string `json:"bookEdition,omitempty"`

	// BookFormat see : https://schema.org/bookFormat
	// The format of the book.
	// types : BookFormatType
	BookFormat *BookFormatType `json:"bookFormat,omitempty"`

	// Illustrator see : https://schema.org/illustrator
	// The illustrator of the book.
	// types : Person
	Illustrator *Person `json:"illustrator,omitempty"`

	// Isbn see : https://schema.org/isbn
	// The ISBN of the book.
	// types : Text
	Isbn string `json:"isbn,omitempty"`

	// NumberOfPages see : https://schema.org/numberOfPages
	// The number of pages in the book.
	// types : Integer
	NumberOfPages float64 `json:"numberOfPages,omitempty"`
}

func (v Book) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Book"

	return json.Marshal(v)
}

func (v *Book) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Book"

	return json.Marshal(*v)
}
