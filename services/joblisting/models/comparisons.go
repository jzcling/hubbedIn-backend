package models

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}

// IsEqual checks the equivalence of two Joblisting objects
func (m1 *Joblisting) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*Joblisting) == nil {
		return true
	}

	if (m1 == nil && m2.(*Joblisting) != nil) ||
		(m1 != nil && m2.(*Joblisting) == nil) {
		return false
	}

	if m1.Name != m2.(*Joblisting).Name ||
		m1.RepoURL != m2.(*Joblisting).RepoURL ||
		*m1.CreatedAt != *m2.(*Joblisting).CreatedAt ||
		*m1.UpdatedAt != *m2.(*Joblisting).UpdatedAt ||
		*m1.DeletedAt != *m2.(*Joblisting).DeletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two CandidateJoblisting objects
func (m1 *CandidateJoblisting) IsEqual(m2 interface{}) bool {
	if m1 == nil && m2.(*CandidateJoblisting) == nil {
		return true
	}

	if (m1 == nil && m2.(*CandidateJoblisting) != nil) ||
		(m1 != nil && m2.(*CandidateJoblisting) == nil) {
		return false
	}

	if m1.CandidateID != m2.(*CandidateJoblisting).CandidateID ||
		m1.JoblistingID != m2.(*CandidateJoblisting).JoblistingID {
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

	if m1.JoblistingID != m2.(*Rating).JoblistingID ||
		m1.ReliabilityRating != m2.(*Rating).ReliabilityRating ||
		m1.MaintainabilityRating != m2.(*Rating).MaintainabilityRating ||
		m1.SecurityRating != m2.(*Rating).SecurityRating ||
		m1.SecurityReviewRating != m2.(*Rating).SecurityReviewRating ||
		m1.Coverage != m2.(*Rating).Coverage ||
		m1.Duplications != m2.(*Rating).Duplications ||
		m1.Lines != m2.(*Rating).Lines ||
		*m1.CreatedAt != *m2.(*Rating).CreatedAt {
		return false
	}
	return true
}
