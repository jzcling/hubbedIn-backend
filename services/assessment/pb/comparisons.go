package pb

import "in-backend/helpers"

// IsEqual checks the equivalence of two Assessment objects
func (m1 *Assessment) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Assessment)
	isNil, resolve := helpers.CheckNil(m1, m2)
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
		m1.NumQuestions != convertedM2.NumQuestions ||
		m1.CanGoBack != convertedM2.CanGoBack {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two AssessmentAttempt objects
func (m1 *AssessmentAttempt) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*AssessmentAttempt)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AssessmentId != convertedM2.AssessmentId ||
		m1.CandidateId != convertedM2.CandidateId ||
		m1.Status != convertedM2.Status ||
		m1.StartedAt.AsTime() != convertedM2.StartedAt.AsTime() ||
		m1.CompletedAt.AsTime() != convertedM2.CompletedAt.AsTime() ||
		m1.CurrentQuestion != convertedM2.CurrentQuestion ||
		m1.Score != convertedM2.Score {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Question objects
func (m1 *Question) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Question)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CreatedBy != convertedM2.CreatedBy ||
		m1.Type != convertedM2.Type ||
		m1.Text != convertedM2.Text ||
		m1.MediaUrl != convertedM2.MediaUrl ||
		m1.Code != convertedM2.Code ||
		!helpers.Equal(m1.Options, convertedM2.Options) ||
		m1.Answer != convertedM2.Answer ||
		m1.Type != convertedM2.Type {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Tag objects
func (m1 *Tag) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Tag)
	isNil, resolve := helpers.CheckNil(m1, m2)
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
	isNil, resolve := helpers.CheckNil(m1, m2)
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
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AttemptId != convertedM2.AttemptId ||
		m1.QuestionId != convertedM2.QuestionId ||
		m1.CandidateId != convertedM2.CandidateId ||
		m1.Selection != convertedM2.Selection ||
		m1.Text != convertedM2.Text ||
		m1.CmMode != convertedM2.CmMode ||
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
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.AssessmentId != convertedM2.AssessmentId ||
		m1.QuestionId != convertedM2.QuestionId {
		return false
	}
	return true
}
