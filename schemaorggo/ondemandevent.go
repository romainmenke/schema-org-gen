package schemaorggo

import "encoding/json"

// OnDemandEvent see : https://schema.org/OnDemandEvent
type OnDemandEvent struct {
	PublicationEvent

	typeContext

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty"`

	// PublishedBy see : http://bib.schema.org/publishedBy
	// An agent associated with the publication event.
	PublishedBy interface{} `json:"publishedBy,omitempty"` // types : Organization Person

	// PublishedOn see : https://schema.org/publishedOn
	// A broadcast service associated with the publication event.
	PublishedOn *BroadcastService `json:"publishedOn,omitempty"`
}

func (v OnDemandEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OnDemandEvent"

	return json.Marshal(v)
}

func (v *OnDemandEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "OnDemandEvent"

	return json.Marshal(*v)
}
