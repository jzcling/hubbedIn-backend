package interfaces

import (
	"context"
	"in-backend/services/assessment/models"
)

// Service describes the assessment Service
type Service interface {
	/* --------------- Assessment --------------- */

	// CreateAssessment creates a new Assessment
	CreateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error)

	// GetAllAssessments returns all Assessments
	GetAllAssessments(ctx context.Context, f models.AssessmentFilters) ([]*models.Assessment, error)

	// GetAssessmentByID finds and returns a Assessment by ID
	GetAssessmentByID(ctx context.Context, id uint64) (*models.Assessment, error)

	// UpdateAssessment updates a Assessment
	UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error)

	// DeleteAssessment deletes a Assessment by ID
	DeleteAssessment(ctx context.Context, id uint64) error

	/* --------------- AssessmentStatus --------------- */

	// CreateAssessmentStatus creates a new AssessmentStatus
	CreateAssessmentStatus(ctx context.Context, m *models.AssessmentStatus) (*models.AssessmentStatus, error)

	// UpdateAssessmentStatus updates a AssessmentStatus
	UpdateAssessmentStatus(ctx context.Context, m *models.AssessmentStatus) (*models.AssessmentStatus, error)

	// DeleteAssessmentStatus deletes a AssessmentStatus by ID
	DeleteAssessmentStatus(ctx context.Context, id uint64) error

	/* --------------- Question --------------- */

	// CreateQuestion creates a new Question
	CreateQuestion(ctx context.Context, m *models.Question) (*models.Question, error)

	// GetAllQuestions returns all Questions
	GetAllQuestions(ctx context.Context, f models.QuestionFilters) ([]*models.Question, error)

	// GetQuestionByID finds and returns a Question by ID
	GetQuestionByID(ctx context.Context, id uint64) (*models.Question, error)

	// UpdateQuestion updates a Question
	UpdateQuestion(ctx context.Context, m *models.Question) (*models.Question, error)

	// DeleteQuestion deletes a Question by ID
	DeleteQuestion(ctx context.Context, id uint64) error

	/* --------------- Tag --------------- */

	// CreateTag creates a new Tag
	CreateTag(ctx context.Context, m *models.Tag) (*models.Tag, error)

	// DeleteTag deletes a Tag by ID
	DeleteTag(ctx context.Context, id uint64) error

	/* --------------- Response --------------- */

	// CreateResponse creates a new Response
	CreateResponse(ctx context.Context, m *models.Response) (*models.Response, error)

	// DeleteResponse deletes a Response by ID
	DeleteResponse(ctx context.Context, id uint64) error
}
