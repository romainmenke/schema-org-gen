package schemaorggo

import "encoding/json"

// LiveBlogPosting see : https://schema.org/LiveBlogPosting
type LiveBlogPosting struct {
	BlogPosting

	typeContext

	// CoverageEndTime see : https://schema.org/coverageEndTime
	// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	CoverageEndTime DateTime `json:"coverageEndTime,omitempty"` // types : DateTime

	// CoverageStartTime see : https://schema.org/coverageStartTime
	// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event's start time. The LiveBlogPosting may also be created before coverage begins.
	CoverageStartTime DateTime `json:"coverageStartTime,omitempty"` // types : DateTime

	// LiveBlogUpdate see : https://schema.org/liveBlogUpdate
	// An update to the LiveBlog.
	LiveBlogUpdate *BlogPosting `json:"liveBlogUpdate,omitempty"` // types : BlogPosting

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
