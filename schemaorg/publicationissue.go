package schemaorg

import "encoding/json"

// PublicationIssue see : https://schema.org/PublicationIssue
type PublicationIssue struct {

typeContext

CreativeWork

// IssueNumber see : /issueNumber
// Identifies the issue of publication; for example, "iii" or "2".
IssueNumber interface{} `json:"issueNumber"` // types : Integer Text

// PageEnd see : /pageEnd
// The page on which the work ends; for example "138" or "xvi".
PageEnd interface{} `json:"pageEnd"` // types : Integer Text

// PageStart see : /pageStart
// The page on which the work starts; for example "135" or "xiii".
PageStart interface{} `json:"pageStart"` // types : Integer Text

// Pagination see : /pagination
// Any description of pages that is not separated into pageStart and pageEnd; for example, "1-6, 9, 55" or "10-12, 46-49".
Pagination string `json:"pagination"`

}

func (v *PublicationIssue) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PublicationIssue"

	return json.Marshal(v)
}
