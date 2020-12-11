package models

import (
	"in-backend/helpers"
	"in-backend/services/assessment/pb"
)

// ToProto maps the ORM Assessment model to the proto model
func (m *Assessment) ToProto() *pb.Assessment {
	if m == nil {
		return nil
	}

	var questions []*pb.Question
	q := m.Questions
	for _, question := range q {
		questions = append(questions, question.ToProto())
	}

	var attempts []*pb.AssessmentAttempt
	s := m.Attempts
	for _, attempt := range s {
		attempts = append(attempts, attempt.ToProto())
	}

	return &pb.Assessment{
		Id:           m.ID,
		Name:         m.Name,
		Description:  m.Description,
		Notes:        m.Notes,
		ImageUrl:     m.ImageURL,
		Difficulty:   m.Difficulty,
		TimeAllowed:  m.TimeAllowed,
		Type:         m.Type,
		Randomise:    m.Randomise,
		NumQuestions: m.NumQuestions,
		Questions:    questions,
		Attempts:     attempts,
	}
}

// ToProto maps the ORM AssessmentAttempt model to the proto model
func (m *AssessmentAttempt) ToProto() *pb.AssessmentAttempt {
	if m == nil {
		return nil
	}
	startedAt := helpers.TimeToProto(m.StartedAt)
	completedAt := helpers.TimeToProto(m.CompletedAt)
	return &pb.AssessmentAttempt{
		Id:           m.ID,
		AssessmentId: m.AssessmentID,
		CandidateId:  m.CandidateID,
		Status:       m.Status,
		StartedAt:    startedAt,
		CompletedAt:  completedAt,
		Score:        m.Score,
	}
}

// ToProto maps the ORM Question model to the proto model
func (m *Question) ToProto() *pb.Question {
	if m == nil {
		return nil
	}

	var tags []*pb.Tag
	t := m.Tags
	for _, tag := range t {
		tags = append(tags, tag.ToProto())
	}

	var assessments []*pb.Assessment
	a := m.Assessments
	for _, assessment := range a {
		assessments = append(assessments, assessment.ToProto())
	}

	var responses []*pb.Response
	r := m.Responses
	for _, response := range r {
		responses = append(responses, response.ToProto())
	}

	return &pb.Question{
		Id:          m.ID,
		CreatedBy:   m.CreatedBy,
		Type:        m.Type,
		Text:        m.Text,
		ImageUrl:    m.ImageURL,
		Options:     m.Options,
		Answer:      m.Answer,
		Tags:        tags,
		Assessments: assessments,
		Responses:   responses,
	}
}

// ToProto maps the ORM Tag model to the proto model
func (m *Tag) ToProto() *pb.Tag {
	if m == nil {
		return nil
	}
	return &pb.Tag{
		Id:   m.ID,
		Name: m.Name,
	}
}

// ToProto maps the ORM QuestionTag model to the proto model
func (m *QuestionTag) ToProto() *pb.QuestionTag {
	if m == nil {
		return nil
	}
	return &pb.QuestionTag{
		Id:         m.ID,
		QuestionId: m.QuestionID,
		TagId:      m.TagID,
	}
}

// ToProto maps the ORM Response model to the proto model
func (m *Response) ToProto() *pb.Response {
	if m == nil {
		return nil
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)

	return &pb.Response{
		Id:          m.ID,
		QuestionId:  m.QuestionID,
		CandidateId: m.CandidateID,
		Selection:   m.Selection,
		Text:        m.Text,
		Score:       m.Score,
		TimeTaken:   m.TimeTaken,
		CreatedAt:   createdAt,
	}
}

// ToProto maps the ORM AssessmentQuestion model to the proto model
func (m *AssessmentQuestion) ToProto() *pb.AssessmentQuestion {
	if m == nil {
		return nil
	}
	return &pb.AssessmentQuestion{
		Id:           m.ID,
		AssessmentId: m.AssessmentID,
		QuestionId:   m.QuestionID,
	}
}
