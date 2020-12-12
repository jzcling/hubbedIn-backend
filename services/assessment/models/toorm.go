package models

import (
	"in-backend/helpers"
	"in-backend/services/assessment/pb"
)

// AssessmentToORM maps the proto Assessment model to the ORM model
func AssessmentToORM(m *pb.Assessment) *Assessment {
	if m == nil {
		return nil
	}

	var questions []*Question
	q := m.Questions
	for _, question := range q {
		questions = append(questions, QuestionToORM(question))
	}

	var attempts []*AssessmentAttempt
	s := m.Attempts
	for _, attempt := range s {
		attempts = append(attempts, AssessmentAttemptToORM(attempt))
	}

	return &Assessment{
		ID:           m.Id,
		Name:         m.Name,
		Description:  m.Description,
		Notes:        m.Notes,
		ImageURL:     m.ImageUrl,
		Difficulty:   m.Difficulty,
		TimeAllowed:  m.TimeAllowed,
		Type:         m.Type,
		Randomise:    m.Randomise,
		NumQuestions: m.NumQuestions,
		Questions:    questions,
		Attempts:     attempts,
	}
}

// AssessmentAttemptToORM maps the proto AssessmentAttempt model to the ORM model
func AssessmentAttemptToORM(m *pb.AssessmentAttempt) *AssessmentAttempt {
	if m == nil {
		return nil
	}

	var questions []*Question
	q := m.Questions
	for _, question := range q {
		questions = append(questions, QuestionToORM(question))
	}

	var questionAttempts []*AttemptQuestion
	qa := m.QuestionAttempts
	for _, attempt := range qa {
		questionAttempts = append(questionAttempts, AttemptQuestionToORM(attempt))
	}

	startedAt := helpers.ProtoTimeToTime(m.StartedAt)
	completedAt := helpers.ProtoTimeToTime(m.CompletedAt)
	return &AssessmentAttempt{
		ID:               m.Id,
		AssessmentID:     m.AssessmentId,
		CandidateID:      m.CandidateId,
		Status:           m.Status,
		StartedAt:        startedAt,
		CompletedAt:      completedAt,
		Score:            m.Score,
		Assessment:       AssessmentToORM(m.Assessment),
		Questions:        questions,
		QuestionAttempts: questionAttempts,
	}
}

// QuestionToORM maps the proto Question model to the ORM model
func QuestionToORM(m *pb.Question) *Question {
	if m == nil {
		return nil
	}

	var tags []*Tag
	t := m.Tags
	for _, tag := range t {
		tags = append(tags, TagToORM(tag))
	}

	var assessments []*Assessment
	a := m.Assessments
	for _, assessment := range a {
		assessments = append(assessments, AssessmentToORM(assessment))
	}

	var assessmentAttempts []*AssessmentAttempt
	aa := m.AssessmentAttempts
	for _, attempt := range aa {
		assessmentAttempts = append(assessmentAttempts, AssessmentAttemptToORM(attempt))
	}

	var attempts []*AttemptQuestion
	at := m.Attempts
	for _, attempt := range at {
		attempts = append(attempts, AttemptQuestionToORM(attempt))
	}

	return &Question{
		ID:                 m.Id,
		CreatedBy:          m.CreatedBy,
		Type:               m.Type,
		Text:               m.Text,
		ImageURL:           m.ImageUrl,
		Options:            m.Options,
		Answer:             m.Answer,
		Tags:               tags,
		Assessments:        assessments,
		AssessmentAttempts: assessmentAttempts,
		Attempts:           attempts,
	}
}

// TagToORM maps the proto Tag model to the ORM model
func TagToORM(m *pb.Tag) *Tag {
	if m == nil {
		return nil
	}
	return &Tag{
		ID:   m.Id,
		Name: m.Name,
	}
}

// QuestionTagToORM maps the proto QuestionTag model to the ORM model
func QuestionTagToORM(m *pb.QuestionTag) *QuestionTag {
	if m == nil {
		return nil
	}
	return &QuestionTag{
		ID:         m.Id,
		QuestionID: m.QuestionId,
		TagID:      m.TagId,
	}
}

// AttemptQuestionToORM maps the proto AttemptQuestion model to the ORM model
func AttemptQuestionToORM(m *pb.AttemptQuestion) *AttemptQuestion {
	if m == nil {
		return nil
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)

	return &AttemptQuestion{
		ID:          m.Id,
		AttemptID:   m.AttemptId,
		QuestionID:  m.QuestionId,
		CandidateID: m.CandidateId,
		Selection:   m.Selection,
		Text:        m.Text,
		Score:       m.Score,
		TimeTaken:   m.TimeTaken,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// AssessmentQuestionToORM maps the proto AssessmentQuestion model to the ORM model
func AssessmentQuestionToORM(m *pb.AssessmentQuestion) *AssessmentQuestion {
	if m == nil {
		return nil
	}
	return &AssessmentQuestion{
		ID:           m.Id,
		AssessmentID: m.AssessmentId,
		QuestionID:   m.QuestionId,
	}
}
