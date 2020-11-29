package transport

import (
	"context"
	"in-backend/services/joblisting/endpoints"
	"in-backend/services/joblisting/models"
	"in-backend/services/joblisting/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for Joblisting Service.
type grpcServer struct {
	createJoblisting  kitgrpc.Handler
	getAllJoblistings kitgrpc.Handler
	getJoblistingByID kitgrpc.Handler
	updateJoblisting  kitgrpc.Handler
	deleteJoblisting  kitgrpc.Handler

	scanJoblisting kitgrpc.Handler

	createCandidateJoblisting kitgrpc.Handler
	deleteCandidateJoblisting kitgrpc.Handler

	createRating kitgrpc.Handler
	deleteRating kitgrpc.Handler

	logger log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.JoblistingServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createJoblisting: kitgrpc.NewServer(
			endpoints.CreateJoblisting,
			decodeCreateJoblistingRequest,
			encodeCreateJoblistingResponse,
			options...,
		),
		getAllJoblistings: kitgrpc.NewServer(
			endpoints.GetAllJoblistings,
			decodeGetAllJoblistingsRequest,
			encodeGetAllJoblistingsResponse,
			options...,
		),
		getJoblistingByID: kitgrpc.NewServer(
			endpoints.GetJoblistingByID,
			decodeGetJoblistingByIDRequest,
			encodeGetJoblistingByIDResponse,
			options...,
		),
		updateJoblisting: kitgrpc.NewServer(
			endpoints.UpdateJoblisting,
			decodeUpdateJoblistingRequest,
			encodeUpdateJoblistingResponse,
			options...,
		),
		deleteJoblisting: kitgrpc.NewServer(
			endpoints.DeleteJoblisting,
			decodeDeleteJoblistingRequest,
			encodeDeleteJoblistingResponse,
			options...,
		),

		scanJoblisting: kitgrpc.NewServer(
			endpoints.ScanJoblisting,
			decodeScanJoblistingRequest,
			encodeScanJoblistingResponse,
			options...,
		),

		createCandidateJoblisting: kitgrpc.NewServer(
			endpoints.CreateCandidateJoblisting,
			decodeCreateCandidateJoblistingRequest,
			encodeCreateCandidateJoblistingResponse,
			options...,
		),
		deleteCandidateJoblisting: kitgrpc.NewServer(
			endpoints.DeleteCandidateJoblisting,
			decodeDeleteCandidateJoblistingRequest,
			encodeDeleteCandidateJoblistingResponse,
			options...,
		),

		createRating: kitgrpc.NewServer(
			endpoints.CreateRating,
			decodeCreateRatingRequest,
			encodeCreateRatingResponse,
			options...,
		),
		deleteRating: kitgrpc.NewServer(
			endpoints.DeleteRating,
			decodeDeleteRatingRequest,
			encodeDeleteRatingResponse,
			options...,
		),

		logger: logger,
	}
}

/* --------------- Joblisting --------------- */

// CreateJoblisting creates a new Joblisting
func (s *grpcServer) CreateJoblisting(ctx context.Context, req *pb.CreateJoblistingRequest) (*pb.Joblisting, error) {
	_, rep, err := s.createJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Joblisting), nil
}

// decodeCreateJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJoblistingRequest)
	return endpoints.CreateJoblistingRequest{Joblisting: models.JoblistingToORM(req.Joblisting), CandidateID: req.CandidateId}, nil
}

// encodeCreateJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Joblisting.ToProto(), nil
	}
	return nil, err
}

// GetAllJoblistings returns all Joblistings
func (s *grpcServer) GetAllJoblistings(ctx context.Context, req *pb.GetAllJoblistingsRequest) (*pb.GetAllJoblistingsResponse, error) {
	_, rep, err := s.getAllJoblistings.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllJoblistingsResponse), nil
}

// decodeGetAllJoblistingsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllJoblistingsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllJoblistingsRequest)
	decoded := endpoints.GetAllJoblistingsRequest{
		ID:          req.Id,
		CandidateID: req.CandidateId,
		Name:        req.Name,
		RepoURL:     req.RepoUrl,
	}
	return decoded, nil
}

// encodeGetAllJoblistingsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllJoblistingsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllJoblistingsResponse)
	err := getError(res.Err)
	if err == nil {
		var joblistings []*pb.Joblisting
		for _, joblisting := range res.Joblistings {
			joblistings = append(joblistings, joblisting.ToProto())
		}
		return &pb.GetAllJoblistingsResponse{Joblistings: joblistings}, nil
	}
	return nil, err
}

// GetJoblistingByID returns a Joblisting by ID
func (s *grpcServer) GetJoblistingByID(ctx context.Context, req *pb.GetJoblistingByIDRequest) (*pb.Joblisting, error) {
	_, rep, err := s.getJoblistingByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Joblisting), nil
}

// decodeGetJoblistingByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetJoblistingByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetJoblistingByIDRequest)
	return endpoints.GetJoblistingByIDRequest{ID: req.Id}, nil
}

// encodeGetJoblistingByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetJoblistingByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetJoblistingByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Joblisting.ToProto(), nil
	}
	return nil, err
}

// UpdateJoblisting updates a Joblisting
func (s *grpcServer) UpdateJoblisting(ctx context.Context, req *pb.UpdateJoblistingRequest) (*pb.Joblisting, error) {
	_, rep, err := s.updateJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Joblisting), nil
}

// decodeUpdateJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateJoblistingRequest)
	return endpoints.UpdateJoblistingRequest{ID: req.Id, Joblisting: models.JoblistingToORM(req.Joblisting)}, nil
}

// encodeUpdateJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Joblisting.ToProto(), nil
	}
	return nil, err
}

// DeleteJoblisting deletes a Joblisting by ID
func (s *grpcServer) DeleteJoblisting(ctx context.Context, req *pb.DeleteJoblistingRequest) (*pb.DeleteJoblistingResponse, error) {
	_, rep, err := s.deleteJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteJoblistingResponse), nil
}

// decodeDeleteJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteJoblistingRequest)
	return endpoints.DeleteJoblistingRequest{ID: req.Id}, nil
}

// encodeDeleteJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteJoblistingResponse{}, nil
	}
	return nil, err
}

// ScanJoblisting scans a Joblisting using sonarqube
func (s *grpcServer) ScanJoblisting(ctx context.Context, req *pb.ScanJoblistingRequest) (*pb.ScanJoblistingResponse, error) {
	_, rep, err := s.scanJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ScanJoblistingResponse), nil
}

// decodeScanJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeScanJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ScanJoblistingRequest)
	return endpoints.ScanJoblistingRequest{ID: req.Id}, nil
}

// encodeScanJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeScanJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.ScanJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.ScanJoblistingResponse{}, nil
	}
	return nil, err
}

/* --------------- Candidate Joblisting --------------- */

// CreateCandidateJoblisting creates a new CandidateJoblisting
func (s *grpcServer) CreateCandidateJoblisting(ctx context.Context, req *pb.CreateCandidateJoblistingRequest) (*pb.CreateCandidateJoblistingResponse, error) {
	_, rep, err := s.createCandidateJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCandidateJoblistingResponse), nil
}

// decodeCreateCandidateJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCandidateJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCandidateJoblistingRequest)
	return endpoints.CreateCandidateJoblistingRequest{CandidateJoblisting: models.CandidateJoblistingToORM(req.CandidateJoblisting)}, nil
}

// encodeCreateCandidateJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCandidateJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCandidateJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateCandidateJoblistingResponse{}, nil
	}
	return nil, err
}

// DeleteCandidateJoblisting deletes a CandidateJoblisting by Candidate ID and Joblisting ID
func (s *grpcServer) DeleteCandidateJoblisting(ctx context.Context, req *pb.DeleteCandidateJoblistingRequest) (*pb.DeleteCandidateJoblistingResponse, error) {
	_, rep, err := s.deleteCandidateJoblisting.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteCandidateJoblistingResponse), nil
}

// decodeDeleteCandidateJoblistingRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteCandidateJoblistingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteCandidateJoblistingRequest)
	return endpoints.DeleteCandidateJoblistingRequest{ID: req.Id}, nil
}

// encodeDeleteCandidateJoblistingResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteCandidateJoblistingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteCandidateJoblistingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteCandidateJoblistingResponse{}, nil
	}
	return nil, err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (s *grpcServer) CreateRating(ctx context.Context, req *pb.CreateRatingRequest) (*pb.CreateRatingResponse, error) {
	_, rep, err := s.createRating.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateRatingResponse), nil
}

// decodeCreateRatingRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateRatingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateRatingRequest)
	return endpoints.CreateRatingRequest{Rating: models.RatingToORM(req.Rating)}, nil
}

// encodeCreateRatingResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateRatingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateRatingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateRatingResponse{}, nil
	}
	return nil, err
}

// DeleteRating deletes a Joblisting Rating
func (s *grpcServer) DeleteRating(ctx context.Context, req *pb.DeleteRatingRequest) (*pb.DeleteRatingResponse, error) {
	_, rep, err := s.deleteRating.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteRatingResponse), nil
}

// decodeDeleteRatingRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteRatingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteRatingRequest)
	return endpoints.DeleteRatingRequest{ID: req.Id}, nil
}

// encodeDeleteRatingResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteRatingResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteRatingResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteRatingResponse{}, nil
	}
	return nil, err
}

func getError(err error) error {
	switch err {
	case nil:
		return nil
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
