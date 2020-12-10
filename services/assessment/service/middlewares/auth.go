package middlewares

import (
	"context"
	"errors"
	"in-backend/helpers"
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

	namespace = "https://hubbedin/"
	idKey     = namespace + "id"
	rolesKey  = namespace + "roles"
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

	roles := claims[rolesKey].([]string)
	if helpers.IsStringInSlice("Admin", roles) {
		return true, nil
	}

	id, err := strconv.ParseUint(claims[idKey].(string), 10, 64)
	if err != nil {
		return false, err
	}
	if ownerID != nil && id == *ownerID {
		return true, nil
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
func (mw authMiddleware) GetAllAssessments(ctx context.Context, f models.AssessmentFilters) ([]*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.GetAllAssessments(ctx, f)
}

// GetAssessmentByID returns a Assessment by ID
func (mw authMiddleware) GetAssessmentByID(ctx context.Context, id uint64) (*models.Assessment, error) {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, errAuth
	}
	return mw.next.GetAssessmentByID(ctx, id)
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

/* --------------- Assessment Status --------------- */

// CreateAssessmentStatus creates a new AssessmentStatus
func (mw authMiddleware) CreateAssessmentStatus(ctx context.Context, m *models.AssessmentStatus) (*models.AssessmentStatus, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.CreateAssessmentStatus(ctx, m)
}

// UpdateAssessmentStatus updates a AssessmentStatus
func (mw authMiddleware) UpdateAssessmentStatus(ctx context.Context, m *models.AssessmentStatus) (*models.AssessmentStatus, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.UpdateAssessmentStatus(ctx, m)
}

// DeleteAssessmentStatus deletes a AssessmentStatus by ID
func (mw authMiddleware) DeleteAssessmentStatus(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteAssessmentStatus(ctx, id)
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

/* --------------- Response --------------- */

// CreateResponse creates a new Response
func (mw authMiddleware) CreateResponse(ctx context.Context, m *models.Response) (*models.Response, error) {
	owns, err := checkAdminOrOwner(ctx, &m.CandidateID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, errAuth
	}
	return mw.next.CreateResponse(ctx, m)
}

// DeleteResponse deletes a Response by ID
func (mw authMiddleware) DeleteResponse(ctx context.Context, id uint64) error {
	isAdmin, err := checkAdminOrOwner(ctx, nil)
	if err != nil {
		return err
	}
	if !isAdmin {
		return errAuth
	}
	return mw.next.DeleteResponse(ctx, id)
}
