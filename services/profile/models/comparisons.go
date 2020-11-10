package models

// IsEqual checks the equivalence of two Candidate objects
func (c1 *Candidate) IsEqual(c2 *Candidate) bool {
	if c1 == nil && c2 == nil {
		return true
	}

	if (c1 == nil && c2 != nil) ||
		(c1 != nil && c2 == nil) {
		return false
	}

	if c1.FirstName != c2.FirstName ||
		c1.LastName != c2.LastName ||
		c1.Email != c2.Email ||
		c1.ContactNumber != c2.ContactNumber ||
		c1.Gender != c2.Gender ||
		c1.Nationality != c2.Nationality ||
		c1.ResidenceCity != c2.ResidenceCity ||
		c1.ExpectedSalaryCurrency != c2.ExpectedSalaryCurrency ||
		c1.ExpectedSalary != c2.ExpectedSalary ||
		c1.LinkedInURL != c2.LinkedInURL ||
		c1.SCMURL != c2.SCMURL ||
		c1.EducationLevel != c2.EducationLevel ||
		c1.Birthday != c2.Birthday ||
		c1.NoticePeriod != c2.NoticePeriod ||
		c1.CreatedAt != c2.CreatedAt ||
		c1.UpdatedAt != c2.UpdatedAt ||
		c1.DeletedAt != c2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Skill objects
func (s1 *Skill) IsEqual(s2 *Skill) bool {
	if s1 == nil && s2 == nil {
		return true
	}

	if (s1 == nil && s2 != nil) ||
		(s1 != nil && s2 == nil) {
		return false
	}

	if s1.Name != s2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Institution objects
func (i1 *Institution) IsEqual(i2 *Institution) bool {
	if i1 == nil && i2 == nil {
		return true
	}

	if (i1 == nil && i2 != nil) ||
		(i1 != nil && i2 == nil) {
		return false
	}

	if i1.Name != i2.Name ||
		i1.Country != i2.Country {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Course objects
func (c1 *Course) IsEqual(c2 *Course) bool {
	if c1 == nil && c2 == nil {
		return true
	}

	if (c1 == nil && c2 != nil) ||
		(c1 != nil && c2 == nil) {
		return false
	}

	if c1.Name != c2.Name ||
		c1.Level != c2.Level {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AcademicHistory objects
func (a1 *AcademicHistory) IsEqual(a2 *AcademicHistory) bool {
	if a1 == nil && a2 == nil {
		return true
	}

	if (a1 == nil && a2 != nil) ||
		(a1 != nil && a2 == nil) {
		return false
	}

	if a1.CandidateID != a2.CandidateID ||
		a1.InstitutionID != a2.InstitutionID ||
		a1.CourseID != a2.CourseID ||
		a1.YearObtained != a2.YearObtained ||
		a1.CreatedAt != a2.CreatedAt ||
		a1.UpdatedAt != a2.UpdatedAt ||
		a1.DeletedAt != a2.DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (c1 *Company) IsEqual(c2 *Company) bool {
	if c1 == nil && c2 == nil {
		return true
	}

	if (c1 == nil && c2 != nil) ||
		(c1 != nil && c2 == nil) {
		return false
	}

	if c1.Name != c2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Department objects
func (d1 *Department) IsEqual(d2 *Department) bool {
	if d1 == nil && d2 == nil {
		return true
	}

	if (d1 == nil && d2 != nil) ||
		(d1 != nil && d2 == nil) {
		return false
	}

	if d1.Name != d2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobHistory objects
func (j1 *JobHistory) IsEqual(j2 *JobHistory) bool {
	if j1 == nil && j2 == nil {
		return true
	}

	if (j1 == nil && j2 != nil) ||
		(j1 != nil && j2 == nil) {
		return false
	}

	if j1.CandidateID != j2.CandidateID ||
		j1.CompanyID != j2.CompanyID ||
		j1.DepartmentID != j2.DepartmentID ||
		j1.Country != j2.Country ||
		j1.City != j2.City ||
		j1.Title != j2.Title ||
		j1.StartDate != j2.StartDate ||
		j1.EndDate != j2.EndDate ||
		j1.SalaryCurrency != j2.SalaryCurrency ||
		j1.Salary != j2.Salary ||
		j1.Description != j2.Description ||
		j1.CreatedAt != j2.CreatedAt ||
		j1.UpdatedAt != j2.UpdatedAt ||
		j1.DeletedAt != j2.DeletedAt {
		return false
	}
	return true
}
