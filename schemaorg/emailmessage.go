package schemaorg

import "encoding/json"

// EmailMessage see : https://schema.org/EmailMessage
type EmailMessage struct {

typeContext

Message

// BccRecipient see : /bccRecipient
// A sub property of recipient. The recipient blind copied on a message.
BccRecipient interface{} `json:"bccRecipient"` // types : ContactPoint Organization Person

// CcRecipient see : /ccRecipient
// A sub property of recipient. The recipient copied on a message.
CcRecipient interface{} `json:"ccRecipient"` // types : ContactPoint Organization Person

// DateRead see : /dateRead
// The date/time at which the message has been read by the recipient if a single recipient exists.
DateRead interface{} `json:"dateRead"`

// DateReceived see : /dateReceived
// The date/time the message was received if a single recipient exists.
DateReceived interface{} `json:"dateReceived"`

// DateSent see : /dateSent
// The date/time at which the message was sent.
DateSent interface{} `json:"dateSent"`

// MessageAttachment see : /messageAttachment
// A CreativeWork attached to the message.
MessageAttachment *CreativeWork `json:"messageAttachment"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

// Sender see : /sender
// A sub property of participant. The participant who is at the sending end of the action.
Sender interface{} `json:"sender"` // types : Audience Organization Person

// ToRecipient see : /toRecipient
// A sub property of recipient. The recipient who was directly sent the message.
ToRecipient interface{} `json:"toRecipient"` // types : Audience ContactPoint Organization Person

}

func (v *EmailMessage) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EmailMessage"

	return json.Marshal(v)
}
