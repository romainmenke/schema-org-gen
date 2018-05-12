package schemaorggo

import "encoding/json"

// ImageObject see : https://schema.org/ImageObject
type ImageObject struct {
	MediaObject

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
