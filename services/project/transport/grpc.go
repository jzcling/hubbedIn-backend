package transport

import (
	"context"
	"in-backend/services/project/endpoints"
	"in-backend/services/project/models"
	"in-backend/services/project/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for Project Service.
type grpcServer struct {
	createProject  kitgrpc.Handler
	getAllProjects kitgrpc.Handler
	getProjectByID kitgrpc.Handler
	updateProject  kitgrpc.Handler
	deleteProject  kitgrpc.Handler

	scanProject kitgrpc.Handler

	createCandidateProject kitgrpc.Handler
	deleteCandidateProject kitgrpc.Handler

	createRating kitgrpc.Handler
	deleteRating kitgrpc.Handler

	logger log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.ProjectServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createProject: kitgrpc.NewServer(
			endpoints.CreateProject,
			decodeCreateProjectRequest,
			encodeCreateProjectResponse,
			options...,
		),
		getAllProjects: kitgrpc.NewServer(
			endpoints.GetAllProjects,
			decodeGetAllProjectsRequest,
			encodeGetAllProjectsResponse,
			options...,
		),
		getProjectByID: kitgrpc.NewServer(
			endpoints.GetProjectByID,
			decodeGetProjectByIDRequest,
			encodeGetProjectByIDResponse,
			options...,
		),
		updateProject: kitgrpc.NewServer(
			endpoints.UpdateProject,
			decodeUpdateProjectRequest,
			encodeUpdateProjectResponse,
			options...,
		),
		deleteProject: kitgrpc.NewServer(
			endpoints.DeleteProject,
			decodeDeleteProjectRequest,
			encodeDeleteProjectResponse,
			options...,
		),

		scanProject: kitgrpc.NewServer(
			endpoints.ScanProject,
			decodeScanProjectRequest,
			encodeScanProjectResponse,
			options...,
		),

		createCandidateProject: kitgrpc.NewServer(
			endpoints.CreateCandidateProject,
			decodeCreateCandidateProjectRequest,
			encodeCreateCandidateProjectResponse,
			options...,
		),
		deleteCandidateProject: kitgrpc.NewServer(
			endpoints.DeleteCandidateProject,
			decodeDeleteCandidateProjectRequest,
			encodeDeleteCandidateProjectResponse,
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

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (s *grpcServer) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.Project, error) {
	_, rep, err := s.createProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Project), nil
}

// decodeCreateProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateProjectRequest)
	return endpoints.CreateProjectRequest{Project: models.ProjectToORM(req.Project), CandidateID: req.CandidateId}, nil
}

// encodeCreateProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Project.ToProto(), nil
	}
	return nil, err
}

// GetAllProjects returns all Projects
func (s *grpcServer) GetAllProjects(ctx context.Context, req *pb.GetAllProjectsRequest) (*pb.GetAllProjectsResponse, error) {
	_, rep, err := s.getAllProjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllProjectsResponse), nil
}

// decodeGetAllProjectsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllProjectsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllProjectsRequest)
	decoded := endpoints.GetAllProjectsRequest{
		ID:          req.Id,
		CandidateID: req.CandidateId,
		Name:        req.Name,
		RepoURL:     req.RepoUrl,
	}
	return decoded, nil
}

// encodeGetAllProjectsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllProjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllProjectsResponse)
	err := getError(res.Err)
	if err == nil {
		var projects []*pb.Project
		for _, project := range res.Projects {
			projects = append(projects, project.ToProto())
		}
		return &pb.GetAllProjectsResponse{Projects: projects}, nil
	}
	return nil, err
}

// GetProjectByID returns a Project by ID
func (s *grpcServer) GetProjectByID(ctx context.Context, req *pb.GetProjectByIDRequest) (*pb.Project, error) {
	_, rep, err := s.getProjectByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Project), nil
}

// decodeGetProjectByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetProjectByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetProjectByIDRequest)
	return endpoints.GetProjectByIDRequest{ID: req.Id}, nil
}

// encodeGetProjectByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetProjectByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetProjectByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Project.ToProto(), nil
	}
	return nil, err
}

// UpdateProject updates a Project
func (s *grpcServer) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.Project, error) {
	_, rep, err := s.updateProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Project), nil
}

// decodeUpdateProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateProjectRequest)
	return endpoints.UpdateProjectRequest{ID: req.Id, Project: models.ProjectToORM(req.Project)}, nil
}

// encodeUpdateProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Project.ToProto(), nil
	}
	return nil, err
}

// DeleteProject deletes a Project by ID
func (s *grpcServer) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	_, rep, err := s.deleteProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteProjectResponse), nil
}

// decodeDeleteProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteProjectRequest)
	return endpoints.DeleteProjectRequest{ID: req.Id}, nil
}

// encodeDeleteProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteProjectResponse{}, nil
	}
	return nil, err
}

// ScanProject scans a Project using sonarqube
func (s *grpcServer) ScanProject(ctx context.Context, req *pb.ScanProjectRequest) (*pb.ScanProjectResponse, error) {
	_, rep, err := s.scanProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ScanProjectResponse), nil
}

// decodeScanProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeScanProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ScanProjectRequest)
	return endpoints.ScanProjectRequest{ID: req.Id}, nil
}

// encodeScanProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeScanProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.ScanProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.ScanProjectResponse{}, nil
	}
	return nil, err
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (s *grpcServer) CreateCandidateProject(ctx context.Context, req *pb.CreateCandidateProjectRequest) (*pb.CreateCandidateProjectResponse, error) {
	_, rep, err := s.createCandidateProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCandidateProjectResponse), nil
}

// decodeCreateCandidateProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCandidateProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCandidateProjectRequest)
	return endpoints.CreateCandidateProjectRequest{CandidateProject: models.CandidateProjectToORM(req.CandidateProject)}, nil
}

// encodeCreateCandidateProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCandidateProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateCandidateProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateCandidateProjectResponse{}, nil
	}
	return nil, err
}

// DeleteCandidateProject deletes a CandidateProject by Candidate ID and Project ID
func (s *grpcServer) DeleteCandidateProject(ctx context.Context, req *pb.DeleteCandidateProjectRequest) (*pb.DeleteCandidateProjectResponse, error) {
	_, rep, err := s.deleteCandidateProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteCandidateProjectResponse), nil
}

// decodeDeleteCandidateProjectRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteCandidateProjectRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteCandidateProjectRequest)
	return endpoints.DeleteCandidateProjectRequest{ID: req.Id}, nil
}

// encodeDeleteCandidateProjectResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteCandidateProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteCandidateProjectResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteCandidateProjectResponse{}, nil
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

// DeleteRating deletes a Project Rating
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
