package models

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}

// IsEqual checks the equivalence of two Candidate objects
func (m1 *Candidate) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Candidate) == nil {
		return true
	}

	if (m1 == nil && m2.(*Candidate) != nil) ||
		(m1 != nil && m2.(*Candidate) == nil) {
		return false
	}

	if ((m1.Birthday != nil && m2.(*Candidate).Birthday == nil) || (m1.Birthday == nil && m2.(*Candidate).Birthday != nil)) ||
		((m1.CreatedAt != nil && m2.(*Candidate).CreatedAt == nil) || (m1.CreatedAt == nil && m2.(*Candidate).CreatedAt != nil)) ||
		((m1.UpdatedAt != nil && m2.(*Candidate).UpdatedAt == nil) || (m1.UpdatedAt == nil && m2.(*Candidate).UpdatedAt != nil)) ||
		((m1.DeletedAt != nil && m2.(*Candidate).DeletedAt == nil) || (m1.DeletedAt == nil && m2.(*Candidate).DeletedAt != nil)) {
		return false
	}

	if m1.AuthID != m2.(*Candidate).AuthID ||
		m1.FirstName != m2.(*Candidate).FirstName ||
		m1.LastName != m2.(*Candidate).LastName ||
		m1.Email != m2.(*Candidate).Email ||
		m1.ContactNumber != m2.(*Candidate).ContactNumber ||
		m1.Picture != m2.(*Candidate).Picture ||
		m1.Gender != m2.(*Candidate).Gender ||
		m1.Nationality != m2.(*Candidate).Nationality ||
		m1.ResidenceCity != m2.(*Candidate).ResidenceCity ||
		m1.ExpectedSalaryCurrency != m2.(*Candidate).ExpectedSalaryCurrency ||
		m1.ExpectedSalary != m2.(*Candidate).ExpectedSalary ||
		m1.LinkedInURL != m2.(*Candidate).LinkedInURL ||
		m1.SCMURL != m2.(*Candidate).SCMURL ||
		m1.WebsiteURL != m2.(*Candidate).WebsiteURL ||
		m1.EducationLevel != m2.(*Candidate).EducationLevel ||
		m1.Summary != m2.(*Candidate).Summary ||
		((m1.Birthday != nil && m2.(*Candidate).Birthday != nil) && (*m1.Birthday != *m2.(*Candidate).Birthday)) ||
		m1.NoticePeriod != m2.(*Candidate).NoticePeriod ||
		((m1.CreatedAt != nil && m2.(*Candidate).CreatedAt != nil) && (*m1.CreatedAt != *m2.(*Candidate).CreatedAt)) ||
		((m1.UpdatedAt != nil && m2.(*Candidate).UpdatedAt != nil) && (*m1.UpdatedAt != *m2.(*Candidate).UpdatedAt)) ||
		((m1.DeletedAt != nil && m2.(*Candidate).DeletedAt != nil) && (*m1.DeletedAt != *m2.(*Candidate).DeletedAt)) {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Skill objects
func (m1 *Skill) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Skill) == nil {
		return true
	}

	if (m1 == nil && m2.(*Skill) != nil) ||
		(m1 != nil && m2.(*Skill) == nil) {
		return false
	}

	if m1.Name != m2.(*Skill).Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two UserSkill objects
func (m1 *UserSkill) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*UserSkill) == nil {
		return true
	}

	if (m1 == nil && m2.(*UserSkill) != nil) ||
		(m1 != nil && m2.(*UserSkill) == nil) {
		return false
	}

	if m1.CandidateID != m2.(*UserSkill).CandidateID ||
		m1.SkillID != m2.(*UserSkill).SkillID ||
		*m1.CreatedAt != *m2.(*UserSkill).CreatedAt ||
		*m1.UpdatedAt != *m2.(*UserSkill).UpdatedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Institution objects
func (m1 *Institution) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Institution) == nil {
		return true
	}

	if (m1 == nil && m2.(*Institution) != nil) ||
		(m1 != nil && m2.(*Institution) == nil) {
		return false
	}

	if m1.Name != m2.(*Institution).Name ||
		m1.Country != m2.(*Institution).Country {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Course objects
func (m1 *Course) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Course) == nil {
		return true
	}

	if (m1 == nil && m2.(*Course) != nil) ||
		(m1 != nil && m2.(*Course) == nil) {
		return false
	}

	if m1.Name != m2.(*Course).Name ||
		m1.Level != m2.(*Course).Level {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AcademicHistory objects
func (m1 *AcademicHistory) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*AcademicHistory) == nil {
		return true
	}

	if (m1 == nil && m2.(*AcademicHistory) != nil) ||
		(m1 != nil && m2.(*AcademicHistory) == nil) {
		return false
	}

	if m1.CandidateID != m2.(*AcademicHistory).CandidateID ||
		m1.InstitutionID != m2.(*AcademicHistory).InstitutionID ||
		m1.CourseID != m2.(*AcademicHistory).CourseID ||
		m1.YearObtained != m2.(*AcademicHistory).YearObtained ||
		m1.Grade != m2.(*AcademicHistory).Grade ||
		*m1.CreatedAt != *m2.(*AcademicHistory).CreatedAt ||
		*m1.UpdatedAt != *m2.(*AcademicHistory).UpdatedAt ||
		*m1.DeletedAt != *m2.(*AcademicHistory).DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (m1 *Company) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Company) == nil {
		return true
	}

	if (m1 == nil && m2.(*Company) != nil) ||
		(m1 != nil && m2.(*Company) == nil) {
		return false
	}

	if m1.Name != m2.(*Company).Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Department objects
func (m1 *Department) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Department) == nil {
		return true
	}

	if (m1 == nil && m2.(*Department) != nil) ||
		(m1 != nil && m2.(*Department) == nil) {
		return false
	}

	if m1.Name != m2.(*Department).Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobHistory objects
func (m1 *JobHistory) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*JobHistory) == nil {
		return true
	}

	if (m1 == nil && m2.(*JobHistory) != nil) ||
		(m1 != nil && m2.(*JobHistory) == nil) {
		return false
	}

	if m1.CandidateID != m2.(*JobHistory).CandidateID ||
		m1.CompanyID != m2.(*JobHistory).CompanyID ||
		m1.DepartmentID != m2.(*JobHistory).DepartmentID ||
		m1.Country != m2.(*JobHistory).Country ||
		m1.City != m2.(*JobHistory).City ||
		m1.Title != m2.(*JobHistory).Title ||
		*m1.StartDate != *m2.(*JobHistory).StartDate ||
		*m1.EndDate != *m2.(*JobHistory).EndDate ||
		m1.SalaryCurrency != m2.(*JobHistory).SalaryCurrency ||
		m1.Salary != m2.(*JobHistory).Salary ||
		m1.Description != m2.(*JobHistory).Description ||
		*m1.CreatedAt != *m2.(*JobHistory).CreatedAt ||
		*m1.UpdatedAt != *m2.(*JobHistory).UpdatedAt ||
		*m1.DeletedAt != *m2.(*JobHistory).DeletedAt {
		return false
	}
	return true
}
