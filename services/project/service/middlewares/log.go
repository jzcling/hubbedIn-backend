package middlewares

import (
	"context"
	"fmt"
	"in-backend/services/project"
	"in-backend/services/project/models"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type logMiddleware struct {
	logger log.Logger
	next   project.Service
}

// NewLogMiddleware creates and returns a new Log Middleware that implements the project Service interface
func NewLogMiddleware(logger log.Logger, svc project.Service) project.Service {
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

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (mw logMiddleware) CreateProject(ctx context.Context, input *models.Project, cid uint64) (output *models.Project, err error) {
	defer mw.log("CreateProject", time.Now(), input, &output, &err)
	output, err = mw.next.CreateProject(ctx, input, cid)
	return
}

// GetAllProjects returns all Projects
func (mw logMiddleware) GetAllProjects(ctx context.Context, input models.ProjectFilters) (output []*models.Project, err error) {
	defer mw.log("GetAllProjects", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllProjects(ctx, input)
	return
}

// GetProjectByID returns a Project by ID
func (mw logMiddleware) GetProjectByID(ctx context.Context, input uint64) (output *models.Project, err error) {
	defer mw.log("GetProjectByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetProjectByID(ctx, input)
	return
}

// UpdateProject updates a Project
func (mw logMiddleware) UpdateProject(ctx context.Context, input *models.Project) (output *models.Project, err error) {
	defer mw.log("UpdateProject", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateProject(ctx, input)
	return
}

// DeleteProject deletes a Project by ID
func (mw logMiddleware) DeleteProject(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteProject", time.Now(), input, nil, &err)
	err = mw.next.DeleteProject(ctx, input)
	return
}

// ScanProject scans a Project using sonarqube
func (mw logMiddleware) ScanProject(ctx context.Context, input uint64) (err error) {
	defer mw.log("ScanProject", time.Now(), input, nil, &err)
	err = mw.next.ScanProject(ctx, input)
	return
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (mw logMiddleware) CreateCandidateProject(ctx context.Context, input *models.CandidateProject) (err error) {
	defer mw.log("CreateCandidateProject", time.Now(), input, nil, &err)
	err = mw.next.CreateCandidateProject(ctx, input)
	return
}

// DeleteCandidateProject deletes a CandidateProject by ID
func (mw logMiddleware) DeleteCandidateProject(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteCandidateProject", time.Now(), input, nil, &err)
	err = mw.next.DeleteCandidateProject(ctx, input)
	return
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (mw logMiddleware) CreateRating(ctx context.Context, input *models.Rating) (err error) {
	defer mw.log("CreateRating", time.Now(), input, nil, &err)
	err = mw.next.CreateRating(ctx, input)
	return
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (mw logMiddleware) DeleteRating(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteRating", time.Now(), input, nil, &err)
	err = mw.next.DeleteRating(ctx, input)
	return
}
