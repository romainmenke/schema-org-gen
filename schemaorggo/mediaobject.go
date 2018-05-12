package schemaorggo

import "encoding/json"

// MediaObject see : https://schema.org/MediaObject
type MediaObject struct {
	CreativeWork

	typeContext

	// AssociatedArticle see : https://schema.org/associatedArticle
	// A NewsArticle associated with the Media Object.
	AssociatedArticle *NewsArticle `json:"associatedArticle,omitempty"` // types : NewsArticle

	// Bitrate see : https://schema.org/bitrate
	// The bitrate of the media object.
	Bitrate string `json:"bitrate,omitempty"` // types : Text

	// ContentSize see : https://schema.org/contentSize
	// File size in (mega/kilo) bytes.
	ContentSize string `json:"contentSize,omitempty"` // types : Text

	// ContentUrl see : https://schema.org/contentUrl
	// Actual bytes of the media object, for example the image file or video file.
	ContentUrl string `json:"contentUrl,omitempty"` // types : URL

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	Duration *Duration `json:"duration,omitempty"` // types : Duration

	// EmbedUrl see : https://schema.org/embedUrl
	// A URL pointing to a player for a specific video. In general, this is the information in the src element of an embed tag and should not be the same as the content of the loc tag.
	EmbedUrl string `json:"embedUrl,omitempty"` // types : URL

	// EncodesCreativeWork see : https://schema.org/encodesCreativeWork
	// The CreativeWork encoded by this media object.
	EncodesCreativeWork *CreativeWork `json:"encodesCreativeWork,omitempty"` // types : CreativeWork

	// EncodingFormat see : https://schema.org/encodingFormat
	// mp3, mpeg4, etc.
	EncodingFormat string `json:"encodingFormat,omitempty"` // types : Text

	// Height see : https://schema.org/height
	// The height of the item.
	Height interface{} `json:"height,omitempty"` // types : Distance QuantitativeValue

	// PlayerType see : https://schema.org/playerType
	// Player type required—for example, Flash or Silverlight.
	PlayerType string `json:"playerType,omitempty"` // types : Text

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	ProductionCompany *Organization `json:"productionCompany,omitempty"` // types : Organization

	// RegionsAllowed see : https://schema.org/regionsAllowed
	// The regions where the media is allowed. If not specified, then it's assumed to be allowed everywhere. Specify the countries in ISO 3166 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166).
	RegionsAllowed *Place `json:"regionsAllowed,omitempty"` // types : Place

	// RequiresSubscription see : https://schema.org/requiresSubscription
	// Indicates if use of the media require a subscription  (either paid or free). Allowed values are true or false (note that an earlier version had 'yes', 'no').
	RequiresSubscription bool `json:"requiresSubscription,omitempty"` // types : Boolean

	// UploadDate see : https://schema.org/uploadDate
	// Date when this media object was uploaded to this site.
	UploadDate Date `json:"uploadDate,omitempty"` // types : Date

	// Width see : https://schema.org/width
	// The width of the item.
	Width interface{} `json:"width,omitempty"` // types : Distance QuantitativeValue

}

func (v MediaObject) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MediaObject"

	return json.Marshal(v)
}

func (v *MediaObject) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MediaObject"

	return json.Marshal(*v)
}
