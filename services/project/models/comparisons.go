package models

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}

// IsEqual checks the equivalence of two Project objects
func (m1 *Project) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Project) == nil {
		return true
	}

	if (m1 == nil && m2.(*Project) != nil) ||
		(m1 != nil && m2.(*Project) == nil) {
		return false
	}

	if m1.Name != m2.(*Project).Name ||
		m1.RepoURL != m2.(*Project).RepoURL ||
		m1.CreatedAt != m2.(*Project).CreatedAt ||
		m1.UpdatedAt != m2.(*Project).UpdatedAt ||
		m1.DeletedAt != m2.(*Project).DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two CandidateProject objects
func (m1 *CandidateProject) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*CandidateProject) == nil {
		return true
	}

	if (m1 == nil && m2.(*CandidateProject) != nil) ||
		(m1 != nil && m2.(*CandidateProject) == nil) {
		return false
	}

	if m1.CandidateID != m2.(*CandidateProject).CandidateID ||
		m1.ProjectID != m2.(*CandidateProject).ProjectID {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Rating objects
func (m1 *Rating) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Rating) == nil {
		return true
	}

	if (m1 == nil && m2.(*Rating) != nil) ||
		(m1 != nil && m2.(*Rating) == nil) {
		return false
	}

	if m1.ProjectID != m2.(*Rating).ProjectID ||
		m1.ReliabilityRating != m2.(*Rating).ReliabilityRating ||
		m1.MaintainabilityRating != m2.(*Rating).MaintainabilityRating ||
		m1.SecurityRating != m2.(*Rating).SecurityRating ||
		m1.SecurityReviewRating != m2.(*Rating).SecurityReviewRating ||
		m1.Coverage != m2.(*Rating).Coverage ||
		m1.Duplications != m2.(*Rating).Duplications ||
		m1.Lines != m2.(*Rating).Lines ||
		m1.CreatedAt != m2.(*Rating).CreatedAt {
		return false
	}
	return true
}
