package schemaorg

import "encoding/json"

// OnDemandEvent see : https://schema.org/OnDemandEvent
type OnDemandEvent struct {

typeContext

PublicationEvent

// IsAccessibleForFree see : /isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree bool `json:"isAccessibleForFree"`

// PublishedBy see : http://bib.schema.org/publishedBy
// An agent associated with the publication event.
PublishedBy interface{} `json:"publishedBy"` // types : Organization Person

// PublishedOn see : /publishedOn
// A broadcast service associated with the publication event.
PublishedOn *BroadcastService `json:"publishedOn"`

}

func (v *OnDemandEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "OnDemandEvent"

	return json.Marshal(v)
}
