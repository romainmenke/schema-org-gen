package schemaorggo

import "encoding/json"

// MediaObject see : https://schema.org/MediaObject
type MediaObject struct {
	CreativeWork

	typeContext

	// AssociatedArticle see : https://schema.org/associatedArticle
	// A NewsArticle associated with the Media Object.
	// types : NewsArticle
	AssociatedArticle []*NewsArticle `json:"associatedArticle,omitempty"`

	// Bitrate see : https://schema.org/bitrate
	// The bitrate of the media object.
	// types : Text
	Bitrate []string `json:"bitrate,omitempty"`

	// ContentSize see : https://schema.org/contentSize
	// File size in (mega/kilo) bytes.
	// types : Text
	ContentSize []string `json:"contentSize,omitempty"`

	// ContentUrl see : https://schema.org/contentUrl
	// Actual bytes of the media object, for example the image file or video file.
	// types : URL
	ContentUrl []string `json:"contentUrl,omitempty"`

	// Duration see : https://pending.schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration []*Duration `json:"duration,omitempty"`

	// EmbedUrl see : https://schema.org/embedUrl
	// A URL pointing to a player for a specific video. In general, this is the information in the src element of an embed tag and should not be the same as the content of the loc tag.
	// types : URL
	EmbedUrl []string `json:"embedUrl,omitempty"`

	// EncodesCreativeWork see : https://schema.org/encodesCreativeWork
	// The CreativeWork encoded by this media object.
	// types : CreativeWork
	EncodesCreativeWork []*CreativeWork `json:"encodesCreativeWork,omitempty"`

	// EncodingFormat see : https://schema.org/encodingFormat
	// Media type typically expressed using a MIME format (see IANA site (see: https://schema.orghttp://www.iana.org/assignments/media-types/media-types.xhtml) and MDN reference (see: https://schema.orghttps://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)) e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.).
	//
	// In cases where a CreativeWork (see: https://schema.org/CreativeWork) has several media type representations, encoding (see: https://schema.org/encoding) can be used to indicate each MediaObject (see: https://schema.org/MediaObject) alongside particular encodingFormat (see: https://schema.org/encodingFormat) information.
	//
	// Unregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry. Supersedes fileFormat (see: https://schema.org/fileFormat).
	// types : Text URL
	EncodingFormat []string `json:"encodingFormat,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// PlayerType see : https://schema.org/playerType
	// Player type required—for example, Flash or Silverlight.
	// types : Text
	PlayerType []string `json:"playerType,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany []*Organization `json:"productionCompany,omitempty"`

	// RegionsAllowed see : https://schema.org/regionsAllowed
	// The regions where the media is allowed. If not specified, then it&#39;s assumed to be allowed everywhere. Specify the countries in ISO 3166 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166).
	// types : Place
	RegionsAllowed []*Place `json:"regionsAllowed,omitempty"`

	// RequiresSubscription see : https://pending.schema.org/requiresSubscription
	// Indicates if use of the media require a subscription  (either paid or free). Allowed values are true or false (note that an earlier version had &#39;yes&#39;, &#39;no&#39;).
	// types : Boolean MediaSubscription
	RequiresSubscription []interface{} `json:"requiresSubscription,omitempty"`

	// UploadDate see : https://schema.org/uploadDate
	// Date when this media object was uploaded to this site.
	// types : Date
	UploadDate []Date `json:"uploadDate,omitempty"`

	// Width see : https://schema.org/width
	// The width of the item.
	// types : Distance QuantitativeValue
	Width []interface{} `json:"width,omitempty"`
}

func (v MediaObject) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AssociatedArticle
		if len(v.AssociatedArticle) == 1 {
			value = v.AssociatedArticle[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["associatedArticle"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Bitrate
		if len(v.Bitrate) == 1 {
			value = v.Bitrate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bitrate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentSize
		if len(v.ContentSize) == 1 {
			value = v.ContentSize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentSize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentUrl
		if len(v.ContentUrl) == 1 {
			value = v.ContentUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentUrl"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Duration
		if len(v.Duration) == 1 {
			value = v.Duration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["duration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EmbedUrl
		if len(v.EmbedUrl) == 1 {
			value = v.EmbedUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["embedUrl"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EncodesCreativeWork
		if len(v.EncodesCreativeWork) == 1 {
			value = v.EncodesCreativeWork[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["encodesCreativeWork"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EncodingFormat
		if len(v.EncodingFormat) == 1 {
			value = v.EncodingFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["encodingFormat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Height
		if len(v.Height) == 1 {
			value = v.Height[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["height"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PlayerType
		if len(v.PlayerType) == 1 {
			value = v.PlayerType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["playerType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductionCompany
		if len(v.ProductionCompany) == 1 {
			value = v.ProductionCompany[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productionCompany"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RegionsAllowed
		if len(v.RegionsAllowed) == 1 {
			value = v.RegionsAllowed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["regionsAllowed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RequiresSubscription
		if len(v.RequiresSubscription) == 1 {
			value = v.RequiresSubscription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["requiresSubscription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UploadDate
		if len(v.UploadDate) == 1 {
			value = v.UploadDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["uploadDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Width
		if len(v.Width) == 1 {
			value = v.Width[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["width"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MediaObject) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MediaObject"

	return data, nil
}

func (v MediaObject) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
