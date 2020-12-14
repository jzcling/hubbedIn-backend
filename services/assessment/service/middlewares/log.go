package middlewares

import (
	"context"
	"fmt"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	var logger log.Logger
	if err != nil {
		logger = level.Error(mw.logger)
	} else {
		logger = level.Info(mw.logger)
	}
	logger.Log(
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
func (mw logMiddleware) GetAllAssessments(ctx context.Context, input models.AssessmentFilters, role *string, cid *uint64) (output []*models.Assessment, err error) {
	defer mw.log("GetAllAssessments", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllAssessments(ctx, input, role, cid)
	return
}

// GetAssessmentByID returns a Assessment by ID
func (mw logMiddleware) GetAssessmentByID(ctx context.Context, input uint64, role *string, cid *uint64) (output *models.Assessment, err error) {
	defer mw.log("GetAssessmentByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetAssessmentByID(ctx, input, role, cid)
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

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (mw logMiddleware) CreateAssessmentAttempt(ctx context.Context, input *models.AssessmentAttempt) (output *models.AssessmentAttempt, err error) {
	defer mw.log("CreateAssessmentAttempt", time.Now(), input, &output, &err)
	output, err = mw.next.CreateAssessmentAttempt(ctx, input)
	return
}

// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (mw logMiddleware) GetAssessmentAttemptByID(ctx context.Context, input uint64) (output *models.AssessmentAttempt, err error) {
	defer mw.log("GetAssessmentAttemptByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetAssessmentAttemptByID(ctx, input)
	return
}

// LocalGetAssessmentAttemptByID returns a AssessmentAttempt by ID
// This method is only for local server to server communication
func (mw logMiddleware) LocalGetAssessmentAttemptByID(ctx context.Context, input uint64) (output *models.AssessmentAttempt, err error) {
	defer mw.log("GetAssessmentAttemptByID", time.Now(), input, &output, &err)
	output, err = mw.next.LocalGetAssessmentAttemptByID(ctx, input)
	return
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (mw logMiddleware) UpdateAssessmentAttempt(ctx context.Context, input *models.AssessmentAttempt) (output *models.AssessmentAttempt, err error) {
	defer mw.log("UpdateAssessmentAttempt", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateAssessmentAttempt(ctx, input)
	return
}

// LocalUpdateAssessmentAttempt updates a AssessmentAttempt
// This method is only for local server to server communication
func (mw logMiddleware) LocalUpdateAssessmentAttempt(ctx context.Context, input *models.AssessmentAttempt) (output *models.AssessmentAttempt, err error) {
	defer mw.log("UpdateAssessmentAttempt", time.Now(), input, &output, &err)
	output, err = mw.next.LocalUpdateAssessmentAttempt(ctx, input)
	return
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (mw logMiddleware) DeleteAssessmentAttempt(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteAssessmentAttempt", time.Now(), input, nil, &err)
	err = mw.next.DeleteAssessmentAttempt(ctx, input)
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

/* --------------- Attempt Question --------------- */

// UpdateAttemptQuestion updates a AttemptQuestion
func (mw logMiddleware) UpdateAttemptQuestion(ctx context.Context, input *models.AttemptQuestion) (output *models.AttemptQuestion, err error) {
	defer mw.log("UpdateAttemptQuestion", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateAttemptQuestion(ctx, input)
	return
}
