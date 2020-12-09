package models

// Comparator compares whether two models are equal
type Comparator interface {
	IsEqual(m2 interface{}) bool
}

// IsEqual checks the equivalence of two Assessment objects
func (m1 *Assessment) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Assessment)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.Description != convertedM2.Description ||
		m1.Notes != convertedM2.Notes ||
		m1.ImageURL != convertedM2.ImageURL ||
		m1.Difficulty != convertedM2.Difficulty ||
		m1.TimeAllowed != convertedM2.TimeAllowed ||
		m1.Type != convertedM2.Type ||
		m1.Randomise != convertedM2.Randomise ||
		m1.NumQuestions != convertedM2.NumQuestions {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AssessmentStatus objects
func (m1 *AssessmentStatus) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AssessmentStatus)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.StartedAt == nil) != (convertedM2.StartedAt == nil) ||
		(m1.CompletedAt == nil) != (convertedM2.CompletedAt == nil) {
		return false
	}

	if m1.AssessmentID != convertedM2.AssessmentID ||
		m1.CandidateID != convertedM2.CandidateID ||
		m1.Status != convertedM2.Status ||
		m1.Score != convertedM2.Score ||
		*m1.StartedAt != *convertedM2.StartedAt ||
		*m1.CompletedAt != *convertedM2.CompletedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Question objects
func (m1 *Question) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Question)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CreatedBy != convertedM2.CreatedBy ||
		m1.Type != convertedM2.Type ||
		m1.Text != convertedM2.Text ||
		m1.ImageURL != convertedM2.ImageURL ||
		testSliceEqual(m1.Options, convertedM2.Options) ||
		m1.Answer != convertedM2.Answer ||
		m1.Type != convertedM2.Type {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Tag objects
func (m1 *Tag) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Tag)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two QuestionTag objects
func (m1 *QuestionTag) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*QuestionTag)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.QuestionID != convertedM2.QuestionID ||
		m1.TagID != convertedM2.TagID {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Response objects
func (m1 *Response) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Response)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) {
		return false
	}

	if m1.QuestionID != convertedM2.QuestionID ||
		m1.CandidateID != convertedM2.CandidateID ||
		m1.Selection != convertedM2.Selection ||
		m1.Text != convertedM2.Text ||
		m1.Score != convertedM2.Score ||
		m1.TimeTaken != convertedM2.TimeTaken ||
		*m1.CreatedAt != *convertedM2.CreatedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AssessmentQuestion objects
func (m1 *AssessmentQuestion) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AssessmentQuestion)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AssessmentID != convertedM2.AssessmentID ||
		m1.QuestionID != convertedM2.QuestionID {
		return false
	}
	return true
}

func testSliceEqual(a, b []string) bool {
	// if one is nil, the other must also be nil
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func checkNil(m1, m2 interface{}) (isNil bool, resolve bool) {
	// if both nil, return true and resolve
	if m1 == nil && m2 == nil {
		return true, true
	}
	// if one is nil and the other not, return false and resolve
	if (m1 == nil) != (m2 == nil) {
		return false, true
	}
	// both are not nil, return false and don't resolve
	return false, false
}
