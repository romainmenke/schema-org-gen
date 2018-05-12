package schemaorg

import "encoding/json"

// Question see : https://schema.org/Question
type Question struct {

CreativeWork

typeContext

// AcceptedAnswer see : https://schema.org/acceptedAnswer
// The answer that has been accepted as best, typically on a Question/Answer site. Sites vary in their selection mechanisms, e.g. drawing on community opinion and/or the view of the Question author.
AcceptedAnswer *Answer `json:"acceptedAnswer"`

// AnswerCount see : https://schema.org/answerCount
// The number of answers this question has received.
AnswerCount int `json:"answerCount"`

// DownvoteCount see : https://schema.org/downvoteCount
// The number of downvotes this question, answer or comment has received from the community.
DownvoteCount int `json:"downvoteCount"`

// SuggestedAnswer see : https://schema.org/suggestedAnswer
// An answer (possibly one of several, possibly incorrect) to a Question, e.g. on a Question/Answer site.
SuggestedAnswer *Answer `json:"suggestedAnswer"`

// UpvoteCount see : https://schema.org/upvoteCount
// The number of upvotes this question, answer or comment has received from the community.
UpvoteCount int `json:"upvoteCount"`

}

func (v *Question) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Question"

	return json.Marshal(v)
}
