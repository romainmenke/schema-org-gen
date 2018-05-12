package schemaorg

import "encoding/json"

// Barcode see : https://schema.org/Barcode
type Barcode struct {

ImageObject

typeContext

// Caption see : https://schema.org/caption
// The caption for this object.
Caption string `json:"caption"`

// ExifData see : https://schema.org/exifData
// exif data for this object.
ExifData interface{} `json:"exifData"` // types : PropertyValue Text

// RepresentativeOfPage see : https://schema.org/representativeOfPage
// Indicates whether this image is representative of the content of the page.
RepresentativeOfPage bool `json:"representativeOfPage"`

// Thumbnail see : https://schema.org/thumbnail
// Thumbnail image for an image or video.
Thumbnail *ImageObject `json:"thumbnail"`

}

func (v *Barcode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Barcode"

	return json.Marshal(v)
}
