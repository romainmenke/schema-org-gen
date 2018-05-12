package schemaorggo

import "encoding/json"

// CourseInstance see : https://schema.org/CourseInstance
type CourseInstance struct {
	Event

	typeContext

	// CourseMode see : https://schema.org/courseMode
	// The medium or means of delivery of the course instance or the mode of study, either as a text label (e.g. &quot;online&quot;, &quot;onsite&quot; or &quot;blended&quot;; &quot;synchronous&quot; or &quot;asynchronous&quot;; &quot;full-time&quot; or &quot;part-time&quot;) or as a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous ).
	// types : Text URL
	CourseMode string `json:"courseMode,omitempty"`

	// Instructor see : https://schema.org/instructor
	// A person assigned to instruct or provide instructional assistance for the CourseInstance (see: https://schema.org/CourseInstance).
	// types : Person
	Instructor *Person `json:"instructor,omitempty"`
}

func (v CourseInstance) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CourseInstance"

	return json.Marshal(v)
}

func (v *CourseInstance) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CourseInstance"

	return json.Marshal(*v)
}
