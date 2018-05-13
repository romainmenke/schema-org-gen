package schemaorggo

import "encoding/json"

// Question see : https://schema.org/Question
type Question struct {
	CreativeWork

	typeContext

	// AcceptedAnswer see : https://schema.org/acceptedAnswer
	// The answer that has been accepted as best, typically on a Question/Answer site. Sites vary in their selection mechanisms, e.g. drawing on community opinion and/or the view of the Question author.
	// types : Answer
	AcceptedAnswer []*Answer `json:"acceptedAnswer,omitempty"`

	// AnswerCount see : https://schema.org/answerCount
	// The number of answers this question has received.
	// types : Integer
	AnswerCount []float64 `json:"answerCount,omitempty"`

	// DownvoteCount see : https://schema.org/downvoteCount
	// The number of downvotes this question, answer or comment has received from the community.
	// types : Integer
	DownvoteCount []float64 `json:"downvoteCount,omitempty"`

	// SuggestedAnswer see : https://schema.org/suggestedAnswer
	// An answer (possibly one of several, possibly incorrect) to a Question, e.g. on a Question/Answer site.
	// types : Answer
	SuggestedAnswer []*Answer `json:"suggestedAnswer,omitempty"`

	// UpvoteCount see : https://schema.org/upvoteCount
	// The number of upvotes this question, answer or comment has received from the community.
	// types : Integer
	UpvoteCount []float64 `json:"upvoteCount,omitempty"`
}

func (v Question) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcceptedAnswer
		if len(v.AcceptedAnswer) == 1 {
			value = v.AcceptedAnswer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acceptedAnswer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AnswerCount
		if len(v.AnswerCount) == 1 {
			value = v.AnswerCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["answerCount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DownvoteCount
		if len(v.DownvoteCount) == 1 {
			value = v.DownvoteCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["downvoteCount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuggestedAnswer
		if len(v.SuggestedAnswer) == 1 {
			value = v.SuggestedAnswer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suggestedAnswer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UpvoteCount
		if len(v.UpvoteCount) == 1 {
			value = v.UpvoteCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["upvoteCount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Question) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Question"

	return data, nil
}

func (v Question) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
