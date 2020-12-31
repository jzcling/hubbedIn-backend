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
	next       interfaces.Service
	repository interfaces.Repository
}

var (
	errAuth = errors.New("Forbidden")

	idKey    = "https://hubbedin/id"
	rolesKey = "https://hubbedin/roles"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the assessment Service interface
func NewAuthMiddleware(svc interfaces.Service, r interfaces.Repository) interfaces.Service {
	return &authMiddleware{
		next:       svc,
		repository: r,
	}
}

func getRoleAndID(ctx context.Context, ownerID *uint64) (*string, *uint64, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, nil, err
	}

	var role string = ""
	var id uint64 = 0

	// this should come first so that role gets overwritten if owner is also an admin
	if claims[idKey] != nil {
		id, err = strconv.ParseUint(claims[idKey].(string), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		if ownerID != nil && id == *ownerID {
			role = "Owner"
		}
	}

	if claims[rolesKey] != nil {
		for _, r := range claims[rolesKey].([]interface{}) {
			roleCast := r.(string)
			if roleCast == "Admin" {
				role = "Admin"
			}
		}
	}

	return &role, &id, nil
}

func getClaims(ctx context.Context) (jwt.MapClaims, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errAuth
	}

	if len(headers["authorization"]) == 0 {
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
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.CreateAssessment(ctx, m)
}

// GetAllAssessments returns all Assessments
func (mw authMiddleware) GetAllAssessments(ctx context.Context, f models.AssessmentFilters, _ *string, _ *uint64) ([]*models.Assessment, error) {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mw.next.GetAllAssessments(ctx, f, role, cid)
}

// GetAssessmentByID returns a Assessment by ID
func (mw authMiddleware) GetAssessmentByID(ctx context.Context, id uint64, _ *string, _ *uint64) (*models.Assessment, error) {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mw.next.GetAssessmentByID(ctx, id, role, cid)
}

// UpdateAssessment updates a Assessment
func (mw authMiddleware) UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.UpdateAssessment(ctx, m)
}

// DeleteAssessment deletes a Assessment by ID
func (mw authMiddleware) DeleteAssessment(ctx context.Context, id uint64) error {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if *role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteAssessment(ctx, id)
}

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (mw authMiddleware) CreateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	role, _, err := getRoleAndID(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.CreateAssessmentAttempt(ctx, m)
}

// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (mw authMiddleware) GetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error) {
	aa, err := mw.repository.GetAssessmentAttemptByID(ctx, id)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &aa.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.GetAssessmentAttemptByID(ctx, id)
}

// LocalGetAssessmentAttemptByID returns a AssessmentAttempt by ID
// This method is only for local server to server communication
func (mw authMiddleware) LocalGetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error) {
	return mw.next.LocalGetAssessmentAttemptByID(ctx, id)
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (mw authMiddleware) UpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	aa, err := mw.repository.GetAssessmentAttemptByID(ctx, m.ID)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &aa.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.UpdateAssessmentAttempt(ctx, m)
}

// LocalUpdateAssessmentAttempt updates a AssessmentAttempt
// This method is only for local server to server communication
func (mw authMiddleware) LocalUpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	return mw.next.LocalUpdateAssessmentAttempt(ctx, m)
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (mw authMiddleware) DeleteAssessmentAttempt(ctx context.Context, id uint64) error {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if *role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteAssessmentAttempt(ctx, id)
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (mw authMiddleware) CreateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.CreateQuestion(ctx, m)
}

// BulkCreateQuestion creates a new Question
func (mw authMiddleware) BulkCreateQuestion(ctx context.Context, m []*models.Question) ([]*models.Question, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.BulkCreateQuestion(ctx, m)
}

// GetAllQuestions returns all Questions
func (mw authMiddleware) GetAllQuestions(ctx context.Context, f models.QuestionFilters) ([]*models.Question, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.GetAllQuestions(ctx, f)
}

// GetQuestionByID returns a Question by ID
func (mw authMiddleware) GetQuestionByID(ctx context.Context, id uint64) (*models.Question, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.GetQuestionByID(ctx, id)
}

// UpdateQuestion updates a Question
func (mw authMiddleware) UpdateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" {
		return nil, errAuth
	}
	return mw.next.UpdateQuestion(ctx, m)
}

// DeleteQuestion deletes a Question by ID
func (mw authMiddleware) DeleteQuestion(ctx context.Context, id uint64) error {
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if *role != "Admin" {
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
	role, _, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	if *role != "Admin" {
		return errAuth
	}
	return mw.next.DeleteTag(ctx, id)
}

/* --------------- Attempt Question --------------- */

// UpdateAttemptQuestion updates a AttemptQuestion
func (mw authMiddleware) UpdateAttemptQuestion(ctx context.Context, m *models.AttemptQuestion) (*models.AttemptQuestion, error) {
	aq, err := mw.repository.GetAttemptQuestionByID(ctx, m.ID)
	if err != nil {
		return nil, err
	}
	role, _, err := getRoleAndID(ctx, &aq.CandidateID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.UpdateAttemptQuestion(ctx, m)
}
