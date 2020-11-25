package database

const (
	relCandidateSkill               string = "Skills"
	relCandidateAcademic            string = "Academics"
	relCandidateAcademicInstitution string = "Academics.Institution"
	relCandidateAcademicCourse      string = "Academics.Course"
	relCandidateJob                 string = "Jobs"
	relCandidateJobCompany          string = "Jobs.Company"
	relCandidateJobDepartment       string = "Jobs.Department"

	filCandidateID   string = "c.id = ?"
	filSkillID       string = "s.id = ?"
	filUserSkillID   string = "us.id = ?"
	filInstitutionID string = "i.id = ?"
	filCourseID      string = "cr.id = ?"
	filAcademicID    string = "ah.id = ?"
	filCompanyID     string = "co.id = ?"
	filDepartmentID  string = "d.id = ?"
	filJobID         string = "jh.id = ?"
)
