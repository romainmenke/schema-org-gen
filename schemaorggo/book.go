package schemaorggo

import "encoding/json"

// Book see : https://schema.org/Book
type Book struct {
	CreativeWork

	typeContext

	// Abridged see : http://bib.schema.org/abridged
	// Indicates whether the book is an abridged edition.
	Abridged bool `json:"abridged,omitempty"` // types : Boolean

	// BookEdition see : https://schema.org/bookEdition
	// The edition of the book.
	BookEdition string `json:"bookEdition,omitempty"` // types : Text

	// BookFormat see : https://schema.org/bookFormat
	// The format of the book.
	BookFormat *BookFormatType `json:"bookFormat,omitempty"` // types : BookFormatType

	// Illustrator see : https://schema.org/illustrator
	// The illustrator of the book.
	Illustrator *Person `json:"illustrator,omitempty"` // types : Person

	// Isbn see : https://schema.org/isbn
	// The ISBN of the book.
	Isbn string `json:"isbn,omitempty"` // types : Text

	// NumberOfPages see : https://schema.org/numberOfPages
	// The number of pages in the book.
	NumberOfPages float64 `json:"numberOfPages,omitempty"` // types : Integer

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
