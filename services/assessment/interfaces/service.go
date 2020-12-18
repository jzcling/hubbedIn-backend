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
	GetAllAssessments(ctx context.Context, f models.AssessmentFilters, role *string, cid *uint64) ([]*models.Assessment, error)

	// GetAssessmentByID finds and returns a Assessment by ID
	GetAssessmentByID(ctx context.Context, id uint64, role *string, cid *uint64) (*models.Assessment, error)

	// UpdateAssessment updates a Assessment
	UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error)

	// DeleteAssessment deletes a Assessment by ID
	DeleteAssessment(ctx context.Context, id uint64) error

	/* --------------- Assessment Attempt --------------- */

	// CreateAssessmentAttempt creates a new AssessmentAttempt
	CreateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error)

	// GetAssessmentAttemptByID finds and returns a AssessmentAttempt by ID
	GetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error)

	// LocalGetAssessmentAttemptByID returns a AssessmentAttempt by ID
	// This method is only for local server to server communication
	LocalGetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error)

	// UpdateAssessmentAttempt updates a AssessmentAttempt
	UpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error)

	// LocalUpdateAssessmentAttempt updates a AssessmentAttempt
	// This method is only for local server to server communication
	LocalUpdateAssessmentAttempt(ctx context.Context, model *models.AssessmentAttempt) (*models.AssessmentAttempt, error)

	// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
	DeleteAssessmentAttempt(ctx context.Context, id uint64) error

	/* --------------- Question --------------- */

	// CreateQuestion creates a new Question
	CreateQuestion(ctx context.Context, m *models.Question) (*models.Question, error)

	// BulkCreateQuestion creates a new Question
	BulkCreateQuestion(ctx context.Context, m []*models.Question) ([]*models.Question, error)

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

	/* --------------- Attempt Question --------------- */

	// UpdateAttemptQuestion updates a AttemptQuestion
	UpdateAttemptQuestion(ctx context.Context, m *models.AttemptQuestion) (*models.AttemptQuestion, error)
}
