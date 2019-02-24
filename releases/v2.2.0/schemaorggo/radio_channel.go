package schemaorggo

import "encoding/json"

// RadioChannel see : https://schema.org/RadioChannel
type RadioChannel struct {

	// With properties from BroadcastChannel see : https://schema.org/BroadcastChannel
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// BroadcastChannelId see : https://schema.org/broadcastChannelId
	// The unique address by which the BroadcastService can be identified in a provider lineup. In US, this is typically a number.
	// types : Text
	BroadcastChannelId []string `json:"broadcastChannelId,omitempty"`

	// BroadcastServiceTier see : https://schema.org/broadcastServiceTier
	// The type of service required to have access to the channel (e.g. Standard or Premium).
	// types : Text
	BroadcastServiceTier []string `json:"broadcastServiceTier,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// InBroadcastLineup see : https://schema.org/inBroadcastLineup
	// The CableOrSatelliteService offering the channel.
	// types : CableOrSatelliteService
	InBroadcastLineup []*CableOrSatelliteService `json:"inBroadcastLineup,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// ProvidesBroadcastService see : https://schema.org/providesBroadcastService
	// The BroadcastService offered on this channel.
	// types : BroadcastService
	ProvidesBroadcastService []*BroadcastService `json:"providesBroadcastService,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v RadioChannel) MarshalJSON() ([]byte, error) {
	type Alias RadioChannel

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"RadioChannel\","), b[1:]...), nil
}
