package schemaorg

import "encoding/json"

// Book see : https://schema.org/Book
type Book struct {

typeContext

CreativeWork

// Abridged see : https://schema.orghttp://bib.schema.org/abridged
// Indicates whether the book is an abridged edition.
Abridged bool `json:"abridged"`

// BookEdition see : https://schema.org/bookEdition
// The edition of the book.
BookEdition string `json:"bookEdition"`

// BookFormat see : https://schema.org/bookFormat
// The format of the book.
BookFormat *BookFormatType `json:"bookFormat"`

// Illustrator see : https://schema.org/illustrator
// The illustrator of the book.
Illustrator *Person `json:"illustrator"`

// Isbn see : https://schema.org/isbn
// The ISBN of the book.
Isbn string `json:"isbn"`

// NumberOfPages see : https://schema.org/numberOfPages
// The number of pages in the book.
NumberOfPages int `json:"numberOfPages"`

}

func (v *Book) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Book"

	return json.Marshal(v)
}
