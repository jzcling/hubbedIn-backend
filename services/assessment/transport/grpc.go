package transport

import (
	"context"
	"in-backend/services/assessment/endpoints"
	"in-backend/services/assessment/models"
	"in-backend/services/assessment/pb"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpc transport service for assessment Service.
type grpcServer struct {
	createAssessment  kitgrpc.Handler
	getAllAssessments kitgrpc.Handler
	getAssessmentByID kitgrpc.Handler
	updateAssessment  kitgrpc.Handler
	deleteAssessment  kitgrpc.Handler

	createAssessmentAttempt       kitgrpc.Handler
	getAssessmentAttemptByID      kitgrpc.Handler
	localGetAssessmentAttemptByID kitgrpc.Handler
	updateAssessmentAttempt       kitgrpc.Handler
	localUpdateAssessmentAttempt  kitgrpc.Handler
	deleteAssessmentAttempt       kitgrpc.Handler

	createQuestion  kitgrpc.Handler
	getAllQuestions kitgrpc.Handler
	getQuestionByID kitgrpc.Handler
	updateQuestion  kitgrpc.Handler
	deleteQuestion  kitgrpc.Handler

	createTag kitgrpc.Handler
	deleteTag kitgrpc.Handler

	updateAttemptQuestion kitgrpc.Handler

	logger log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints endpoints.Endpoints,
	options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.AssessmentServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createAssessment: kitgrpc.NewServer(
			endpoints.CreateAssessment,
			decodeCreateAssessmentRequest,
			encodeCreateAssessmentResponse,
			options...,
		),
		getAllAssessments: kitgrpc.NewServer(
			endpoints.GetAllAssessments,
			decodeGetAllAssessmentsRequest,
			encodeGetAllAssessmentsResponse,
			options...,
		),
		getAssessmentByID: kitgrpc.NewServer(
			endpoints.GetAssessmentByID,
			decodeGetAssessmentByIDRequest,
			encodeGetAssessmentByIDResponse,
			options...,
		),
		updateAssessment: kitgrpc.NewServer(
			endpoints.UpdateAssessment,
			decodeUpdateAssessmentRequest,
			encodeUpdateAssessmentResponse,
			options...,
		),
		deleteAssessment: kitgrpc.NewServer(
			endpoints.DeleteAssessment,
			decodeDeleteAssessmentRequest,
			encodeDeleteAssessmentResponse,
			options...,
		),

		createAssessmentAttempt: kitgrpc.NewServer(
			endpoints.CreateAssessmentAttempt,
			decodeCreateAssessmentAttemptRequest,
			encodeCreateAssessmentAttemptResponse,
			options...,
		),
		getAssessmentAttemptByID: kitgrpc.NewServer(
			endpoints.GetAssessmentAttemptByID,
			decodeGetAssessmentAttemptByIDRequest,
			encodeGetAssessmentAttemptByIDResponse,
			options...,
		),
		localGetAssessmentAttemptByID: kitgrpc.NewServer(
			endpoints.LocalGetAssessmentAttemptByID,
			decodeGetAssessmentAttemptByIDRequest,
			encodeGetAssessmentAttemptByIDResponse,
			options...,
		),
		updateAssessmentAttempt: kitgrpc.NewServer(
			endpoints.UpdateAssessmentAttempt,
			decodeUpdateAssessmentAttemptRequest,
			encodeUpdateAssessmentAttemptResponse,
			options...,
		),
		localUpdateAssessmentAttempt: kitgrpc.NewServer(
			endpoints.LocalUpdateAssessmentAttempt,
			decodeUpdateAssessmentAttemptRequest,
			encodeUpdateAssessmentAttemptResponse,
			options...,
		),
		deleteAssessmentAttempt: kitgrpc.NewServer(
			endpoints.DeleteAssessmentAttempt,
			decodeDeleteAssessmentAttemptRequest,
			encodeDeleteAssessmentAttemptResponse,
			options...,
		),

		createQuestion: kitgrpc.NewServer(
			endpoints.CreateQuestion,
			decodeCreateQuestionRequest,
			encodeCreateQuestionResponse,
			options...,
		),
		getAllQuestions: kitgrpc.NewServer(
			endpoints.GetAllQuestions,
			decodeGetAllQuestionsRequest,
			encodeGetAllQuestionsResponse,
			options...,
		),
		getQuestionByID: kitgrpc.NewServer(
			endpoints.GetQuestionByID,
			decodeGetQuestionByIDRequest,
			encodeGetQuestionByIDResponse,
			options...,
		),
		updateQuestion: kitgrpc.NewServer(
			endpoints.UpdateQuestion,
			decodeUpdateQuestionRequest,
			encodeUpdateQuestionResponse,
			options...,
		),
		deleteQuestion: kitgrpc.NewServer(
			endpoints.DeleteQuestion,
			decodeDeleteQuestionRequest,
			encodeDeleteQuestionResponse,
			options...,
		),

		createTag: kitgrpc.NewServer(
			endpoints.CreateTag,
			decodeCreateTagRequest,
			encodeCreateTagResponse,
			options...,
		),
		deleteTag: kitgrpc.NewServer(
			endpoints.DeleteTag,
			decodeDeleteTagRequest,
			encodeDeleteTagResponse,
			options...,
		),

		updateAttemptQuestion: kitgrpc.NewServer(
			endpoints.UpdateAttemptQuestion,
			decodeUpdateAttemptQuestionRequest,
			encodeUpdateAttemptQuestionResponse,
			options...,
		),

		logger: logger,
	}
}

/* --------------- Assessment --------------- */

// CreateAssessment creates a new Assessment
func (s *grpcServer) CreateAssessment(ctx context.Context, req *pb.CreateAssessmentRequest) (*pb.Assessment, error) {
	_, rep, err := s.createAssessment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Assessment), nil
}

// decodeCreateAssessmentRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateAssessmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateAssessmentRequest)
	return endpoints.CreateAssessmentRequest{Assessment: models.AssessmentToORM(req.Assessment)}, nil
}

// encodeCreateAssessmentResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateAssessmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateAssessmentResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Assessment.ToProto(), nil
	}
	return nil, err
}

// GetAllAssessments returns all Assessments
func (s *grpcServer) GetAllAssessments(ctx context.Context, req *pb.GetAllAssessmentsRequest) (*pb.GetAllAssessmentsResponse, error) {
	_, rep, err := s.getAllAssessments.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllAssessmentsResponse), nil
}

// decodeGetAllAssessmentsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllAssessmentsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllAssessmentsRequest)
	decoded := endpoints.GetAllAssessmentsRequest{
		ID:          req.Id,
		Name:        req.Name,
		Difficulty:  req.Difficulty,
		Type:        req.Type,
		CandidateID: req.CandidateId,
		Status:      req.Status,
		MinScore:    req.MinScore,
	}
	return decoded, nil
}

// encodeGetAllAssessmentsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllAssessmentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllAssessmentsResponse)
	err := getError(res.Err)
	if err == nil {
		var candidates []*pb.Assessment
		for _, candidate := range res.Assessments {
			candidates = append(candidates, candidate.ToProto())
		}
		return &pb.GetAllAssessmentsResponse{Assessments: candidates}, nil
	}
	return nil, err
}

// GetAssessmentByID returns a Assessment by ID
func (s *grpcServer) GetAssessmentByID(ctx context.Context, req *pb.GetAssessmentByIDRequest) (*pb.Assessment, error) {
	_, rep, err := s.getAssessmentByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Assessment), nil
}

// decodeGetAssessmentByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAssessmentByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAssessmentByIDRequest)
	return endpoints.GetAssessmentByIDRequest{ID: req.Id}, nil
}

// encodeGetAssessmentByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAssessmentByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAssessmentByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Assessment.ToProto(), nil
	}
	return nil, err
}

// UpdateAssessment updates a Assessment
func (s *grpcServer) UpdateAssessment(ctx context.Context, req *pb.UpdateAssessmentRequest) (*pb.Assessment, error) {
	_, rep, err := s.updateAssessment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Assessment), nil
}

// decodeUpdateAssessmentRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateAssessmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAssessmentRequest)
	return endpoints.UpdateAssessmentRequest{ID: req.Id, Assessment: models.AssessmentToORM(req.Assessment)}, nil
}

// encodeUpdateAssessmentResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateAssessmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateAssessmentResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Assessment.ToProto(), nil
	}
	return nil, err
}

// DeleteAssessment deletes a Assessment by ID
func (s *grpcServer) DeleteAssessment(ctx context.Context, req *pb.DeleteAssessmentRequest) (*pb.DeleteAssessmentResponse, error) {
	_, rep, err := s.deleteAssessment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteAssessmentResponse), nil
}

// decodeDeleteAssessmentRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteAssessmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteAssessmentRequest)
	return endpoints.DeleteAssessmentRequest{ID: req.Id}, nil
}

// encodeDeleteAssessmentResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteAssessmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteAssessmentResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteAssessmentResponse{}, nil
	}
	return nil, err
}

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (s *grpcServer) CreateAssessmentAttempt(ctx context.Context, req *pb.CreateAssessmentAttemptRequest) (*pb.AssessmentAttempt, error) {
	_, rep, err := s.createAssessmentAttempt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssessmentAttempt), nil
}

// decodeCreateAssessmentAttemptRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateAssessmentAttemptRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateAssessmentAttemptRequest)
	return endpoints.CreateAssessmentAttemptRequest{AssessmentAttempt: models.AssessmentAttemptToORM(req.AssessmentAttempt)}, nil
}

// encodeCreateAssessmentAttemptResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateAssessmentAttemptResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateAssessmentAttemptResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AssessmentAttempt.ToProto(), nil
	}
	return nil, err
}

// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (s *grpcServer) GetAssessmentAttemptByID(ctx context.Context, req *pb.GetAssessmentAttemptByIDRequest) (*pb.AssessmentAttempt, error) {
	_, rep, err := s.getAssessmentAttemptByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssessmentAttempt), nil
}

// LocalGetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (s *grpcServer) LocalGetAssessmentAttemptByID(ctx context.Context, req *pb.GetAssessmentAttemptByIDRequest) (*pb.AssessmentAttempt, error) {
	_, rep, err := s.localGetAssessmentAttemptByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssessmentAttempt), nil
}

// decodeGetAssessmentAttemptByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAssessmentAttemptByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAssessmentAttemptByIDRequest)
	return endpoints.GetAssessmentAttemptByIDRequest{ID: req.Id}, nil
}

// encodeGetAssessmentAttemptByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAssessmentAttemptByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAssessmentAttemptByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AssessmentAttempt.ToProto(), nil
	}
	return nil, err
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (s *grpcServer) UpdateAssessmentAttempt(ctx context.Context, req *pb.UpdateAssessmentAttemptRequest) (*pb.AssessmentAttempt, error) {
	_, rep, err := s.updateAssessmentAttempt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssessmentAttempt), nil
}

// LocalUpdateAssessmentAttempt updates a AssessmentAttempt
func (s *grpcServer) LocalUpdateAssessmentAttempt(ctx context.Context, req *pb.UpdateAssessmentAttemptRequest) (*pb.AssessmentAttempt, error) {
	_, rep, err := s.localUpdateAssessmentAttempt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AssessmentAttempt), nil
}

// decodeUpdateAssessmentAttemptRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateAssessmentAttemptRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAssessmentAttemptRequest)
	return endpoints.UpdateAssessmentAttemptRequest{ID: req.Id, AssessmentAttempt: models.AssessmentAttemptToORM(req.AssessmentAttempt)}, nil
}

// encodeUpdateAssessmentAttemptResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateAssessmentAttemptResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateAssessmentAttemptResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AssessmentAttempt.ToProto(), nil
	}
	return nil, err
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (s *grpcServer) DeleteAssessmentAttempt(ctx context.Context, req *pb.DeleteAssessmentAttemptRequest) (*pb.DeleteAssessmentAttemptResponse, error) {
	_, rep, err := s.deleteAssessmentAttempt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteAssessmentAttemptResponse), nil
}

// decodeDeleteAssessmentAttemptRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteAssessmentAttemptRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteAssessmentAttemptRequest)
	return endpoints.DeleteAssessmentAttemptRequest{ID: req.Id}, nil
}

// encodeDeleteAssessmentAttemptResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteAssessmentAttemptResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteAssessmentAttemptResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteAssessmentAttemptResponse{}, nil
	}
	return nil, err
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (s *grpcServer) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.Question, error) {
	_, rep, err := s.createQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Question), nil
}

// decodeCreateQuestionRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateQuestionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateQuestionRequest)
	return endpoints.CreateQuestionRequest{Question: models.QuestionToORM(req.Question)}, nil
}

// encodeCreateQuestionResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateQuestionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Question.ToProto(), nil
	}
	return nil, err
}

// GetAllQuestions returns all Questions
func (s *grpcServer) GetAllQuestions(ctx context.Context, req *pb.GetAllQuestionsRequest) (*pb.GetAllQuestionsResponse, error) {
	_, rep, err := s.getAllQuestions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllQuestionsResponse), nil
}

// decodeGetAllQuestionsRequest decodes the incoming grpc payload to our go kit payload
func decodeGetAllQuestionsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetAllQuestionsRequest)
	decoded := endpoints.GetAllQuestionsRequest{
		ID:   req.Id,
		Tags: req.Tags,
	}
	return decoded, nil
}

// encodeGetAllQuestionsResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetAllQuestionsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetAllQuestionsResponse)
	err := getError(res.Err)
	if err == nil {
		var candidates []*pb.Question
		for _, candidate := range res.Questions {
			candidates = append(candidates, candidate.ToProto())
		}
		return &pb.GetAllQuestionsResponse{Questions: candidates}, nil
	}
	return nil, err
}

// GetQuestionByID returns a Question by ID
func (s *grpcServer) GetQuestionByID(ctx context.Context, req *pb.GetQuestionByIDRequest) (*pb.Question, error) {
	_, rep, err := s.getQuestionByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Question), nil
}

// decodeGetQuestionByIDRequest decodes the incoming grpc payload to our go kit payload
func decodeGetQuestionByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetQuestionByIDRequest)
	return endpoints.GetQuestionByIDRequest{ID: req.Id}, nil
}

// encodeGetQuestionByIDResponse encodes the outgoing go kit payload to the grpc payload
func encodeGetQuestionByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetQuestionByIDResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Question.ToProto(), nil
	}
	return nil, err
}

// UpdateQuestion updates a Question
func (s *grpcServer) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Question, error) {
	_, rep, err := s.updateQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Question), nil
}

// decodeUpdateQuestionRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateQuestionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateQuestionRequest)
	return endpoints.UpdateQuestionRequest{ID: req.Id, Question: models.QuestionToORM(req.Question)}, nil
}

// encodeUpdateQuestionResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateQuestionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Question.ToProto(), nil
	}
	return nil, err
}

// DeleteQuestion deletes a Question by ID
func (s *grpcServer) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.DeleteQuestionResponse, error) {
	_, rep, err := s.deleteQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteQuestionResponse), nil
}

// decodeDeleteQuestionRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteQuestionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteQuestionRequest)
	return endpoints.DeleteQuestionRequest{ID: req.Id}, nil
}

// encodeDeleteQuestionResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteQuestionResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteQuestionResponse{}, nil
	}
	return nil, err
}

/* --------------- Tag --------------- */

// CreateTag creates a new Tag
func (s *grpcServer) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.Tag, error) {
	_, rep, err := s.createTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Tag), nil
}

// decodeCreateTagRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateTagRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateTagRequest)
	return endpoints.CreateTagRequest{Tag: models.TagToORM(req.Tag)}, nil
}

// encodeCreateTagResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateTagResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.CreateTagResponse)
	err := getError(res.Err)
	if err == nil {
		return res.Tag.ToProto(), nil
	}
	return nil, err
}

// DeleteTag deletes a Tag by ID
func (s *grpcServer) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagResponse, error) {
	_, rep, err := s.deleteTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTagResponse), nil
}

// decodeDeleteTagRequest decodes the incoming grpc payload to our go kit payload
func decodeDeleteTagRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteTagRequest)
	return endpoints.DeleteTagRequest{ID: req.Id}, nil
}

// encodeDeleteTagResponse encodes the outgoing go kit payload to the grpc payload
func encodeDeleteTagResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.DeleteTagResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.DeleteTagResponse{}, nil
	}
	return nil, err
}

/* --------------- Attempt Question --------------- */

// UpdateAttemptQuestion updates a AttemptQuestion
func (s *grpcServer) UpdateAttemptQuestion(ctx context.Context, req *pb.UpdateAttemptQuestionRequest) (*pb.AttemptQuestion, error) {
	_, rep, err := s.updateAttemptQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AttemptQuestion), nil
}

// decodeUpdateAttemptQuestionRequest decodes the incoming grpc payload to our go kit payload
func decodeUpdateAttemptQuestionRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAttemptQuestionRequest)
	return endpoints.UpdateAttemptQuestionRequest{ID: req.Id, AttemptQuestion: models.AttemptQuestionToORM(req.AttemptQuestion)}, nil
}

// encodeUpdateAttemptQuestionResponse encodes the outgoing go kit payload to the grpc payload
func encodeUpdateAttemptQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.UpdateAttemptQuestionResponse)
	err := getError(res.Err)
	if err == nil {
		return res.AttemptQuestion.ToProto(), nil
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
