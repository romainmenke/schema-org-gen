package schemaorggo

import "encoding/json"

// Comment see : https://schema.org/Comment
type Comment struct {
	CreativeWork

	typeContext

	// DownvoteCount see : https://schema.org/downvoteCount
	// The number of downvotes this question, answer or comment has received from the community.
	// types : Integer
	DownvoteCount []float64 `json:"downvoteCount,omitempty"`

	// ParentItem see : https://schema.org/parentItem
	// The parent of a question, answer or item in general.
	// types : Question
	ParentItem []*Question `json:"parentItem,omitempty"`

	// UpvoteCount see : https://schema.org/upvoteCount
	// The number of upvotes this question, answer or comment has received from the community.
	// types : Integer
	UpvoteCount []float64 `json:"upvoteCount,omitempty"`
}

func (v Comment) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

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
		var value interface{} = v.ParentItem
		if len(v.ParentItem) == 1 {
			value = v.ParentItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["parentItem"] = json.RawMessage(b)
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

func (v Comment) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Comment"

	return data, nil
}

func (v Comment) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
