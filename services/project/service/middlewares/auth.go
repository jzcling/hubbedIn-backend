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
	next       project.Service
	repository project.Repository
}

var (
	errAuth = errors.New("Forbidden")

	idKey    = "https://hubbedin/id"
	rolesKey = "https://hubbedin/roles"
)

// NewAuthMiddleware creates and returns a new Auth Middleware that implements the project Service interface
func NewAuthMiddleware(svc project.Service, r project.Repository) project.Service {
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

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (mw authMiddleware) CreateProject(ctx context.Context, m *models.Project, cid uint64) (*models.Project, error) {
	role, _, err := getRoleAndID(ctx, &cid)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.CreateProject(ctx, m, cid)
}

// GetAllProjects returns all Projects
func (mw authMiddleware) GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error) {
	var role *string
	var err error
	if f.CandidateID > 0 {
		role, _, err = getRoleAndID(ctx, &f.CandidateID)
	} else {
		role, _, err = getRoleAndID(ctx, nil)
	}
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && *role != "Owner" {
		return nil, errAuth
	}
	return mw.next.GetAllProjects(ctx, f)
}

// GetProjectByID returns a Project by ID
func (mw authMiddleware) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	cp, err := mw.repository.GetCandidateProject(ctx, *cid, id)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && cp.CandidateID != *cid {
		return nil, errAuth
	}
	return mw.next.GetProjectByID(ctx, id)
}

// UpdateProject updates a Project
func (mw authMiddleware) UpdateProject(ctx context.Context, m *models.Project) (*models.Project, error) {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return nil, err
	}
	cp, err := mw.repository.GetCandidateProject(ctx, *cid, m.ID)
	if err != nil {
		return nil, err
	}
	if *role != "Admin" && cp.CandidateID != *cid {
		return nil, errAuth
	}
	return mw.next.UpdateProject(ctx, m)
}

// DeleteProject deletes a Project by ID
func (mw authMiddleware) DeleteProject(ctx context.Context, id uint64) error {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	cp, err := mw.repository.GetCandidateProject(ctx, *cid, id)
	if err != nil {
		return err
	}
	if *role != "Admin" && cp.CandidateID != *cid {
		return errAuth
	}
	return mw.next.DeleteProject(ctx, id)
}

// ScanProject scans a Project using sonarqube
func (mw authMiddleware) ScanProject(ctx context.Context, id uint64) error {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	cp, err := mw.repository.GetCandidateProject(ctx, *cid, id)
	if err != nil {
		return err
	}
	if *role != "Admin" && cp.CandidateID != *cid {
		return errAuth
	}
	return mw.next.ScanProject(ctx, id)
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (mw authMiddleware) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	role, _, err := getRoleAndID(ctx, &m.CandidateID)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
	}
	return mw.next.CreateCandidateProject(ctx, m)
}

// DeleteCandidateProject deletes a CandidateProject by ID
func (mw authMiddleware) DeleteCandidateProject(ctx context.Context, id uint64) error {
	cp, err := mw.repository.GetCandidateProjectByID(ctx, id)
	if err != nil {
		return err
	}
	role, _, err := getRoleAndID(ctx, &cp.CandidateID)
	if err != nil {
		return err
	}
	if *role != "Admin" && *role != "Owner" {
		return errAuth
	}
	return mw.next.DeleteCandidateProject(ctx, id)
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (mw authMiddleware) CreateRating(ctx context.Context, m *models.Rating) error {
	role, cid, err := getRoleAndID(ctx, nil)
	if err != nil {
		return err
	}
	cp, err := mw.repository.GetCandidateProject(ctx, *cid, m.ProjectID)
	if err != nil {
		return err
	}
	if *role != "Admin" && cp.CandidateID != *cid {
		return errAuth
	}
	return mw.next.CreateRating(ctx, m)
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (mw authMiddleware) DeleteRating(ctx context.Context, id uint64) error {
	return mw.next.DeleteRating(ctx, id)
}
