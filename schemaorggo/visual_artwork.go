package schemaorggo

import "encoding/json"

// VisualArtwork see : https://schema.org/VisualArtwork
type VisualArtwork struct {
	CreativeWork

	typeContext

	// ArtEdition see : https://schema.org/artEdition
	// The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, &#39;artEdition&#39; refers to the total number of copies (in this example &quot;20&quot;).
	// types : Integer Text
	ArtEdition []interface{} `json:"artEdition,omitempty"`

	// ArtMedium see : https://schema.org/artMedium
	// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
	// types : Text URL
	ArtMedium []string `json:"artMedium,omitempty"`

	// Artform see : https://schema.org/artform
	// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
	// types : Text URL
	Artform []string `json:"artform,omitempty"`

	// Artist see : http://bib.schema.org/artist
	// The primary artist for a work
	//     in a medium other than pencils or digital line art--for example, if the
	//     primary artwork is done in watercolors or digital paints.
	// types : Person
	Artist []*Person `json:"artist,omitempty"`

	// ArtworkSurface see : https://schema.org/artworkSurface
	// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc. Supersedes surface (see: https://schema.org/surface).
	// types : Text URL
	ArtworkSurface []string `json:"artworkSurface,omitempty"`

	// Colorist see : http://bib.schema.org/colorist
	// The individual who adds color to inked drawings.
	// types : Person
	Colorist []*Person `json:"colorist,omitempty"`

	// Depth see : https://schema.org/depth
	// The depth of the item.
	// types : Distance QuantitativeValue
	Depth []interface{} `json:"depth,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// Inker see : http://bib.schema.org/inker
	// The individual who traces over the pencil drawings in ink after pencils are complete.
	// types : Person
	Inker []*Person `json:"inker,omitempty"`

	// Letterer see : http://bib.schema.org/letterer
	// The individual who adds lettering, including speech balloons and sound effects, to artwork.
	// types : Person
	Letterer []*Person `json:"letterer,omitempty"`

	// Penciler see : http://bib.schema.org/penciler
	// The individual who draws the primary narrative artwork.
	// types : Person
	Penciler []*Person `json:"penciler,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width []interface{} `json:"width,omitempty"`
}

func (v VisualArtwork) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ArtEdition
		if len(v.ArtEdition) == 1 {
			value = v.ArtEdition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["artEdition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArtMedium
		if len(v.ArtMedium) == 1 {
			value = v.ArtMedium[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["artMedium"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Artform
		if len(v.Artform) == 1 {
			value = v.Artform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["artform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Artist
		if len(v.Artist) == 1 {
			value = v.Artist[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["artist"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ArtworkSurface
		if len(v.ArtworkSurface) == 1 {
			value = v.ArtworkSurface[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["artworkSurface"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Colorist
		if len(v.Colorist) == 1 {
			value = v.Colorist[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["colorist"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Depth
		if len(v.Depth) == 1 {
			value = v.Depth[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["depth"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Height
		if len(v.Height) == 1 {
			value = v.Height[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["height"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Inker
		if len(v.Inker) == 1 {
			value = v.Inker[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inker"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Letterer
		if len(v.Letterer) == 1 {
			value = v.Letterer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["letterer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Penciler
		if len(v.Penciler) == 1 {
			value = v.Penciler[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["penciler"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Width
		if len(v.Width) == 1 {
			value = v.Width[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["width"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v VisualArtwork) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "VisualArtwork"

	return data, nil
}

func (v VisualArtwork) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
