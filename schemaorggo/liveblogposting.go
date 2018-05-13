package schemaorggo

import "encoding/json"

// LiveBlogPosting see : https://schema.org/LiveBlogPosting
type LiveBlogPosting struct {
	BlogPosting

	typeContext

	// CoverageEndTime see : https://schema.org/coverageEndTime
	// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	// types : DateTime
	CoverageEndTime []DateTime `json:"coverageEndTime,omitempty"`

	// CoverageStartTime see : https://schema.org/coverageStartTime
	// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event&#39;s start time. The LiveBlogPosting may also be created before coverage begins.
	// types : DateTime
	CoverageStartTime []DateTime `json:"coverageStartTime,omitempty"`

	// LiveBlogUpdate see : https://schema.org/liveBlogUpdate
	// An update to the LiveBlog.
	// types : BlogPosting
	LiveBlogUpdate []*BlogPosting `json:"liveBlogUpdate,omitempty"`
}

func (v LiveBlogPosting) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.BlogPosting.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CoverageEndTime
		if len(v.CoverageEndTime) == 1 {
			value = v.CoverageEndTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["coverageEndTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CoverageStartTime
		if len(v.CoverageStartTime) == 1 {
			value = v.CoverageStartTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["coverageStartTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LiveBlogUpdate
		if len(v.LiveBlogUpdate) == 1 {
			value = v.LiveBlogUpdate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["liveBlogUpdate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LiveBlogPosting) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LiveBlogPosting"

	return data, nil
}

func (v LiveBlogPosting) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
