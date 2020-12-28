package pb

import "in-backend/helpers"

// IsEqual checks the equivalence of two User objects
func (m1 *User) IsEqual(m2 *User) bool {
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AuthId != m2.AuthId ||
		m1.FirstName != m2.FirstName ||
		m1.LastName != m2.LastName ||
		m1.Email != m2.Email ||
		m1.ContactNumber != m2.ContactNumber ||
		m1.Picture != m2.Picture ||
		m1.Gender != m2.Gender ||
		helpers.IsStringSliceEqual(m1.Roles, m2.Roles) ||
		m1.CandidateId != m2.CandidateId ||
		m1.JobCompanyId != m2.JobCompanyId ||
		m1.CreatedAt.AsTime() != m2.CreatedAt.AsTime() ||
		m1.UpdatedAt.AsTime() != m2.UpdatedAt.AsTime() ||
		m1.DeletedAt.AsTime() != m2.DeletedAt.AsTime() {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Candidate objects
func (c1 *Candidate) IsEqual(c2 *Candidate) bool {
	isNil, resolve := helpers.CheckNil(c1, c2)
	if resolve {
		return isNil
	}

	if c1.Nationality != c2.Nationality ||
		c1.ResidenceCity != c2.ResidenceCity ||
		c1.ExpectedSalaryCurrency != c2.ExpectedSalaryCurrency ||
		c1.ExpectedSalary != c2.ExpectedSalary ||
		c1.LinkedInUrl != c2.LinkedInUrl ||
		c1.ScmUrl != c2.ScmUrl ||
		c1.WebsiteUrl != c2.WebsiteUrl ||
		c1.EducationLevel != c2.EducationLevel ||
		c1.Summary != c2.Summary ||
		c1.Birthday.AsTime() != c2.Birthday.AsTime() ||
		c1.NoticePeriod != c2.NoticePeriod ||
		!helpers.IsStringSliceEqual(c1.PreferredRoles, c2.PreferredRoles) ||
		c1.CreatedAt.AsTime() != c2.CreatedAt.AsTime() ||
		c1.UpdatedAt.AsTime() != c2.UpdatedAt.AsTime() ||
		c1.DeletedAt.AsTime() != c2.DeletedAt.AsTime() {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Skill objects
func (s1 *Skill) IsEqual(s2 *Skill) bool {
	isNil, resolve := helpers.CheckNil(s1, s2)
	if resolve {
		return isNil
	}

	if s1.Name != s2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two UserSkill objects
func (us1 *UserSkill) IsEqual(us2 *UserSkill) bool {
	isNil, resolve := helpers.CheckNil(us1, us2)
	if resolve {
		return isNil
	}

	if us1.CandidateId != us2.CandidateId ||
		us1.SkillId != us2.SkillId ||
		us1.CreatedAt != us2.CreatedAt ||
		us1.UpdatedAt != us2.UpdatedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Institution objects
func (i1 *Institution) IsEqual(i2 *Institution) bool {
	isNil, resolve := helpers.CheckNil(i1, i2)
	if resolve {
		return isNil
	}

	if i1.Name != i2.Name ||
		i1.Country != i2.Country {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Course objects
func (c1 *Course) IsEqual(c2 *Course) bool {
	isNil, resolve := helpers.CheckNil(c1, c2)
	if resolve {
		return isNil
	}

	if c1.Name != c2.Name ||
		c1.Level != c2.Level {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AcademicHistory objects
func (a1 *AcademicHistory) IsEqual(a2 *AcademicHistory) bool {
	isNil, resolve := helpers.CheckNil(a1, a2)
	if resolve {
		return isNil
	}

	if a1.CandidateId != a2.CandidateId ||
		a1.InstitutionId != a2.InstitutionId ||
		a1.CourseId != a2.CourseId ||
		a1.YearObtained != a2.YearObtained ||
		a1.Grade != a2.Grade ||
		a1.CreatedAt.AsTime() != a2.CreatedAt.AsTime() ||
		a1.UpdatedAt.AsTime() != a2.UpdatedAt.AsTime() ||
		a1.DeletedAt.AsTime() != a2.DeletedAt.AsTime() {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (c1 *Company) IsEqual(c2 *Company) bool {
	isNil, resolve := helpers.CheckNil(c1, c2)
	if resolve {
		return isNil
	}

	if c1.Name != c2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Department objects
func (d1 *Department) IsEqual(d2 *Department) bool {
	isNil, resolve := helpers.CheckNil(d1, d2)
	if resolve {
		return isNil
	}

	if d1.Name != d2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobHistory objects
func (j1 *JobHistory) IsEqual(j2 *JobHistory) bool {
	isNil, resolve := helpers.CheckNil(j1, j2)
	if resolve {
		return isNil
	}

	if j1.CandidateId != j2.CandidateId ||
		j1.CompanyId != j2.CompanyId ||
		j1.DepartmentId != j2.DepartmentId ||
		j1.Country != j2.Country ||
		j1.City != j2.City ||
		j1.Title != j2.Title ||
		j1.StartDate.AsTime() != j2.StartDate.AsTime() ||
		j1.EndDate.AsTime() != j2.EndDate.AsTime() ||
		j1.SalaryCurrency != j2.SalaryCurrency ||
		j1.Salary != j2.Salary ||
		j1.Description != j2.Description ||
		j1.CreatedAt.AsTime() != j2.CreatedAt.AsTime() ||
		j1.UpdatedAt.AsTime() != j2.UpdatedAt.AsTime() ||
		j1.DeletedAt.AsTime() != j2.DeletedAt.AsTime() {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JoblistingCompany objects
func (m1 *JoblistingCompany) IsEqual(m2 *JoblistingCompany) bool {
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != m2.Name ||
		m1.LogoUrl != m2.LogoUrl ||
		m1.Size != m2.Size {
		return false
	}
	return true
}
