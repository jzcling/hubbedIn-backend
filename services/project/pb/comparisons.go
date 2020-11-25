package pb

// IsEqual checks the equivalence of two Project objects
func (m1 *Project) IsEqual(m2 *Project) bool {
	if m1 == nil && m2 == nil {
		return true
	}

	if (m1 == nil && m2 != nil) ||
		(m1 != nil && m2 == nil) {
		return false
	}

	if m1.Name != m2.Name ||
		m1.RepoUrl != m2.RepoUrl ||
		m1.CreatedAt.AsTime() != m2.CreatedAt.AsTime() ||
		m1.UpdatedAt.AsTime() != m2.UpdatedAt.AsTime() ||
		m1.DeletedAt.AsTime() != m2.DeletedAt.AsTime() {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two CandidateProject objects
func (m1 *CandidateProject) IsEqual(m2 *CandidateProject) bool {
	if m1 == nil && m2 == nil {
		return true
	}

	if (m1 == nil && m2 != nil) ||
		(m1 != nil && m2 == nil) {
		return false
	}

	if m1.CandidateId != m2.CandidateId ||
		m1.ProjectId != m2.ProjectId {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Rating objects
func (m1 *Rating) IsEqual(m2 *Rating) bool {
	if m1 == nil && m2 == nil {
		return true
	}

	if (m1 == nil && m2 != nil) ||
		(m1 != nil && m2 == nil) {
		return false
	}

	if m1.ProjectId != m2.ProjectId ||
		m1.ReliabilityRating != m2.ReliabilityRating ||
		m1.MaintainabilityRating != m2.MaintainabilityRating ||
		m1.SecurityRating != m2.SecurityRating ||
		m1.SecurityReviewRating != m2.SecurityReviewRating ||
		m1.Coverage != m2.Coverage ||
		m1.Duplications != m2.Duplications ||
		m1.Lines != m2.Lines ||
		m1.CreatedAt != m2.CreatedAt {
		return false
	}
	return true
}
