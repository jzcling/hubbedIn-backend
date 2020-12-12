package middlewares

import (
	"context"
	"errors"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type authMiddleware struct {
	next interfaces.Service
}

var (
	errAuth = errors.New("Forbidden")

	idKey    = "https://hubbedin/id"
	rolesKey = "https://hubbedin/roles"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the assessment Service interface
func NewAuthMiddleware(svc interfaces.Service) interfaces.Service {
	return &authMiddleware{
		next: svc,
	}
}

func checkAdminOrOwner(ctx context.Context, ownerID *uint64) (bool, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return false, err
	}

	if claims[rolesKey] != nil {
		for _, r := range claims[rolesKey].([]interface{}) {
			role := r.(string)
			if role == "Admin" {
				return true, nil
			}
		}
	}

	if claims[idKey] != nil {
		id, err := strconv.ParseUint(claims[idKey].(string), 10, 64)
		if err != nil {
			return false, err
		}
		if ownerID != nil && id == *ownerID {
			return true, nil
		}
	}

	return false, nil
}

func getClaims(ctx context.Context) (jwt.MapClaims, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errAuth
	}
	tokenString := strings.Split(headers["authorization"][0], " ")[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, errAuth
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errAuth
	}
	return claims, nil
}

/* --------------- Assessment --------------- */

// CreateAssessment creates a new Assessment
func (mw authMiddleware) CreateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.CreateAssessment(ctx, m)
}

// GetAllAssessments returns all Assessments
func (mw authMiddleware) GetAllAssessments(ctx context.Context, f models.AssessmentFilters, _ *bool) ([]*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mw.next.GetAllAssessments(ctx, f, &isAdmin)
}

// GetAssessmentByID returns a Assessment by ID
func (mw authMiddleware) GetAssessmentByID(ctx context.Context, id uint64, _ *bool) (*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mw.next.GetAssessmentByID(ctx, id, &isAdmin)
}

// UpdateAssessment updates a Assessment
func (mw authMiddleware) UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.UpdateAssessment(ctx, m)
}

// DeleteAssessment deletes a Assessment by ID
func (mw authMiddleware) DeleteAssessment(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteAssessment(ctx, id)
}

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (mw authMiddleware) CreateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.CreateAssessmentAttempt(ctx, m)
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (mw authMiddleware) UpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.UpdateAssessmentAttempt(ctx, m)
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (mw authMiddleware) DeleteAssessmentAttempt(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteAssessmentAttempt(ctx, id)
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (mw authMiddleware) CreateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.CreateQuestion(ctx, m)
}

// GetAllQuestions returns all Questions
func (mw authMiddleware) GetAllQuestions(ctx context.Context, f models.QuestionFilters) ([]*models.Question, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.GetAllQuestions(ctx, f)
}

// GetQuestionByID returns a Question by ID
func (mw authMiddleware) GetQuestionByID(ctx context.Context, id uint64) (*models.Question, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.GetQuestionByID(ctx, id)
}

// UpdateQuestion updates a Question
func (mw authMiddleware) UpdateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.UpdateQuestion(ctx, m)
}

// DeleteQuestion deletes a Question by ID
func (mw authMiddleware) DeleteQuestion(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteQuestion(ctx, id)
}

/* --------------- Tag --------------- */

// CreateTag creates a new Tag
func (mw authMiddleware) CreateTag(ctx context.Context, m *models.Tag) (*models.Tag, error) {
	return mw.next.CreateTag(ctx, m)
}

// DeleteTag deletes a Tag by ID
func (mw authMiddleware) DeleteTag(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteTag(ctx, id)
}

/* --------------- Attempt Question --------------- */

// UpdateAttemptQuestion updates a AttemptQuestion
func (mw authMiddleware) UpdateAttemptQuestion(ctx context.Context, m *models.AttemptQuestion) (*models.AttemptQuestion, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.UpdateAttemptQuestion(ctx, m)
}
