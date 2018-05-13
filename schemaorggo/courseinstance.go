package schemaorggo

import "encoding/json"

// CourseInstance see : https://schema.org/CourseInstance
type CourseInstance struct {
	Event

	typeContext

	// CourseMode see : https://schema.org/courseMode
	// The medium or means of delivery of the course instance or the mode of study, either as a text label (e.g. &quot;online&quot;, &quot;onsite&quot; or &quot;blended&quot;; &quot;synchronous&quot; or &quot;asynchronous&quot;; &quot;full-time&quot; or &quot;part-time&quot;) or as a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous ).
	// types : Text URL
	CourseMode []string `json:"courseMode,omitempty"`

	// Instructor see : https://schema.org/instructor
	// A person assigned to instruct or provide instructional assistance for the CourseInstance (see: https://schema.org/CourseInstance).
	// types : Person
	Instructor []*Person `json:"instructor,omitempty"`
}

func (v CourseInstance) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CourseMode
		if len(v.CourseMode) == 1 {
			value = v.CourseMode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["courseMode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Instructor
		if len(v.Instructor) == 1 {
			value = v.Instructor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["instructor"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v CourseInstance) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CourseInstance"

	return data, nil
}

func (v CourseInstance) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
