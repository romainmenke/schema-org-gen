package schemaorggo

import "encoding/json"

// EmailMessage see : https://schema.org/EmailMessage
type EmailMessage struct {
	Message

	typeContext

	// BccRecipient see : https://schema.org/bccRecipient
	// A sub property of recipient. The recipient blind copied on a message.
	// types : ContactPoint Organization Person
	BccRecipient interface{} `json:"bccRecipient,omitempty"`

	// CcRecipient see : https://schema.org/ccRecipient
	// A sub property of recipient. The recipient copied on a message.
	// types : ContactPoint Organization Person
	CcRecipient interface{} `json:"ccRecipient,omitempty"`

	// DateRead see : https://schema.org/dateRead
	// The date/time at which the message has been read by the recipient if a single recipient exists.
	// types : DateTime
	DateRead DateTime `json:"dateRead,omitempty"`

	// DateReceived see : https://schema.org/dateReceived
	// The date/time the message was received if a single recipient exists.
	// types : DateTime
	DateReceived DateTime `json:"dateReceived,omitempty"`

	// DateSent see : https://schema.org/dateSent
	// The date/time at which the message was sent.
	// types : DateTime
	DateSent DateTime `json:"dateSent,omitempty"`

	// MessageAttachment see : https://schema.org/messageAttachment
	// A CreativeWork attached to the message.
	// types : CreativeWork
	MessageAttachment *CreativeWork `json:"messageAttachment,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient interface{} `json:"recipient,omitempty"`

	// Sender see : https://schema.org/sender
	// A sub property of participant. The participant who is at the sending end of the action.
	// types : Audience Organization Person
	Sender interface{} `json:"sender,omitempty"`

	// ToRecipient see : https://schema.org/toRecipient
	// A sub property of recipient. The recipient who was directly sent the message.
	// types : Audience ContactPoint Organization Person
	ToRecipient interface{} `json:"toRecipient,omitempty"`
}

func (v EmailMessage) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EmailMessage"

	return json.Marshal(v)
}

func (v *EmailMessage) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EmailMessage"

	return json.Marshal(*v)
}
