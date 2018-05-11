package schemaorg

import "encoding/json"

// Barcode see : https://schema.org/Barcode
type Barcode struct {

typeContext

ImageObject

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

func (v *Barcode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Barcode"

	return json.Marshal(v)
}
