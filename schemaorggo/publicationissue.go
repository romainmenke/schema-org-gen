package schemaorggo

import "encoding/json"

// PublicationIssue see : https://schema.org/PublicationIssue
type PublicationIssue struct {
	CreativeWork

	typeContext

	// IssueNumber see : https://schema.org/issueNumber
	// Identifies the issue of publication; for example, &quot;iii&quot; or &quot;2&quot;.
	// types : Integer Text
	IssueNumber interface{} `json:"issueNumber,omitempty"`

	// PageEnd see : https://schema.org/pageEnd
	// The page on which the work ends; for example &quot;138&quot; or &quot;xvi&quot;.
	// types : Integer Text
	PageEnd interface{} `json:"pageEnd,omitempty"`

	// PageStart see : https://schema.org/pageStart
	// The page on which the work starts; for example &quot;135&quot; or &quot;xiii&quot;.
	// types : Integer Text
	PageStart interface{} `json:"pageStart,omitempty"`

	// Pagination see : https://schema.org/pagination
	// Any description of pages that is not separated into pageStart and pageEnd; for example, &quot;1-6, 9, 55&quot; or &quot;10-12, 46-49&quot;.
	// types : Text
	Pagination string `json:"pagination,omitempty"`
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
