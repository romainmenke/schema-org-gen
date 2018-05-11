package schemaorg

import "encoding/json"

// VisualArtwork see : https://schema.org/VisualArtwork
type VisualArtwork struct {

typeContext

CreativeWork

// ArtEdition see : /artEdition
// The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, 'artEdition' refers to the total number of copies (in this example "20").
ArtEdition interface{} `json:"artEdition"` // types : Integer Text

// ArtMedium see : /artMedium
// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
ArtMedium interface{} `json:"artMedium"` // types : Text URL

// Artform see : /artform
// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
Artform interface{} `json:"artform"` // types : Text URL

// Artist see : http://bib.schema.org/artist
// The primary artist for a work
//     in a medium other than pencils or digital line art--for example, if the
//     primary artwork is done in watercolors or digital paints.
Artist *Person `json:"artist"`

// ArtworkSurface see : /artworkSurface
// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc. Supersedes surface (see: https://schema.org/surface).
ArtworkSurface interface{} `json:"artworkSurface"` // types : Text URL

// Colorist see : http://bib.schema.org/colorist
// The individual who adds color to inked drawings.
Colorist *Person `json:"colorist"`

// Depth see : /depth
// The depth of the item.
Depth interface{} `json:"depth"` // types : Distance QuantitativeValue

// Height see : /height
// The height of the item.
Height interface{} `json:"height"` // types : Distance QuantitativeValue

// Inker see : http://bib.schema.org/inker
// The individual who traces over the pencil drawings in ink after pencils are complete.
Inker *Person `json:"inker"`

// Letterer see : http://bib.schema.org/letterer
// The individual who adds lettering, including speech balloons and sound effects, to artwork.
Letterer *Person `json:"letterer"`

// Penciler see : http://bib.schema.org/penciler
// The individual who draws the primary narrative artwork.
Penciler *Person `json:"penciler"`

// Width see : /width
// The width of the item.
Width interface{} `json:"width"` // types : Distance QuantitativeValue

}

func (v *VisualArtwork) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VisualArtwork"

	return json.Marshal(v)
}
