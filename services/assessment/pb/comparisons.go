package pb

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
		m1.ImageUrl != convertedM2.ImageUrl ||
		m1.Difficulty != convertedM2.Difficulty ||
		m1.TimeAllowed != convertedM2.TimeAllowed ||
		m1.Type != convertedM2.Type ||
		m1.Randomise != convertedM2.Randomise ||
		m1.NumQuestions != convertedM2.NumQuestions {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AssessmentAttempt objects
func (m1 *AssessmentAttempt) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AssessmentAttempt)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AssessmentId != convertedM2.AssessmentId ||
		m1.CandidateId != convertedM2.CandidateId ||
		m1.Status != convertedM2.Status ||
		m1.Score != convertedM2.Score ||
		m1.StartedAt.AsTime() != convertedM2.StartedAt.AsTime() ||
		m1.CompletedAt.AsTime() != convertedM2.CompletedAt.AsTime() {
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
		m1.ImageUrl != convertedM2.ImageUrl ||
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

	if m1.QuestionId != convertedM2.QuestionId ||
		m1.TagId != convertedM2.TagId {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AttemptQuestion objects
func (m1 *AttemptQuestion) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AttemptQuestion)
	isNil, resolve := checkNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AttemptId != convertedM2.AttemptId ||
		m1.QuestionId != convertedM2.QuestionId ||
		m1.CandidateId != convertedM2.CandidateId ||
		m1.Selection != convertedM2.Selection ||
		m1.Text != convertedM2.Text ||
		m1.Score != convertedM2.Score ||
		m1.TimeTaken != convertedM2.TimeTaken ||
		m1.CreatedAt.AsTime() != convertedM2.CreatedAt.AsTime() ||
		m1.UpdatedAt.AsTime() != convertedM2.UpdatedAt.AsTime() {
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

	if m1.AssessmentId != convertedM2.AssessmentId ||
		m1.QuestionId != convertedM2.QuestionId {
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
