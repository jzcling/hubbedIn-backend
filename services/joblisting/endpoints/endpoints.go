package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/joblisting"
	"in-backend/services/joblisting/models"
)

// Endpoints holds all Go kit endpoints for the Joblisting Service.
type Endpoints struct {
	CreateJoblisting  endpoint.Endpoint
	GetAllJoblistings endpoint.Endpoint
	GetJoblistingByID endpoint.Endpoint
	UpdateJoblisting  endpoint.Endpoint
	DeleteJoblisting  endpoint.Endpoint

	ScanJoblisting endpoint.Endpoint

	CreateCandidateJoblisting endpoint.Endpoint
	DeleteCandidateJoblisting endpoint.Endpoint

	CreateRating endpoint.Endpoint
	DeleteRating endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Joblisting service.
func MakeEndpoints(s joblisting.Service) Endpoints {
	return Endpoints{
		CreateJoblisting:  makeCreateJoblistingEndpoint(s),
		GetAllJoblistings: makeGetAllJoblistingsEndpoint(s),
		GetJoblistingByID: makeGetJoblistingByIDEndpoint(s),
		UpdateJoblisting:  makeUpdateJoblistingEndpoint(s),
		DeleteJoblisting:  makeDeleteJoblistingEndpoint(s),

		ScanJoblisting: makeScanJoblistingEndpoint(s),

		CreateCandidateJoblisting: makeCreateCandidateJoblistingEndpoint(s),
		DeleteCandidateJoblisting: makeDeleteCandidateJoblistingEndpoint(s),

		CreateRating: makeCreateRatingEndpoint(s),
		DeleteRating: makeDeleteRatingEndpoint(s),
	}
}

/* -------------- Joblisting -------------- */

func makeCreateJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateJoblistingRequest)
		m, err := s.CreateJoblisting(ctx, req.Joblisting, req.CandidateID)
		return CreateJoblistingResponse{Joblisting: m, Err: err}, nil
	}
}

// CreateJoblistingRequest declares the inputs required for creating a Joblisting
type CreateJoblistingRequest struct {
	Joblisting  *models.Joblisting
	CandidateID uint64
}

// CreateJoblistingResponse declares the outputs after attempting to create a Joblisting
type CreateJoblistingResponse struct {
	Joblisting *models.Joblisting
	Err        error
}

func makeGetAllJoblistingsEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllJoblistingsRequest)
		f := models.JoblistingFilters(req)
		m, err := s.GetAllJoblistings(ctx, f)
		return GetAllJoblistingsResponse{Joblistings: m, Err: err}, nil
	}
}

// GetAllJoblistingsRequest declares the inputs required for getting all Joblistings
type GetAllJoblistingsRequest struct {
	ID          []uint64
	CandidateID uint64
	Name        string
	RepoURL     string
}

// GetAllJoblistingsResponse declares the outputs after attempting to get all Joblistings
type GetAllJoblistingsResponse struct {
	Joblistings []*models.Joblisting
	Err         error
}

func makeGetJoblistingByIDEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetJoblistingByIDRequest)
		m, err := s.GetJoblistingByID(ctx, req.ID)
		return GetJoblistingByIDResponse{Joblisting: m, Err: err}, nil
	}
}

// GetJoblistingByIDRequest declares the inputs required for getting a single Joblisting by ID
type GetJoblistingByIDRequest struct {
	ID uint64
}

// GetJoblistingByIDResponse declares the outputs after attempting to get a single Joblisting by ID
type GetJoblistingByIDResponse struct {
	Joblisting *models.Joblisting
	Err        error
}

func makeUpdateJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateJoblistingRequest)
		m, err := s.UpdateJoblisting(ctx, req.Joblisting)
		return UpdateJoblistingResponse{Joblisting: m, Err: err}, nil
	}
}

// UpdateJoblistingRequest declares the inputs required for updating a Joblisting
type UpdateJoblistingRequest struct {
	ID         uint64
	Joblisting *models.Joblisting
}

// UpdateJoblistingResponse declares the outputs after attempting to update a Joblisting
type UpdateJoblistingResponse struct {
	Joblisting *models.Joblisting
	Err        error
}

func makeDeleteJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteJoblistingRequest)
		err := s.DeleteJoblisting(ctx, req.ID)
		return DeleteJoblistingResponse{Err: err}, nil
	}
}

// DeleteJoblistingRequest declares the inputs required for deleting a Joblisting
type DeleteJoblistingRequest struct {
	ID uint64
}

// DeleteJoblistingResponse declares the outputs after attempting to delete a Joblisting
type DeleteJoblistingResponse struct {
	Err error
}

func makeScanJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ScanJoblistingRequest)
		err := s.ScanJoblisting(ctx, req.ID)
		return ScanJoblistingResponse{Err: err}, nil
	}
}

// ScanJoblistingRequest declares the inputs required for scanning a Joblisting
type ScanJoblistingRequest struct {
	ID uint64
}

// ScanJoblistingResponse declares the outputs after attempting to scanning a Joblisting
type ScanJoblistingResponse struct {
	Err error
}

/* -------------- Candidate Joblisting -------------- */

func makeCreateCandidateJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCandidateJoblistingRequest)
		err := s.CreateCandidateJoblisting(ctx, req.CandidateJoblisting)
		return CreateCandidateJoblistingResponse{Err: err}, nil
	}
}

// CreateCandidateJoblistingRequest declares the inputs required for creating a Candidate Joblisting
type CreateCandidateJoblistingRequest struct {
	CandidateJoblisting *models.CandidateJoblisting
}

// CreateCandidateJoblistingResponse declares the outputs after attempting to create a Candidate Joblisting
type CreateCandidateJoblistingResponse struct {
	Err error
}

func makeDeleteCandidateJoblistingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCandidateJoblistingRequest)
		err := s.DeleteCandidateJoblisting(ctx, req.ID)
		return DeleteCandidateJoblistingResponse{Err: err}, nil
	}
}

// DeleteCandidateJoblistingRequest declares the inputs required for deleting a Candidate Joblisting
type DeleteCandidateJoblistingRequest struct {
	ID uint64
}

// DeleteCandidateJoblistingResponse declares the outputs after attempting to delete a Candidate Joblisting
type DeleteCandidateJoblistingResponse struct {
	Err error
}

/* -------------- Rating -------------- */

func makeCreateRatingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRatingRequest)
		err := s.CreateRating(ctx, req.Rating)
		return CreateRatingResponse{Err: err}, nil
	}
}

// CreateRatingRequest declares the inputs required for creating a Joblisting Rating
type CreateRatingRequest struct {
	Rating *models.Rating
}

// CreateRatingResponse declares the outputs after attempting to create a Joblisting Rating
type CreateRatingResponse struct {
	Err error
}

func makeDeleteRatingEndpoint(s joblisting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRatingRequest)
		err := s.DeleteRating(ctx, req.ID)
		return DeleteRatingResponse{Err: err}, nil
	}
}

// DeleteRatingRequest declares the inputs required for deleting a Joblisting Rating
type DeleteRatingRequest struct {
	ID uint64
}

// DeleteRatingResponse declares the outputs after attempting to delete a Joblisting Rating
type DeleteRatingResponse struct {
	Err error
}
