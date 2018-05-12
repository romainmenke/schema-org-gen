package schemaorggo

import "encoding/json"

// PublicationIssue see : https://schema.org/PublicationIssue
type PublicationIssue struct {
	CreativeWork

	typeContext

	// IssueNumber see : https://schema.org/issueNumber
	// Identifies the issue of publication; for example, "iii" or "2".
	IssueNumber interface{} `json:"issueNumber,omitempty"` // types : Integer Text

	// PageEnd see : https://schema.org/pageEnd
	// The page on which the work ends; for example "138" or "xvi".
	PageEnd interface{} `json:"pageEnd,omitempty"` // types : Integer Text

	// PageStart see : https://schema.org/pageStart
	// The page on which the work starts; for example "135" or "xiii".
	PageStart interface{} `json:"pageStart,omitempty"` // types : Integer Text

	// Pagination see : https://schema.org/pagination
	// Any description of pages that is not separated into pageStart and pageEnd; for example, "1-6, 9, 55" or "10-12, 46-49".
	Pagination string `json:"pagination,omitempty"` // types : Text

}

func (v PublicationIssue) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PublicationIssue"

	return json.Marshal(v)
}

func (v *PublicationIssue) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PublicationIssue"

	return json.Marshal(*v)
}
