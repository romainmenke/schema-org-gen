package schemaorggo

import "encoding/json"

// Message see : https://schema.org/Message
type Message struct {
	CreativeWork

	typeContext

	// BccRecipient see : https://schema.org/bccRecipient
	// A sub property of recipient. The recipient blind copied on a message.
	// types : ContactPoint Organization Person
	BccRecipient []interface{} `json:"bccRecipient,omitempty"`

	// CcRecipient see : https://schema.org/ccRecipient
	// A sub property of recipient. The recipient copied on a message.
	// types : ContactPoint Organization Person
	CcRecipient []interface{} `json:"ccRecipient,omitempty"`

	// DateRead see : https://schema.org/dateRead
	// The date/time at which the message has been read by the recipient if a single recipient exists.
	// types : DateTime
	DateRead []DateTime `json:"dateRead,omitempty"`

	// DateReceived see : https://schema.org/dateReceived
	// The date/time the message was received if a single recipient exists.
	// types : DateTime
	DateReceived []DateTime `json:"dateReceived,omitempty"`

	// DateSent see : https://schema.org/dateSent
	// The date/time at which the message was sent.
	// types : DateTime
	DateSent []DateTime `json:"dateSent,omitempty"`

	// MessageAttachment see : https://schema.org/messageAttachment
	// A CreativeWork attached to the message.
	// types : CreativeWork
	MessageAttachment []*CreativeWork `json:"messageAttachment,omitempty"`

	// Recipient see : https://schema.org/recipient
	// A sub property of participant. The participant who is at the receiving end of the action.
	// types : Audience ContactPoint Organization Person
	Recipient []interface{} `json:"recipient,omitempty"`

	// Sender see : https://schema.org/sender
	// A sub property of participant. The participant who is at the sending end of the action.
	// types : Audience Organization Person
	Sender []interface{} `json:"sender,omitempty"`

	// ToRecipient see : https://schema.org/toRecipient
	// A sub property of recipient. The recipient who was directly sent the message.
	// types : Audience ContactPoint Organization Person
	ToRecipient []interface{} `json:"toRecipient,omitempty"`
}

func (v Message) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.BccRecipient
		if len(v.BccRecipient) == 1 {
			value = v.BccRecipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bccRecipient"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CcRecipient
		if len(v.CcRecipient) == 1 {
			value = v.CcRecipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ccRecipient"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateRead
		if len(v.DateRead) == 1 {
			value = v.DateRead[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateRead"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateReceived
		if len(v.DateReceived) == 1 {
			value = v.DateReceived[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateReceived"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateSent
		if len(v.DateSent) == 1 {
			value = v.DateSent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateSent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MessageAttachment
		if len(v.MessageAttachment) == 1 {
			value = v.MessageAttachment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["messageAttachment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Recipient
		if len(v.Recipient) == 1 {
			value = v.Recipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipient"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Sender
		if len(v.Sender) == 1 {
			value = v.Sender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ToRecipient
		if len(v.ToRecipient) == 1 {
			value = v.ToRecipient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["toRecipient"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Message) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Message"

	return data, nil
}

func (v Message) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
