package models

import "in-backend/helpers"

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}

// IsEqual checks the equivalence of two User objects
func (m1 *User) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*User)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) ||
		(m1.DeletedAt == nil) != (convertedM2.DeletedAt == nil) {
		return false
	}

	if m1.AuthID != convertedM2.AuthID ||
		m1.FirstName != convertedM2.FirstName ||
		m1.LastName != convertedM2.LastName ||
		m1.Email != convertedM2.Email ||
		m1.ContactNumber != convertedM2.ContactNumber ||
		m1.Picture != convertedM2.Picture ||
		m1.Gender != convertedM2.Gender ||
		helpers.IsStringSliceEqual(m1.Roles, convertedM2.Roles) ||
		m1.CandidateID != convertedM2.CandidateID ||
		m1.JobCompanyID != convertedM2.JobCompanyID ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt ||
		*m1.DeletedAt != *convertedM2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Candidate objects
func (m1 *Candidate) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Candidate)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.Birthday == nil) != (convertedM2.Birthday == nil) ||
		(m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) ||
		(m1.DeletedAt == nil) != (convertedM2.DeletedAt == nil) {
		return false
	}

	if m1.Nationality != convertedM2.Nationality ||
		m1.ResidenceCity != convertedM2.ResidenceCity ||
		m1.ExpectedSalaryCurrency != convertedM2.ExpectedSalaryCurrency ||
		m1.ExpectedSalary != convertedM2.ExpectedSalary ||
		m1.LinkedInURL != convertedM2.LinkedInURL ||
		m1.SCMURL != convertedM2.SCMURL ||
		m1.WebsiteURL != convertedM2.WebsiteURL ||
		m1.EducationLevel != convertedM2.EducationLevel ||
		m1.Summary != convertedM2.Summary ||
		*m1.Birthday != *convertedM2.Birthday ||
		m1.NoticePeriod != convertedM2.NoticePeriod ||
		!helpers.IsStringSliceEqual(m1.PreferredRoles, convertedM2.PreferredRoles) ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt ||
		*m1.DeletedAt != *convertedM2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Skill objects
func (m1 *Skill) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Skill)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two UserSkill objects
func (m1 *UserSkill) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*UserSkill)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CandidateID != convertedM2.CandidateID ||
		m1.SkillID != convertedM2.SkillID ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Institution objects
func (m1 *Institution) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Institution)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.Country != convertedM2.Country {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Course objects
func (m1 *Course) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Course)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.Level != convertedM2.Level {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AcademicHistory objects
func (m1 *AcademicHistory) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AcademicHistory)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) ||
		(m1.DeletedAt == nil) != (convertedM2.DeletedAt == nil) {
		return false
	}

	if m1.CandidateID != convertedM2.CandidateID ||
		m1.InstitutionID != convertedM2.InstitutionID ||
		m1.CourseID != convertedM2.CourseID ||
		m1.YearObtained != convertedM2.YearObtained ||
		m1.Grade != convertedM2.Grade ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt ||
		*m1.DeletedAt != *convertedM2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (m1 *Company) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Company)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Department objects
func (m1 *Department) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Department)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobHistory objects
func (m1 *JobHistory) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobHistory)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.StartDate == nil) != (convertedM2.StartDate == nil) ||
		(m1.EndDate == nil) != (convertedM2.EndDate == nil) ||
		(m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) ||
		(m1.DeletedAt == nil) != (convertedM2.DeletedAt == nil) {
		return false
	}

	if m1.CandidateID != convertedM2.CandidateID ||
		m1.CompanyID != convertedM2.CompanyID ||
		m1.DepartmentID != convertedM2.DepartmentID ||
		m1.Country != convertedM2.Country ||
		m1.City != convertedM2.City ||
		m1.Title != convertedM2.Title ||
		*m1.StartDate != *convertedM2.StartDate ||
		*m1.EndDate != *convertedM2.EndDate ||
		m1.SalaryCurrency != convertedM2.SalaryCurrency ||
		m1.Salary != convertedM2.Salary ||
		m1.Description != convertedM2.Description ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt ||
		*m1.DeletedAt != *convertedM2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobCompany objects
func (m1 *JobCompany) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobCompany)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.LogoURL != convertedM2.LogoURL ||
		m1.Size != convertedM2.Size {
		return false
	}
	return true
}
