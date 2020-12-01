package middlewares

import (
	"context"
	"errors"
	"in-backend/services/project"
	"in-backend/services/project/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type authMiddleware struct {
	next project.Service
}

var (
	errAuth = errors.New("Forbidden")

	namespace = "https://hubbedin/id"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the project Service interface
func NewAuthMiddleware(svc project.Service) project.Service {
	return &authMiddleware{
		next: svc,
	}
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

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (mw authMiddleware) CreateProject(ctx context.Context, model *models.Project, cid uint64) (*models.Project, error) {
	claims, err := getClaims(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	if id != cid {
		return nil, errAuth
	}
	return mw.next.CreateProject(ctx, model, cid)
}

// GetAllProjects returns all Projects
func (mw authMiddleware) GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error) {
	return mw.next.GetAllProjects(ctx, f)
}

// GetProjectByID returns a Project by ID
func (mw authMiddleware) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	return mw.next.GetProjectByID(ctx, id)
}

// UpdateProject updates a Project
func (mw authMiddleware) UpdateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	return mw.next.UpdateProject(ctx, model)
}

// DeleteProject deletes a Project by ID
func (mw authMiddleware) DeleteProject(ctx context.Context, id uint64) error {
	return mw.next.DeleteProject(ctx, id)
}

// ScanProject scans a Project using sonarqube
func (mw authMiddleware) ScanProject(ctx context.Context, id uint64) error {
	return mw.next.ScanProject(ctx, id)
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (mw authMiddleware) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	claims, err := getClaims(ctx)
	if err != nil {
		return err
	}
	id, err := strconv.ParseUint(claims[namespace].(string), 10, 64)
	if err != nil {
		return err
	}
	if id != m.CandidateID {
		return errAuth
	}
	return mw.next.CreateCandidateProject(ctx, m)
}

// DeleteCandidateProject deletes a CandidateProject by ID
func (mw authMiddleware) DeleteCandidateProject(ctx context.Context, id uint64) error {
	return mw.next.DeleteCandidateProject(ctx, id)
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (mw authMiddleware) CreateRating(ctx context.Context, m *models.Rating) error {
	return mw.next.CreateRating(ctx, m)
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (mw authMiddleware) DeleteRating(ctx context.Context, id uint64) error {
	return mw.next.DeleteRating(ctx, id)
}
