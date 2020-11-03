package transport

import (
	"context"
	"in-backend/services/profile/endpoints"
	"in-backend/services/profile/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for Profile Service.
type grpcServer struct {
	createCandidate  kitgrpc.Handler
	getAllCandidates kitgrpc.Handler
	getCandidateByID kitgrpc.Handler
	updateCandidate  kitgrpc.Handler
	deleteCandidate  kitgrpc.Handler
	logger           log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.ProfileServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createCandidate: kitgrpc.NewServer(
			endpoints.CreateCandidate,
			decodeCreateCandidateRequest,
			encodeCreateCandidateResponse,
			options...,
		),
		getAllCandidates: kitgrpc.NewServer(
			endpoints.GetAllCandidates,
			decodeGetAllCandidatesRequest,
			encodeGetAllCandidatesResponse,
			options...,
		),
		getCandidateByID: kitgrpc.NewServer(
			endpoints.GetCandidateByID,
			decodeGetCandidateByIDRequest,
			encodeGetCandidateByIDResponse,
			options...,
		),
		updateCandidate: kitgrpc.NewServer(
			endpoints.UpdateCandidate,
			decodeUpdateCandidateRequest,
			encodeUpdateCandidateResponse,
			options...,
		),
		deleteCandidate: kitgrpc.NewServer(
			endpoints.DeleteCandidate,
			decodeDeleteCandidateRequest,
			encodeDeleteCandidateResponse,
			options...,
		),
		logger: logger,
	}
}

// CreateCandidate creates a new Candidate
func (s *grpcServer) CreateCandidate(ctx context.Context, req *pb.CreateCandidateRequest) (*pb.Candidate, error) {
	_, rep, err := s.createCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeCreateCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCandidateRequest)
	return endpoints.CreateCandidateRequest{Candidate: req.Candidate.ToORM()}, nil
}

// encodeCreateCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.Candidate{}, nil
	}
	return nil, err
}

// GetAllCandidates returns all Candidates
func (s *grpcServer) GetAllCandidates(ctx context.Context, req *pb.GetAllCandidatesRequest) (*pb.GetAllCandidatesResponse, error) {
	_, rep, err := s.getAllCandidates.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllCandidatesResponse), nil
}

// decodeGetAllCandidatesRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllCandidatesRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetAllCandidatesRequest)
	return endpoints.GetAllCandidatesRequest{}, nil
}

// encodeGetAllCandidatesResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllCandidatesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllCandidatesResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.GetAllCandidatesResponse{}, nil
	}
	return nil, err
}

// GetCandidateByID returns a Candidate by ID
func (s *grpcServer) GetCandidateByID(ctx context.Context, req *pb.GetCandidateByIDRequest) (*pb.Candidate, error) {
	_, rep, err := s.getCandidateByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeGetCandidateByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetCandidateByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCandidateByIDRequest)
	return endpoints.GetCandidateByIDRequest{ID: req.Id}, nil
}

// encodeGetCandidateByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetCandidateByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetCandidateByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.Candidate{}, nil
	}
	return nil, err
}

// UpdateCandidate updates a Candidate
func (s *grpcServer) UpdateCandidate(ctx context.Context, req *pb.UpdateCandidateRequest) (*pb.Candidate, error) {
	_, rep, err := s.updateCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Candidate), nil
}

// decodeUpdateCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateCandidateRequest)
	return endpoints.UpdateCandidateRequest{Candidate: req.Candidate.ToORM()}, nil
}

// encodeUpdateCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.Candidate{}, nil
	}
	return nil, err
}

// DeleteCandidate deletes a Candidate by ID
func (s *grpcServer) DeleteCandidate(ctx context.Context, req *pb.DeleteCandidateRequest) (*pb.DeleteCandidateResponse, error) {
	_, rep, err := s.deleteCandidate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteCandidateResponse), nil
}

// decodeDeleteCandidateRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteCandidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteCandidateRequest)
	return endpoints.DeleteCandidateRequest{ID: req.Id}, nil
}

// encodeDeleteCandidateResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteCandidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteCandidateResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteCandidateResponse{}, nil
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
