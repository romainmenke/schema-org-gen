package schemaorg

import "encoding/json"

// EmailMessage see : https://schema.org/EmailMessage
type EmailMessage struct {

Message

typeContext

// BccRecipient see : https://schema.org/bccRecipient
// A sub property of recipient. The recipient blind copied on a message.
BccRecipient interface{} `json:"bccRecipient"` // types : ContactPoint Organization Person

// CcRecipient see : https://schema.org/ccRecipient
// A sub property of recipient. The recipient copied on a message.
CcRecipient interface{} `json:"ccRecipient"` // types : ContactPoint Organization Person

// DateRead see : https://schema.org/dateRead
// The date/time at which the message has been read by the recipient if a single recipient exists.
DateRead interface{} `json:"dateRead"`

// DateReceived see : https://schema.org/dateReceived
// The date/time the message was received if a single recipient exists.
DateReceived interface{} `json:"dateReceived"`

// DateSent see : https://schema.org/dateSent
// The date/time at which the message was sent.
DateSent interface{} `json:"dateSent"`

// MessageAttachment see : https://schema.org/messageAttachment
// A CreativeWork attached to the message.
MessageAttachment *CreativeWork `json:"messageAttachment"`

// Recipient see : https://schema.org/recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

// Sender see : https://schema.org/sender
// A sub property of participant. The participant who is at the sending end of the action.
Sender interface{} `json:"sender"` // types : Audience Organization Person

// ToRecipient see : https://schema.org/toRecipient
// A sub property of recipient. The recipient who was directly sent the message.
ToRecipient interface{} `json:"toRecipient"` // types : Audience ContactPoint Organization Person

}

func (v *EmailMessage) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EmailMessage"

	return json.Marshal(v)
}
