package interfaces

import (
	"context"
	"in-backend/services/assessment/models"
)

// Repository declares the repository for assessments
type Repository interface {
	/* --------------- Assessment --------------- */

	// CreateAssessment creates a new Assessment
	CreateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error)

	// GetAllAssessments returns all Assessments
	GetAllAssessments(ctx context.Context, f models.AssessmentFilters, admin *bool) ([]*models.Assessment, error)

	// GetAssessmentByID finds and returns a Assessment by ID
	GetAssessmentByID(ctx context.Context, id uint64, admin *bool) (*models.Assessment, error)

	// UpdateAssessment updates a Assessment
	UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error)

	// DeleteAssessment deletes a Assessment by ID
	DeleteAssessment(ctx context.Context, id uint64) error

	/* --------------- Assessment Attempt --------------- */

	// CreateAssessmentAttempt creates a new AssessmentAttempt
	CreateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error)

	// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
	GetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error)

	// UpdateAssessmentAttempt updates a AssessmentAttempt
	UpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error)

	// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
	DeleteAssessmentAttempt(ctx context.Context, id uint64) error

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

	/* --------------- Attempt Question --------------- */

	// GetAttemptQuestionByID returns a AttemptQuestion by ID
	GetAttemptQuestionByID(ctx context.Context, id uint64) (*models.AttemptQuestion, error)

	// UpdateAttemptQuestion updates a AttemptQuestion
	UpdateAttemptQuestion(ctx context.Context, m *models.AttemptQuestion) (*models.AttemptQuestion, error)
}
