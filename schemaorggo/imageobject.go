package schemaorggo

import "encoding/json"

// ImageObject see : https://schema.org/ImageObject
type ImageObject struct {
	MediaObject

	typeContext

	// Caption see : https://schema.org/caption
	// The caption for this object.
	Caption string `json:"caption,omitempty"`

	// ExifData see : https://schema.org/exifData
	// exif data for this object.
	ExifData interface{} `json:"exifData,omitempty"` // types : PropertyValue Text

	// RepresentativeOfPage see : https://schema.org/representativeOfPage
	// Indicates whether this image is representative of the content of the page.
	RepresentativeOfPage bool `json:"representativeOfPage,omitempty"`

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	Thumbnail *ImageObject `json:"thumbnail,omitempty"`
}

func (v ImageObject) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ImageObject"

	return json.Marshal(v)
}

func (v *ImageObject) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ImageObject"

	return json.Marshal(*v)
}
