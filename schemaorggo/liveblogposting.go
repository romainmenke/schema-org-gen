package schemaorggo

import "encoding/json"

// LiveBlogPosting see : https://schema.org/LiveBlogPosting
type LiveBlogPosting struct {
	BlogPosting

	typeContext

	// CoverageEndTime see : https://schema.org/coverageEndTime
	// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	// types : DateTime
	CoverageEndTime DateTime `json:"coverageEndTime,omitempty"`

	// CoverageStartTime see : https://schema.org/coverageStartTime
	// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event&#39;s start time. The LiveBlogPosting may also be created before coverage begins.
	// types : DateTime
	CoverageStartTime DateTime `json:"coverageStartTime,omitempty"`

	// LiveBlogUpdate see : https://schema.org/liveBlogUpdate
	// An update to the LiveBlog.
	// types : BlogPosting
	LiveBlogUpdate *BlogPosting `json:"liveBlogUpdate,omitempty"`
}

func (v LiveBlogPosting) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LiveBlogPosting"

	return json.Marshal(v)
}

func (v *LiveBlogPosting) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LiveBlogPosting"

	return json.Marshal(*v)
}
