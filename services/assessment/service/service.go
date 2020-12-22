package service

import (
	"context"
	"errors"
	"time"

	"in-backend/helpers"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"

	"github.com/gocraft/work"
)

// Service implements the assessment Service interface
type service struct {
	repository interfaces.Repository
	enqueuer   *work.Enqueuer
}

// New creates and returns a new Service that implements the assessment Service interface
func New(r interfaces.Repository, e *work.Enqueuer) interfaces.Service {
	return &service{
		repository: r,
		enqueuer:   e,
	}
}

/* --------------- Assessment --------------- */

// CreateAssessment creates a new Assessment
func (s *service) CreateAssessment(ctx context.Context, model *models.Assessment) (*models.Assessment, error) {
	m, err := s.repository.CreateAssessment(ctx, model)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetAllAssessments returns all Assessments
func (s *service) GetAllAssessments(ctx context.Context, f models.AssessmentFilters, role *string, cid *uint64) ([]*models.Assessment, error) {
	m, err := s.repository.GetAllAssessments(ctx, f, role, cid)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAssessmentByID returns a Assessment by ID
func (s *service) GetAssessmentByID(ctx context.Context, id uint64, role *string, cid *uint64) (*models.Assessment, error) {
	m, err := s.repository.GetAssessmentByID(ctx, id, role, cid)
	if err != nil {
		return nil, err
	}
	return m, err
}

// UpdateAssessment updates a Assessment
func (s *service) UpdateAssessment(ctx context.Context, model *models.Assessment) (*models.Assessment, error) {
	m, err := s.repository.UpdateAssessment(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteAssessment deletes a Assessment by ID
func (s *service) DeleteAssessment(ctx context.Context, id uint64) error {
	err := s.repository.DeleteAssessment(ctx, id)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (s *service) CreateAssessmentAttempt(ctx context.Context, model *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	aa, err := s.repository.GetLatestAssessmentAttemptByCandidate(ctx, model.CandidateID)
	if err != nil {
		return nil, err
	}

	_, months, _, _, _, _ := helpers.TimeDiff(time.Now(), *aa.StartedAt)
	if months < 3 {
		return nil, errors.New("Minimum of 3 months between attempts")
	}

	m, err := s.repository.CreateAssessmentAttempt(ctx, model)
	if err != nil {
		return nil, err
	}

	role := "Admin"
	a, err := s.repository.GetAssessmentByID(ctx, m.AssessmentID, &role, nil)
	if err != nil {
		return nil, err
	}
	s.scheduleAssessmentAttemptEnd(int64(m.ID), int64(a.TimeAllowed))

	return m, err
}

func (s *service) scheduleAssessmentAttemptEnd(id int64, ta int64) error {
	_, err := s.enqueuer.EnqueueIn("end_assessment_attempt", ta, work.Q{"id": id})
	return err
}

// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (s *service) GetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error) {
	m, err := s.repository.GetAssessmentAttemptByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

// LocalGetAssessmentAttemptByID returns a AssessmentAttempt by ID
// This method is only for local server to server communication
func (s *service) LocalGetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error) {
	m, err := s.repository.GetAssessmentAttemptByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (s *service) UpdateAssessmentAttempt(ctx context.Context, model *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	m, err := s.repository.UpdateAssessmentAttempt(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// LocalUpdateAssessmentAttempt updates a AssessmentAttempt
// This method is only for local server to server communication
func (s *service) LocalUpdateAssessmentAttempt(ctx context.Context, model *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	m, err := s.repository.UpdateAssessmentAttempt(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (s *service) DeleteAssessmentAttempt(ctx context.Context, id uint64) error {
	err := s.repository.DeleteAssessmentAttempt(ctx, id)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (s *service) CreateQuestion(ctx context.Context, model *models.Question) (*models.Question, error) {
	m, err := s.repository.CreateQuestion(ctx, model)
	if err != nil {
		return nil, err
	}

	return m, err
}

// BulkCreateQuestion creates a new Question
func (s *service) BulkCreateQuestion(ctx context.Context, m []*models.Question) ([]*models.Question, error) {
	m, err := s.repository.BulkCreateQuestion(ctx, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetAllQuestions returns all Questions
func (s *service) GetAllQuestions(ctx context.Context, f models.QuestionFilters) ([]*models.Question, error) {
	m, err := s.repository.GetAllQuestions(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetQuestionByID returns a Question by ID
func (s *service) GetQuestionByID(ctx context.Context, id uint64) (*models.Question, error) {
	m, err := s.repository.GetQuestionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

// UpdateQuestion updates a Question
func (s *service) UpdateQuestion(ctx context.Context, model *models.Question) (*models.Question, error) {
	m, err := s.repository.UpdateQuestion(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteQuestion deletes a Question by ID
func (s *service) DeleteQuestion(ctx context.Context, id uint64) error {
	err := s.repository.DeleteQuestion(ctx, id)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Tag --------------- */

// CreateTag creates a new Tag
func (s *service) CreateTag(ctx context.Context, model *models.Tag) (*models.Tag, error) {
	m, err := s.repository.CreateTag(ctx, model)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteTag deletes a Tag by ID
func (s *service) DeleteTag(ctx context.Context, id uint64) error {
	err := s.repository.DeleteTag(ctx, id)
	if err != nil {
		return err
	}
	return err
}

/* --------------- Attempt Question --------------- */

// UpdateAttemptQuestion updates a AttemptQuestion
func (s *service) UpdateAttemptQuestion(ctx context.Context, model *models.AttemptQuestion) (*models.AttemptQuestion, error) {
	m, err := s.repository.UpdateAttemptQuestion(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}
