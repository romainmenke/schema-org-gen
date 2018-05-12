package schemaorggo

import "encoding/json"

// Course see : https://schema.org/Course
type Course struct {
	CreativeWork

	typeContext

	// CourseCode see : https://schema.org/courseCode
	// The identifier for the Course (see: https://schema.org/Course) used by the course provider (see: https://schema.org/provider) (e.g. CS101 or 6.001).
	CourseCode string `json:"courseCode"`

	// CoursePrerequisites see : https://schema.org/coursePrerequisites
	// Requirements for taking the Course. May be completion of another Course (see: https://schema.org/Course) or a textual description like "permission of instructor". Requirements may be a pre-requisite competency, referenced using AlignmentObject (see: https://schema.org/AlignmentObject).
	CoursePrerequisites interface{} `json:"coursePrerequisites"` // types : AlignmentObject Course Text

	// EducationalCredentialAwarded see : http://pending.schema.org/educationalCredentialAwarded
	// A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course.
	EducationalCredentialAwarded interface{} `json:"educationalCredentialAwarded"` // types : Text URL

	// HasCourseInstance see : https://schema.org/hasCourseInstance
	// An offering of the course at a specific time and place or through specific media or mode of study or to a specific section of students.
	HasCourseInstance *CourseInstance `json:"hasCourseInstance"`
}

func (v *Course) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Course"

	return json.Marshal(v)
}
