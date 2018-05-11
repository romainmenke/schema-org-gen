package schemaorg

import "encoding/json"

// LiveBlogPosting see : https://schema.org/LiveBlogPosting
type LiveBlogPosting struct {

typeContext

BlogPosting

// CoverageEndTime see : /coverageEndTime
// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
CoverageEndTime interface{} `json:"coverageEndTime"`

// CoverageStartTime see : /coverageStartTime
// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event's start time. The LiveBlogPosting may also be created before coverage begins.
CoverageStartTime interface{} `json:"coverageStartTime"`

// LiveBlogUpdate see : /liveBlogUpdate
// An update to the LiveBlog.
LiveBlogUpdate *BlogPosting `json:"liveBlogUpdate"`

}

func (v *LiveBlogPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LiveBlogPosting"

	return json.Marshal(v)
}
