package schemaorg

import "encoding/json"

// MediaObject see : https://schema.org/MediaObject
type MediaObject struct {

typeContext

CreativeWork

// AssociatedArticle see : /associatedArticle
// A NewsArticle associated with the Media Object.
AssociatedArticle *NewsArticle `json:"associatedArticle"`

// Bitrate see : /bitrate
// The bitrate of the media object.
Bitrate string `json:"bitrate"`

// ContentSize see : /contentSize
// File size in (mega/kilo) bytes.
ContentSize string `json:"contentSize"`

// ContentUrl see : /contentUrl
// Actual bytes of the media object, for example the image file or video file.
ContentUrl string `json:"contentUrl"`

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration *Duration `json:"duration"`

// EmbedUrl see : /embedUrl
// A URL pointing to a player for a specific video. In general, this is the information in the src element of an embed tag and should not be the same as the content of the loc tag.
EmbedUrl string `json:"embedUrl"`

// EncodesCreativeWork see : /encodesCreativeWork
// The CreativeWork encoded by this media object.
EncodesCreativeWork *CreativeWork `json:"encodesCreativeWork"`

// EncodingFormat see : /encodingFormat
// mp3, mpeg4, etc.
EncodingFormat string `json:"encodingFormat"`

// Height see : /height
// The height of the item.
Height interface{} `json:"height"` // types : Distance QuantitativeValue

// PlayerType see : /playerType
// Player type required—for example, Flash or Silverlight.
PlayerType string `json:"playerType"`

// ProductionCompany see : /productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// RegionsAllowed see : /regionsAllowed
// The regions where the media is allowed. If not specified, then it's assumed to be allowed everywhere. Specify the countries in ISO 3166 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166).
RegionsAllowed *Place `json:"regionsAllowed"`

// RequiresSubscription see : /requiresSubscription
// Indicates if use of the media require a subscription  (either paid or free). Allowed values are true or false (note that an earlier version had 'yes', 'no').
RequiresSubscription bool `json:"requiresSubscription"`

// UploadDate see : /uploadDate
// Date when this media object was uploaded to this site.
UploadDate interface{} `json:"uploadDate"`

// Width see : /width
// The width of the item.
Width interface{} `json:"width"` // types : Distance QuantitativeValue

}

func (v *MediaObject) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MediaObject"

	return json.Marshal(v)
}
