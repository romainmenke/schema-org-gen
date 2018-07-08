package schemaorggo

import "encoding/json"

// Course see : https://schema.org/Course
type Course struct {
	CreativeWork

	typeContext

	// CourseCode see : https://schema.org/courseCode
	// The identifier for the Course (see: https://schema.org/Course) used by the course provider (see: https://schema.org/provider) (e.g. CS101 or 6.001).
	// types : Text
	CourseCode []string `json:"courseCode,omitempty"`

	// CoursePrerequisites see : https://schema.org/coursePrerequisites
	// Requirements for taking the Course. May be completion of another Course (see: https://schema.org/Course) or a textual description like &quot;permission of instructor&quot;. Requirements may be a pre-requisite competency, referenced using AlignmentObject (see: https://schema.org/AlignmentObject).
	// types : AlignmentObject Course Text
	CoursePrerequisites []interface{} `json:"coursePrerequisites,omitempty"`

	// EducationalCredentialAwarded see : https://pending.schema.org/educationalCredentialAwarded
	// A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course.
	// types : Text URL
	EducationalCredentialAwarded []string `json:"educationalCredentialAwarded,omitempty"`

	// HasCourseInstance see : https://schema.org/hasCourseInstance
	// An offering of the course at a specific time and place or through specific media or mode of study or to a specific section of students.
	// types : CourseInstance
	HasCourseInstance []*CourseInstance `json:"hasCourseInstance,omitempty"`
}

func (v Course) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.CourseCode
		if len(v.CourseCode) == 1 {
			value = v.CourseCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["courseCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CoursePrerequisites
		if len(v.CoursePrerequisites) == 1 {
			value = v.CoursePrerequisites[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["coursePrerequisites"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EducationalCredentialAwarded
		if len(v.EducationalCredentialAwarded) == 1 {
			value = v.EducationalCredentialAwarded[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationalCredentialAwarded"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasCourseInstance
		if len(v.HasCourseInstance) == 1 {
			value = v.HasCourseInstance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasCourseInstance"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Course) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Course"

	return data, nil
}

func (v Course) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
