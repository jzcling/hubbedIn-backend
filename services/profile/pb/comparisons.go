package pb

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
		c1.Picture != c2.Picture ||
		c1.Gender != c2.Gender ||
		c1.Nationality != c2.Nationality ||
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
		c1.CreatedAt.AsTime() != c2.CreatedAt.AsTime() ||
		c1.UpdatedAt.AsTime() != c2.UpdatedAt.AsTime() ||
		c1.DeletedAt.AsTime() != c2.DeletedAt.AsTime() {
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

// IsEqual checks the equivalence of two UserSkill objects
func (us1 *UserSkill) IsEqual(us2 *UserSkill) bool {
	if us1 == nil && us2 == nil {
		return true
	}

	if (us1 == nil && us2 != nil) ||
		(us1 != nil && us2 == nil) {
		return false
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

	if a1.CandidateId != a2.CandidateId ||
		a1.InstitutionId != a2.InstitutionId ||
		a1.CourseId != a2.CourseId ||
		a1.YearObtained != a2.YearObtained ||
		a1.CreatedAt.AsTime() != a2.CreatedAt.AsTime() ||
		a1.UpdatedAt.AsTime() != a2.UpdatedAt.AsTime() ||
		a1.DeletedAt.AsTime() != a2.DeletedAt.AsTime() {
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
