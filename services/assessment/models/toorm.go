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

	var statuses []*AssessmentStatus
	s := m.CandidateStatuses
	for _, status := range s {
		statuses = append(statuses, AssessmentStatusToORM(status))
	}

	return &Assessment{
		ID:                m.Id,
		Name:              m.Name,
		Description:       m.Description,
		Notes:             m.Notes,
		ImageURL:          m.ImageUrl,
		Difficulty:        m.Difficulty,
		TimeAllowed:       m.TimeAllowed,
		Type:              m.Type,
		Randomise:         m.Randomise,
		NumQuestions:      m.NumQuestions,
		Questions:         questions,
		CandidateStatuses: statuses,
	}
}

// AssessmentStatusToORM maps the proto AssessmentStatus model to the ORM model
func AssessmentStatusToORM(m *pb.AssessmentStatus) *AssessmentStatus {
	if m == nil {
		return nil
	}
	startedAt := helpers.ProtoTimeToTime(m.StartedAt)
	completedAt := helpers.ProtoTimeToTime(m.CompletedAt)
	return &AssessmentStatus{
		ID:           m.Id,
		AssessmentID: m.AssessmentId,
		CandidateID:  m.CandidateId,
		Status:       m.Status,
		StartedAt:    startedAt,
		CompletedAt:  completedAt,
		Score:        m.Score,
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

	var responses []*Response
	r := m.Responses
	for _, response := range r {
		responses = append(responses, ResponseToORM(response))
	}

	return &Question{
		ID:          m.Id,
		CreatedBy:   m.CreatedBy,
		Type:        m.Type,
		Text:        m.Text,
		ImageURL:    m.ImageUrl,
		Options:     m.Options,
		Answer:      m.Answer,
		Tags:        tags,
		Assessments: assessments,
		Responses:   responses,
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

// ResponseToORM maps the proto Response model to the ORM model
func ResponseToORM(m *pb.Response) *Response {
	if m == nil {
		return nil
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)

	return &Response{
		ID:          m.Id,
		QuestionID:  m.QuestionId,
		CandidateID: m.CandidateId,
		Selection:   m.Selection,
		Text:        m.Text,
		Score:       m.Score,
		TimeTaken:   m.TimeTaken,
		CreatedAt:   createdAt,
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
