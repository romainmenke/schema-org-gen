package schemaorggo

import "encoding/json"

// Barcode see : https://schema.org/Barcode
type Barcode struct {
	ImageObject

	typeContext

	// Caption see : https://schema.org/caption
	// The caption for this object.
	// types : Text
	Caption []string `json:"caption,omitempty"`

	// ExifData see : https://schema.org/exifData
	// exif data for this object.
	// types : PropertyValue Text
	ExifData []interface{} `json:"exifData,omitempty"`

	// RepresentativeOfPage see : https://schema.org/representativeOfPage
	// Indicates whether this image is representative of the content of the page.
	// types : Boolean
	RepresentativeOfPage []bool `json:"representativeOfPage,omitempty"`

	// Thumbnail see : https://schema.org/thumbnail
	// Thumbnail image for an image or video.
	// types : ImageObject
	Thumbnail []*ImageObject `json:"thumbnail,omitempty"`
}

func (v Barcode) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ImageObject.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Caption
		if len(v.Caption) == 1 {
			value = v.Caption[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["caption"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExifData
		if len(v.ExifData) == 1 {
			value = v.ExifData[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exifData"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RepresentativeOfPage
		if len(v.RepresentativeOfPage) == 1 {
			value = v.RepresentativeOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["representativeOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Thumbnail
		if len(v.Thumbnail) == 1 {
			value = v.Thumbnail[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["thumbnail"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Barcode) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Barcode"

	return data, nil
}

func (v Barcode) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
