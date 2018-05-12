package schemaorggo

import "encoding/json"

// LiveBlogPosting see : https://schema.org/LiveBlogPosting
type LiveBlogPosting struct {
	BlogPosting

	typeContext

	// CoverageEndTime see : https://schema.org/coverageEndTime
	// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	CoverageEndTime interface{} `json:"coverageEndTime"`

	// CoverageStartTime see : https://schema.org/coverageStartTime
	// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event's start time. The LiveBlogPosting may also be created before coverage begins.
	CoverageStartTime interface{} `json:"coverageStartTime"`

	// LiveBlogUpdate see : https://schema.org/liveBlogUpdate
	// An update to the LiveBlog.
	LiveBlogUpdate *BlogPosting `json:"liveBlogUpdate"`
}

func (v *LiveBlogPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LiveBlogPosting"

	return json.Marshal(v)
}
