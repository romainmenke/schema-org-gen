package schemaorg

import "encoding/json"

// ImageObject see : https://schema.org/ImageObject
type ImageObject struct {

typeContext

MediaObject

// Caption see : /caption
// The caption for this object.
Caption string `json:"caption"`

// ExifData see : /exifData
// exif data for this object.
ExifData interface{} `json:"exifData"` // types : PropertyValue Text

// RepresentativeOfPage see : /representativeOfPage
// Indicates whether this image is representative of the content of the page.
RepresentativeOfPage bool `json:"representativeOfPage"`

// Thumbnail see : /thumbnail
// Thumbnail image for an image or video.
Thumbnail *ImageObject `json:"thumbnail"`

}

func (v *ImageObject) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ImageObject"

	return json.Marshal(v)
}
