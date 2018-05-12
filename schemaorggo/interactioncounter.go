package schemaorggo

import "encoding/json"

// InteractionCounter see : https://schema.org/InteractionCounter
type InteractionCounter struct {
	StructuredValue

	typeContext

	// InteractionService see : https://schema.org/interactionService
	// The WebSite or SoftwareApplication where the interactions took place.
	InteractionService interface{} `json:"interactionService,omitempty"` // types : SoftwareApplication WebSite

	// InteractionType see : https://schema.org/interactionType
	// The Action representing the type of interaction. For up votes, +1s, etc. use LikeAction (see: https://schema.org/LikeAction). For down votes use DislikeAction (see: https://schema.org/DislikeAction). Otherwise, use the most specific Action.
	InteractionType *Action `json:"interactionType,omitempty"`

	// UserInteractionCount see : https://schema.org/userInteractionCount
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication.
	UserInteractionCount int `json:"userInteractionCount,omitempty"`
}

func (v InteractionCounter) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InteractionCounter"

	return json.Marshal(v)
}

func (v *InteractionCounter) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "InteractionCounter"

	return json.Marshal(*v)
}
