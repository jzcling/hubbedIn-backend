package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/mitchellh/mapstructure"

	"in-backend/services/project"
	"in-backend/services/project/models"
)

// Service implements the project Service interface
type service struct {
	repository project.Repository
	logger     log.Logger
	client     HTTPClient
}

// HTTPClient describes a default http client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Component declares the model of a metrics response from sonarqube's API
type Component struct {
	Measures []Measure
	Other    map[string]interface{} `mapstructure:",remain"`
}

// Measure declares the model of a measure response from sonarqube's API
type Measure struct {
	Metric    string `mapstructure:"metric"`
	Value     string `mapstructure:"value"`
	BestValue bool   `mapstructure:"bestValue,omitempty"`
}

// New creates and returns a new Service that implements the project Service interface
func New(r project.Repository, l log.Logger, c HTTPClient) project.Service {
	return &service{
		repository: r,
		logger:     l,
		client:     c,
	}
}

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (s *service) CreateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	logger := log.With(s.logger, "method", "CreateProject")

	m, err := s.repository.CreateProject(ctx, model)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return m, err
}

// GetAllProjects returns all Projects
func (s *service) GetAllProjects(ctx context.Context) ([]*models.Project, error) {
	logger := log.With(s.logger, "method", "GetAllProjects")

	m, err := s.repository.GetAllProjects(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// GetProjectByID returns a Project by ID
func (s *service) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	logger := log.With(s.logger, "method", "GetProjectByID")

	m, err := s.repository.GetProjectByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// UpdateProject updates a Project
func (s *service) UpdateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	logger := log.With(s.logger, "method", "UpdateProject")

	m, err := s.repository.UpdateProject(ctx, model)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// DeleteProject deletes a Project by ID
func (s *service) DeleteProject(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteProject")

	err := s.repository.DeleteProject(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

// ScanProject scans a Project using sonarqube
func (s *service) ScanProject(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "ScanProject")

	m, err := s.GetProjectByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	go s.scanAndStoreResult(m, logger)

	return nil
}

func (s *service) scanAndStoreResult(m *models.Project, logger log.Logger) error {
	name := strings.ToLower(m.Name)
	name = strings.ReplaceAll(name, " ", "_")
	name = name + "_" + strconv.FormatUint(m.ID, 10)
	_, err := exec.Command("/bin/sh", "-c", "./scan.sh -u "+m.RepoURL+" -n "+name).Output()
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	jsonBody, err := s.getRatingMeasures(name, logger)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	var component Component
	mapstructure.Decode(jsonBody["component"].(map[string]interface{}), &component)

	now := time.Now()
	r := &models.Rating{
		ProjectID: m.ID,
		CreatedAt: &now,
	}
	for _, measure := range component.Measures {
		switch measure.Metric {
		case "reliability_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.ReliabilityRating = int32(v)
		case "sqale_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.MaintainabilityRating = int32(v)
		case "security_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.SecurityRating = int32(v)
		case "security_review_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.SecurityReviewRating = int32(v)
		case "coverage":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.Coverage = float32(v)
		case "duplicated_lines_density":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
			r.Duplications = float32(v)
		case "ncloc":
			r.Lines, err = strconv.ParseUint(measure.Value, 10, 64)
			if err != nil {
				level.Error(logger).Log("err", err)
				return err
			}
		}
	}

	err = s.CreateRating(context.Background(), r)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

func (s *service) getRatingMeasures(name string, logger log.Logger) (map[string]interface{}, error) {
	metrics := []string{
		"reliability_rating",
		"sqale_rating",
		"security_rating",
		"security_review_rating",
		"coverage",
		"duplicated_lines_density",
		"ncloc",
	}

	payload := url.Values{}
	payload.Add("component", name)
	payload.Add("metricKeys", strings.Join(metrics, ","))
	req, err := http.NewRequest("GET", "http://sonarqube:9000/api/measures/component?"+payload.Encode(), nil)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	jsonBody := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return jsonBody, nil
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (s *service) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	logger := log.With(s.logger, "method", "CreateCandidateProject")

	err := s.repository.CreateCandidateProject(ctx, m)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return err
}

// DeleteCandidateProject deletes a CandidateProject by Candidate ID and Project ID
func (s *service) DeleteCandidateProject(ctx context.Context, cid, pid uint64) error {
	logger := log.With(s.logger, "method", "DeleteCandidateProject")

	err := s.repository.DeleteCandidateProject(ctx, cid, pid)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

// GetAllProjectsByCandidate returns all Projects by a Candidate
func (s *service) GetAllProjectsByCandidate(ctx context.Context, cid uint64) ([]*models.Project, error) {
	logger := log.With(s.logger, "method", "GetAllProjectsByCandidate")

	c, err := s.repository.GetAllProjectsByCandidate(ctx, cid)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return c, err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (s *service) CreateRating(ctx context.Context, m *models.Rating) error {
	logger := log.With(s.logger, "method", "CreateRating")

	err := s.repository.CreateRating(ctx, m)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return err
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (s *service) DeleteRating(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteRating")

	err := s.repository.DeleteRating(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}
