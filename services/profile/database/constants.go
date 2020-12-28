package database

const (
	relCandidate                     string = "Candidate"
	relCandidateSkills               string = "Candidate.Skills"
	relCandidateAcademics            string = "Candidate.Academics"
	relCandidateAcademicsInstitution string = "Candidate.Academics.Institution"
	relCandidateAcademicsCourse      string = "Candidate.Academics.Course"
	relCandidateJobs                 string = "Candidate.Jobs"
	relCandidateJobsCompany          string = "Candidate.Jobs.Company"
	relCandidateJobsDepartment       string = "Candidate.Jobs.Department"

	relSkills               string = "Skills"
	relAcademics            string = "Academics"
	relAcademicsInstitution string = "Academics.Institution"
	relAcademicsCourse      string = "Academics.Course"
	relJobs                 string = "Jobs"
	relJobsCompany          string = "Jobs.Company"
	relJobsDepartment       string = "Jobs.Department"

	filUserID        string = "u.id = ?"
	filSkillID       string = "s.id = ?"
	filUserSkillID   string = "us.id = ?"
	filInstitutionID string = "i.id = ?"
	filCourseID      string = "cr.id = ?"
	filAcademicID    string = "ah.id = ?"
	filCompanyID     string = "co.id = ?"
	filDepartmentID  string = "d.id = ?"
	filJobID         string = "jh.id = ?"

	filNameIn    string = "lower(name) in (?)"
	filLevelIn   string = "lower(level) in (?)"
	filCountryIn string = "lower(country) in (?)"
	filIDIn      string = "id in (?)"

	filNameEquals    string = "lower(name) = ?"
	filLevelEquals   string = "lower(level) = ?"
	filCountryEquals string = "lower(country) = ?"
)
