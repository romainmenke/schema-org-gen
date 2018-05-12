package schemaorg

import "encoding/json"

// PublicationEvent see : https://schema.org/PublicationEvent
type PublicationEvent struct {

Event

typeContext

// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree bool `json:"isAccessibleForFree"`

// PublishedBy see : http://bib.schema.org/publishedBy
// An agent associated with the publication event.
PublishedBy interface{} `json:"publishedBy"` // types : Organization Person

// PublishedOn see : https://schema.org/publishedOn
// A broadcast service associated with the publication event.
PublishedOn *BroadcastService `json:"publishedOn"`

}

func (v *PublicationEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PublicationEvent"

	return json.Marshal(v)
}
