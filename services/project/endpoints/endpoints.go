package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/project"
	"in-backend/services/project/models"
)

// Endpoints holds all Go kit endpoints for the Project Service.
type Endpoints struct {
	CreateProject  endpoint.Endpoint
	GetAllProjects endpoint.Endpoint
	GetProjectByID endpoint.Endpoint
	UpdateProject  endpoint.Endpoint
	DeleteProject  endpoint.Endpoint

	ScanProject endpoint.Endpoint

	CreateCandidateProject    endpoint.Endpoint
	DeleteCandidateProject    endpoint.Endpoint
	GetAllProjectsByCandidate endpoint.Endpoint

	CreateRating endpoint.Endpoint
	DeleteRating endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Project service.
func MakeEndpoints(s project.Service) Endpoints {
	return Endpoints{
		CreateProject:  makeCreateProjectEndpoint(s),
		GetAllProjects: makeGetAllProjectsEndpoint(s),
		GetProjectByID: makeGetProjectByIDEndpoint(s),
		UpdateProject:  makeUpdateProjectEndpoint(s),
		DeleteProject:  makeDeleteProjectEndpoint(s),

		ScanProject: makeScanProjectEndpoint(s),

		CreateCandidateProject:    makeCreateCandidateProjectEndpoint(s),
		DeleteCandidateProject:    makeDeleteCandidateProjectEndpoint(s),
		GetAllProjectsByCandidate: makeGetAllProjectsByCandidateEndpoint(s),

		CreateRating: makeCreateRatingEndpoint(s),
		DeleteRating: makeDeleteRatingEndpoint(s),
	}
}

/* -------------- Project -------------- */

func makeCreateProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProjectRequest)
		m, err := s.CreateProject(ctx, req.Project)
		return CreateProjectResponse{Project: m, Err: err}, nil
	}
}

// CreateProjectRequest declares the inputs required for creating a Project
type CreateProjectRequest struct {
	Project *models.Project
}

// CreateProjectResponse declares the outputs after attempting to create a Project
type CreateProjectResponse struct {
	Project *models.Project
	Err     error
}

func makeGetAllProjectsEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		m, err := s.GetAllProjects(ctx)
		return GetAllProjectsResponse{Projects: m, Err: err}, nil
	}
}

// GetAllProjectsRequest declares the inputs required for getting all Projects
// TODO: filters
type GetAllProjectsRequest struct {
	ID []uint64
}

// GetAllProjectsResponse declares the outputs after attempting to get all Projects
type GetAllProjectsResponse struct {
	Projects []*models.Project
	Err      error
}

func makeGetProjectByIDEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProjectByIDRequest)
		m, err := s.GetProjectByID(ctx, req.ID)
		return GetProjectByIDResponse{Project: m, Err: err}, nil
	}
}

// GetProjectByIDRequest declares the inputs required for getting a single Project by ID
type GetProjectByIDRequest struct {
	ID uint64
}

// GetProjectByIDResponse declares the outputs after attempting to get a single Project by ID
type GetProjectByIDResponse struct {
	Project *models.Project
	Err     error
}

func makeUpdateProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProjectRequest)
		m, err := s.UpdateProject(ctx, req.Project)
		return UpdateProjectResponse{Project: m, Err: err}, nil
	}
}

// UpdateProjectRequest declares the inputs required for updating a Project
type UpdateProjectRequest struct {
	ID      uint64
	Project *models.Project
}

// UpdateProjectResponse declares the outputs after attempting to update a Project
type UpdateProjectResponse struct {
	Project *models.Project
	Err     error
}

func makeDeleteProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProjectRequest)
		err := s.DeleteProject(ctx, req.ID)
		return DeleteProjectResponse{Err: err}, nil
	}
}

// DeleteProjectRequest declares the inputs required for deleting a Project
type DeleteProjectRequest struct {
	ID uint64
}

// DeleteProjectResponse declares the outputs after attempting to delete a Project
type DeleteProjectResponse struct {
	Err error
}

func makeScanProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ScanProjectRequest)
		err := s.ScanProject(ctx, req.ID)
		return ScanProjectResponse{Err: err}, nil
	}
}

// ScanProjectRequest declares the inputs required for scanning a Project
type ScanProjectRequest struct {
	ID uint64
}

// ScanProjectResponse declares the outputs after attempting to scanning a Project
type ScanProjectResponse struct {
	Err error
}

/* -------------- Candidate Project -------------- */

func makeCreateCandidateProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCandidateProjectRequest)
		err := s.CreateCandidateProject(ctx, req.CandidateProject)
		return CreateCandidateProjectResponse{Err: err}, nil
	}
}

// CreateCandidateProjectRequest declares the inputs required for creating a Candidate Project
type CreateCandidateProjectRequest struct {
	CandidateProject *models.CandidateProject
}

// CreateCandidateProjectResponse declares the outputs after attempting to create a Candidate Project
type CreateCandidateProjectResponse struct {
	Err error
}

func makeDeleteCandidateProjectEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCandidateProjectRequest)
		err := s.DeleteCandidateProject(ctx, req.CandidateID, req.ProjectID)
		return DeleteCandidateProjectResponse{Err: err}, nil
	}
}

// DeleteCandidateProjectRequest declares the inputs required for deleting a Candidate Project
type DeleteCandidateProjectRequest struct {
	CandidateID uint64
	ProjectID   uint64
}

// DeleteCandidateProjectResponse declares the outputs after attempting to delete a Candidate Project
type DeleteCandidateProjectResponse struct {
	Err error
}

func makeGetAllProjectsByCandidateEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllProjectsByCandidateRequest)
		m, err := s.GetAllProjectsByCandidate(ctx, req.CandidateID)
		return GetAllProjectsResponse{Projects: m, Err: err}, nil
	}
}

// GetAllProjectsByCandidateRequest declares the inputs required for getting all Projects by a Candidate
type GetAllProjectsByCandidateRequest struct {
	CandidateID uint64
}

// GetAllProjectsByCandidateResponse declares the outputs after attempting to get all Projects by a Candidate
type GetAllProjectsByCandidateResponse struct {
	Projects []*models.Project
	Err      error
}

/* -------------- Rating -------------- */

func makeCreateRatingEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRatingRequest)
		err := s.CreateRating(ctx, req.Rating)
		return CreateRatingResponse{Err: err}, nil
	}
}

// CreateRatingRequest declares the inputs required for creating a Project Rating
type CreateRatingRequest struct {
	Rating *models.Rating
}

// CreateRatingResponse declares the outputs after attempting to create a Project Rating
type CreateRatingResponse struct {
	Err error
}

func makeDeleteRatingEndpoint(s project.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRatingRequest)
		err := s.DeleteRating(ctx, req.ID)
		return DeleteRatingResponse{Err: err}, nil
	}
}

// DeleteRatingRequest declares the inputs required for deleting a Project Rating
type DeleteRatingRequest struct {
	ID uint64
}

// DeleteRatingResponse declares the outputs after attempting to delete a Project Rating
type DeleteRatingResponse struct {
	Err error
}
