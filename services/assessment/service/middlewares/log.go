package middlewares

import (
	"context"
	"fmt"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"time"

	"github.com/go-kit/kit/log"
)

type logMiddleware struct {
	logger log.Logger
	next   interfaces.Service
}

// NewLogMiddleware creates and returns a new Log Middleware that implements the assessment Service interface
func NewLogMiddleware(logger log.Logger, svc interfaces.Service) interfaces.Service {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}

func (mw logMiddleware) log(method string, begin time.Time, input, output interface{}, err *error) {
	mw.logger.Log(
		"method", method,
		"input", fmt.Sprintf("%v", input),
		"output", fmt.Sprintf("%v", output),
		"err", err,
		"took", time.Since(begin),
	)
}

/* --------------- Assessment --------------- */

// CreateAssessment creates a new Assessment
func (mw logMiddleware) CreateAssessment(ctx context.Context, input *models.Assessment) (output *models.Assessment, err error) {
	defer mw.log("CreateAssessment", time.Now(), input, &output, &err)
	output, err = mw.next.CreateAssessment(ctx, input)
	return
}

// GetAllAssessments returns all Assessments
func (mw logMiddleware) GetAllAssessments(ctx context.Context, input models.AssessmentFilters, admin *bool) (output []*models.Assessment, err error) {
	defer mw.log("GetAllAssessments", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllAssessments(ctx, input, admin)
	return
}

// GetAssessmentByID returns a Assessment by ID
func (mw logMiddleware) GetAssessmentByID(ctx context.Context, input uint64, admin *bool) (output *models.Assessment, err error) {
	defer mw.log("GetAssessmentByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetAssessmentByID(ctx, input, admin)
	return
}

// UpdateAssessment updates a Assessment
func (mw logMiddleware) UpdateAssessment(ctx context.Context, input *models.Assessment) (output *models.Assessment, err error) {
	defer mw.log("UpdateAssessment", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateAssessment(ctx, input)
	return
}

// DeleteAssessment deletes a Assessment by ID
func (mw logMiddleware) DeleteAssessment(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteAssessment", time.Now(), input, nil, &err)
	err = mw.next.DeleteAssessment(ctx, input)
	return
}

/* --------------- Assessment Status --------------- */

// CreateAssessmentStatus creates a new AssessmentStatus
func (mw logMiddleware) CreateAssessmentStatus(ctx context.Context, input *models.AssessmentStatus) (output *models.AssessmentStatus, err error) {
	defer mw.log("CreateAssessmentStatus", time.Now(), input, &output, &err)
	output, err = mw.next.CreateAssessmentStatus(ctx, input)
	return
}

// UpdateAssessmentStatus updates a AssessmentStatus
func (mw logMiddleware) UpdateAssessmentStatus(ctx context.Context, input *models.AssessmentStatus) (output *models.AssessmentStatus, err error) {
	defer mw.log("UpdateAssessmentStatus", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateAssessmentStatus(ctx, input)
	return
}

// DeleteAssessmentStatus deletes a AssessmentStatus by ID
func (mw logMiddleware) DeleteAssessmentStatus(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteAssessmentStatus", time.Now(), input, nil, &err)
	err = mw.next.DeleteAssessmentStatus(ctx, input)
	return
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (mw logMiddleware) CreateQuestion(ctx context.Context, input *models.Question) (output *models.Question, err error) {
	defer mw.log("CreateQuestion", time.Now(), input, &output, &err)
	output, err = mw.next.CreateQuestion(ctx, input)
	return
}

// GetAllQuestions returns all Questions
func (mw logMiddleware) GetAllQuestions(ctx context.Context, input models.QuestionFilters) (output []*models.Question, err error) {
	defer mw.log("GetAllQuestions", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllQuestions(ctx, input)
	return
}

// GetQuestionByID returns a Question by ID
func (mw logMiddleware) GetQuestionByID(ctx context.Context, input uint64) (output *models.Question, err error) {
	defer mw.log("GetQuestionByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetQuestionByID(ctx, input)
	return
}

// UpdateQuestion updates a Question
func (mw logMiddleware) UpdateQuestion(ctx context.Context, input *models.Question) (output *models.Question, err error) {
	defer mw.log("UpdateQuestion", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateQuestion(ctx, input)
	return
}

// DeleteQuestion deletes a Question by ID
func (mw logMiddleware) DeleteQuestion(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteQuestion", time.Now(), input, nil, &err)
	err = mw.next.DeleteQuestion(ctx, input)
	return
}

/* --------------- Tag --------------- */

// CreateTag creates a new Tag
func (mw logMiddleware) CreateTag(ctx context.Context, input *models.Tag) (output *models.Tag, err error) {
	defer mw.log("CreateTag", time.Now(), input, &output, &err)
	output, err = mw.next.CreateTag(ctx, input)
	return
}

// DeleteTag deletes a Tag by ID
func (mw logMiddleware) DeleteTag(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteTag", time.Now(), input, nil, &err)
	err = mw.next.DeleteTag(ctx, input)
	return
}

/* --------------- Response --------------- */

// CreateResponse creates a new Response
func (mw logMiddleware) CreateResponse(ctx context.Context, input *models.Response) (output *models.Response, err error) {
	defer mw.log("CreateResponse", time.Now(), input, &output, &err)
	output, err = mw.next.CreateResponse(ctx, input)
	return
}

// DeleteResponse deletes a Response by ID
func (mw logMiddleware) DeleteResponse(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteResponse", time.Now(), input, nil, &err)
	err = mw.next.DeleteResponse(ctx, input)
	return
}
