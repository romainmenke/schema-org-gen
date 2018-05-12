package schemaorggo

import "encoding/json"

// Barcode see : https://schema.org/Barcode
type Barcode struct {
	ImageObject

	typeContext

	// Caption see : https://schema.org/caption
	// The caption for this object.
	Caption string `json:"caption,omitempty"` // types : Text

	// ExifData see : https://schema.org/exifData
	// exif data for this object.
	ExifData interface{} `json:"exifData,omitempty"` // types : PropertyValue Text

	// RepresentativeOfPage see : https://schema.org/representativeOfPage
	// Indicates whether this image is representative of the content of the page.
	RepresentativeOfPage bool `json:"representativeOfPage,omitempty"` // types : Boolean

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	Thumbnail *ImageObject `json:"thumbnail,omitempty"` // types : ImageObject

}

func (v Barcode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Barcode"

	return json.Marshal(v)
}

func (v *Barcode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Barcode"

	return json.Marshal(*v)
}
