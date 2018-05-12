package schemaorggo

import "encoding/json"

// Course see : https://schema.org/Course
type Course struct {
	CreativeWork

	typeContext

	// CourseCode see : https://schema.org/courseCode
	// The identifier for the Course (see: https://schema.org/Course) used by the course provider (see: https://schema.org/provider) (e.g. CS101 or 6.001).
	// types : Text
	CourseCode string `json:"courseCode,omitempty"`

	// CoursePrerequisites see : https://schema.org/coursePrerequisites
	// Requirements for taking the Course. May be completion of another Course (see: https://schema.org/Course) or a textual description like &quot;permission of instructor&quot;. Requirements may be a pre-requisite competency, referenced using AlignmentObject (see: https://schema.org/AlignmentObject).
	// types : AlignmentObject Course Text
	CoursePrerequisites interface{} `json:"coursePrerequisites,omitempty"`

	// EducationalCredentialAwarded see : http://pending.schema.org/educationalCredentialAwarded
	// A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course.
	// types : Text URL
	EducationalCredentialAwarded string `json:"educationalCredentialAwarded,omitempty"`

	// HasCourseInstance see : https://schema.org/hasCourseInstance
	// An offering of the course at a specific time and place or through specific media or mode of study or to a specific section of students.
	// types : CourseInstance
	HasCourseInstance *CourseInstance `json:"hasCourseInstance,omitempty"`
}

func (v Course) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Course"

	return json.Marshal(v)
}

func (v *Course) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Course"

	return json.Marshal(*v)
}
