package schemaorggo

import "encoding/json"

// Barcode see : https://schema.org/Barcode
type Barcode struct {
	ImageObject

	typeContext

	// Caption see : https://schema.org/caption
	// The caption for this object.
	// types : Text
	Caption string `json:"caption,omitempty"`

	// ExifData see : https://schema.org/exifData
	// exif data for this object.
	// types : PropertyValue Text
	ExifData interface{} `json:"exifData,omitempty"`

	// RepresentativeOfPage see : https://schema.org/representativeOfPage
	// Indicates whether this image is representative of the content of the page.
	// types : Boolean
	RepresentativeOfPage bool `json:"representativeOfPage,omitempty"`

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	// types : ImageObject
	Thumbnail *ImageObject `json:"thumbnail,omitempty"`
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
