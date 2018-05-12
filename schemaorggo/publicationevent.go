package schemaorggo

import "encoding/json"

// PublicationEvent see : https://schema.org/PublicationEvent
type PublicationEvent struct {
	Event

	typeContext

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty"`

	// PublishedBy see : http://bib.schema.org/publishedBy
	// An agent associated with the publication event.
	// types : Organization Person
	PublishedBy interface{} `json:"publishedBy,omitempty"`

	// PublishedOn see : https://schema.org/publishedOn
	// A broadcast service associated with the publication event.
	// types : BroadcastService
	PublishedOn *BroadcastService `json:"publishedOn,omitempty"`
}

func (v PublicationEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PublicationEvent"

	return json.Marshal(v)
}

func (v *PublicationEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PublicationEvent"

	return json.Marshal(*v)
}
