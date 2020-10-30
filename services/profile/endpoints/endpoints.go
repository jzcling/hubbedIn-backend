package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
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
		c, err := s.CreateCandidate(ctx, req.Candidate)
		return CreateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

type CreateCandidateRequest struct {
	Candidate models.Candidate
}

type CreateCandidateResponse struct {
	Candidate models.Candidate
	Err       error
}

func makeGetAllCandidatesEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context) (interface{}, error) {
		c, err := s.GetAllCandidates(ctx)
		return GetAllCandidatesResponse{Candidates: c, Err: err}, nil
	}
}

type GetAllCandidatesRequest struct {
}

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

type GetCandidateByIDRequest struct {
	ID string
}

type GetCandidateByIDResponse struct {
	Candidate models.Candidate
	Err       error
}

func makeUpdateCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCandidateRequest)
		c, err := s.UpdateCandidate(ctx, req.Candidate)
		return UpdateCandidateResponse{Candidate: c, Err: err}, nil
	}
}

type UpdateCandidateRequest struct {
	Candidate models.Candidate
}

type UpdateCandidateResponse struct {
	Candidate models.Candidate
	Err       error
}

func DeleteCandidateEndpoint(s profile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCandidateRequest)
		err := s.DeleteCandidate(ctx, req.ID)
		return DeleteCandidateResponse{Err: err}, nil
	}
}

type DeleteCandidateRequest struct {
	ID string
}

type DeleteCandidateResponse struct {
	Err error
}
