package service

import (
	"context"

	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
)

// Service implements the assessment Service interface
type service struct {
	repository interfaces.Repository
}

// New creates and returns a new Service that implements the assessment Service interface
func New(r interfaces.Repository) interfaces.Service {
	return &service{
		repository: r,
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
func (s *service) GetAllAssessments(ctx context.Context, f models.AssessmentFilters) ([]*models.Assessment, error) {
	m, err := s.repository.GetAllAssessments(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAssessmentByID returns a Assessment by ID
func (s *service) GetAssessmentByID(ctx context.Context, id uint64) (*models.Assessment, error) {
	m, err := s.repository.GetAssessmentByID(ctx, id)
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

/* --------------- Assessment Status --------------- */

// CreateAssessmentStatus creates a new AssessmentStatus
func (s *service) CreateAssessmentStatus(ctx context.Context, model *models.AssessmentStatus) (*models.AssessmentStatus, error) {
	m, err := s.repository.CreateAssessmentStatus(ctx, model)
	if err != nil {
		return nil, err
	}

	return m, err
}

// UpdateAssessmentStatus updates a AssessmentStatus
func (s *service) UpdateAssessmentStatus(ctx context.Context, model *models.AssessmentStatus) (*models.AssessmentStatus, error) {
	m, err := s.repository.UpdateAssessmentStatus(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteAssessmentStatus deletes a AssessmentStatus by ID
func (s *service) DeleteAssessmentStatus(ctx context.Context, id uint64) error {
	err := s.repository.DeleteAssessmentStatus(ctx, id)
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

/* --------------- Response --------------- */

// CreateResponse creates a new Response
func (s *service) CreateResponse(ctx context.Context, model *models.Response) (*models.Response, error) {
	m, err := s.repository.CreateResponse(ctx, model)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteResponse deletes a Response by ID
func (s *service) DeleteResponse(ctx context.Context, id uint64) error {
	err := s.repository.DeleteResponse(ctx, id)
	if err != nil {
		return err
	}
	return err
}
