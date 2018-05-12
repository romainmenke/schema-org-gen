package schemaorggo

import "encoding/json"

// VisualArtwork see : https://schema.org/VisualArtwork
type VisualArtwork struct {
	CreativeWork

	typeContext

	// ArtEdition see : https://schema.org/artEdition
	// The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, 'artEdition' refers to the total number of copies (in this example "20").
	ArtEdition interface{} `json:"artEdition,omitempty"` // types : Integer Text

	// ArtMedium see : https://schema.org/artMedium
	// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
	ArtMedium interface{} `json:"artMedium,omitempty"` // types : Text URL

	// Artform see : https://schema.org/artform
	// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
	Artform interface{} `json:"artform,omitempty"` // types : Text URL

	// Artist see : http://bib.schema.org/artist
	// The primary artist for a work
	//     in a medium other than pencils or digital line art--for example, if the
	//     primary artwork is done in watercolors or digital paints.
	Artist *Person `json:"artist,omitempty"`

	// ArtworkSurface see : https://schema.org/artworkSurface
	// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc. Supersedes surface (see: https://schema.org/surface).
	ArtworkSurface interface{} `json:"artworkSurface,omitempty"` // types : Text URL

	// Colorist see : http://bib.schema.org/colorist
	// The individual who adds color to inked drawings.
	Colorist *Person `json:"colorist,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	Depth interface{} `json:"depth,omitempty"` // types : Distance QuantitativeValue

	// Height see : https://schema.org/height
	// The height of the item.
	Height interface{} `json:"height,omitempty"` // types : Distance QuantitativeValue

	// Inker see : http://bib.schema.org/inker
	// The individual who traces over the pencil drawings in ink after pencils are complete.
	Inker *Person `json:"inker,omitempty"`

	// Letterer see : http://bib.schema.org/letterer
	// The individual who adds lettering, including speech balloons and sound effects, to artwork.
	Letterer *Person `json:"letterer,omitempty"`

	// Penciler see : http://bib.schema.org/penciler
	// The individual who draws the primary narrative artwork.
	Penciler *Person `json:"penciler,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	Width interface{} `json:"width,omitempty"` // types : Distance QuantitativeValue

}

func (v VisualArtwork) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VisualArtwork"

	return json.Marshal(v)
}

func (v *VisualArtwork) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "VisualArtwork"

	return json.Marshal(*v)
}
