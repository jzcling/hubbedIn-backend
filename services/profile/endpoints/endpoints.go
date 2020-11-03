package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/profile"
	"in-backend/services/profile/models"
)

// Endpoints holds all Go kit endpoints for the Profile Service.
type Endpoints struct {
	CreateCandidate  endpoint.Endpoint
	GetAllCandidates endpoint.Endpoint
	GetCandidateByID endpoint.Endpoint
	UpdateCandidate  endpoint.Endpoint
	DeleteCandidate  endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Profile service.
func MakeEndpoints(s profile.Service) Endpoints {
	return Endpoints{
		CreateCandidate:  makeCreateCandidateEndpoint(s),
		GetAllCandidates: makeGetAllCandidatesEndpoint(s),
		GetCandidateByID: makeGetCandidateByIDEndpoint(s),
		UpdateCandidate:  makeUpdateCandidateEndpoint(s),
		DeleteCandidate:  makeDeleteCandidateEndpoint(s)}
}

func makeCreateCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCandidateRequest)
		c, err := s.CreateCandidate(ctx, &req.Candidate)
		return CreateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

// CreateCandidateRequest declares the inputs required for creating a candidate
type CreateCandidateRequest struct {
	Candidate models.Candidate
}

// CreateCandidateResponse declares the outputs after attempting to create a candidate
type CreateCandidateResponse struct {
	Candidate models.Candidate
	Err       error
}

func makeGetAllCandidatesEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		c, err := s.GetAllCandidates(ctx)
		return GetAllCandidatesResponse{Candidates: c, Err: err}, nil
	}
}

// GetAllCandidatesRequest declares the inputs required for getting all candidates
type GetAllCandidatesRequest struct {
	// TODO
}

// GetAllCandidatesResponse declares the outputs after attempting to get all candidates
type GetAllCandidatesResponse struct {
	Candidates []models.Candidate
	Err        error
}

func makeGetCandidateByIDEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCandidateByIDRequest)
		c, err := s.GetCandidateByID(ctx, req.ID)
		return GetCandidateByIDResponse{Candidate: c, Err: err}, nil
	}
}

// GetCandidateByIDRequest declares the inputs required for getting a single candidate by ID
type GetCandidateByIDRequest struct {
	ID uint64
}

// GetCandidateByIDResponse declares the outputs after attempting to get a single candidate by ID
type GetCandidateByIDResponse struct {
	Candidate models.Candidate
	Err       error
}

func makeUpdateCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCandidateRequest)
		c, err := s.UpdateCandidate(ctx, &req.Candidate)
		return UpdateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

// UpdateCandidateRequest declares the inputs required for updating a candidate
type UpdateCandidateRequest struct {
	Candidate models.Candidate
}

// UpdateCandidateResponse declares the outputs after attempting to update a candidate
type UpdateCandidateResponse struct {
	Candidate models.Candidate
	Err       error
}

func makeDeleteCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCandidateRequest)
		err := s.DeleteCandidate(ctx, req.ID)
		return DeleteCandidateResponse{Err: err}, nil
	}
}

// DeleteCandidateRequest declares the inputs required for deleting a candidate
type DeleteCandidateRequest struct {
	ID uint64
}

// DeleteCandidateResponse declares the outputs after attempting to delete a candidate
type DeleteCandidateResponse struct {
	Err error
}
