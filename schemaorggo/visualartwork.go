package schemaorggo

import "encoding/json"

// VisualArtwork see : https://schema.org/VisualArtwork
type VisualArtwork struct {
	CreativeWork

	typeContext

	// ArtEdition see : https://schema.org/artEdition
	// The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, &#39;artEdition&#39; refers to the total number of copies (in this example &quot;20&quot;).
	// types : Integer Text
	ArtEdition interface{} `json:"artEdition,omitempty"`

	// ArtMedium see : https://schema.org/artMedium
	// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
	// types : Text URL
	ArtMedium string `json:"artMedium,omitempty"`

	// Artform see : https://schema.org/artform
	// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
	// types : Text URL
	Artform string `json:"artform,omitempty"`

	// Artist see : http://bib.schema.org/artist
	// The primary artist for a work
	//     in a medium other than pencils or digital line art--for example, if the
	//     primary artwork is done in watercolors or digital paints.
	// types : Person
	Artist *Person `json:"artist,omitempty"`

	// ArtworkSurface see : https://schema.org/artworkSurface
	// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc. Supersedes surface (see: https://schema.org/surface).
	// types : Text URL
	ArtworkSurface string `json:"artworkSurface,omitempty"`

	// Colorist see : http://bib.schema.org/colorist
	// The individual who adds color to inked drawings.
	// types : Person
	Colorist *Person `json:"colorist,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	// types : Distance QuantitativeValue
	Depth interface{} `json:"depth,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height interface{} `json:"height,omitempty"`

	// Inker see : http://bib.schema.org/inker
	// The individual who traces over the pencil drawings in ink after pencils are complete.
	// types : Person
	Inker *Person `json:"inker,omitempty"`

	// Letterer see : http://bib.schema.org/letterer
	// The individual who adds lettering, including speech balloons and sound effects, to artwork.
	// types : Person
	Letterer *Person `json:"letterer,omitempty"`

	// Penciler see : http://bib.schema.org/penciler
	// The individual who draws the primary narrative artwork.
	// types : Person
	Penciler *Person `json:"penciler,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width interface{} `json:"width,omitempty"`
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
