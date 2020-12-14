package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
)

// Endpoints holds all Go kit endpoints for the assessment Service.
type Endpoints struct {
	CreateAssessment  endpoint.Endpoint
	GetAllAssessments endpoint.Endpoint
	GetAssessmentByID endpoint.Endpoint
	UpdateAssessment  endpoint.Endpoint
	DeleteAssessment  endpoint.Endpoint

	CreateAssessmentAttempt       endpoint.Endpoint
	GetAssessmentAttemptByID      endpoint.Endpoint
	LocalGetAssessmentAttemptByID endpoint.Endpoint
	UpdateAssessmentAttempt       endpoint.Endpoint
	LocalUpdateAssessmentAttempt  endpoint.Endpoint
	DeleteAssessmentAttempt       endpoint.Endpoint

	CreateQuestion  endpoint.Endpoint
	GetAllQuestions endpoint.Endpoint
	GetQuestionByID endpoint.Endpoint
	UpdateQuestion  endpoint.Endpoint
	DeleteQuestion  endpoint.Endpoint

	CreateTag endpoint.Endpoint
	DeleteTag endpoint.Endpoint

	UpdateAttemptQuestion endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the assessment service.
func MakeEndpoints(s interfaces.Service) Endpoints {
	return Endpoints{
		CreateAssessment:  makeCreateAssessmentEndpoint(s),
		GetAllAssessments: makeGetAllAssessmentsEndpoint(s),
		GetAssessmentByID: makeGetAssessmentByIDEndpoint(s),
		UpdateAssessment:  makeUpdateAssessmentEndpoint(s),
		DeleteAssessment:  makeDeleteAssessmentEndpoint(s),

		CreateAssessmentAttempt:       makeCreateAssessmentAttemptEndpoint(s),
		GetAssessmentAttemptByID:      makeGetAssessmentAttemptByIDEndpoint(s),
		LocalGetAssessmentAttemptByID: makeLocalGetAssessmentAttemptByIDEndpoint(s),
		UpdateAssessmentAttempt:       makeUpdateAssessmentAttemptEndpoint(s),
		LocalUpdateAssessmentAttempt:  makeLocalUpdateAssessmentAttemptEndpoint(s),
		DeleteAssessmentAttempt:       makeDeleteAssessmentAttemptEndpoint(s),

		CreateQuestion:  makeCreateQuestionEndpoint(s),
		GetAllQuestions: makeGetAllQuestionsEndpoint(s),
		GetQuestionByID: makeGetQuestionByIDEndpoint(s),
		UpdateQuestion:  makeUpdateQuestionEndpoint(s),
		DeleteQuestion:  makeDeleteQuestionEndpoint(s),

		CreateTag: makeCreateTagEndpoint(s),
		DeleteTag: makeDeleteTagEndpoint(s),

		UpdateAttemptQuestion: makeUpdateAttemptQuestionEndpoint(s),
	}
}

/* -------------- Assessment -------------- */

func makeCreateAssessmentEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAssessmentRequest)
		m, err := s.CreateAssessment(ctx, req.Assessment)
		return CreateAssessmentResponse{Assessment: m, Err: err}, nil
	}
}

// CreateAssessmentRequest declares the inputs required for creating a assessment
type CreateAssessmentRequest struct {
	Assessment *models.Assessment
}

// CreateAssessmentResponse declares the outputs after attempting to create a assessment
type CreateAssessmentResponse struct {
	Assessment *models.Assessment
	Err        error
}

func makeGetAllAssessmentsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllAssessmentsRequest)
		f := models.AssessmentFilters(req)
		m, err := s.GetAllAssessments(ctx, f, nil, nil)
		return GetAllAssessmentsResponse{Assessments: m, Err: err}, nil
	}
}

// GetAllAssessmentsRequest declares the inputs required for getting all assessments
type GetAllAssessmentsRequest struct {
	ID         []uint64
	Name       string
	Difficulty []string
	Type       []string

	// relation filters
	CandidateID uint64
	Status      []string
	MinScore    int64
}

// GetAllAssessmentsResponse declares the outputs after attempting to get all assessments
type GetAllAssessmentsResponse struct {
	Assessments []*models.Assessment
	Err         error
}

func makeGetAssessmentByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAssessmentByIDRequest)
		m, err := s.GetAssessmentByID(ctx, req.ID, nil, nil)
		return GetAssessmentByIDResponse{Assessment: m, Err: err}, nil
	}
}

// GetAssessmentByIDRequest declares the inputs required for getting a single assessment by ID
type GetAssessmentByIDRequest struct {
	ID uint64
}

// GetAssessmentByIDResponse declares the outputs after attempting to get a single assessment by ID
type GetAssessmentByIDResponse struct {
	Assessment *models.Assessment
	Err        error
}

func makeUpdateAssessmentEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAssessmentRequest)
		m, err := s.UpdateAssessment(ctx, req.Assessment)
		return UpdateAssessmentResponse{Assessment: m, Err: err}, nil
	}
}

// UpdateAssessmentRequest declares the inputs required for updating a assessment
type UpdateAssessmentRequest struct {
	ID         uint64
	Assessment *models.Assessment
}

// UpdateAssessmentResponse declares the outputs after attempting to update a assessment
type UpdateAssessmentResponse struct {
	Assessment *models.Assessment
	Err        error
}

func makeDeleteAssessmentEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAssessmentRequest)
		err := s.DeleteAssessment(ctx, req.ID)
		return DeleteAssessmentResponse{Err: err}, nil
	}
}

// DeleteAssessmentRequest declares the inputs required for deleting a assessment
type DeleteAssessmentRequest struct {
	ID uint64
}

// DeleteAssessmentResponse declares the outputs after attempting to delete a assessment
type DeleteAssessmentResponse struct {
	Err error
}

/* -------------- Assessment Attempt -------------- */

func makeCreateAssessmentAttemptEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAssessmentAttemptRequest)
		m, err := s.CreateAssessmentAttempt(ctx, req.AssessmentAttempt)
		return CreateAssessmentAttemptResponse{AssessmentAttempt: m, Err: err}, nil
	}
}

// CreateAssessmentAttemptRequest declares the inputs required for creating a assessment attempt
type CreateAssessmentAttemptRequest struct {
	AssessmentAttempt *models.AssessmentAttempt
}

// CreateAssessmentAttemptResponse declares the outputs after attempting to create a assessment attempt
type CreateAssessmentAttemptResponse struct {
	AssessmentAttempt *models.AssessmentAttempt
	Err               error
}

func makeGetAssessmentAttemptByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAssessmentAttemptByIDRequest)
		m, err := s.GetAssessmentAttemptByID(ctx, req.ID)
		return GetAssessmentAttemptByIDResponse{AssessmentAttempt: m, Err: err}, nil
	}
}

func makeLocalGetAssessmentAttemptByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAssessmentAttemptByIDRequest)
		m, err := s.LocalGetAssessmentAttemptByID(ctx, req.ID)
		return GetAssessmentAttemptByIDResponse{AssessmentAttempt: m, Err: err}, nil
	}
}

// GetAssessmentAttemptByIDRequest declares the inputs required for getting a single assessment attempt by ID
type GetAssessmentAttemptByIDRequest struct {
	ID uint64
}

// GetAssessmentAttemptByIDResponse declares the outputs after attempting to get a single assessment attempt by ID
type GetAssessmentAttemptByIDResponse struct {
	AssessmentAttempt *models.AssessmentAttempt
	Err               error
}

func makeUpdateAssessmentAttemptEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAssessmentAttemptRequest)
		m, err := s.UpdateAssessmentAttempt(ctx, req.AssessmentAttempt)
		return UpdateAssessmentAttemptResponse{AssessmentAttempt: m, Err: err}, nil
	}
}

func makeLocalUpdateAssessmentAttemptEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAssessmentAttemptRequest)
		m, err := s.LocalUpdateAssessmentAttempt(ctx, req.AssessmentAttempt)
		return UpdateAssessmentAttemptResponse{AssessmentAttempt: m, Err: err}, nil
	}
}

// UpdateAssessmentAttemptRequest declares the inputs required for updating a assessment attempt
type UpdateAssessmentAttemptRequest struct {
	ID                uint64
	AssessmentAttempt *models.AssessmentAttempt
}

// UpdateAssessmentAttemptResponse declares the outputs after attempting to update a assessment attempt
type UpdateAssessmentAttemptResponse struct {
	AssessmentAttempt *models.AssessmentAttempt
	Err               error
}

func makeDeleteAssessmentAttemptEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAssessmentAttemptRequest)
		err := s.DeleteAssessmentAttempt(ctx, req.ID)
		return DeleteAssessmentAttemptResponse{Err: err}, nil
	}
}

// DeleteAssessmentAttemptRequest declares the inputs required for deleting a assessment attempt
type DeleteAssessmentAttemptRequest struct {
	ID uint64
}

// DeleteAssessmentAttemptResponse declares the outputs after attempting to delete a assessment attempt
type DeleteAssessmentAttemptResponse struct {
	Err error
}

/* -------------- Question -------------- */

func makeCreateQuestionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateQuestionRequest)
		m, err := s.CreateQuestion(ctx, req.Question)
		return CreateQuestionResponse{Question: m, Err: err}, nil
	}
}

// CreateQuestionRequest declares the inputs required for creating a question
type CreateQuestionRequest struct {
	Question *models.Question
}

// CreateQuestionResponse declares the outputs after attempting to create a question
type CreateQuestionResponse struct {
	Question *models.Question
	Err      error
}

func makeGetAllQuestionsEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllQuestionsRequest)
		f := models.QuestionFilters(req)
		m, err := s.GetAllQuestions(ctx, f)
		return GetAllQuestionsResponse{Questions: m, Err: err}, nil
	}
}

// GetAllQuestionsRequest declares the inputs required for getting all questions
type GetAllQuestionsRequest struct {
	ID   []uint64
	Tags []string
}

// GetAllQuestionsResponse declares the outputs after attempting to get all questions
type GetAllQuestionsResponse struct {
	Questions []*models.Question
	Err       error
}

func makeGetQuestionByIDEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetQuestionByIDRequest)
		m, err := s.GetQuestionByID(ctx, req.ID)
		return GetQuestionByIDResponse{Question: m, Err: err}, nil
	}
}

// GetQuestionByIDRequest declares the inputs required for getting a single question by ID
type GetQuestionByIDRequest struct {
	ID uint64
}

// GetQuestionByIDResponse declares the outputs after attempting to get a single question by ID
type GetQuestionByIDResponse struct {
	Question *models.Question
	Err      error
}

func makeUpdateQuestionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)
		m, err := s.UpdateQuestion(ctx, req.Question)
		return UpdateQuestionResponse{Question: m, Err: err}, nil
	}
}

// UpdateQuestionRequest declares the inputs required for updating a question
type UpdateQuestionRequest struct {
	ID       uint64
	Question *models.Question
}

// UpdateQuestionResponse declares the outputs after attempting to update a question
type UpdateQuestionResponse struct {
	Question *models.Question
	Err      error
}

func makeDeleteQuestionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteQuestionRequest)
		err := s.DeleteQuestion(ctx, req.ID)
		return DeleteQuestionResponse{Err: err}, nil
	}
}

// DeleteQuestionRequest declares the inputs required for deleting a question
type DeleteQuestionRequest struct {
	ID uint64
}

// DeleteQuestionResponse declares the outputs after attempting to delete a question
type DeleteQuestionResponse struct {
	Err error
}

/* -------------- Tag -------------- */

func makeCreateTagEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTagRequest)
		m, err := s.CreateTag(ctx, req.Tag)
		return CreateTagResponse{Tag: m, Err: err}, nil
	}
}

// CreateTagRequest declares the inputs required for creating a tag
type CreateTagRequest struct {
	Tag *models.Tag
}

// CreateTagResponse declares the outputs after attempting to create a tag
type CreateTagResponse struct {
	Tag *models.Tag
	Err error
}

func makeDeleteTagEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTagRequest)
		err := s.DeleteTag(ctx, req.ID)
		return DeleteTagResponse{Err: err}, nil
	}
}

// DeleteTagRequest declares the inputs required for deleting a tag
type DeleteTagRequest struct {
	ID uint64
}

// DeleteTagResponse declares the outputs after attempting to delete a tag
type DeleteTagResponse struct {
	Err error
}

/* -------------- Attempt Question -------------- */

func makeUpdateAttemptQuestionEndpoint(s interfaces.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAttemptQuestionRequest)
		m, err := s.UpdateAttemptQuestion(ctx, req.AttemptQuestion)
		return UpdateAttemptQuestionResponse{AttemptQuestion: m, Err: err}, nil
	}
}

// UpdateAttemptQuestionRequest declares the inputs required for updating a AttemptQuestion
type UpdateAttemptQuestionRequest struct {
	ID              uint64
	AttemptQuestion *models.AttemptQuestion
}

// UpdateAttemptQuestionResponse declares the outputs after attempting to update a attempt question
type UpdateAttemptQuestionResponse struct {
	AttemptQuestion *models.AttemptQuestion
	Err             error
}
