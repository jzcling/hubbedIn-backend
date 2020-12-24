package models

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
		m1.ImageURL != convertedM2.ImageURL ||
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

	if (m1.StartedAt == nil) != (convertedM2.StartedAt == nil) ||
		(m1.CompletedAt == nil) != (convertedM2.CompletedAt == nil) {
		return false
	}

	if m1.AssessmentID != convertedM2.AssessmentID ||
		m1.CandidateID != convertedM2.CandidateID ||
		m1.Status != convertedM2.Status ||
		*m1.StartedAt != *convertedM2.StartedAt ||
		*m1.CompletedAt != *convertedM2.CompletedAt ||
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
		m1.MediaURL != convertedM2.MediaURL ||
		m1.Code != convertedM2.Code ||
		helpers.IsStringSliceEqual(m1.Options, convertedM2.Options) ||
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

	if m1.QuestionID != convertedM2.QuestionID ||
		m1.TagID != convertedM2.TagID {
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

	if (m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) {
		return false
	}

	if m1.AttemptID != convertedM2.AttemptID ||
		m1.QuestionID != convertedM2.QuestionID ||
		m1.CandidateID != convertedM2.CandidateID ||
		m1.Selection != convertedM2.Selection ||
		m1.Text != convertedM2.Text ||
		m1.CMMode != convertedM2.CMMode ||
		m1.Score != convertedM2.Score ||
		m1.TimeTaken != convertedM2.TimeTaken ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt {
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

	if m1.AssessmentID != convertedM2.AssessmentID ||
		m1.QuestionID != convertedM2.QuestionID {
		return false
	}
	return true
}
