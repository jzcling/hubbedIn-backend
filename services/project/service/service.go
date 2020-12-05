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

	"github.com/mitchellh/mapstructure"

	"in-backend/services/project"
	"in-backend/services/project/models"
)

// Service implements the project Service interface
type service struct {
	repository project.Repository
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
func New(r project.Repository, c HTTPClient) project.Service {
	return &service{
		repository: r,
		client:     c,
	}
}

/* --------------- Project --------------- */

// CreateProject creates a new Project
func (s *service) CreateProject(ctx context.Context, model *models.Project, cid uint64) (*models.Project, error) {
	f := models.ProjectFilters{
		RepoURL: model.RepoURL,
	}
	existing, err := s.repository.GetAllProjects(ctx, f)
	if err != nil {
		return nil, err
	}

	var m *models.Project
	if existing == nil {
		m, err = s.repository.CreateProject(ctx, model)
		if err != nil {
			return nil, err
		}
	} else {
		m = existing[0]
	}

	cp := &models.CandidateProject{
		CandidateID: cid,
		ProjectID:   m.ID,
	}
	err = s.repository.CreateCandidateProject(ctx, cp)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetAllProjects returns all Projects
func (s *service) GetAllProjects(ctx context.Context, f models.ProjectFilters) ([]*models.Project, error) {
	m, err := s.repository.GetAllProjects(ctx, f)
	return m, err
}

// GetProjectByID returns a Project by ID
func (s *service) GetProjectByID(ctx context.Context, id uint64) (*models.Project, error) {
	m, err := s.repository.GetProjectByID(ctx, id)
	return m, err
}

// UpdateProject updates a Project
func (s *service) UpdateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	m, err := s.repository.UpdateProject(ctx, model)
	return m, err
}

// DeleteProject deletes a Project by ID
func (s *service) DeleteProject(ctx context.Context, id uint64) error {
	err := s.repository.DeleteProject(ctx, id)
	return err
}

// ScanProject scans a Project using sonarqube
func (s *service) ScanProject(ctx context.Context, id uint64) error {
	m, err := s.GetProjectByID(ctx, id)
	if err != nil {
		return err
	}

	go s.scanAndStoreResult(m)

	return nil
}

func (s *service) scanAndStoreResult(m *models.Project) error {
	name := strings.ToLower(m.Name)
	name = strings.ReplaceAll(name, " ", "_")
	name = name + "_" + strconv.FormatUint(m.ID, 10)
	_, err := exec.Command("/bin/sh", "-c", "./scan.sh -u "+m.RepoURL+" -n "+name).Output()
	if err != nil {
		return err
	}

	jsonBody, err := s.getRatingMeasures(name)
	if err != nil {
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
				return err
			}
			r.ReliabilityRating = int32(v)
		case "sqale_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				return err
			}
			r.MaintainabilityRating = int32(v)
		case "security_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				return err
			}
			r.SecurityRating = int32(v)
		case "security_review_rating":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				return err
			}
			r.SecurityReviewRating = int32(v)
		case "coverage":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				return err
			}
			r.Coverage = float32(v)
		case "duplicated_lines_density":
			v, err := strconv.ParseFloat(measure.Value, 32)
			if err != nil {
				return err
			}
			r.Duplications = float32(v)
		case "ncloc":
			r.Lines, err = strconv.ParseUint(measure.Value, 10, 64)
			if err != nil {
				return err
			}
		}
	}

	err = s.CreateRating(context.Background(), r)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) getRatingMeasures(name string) (map[string]interface{}, error) {
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
	req, err := http.NewRequest("GET", "http://sonarqube-sonarqube-svc:9000/api/measures/component?"+payload.Encode(), nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, nil
}

/* --------------- Candidate Project --------------- */

// CreateCandidateProject creates a new CandidateProject
func (s *service) CreateCandidateProject(ctx context.Context, m *models.CandidateProject) error {
	err := s.repository.CreateCandidateProject(ctx, m)
	return err
}

// DeleteCandidateProject deletes a CandidateProject by ID
func (s *service) DeleteCandidateProject(ctx context.Context, id uint64) error {
	err := s.repository.DeleteCandidateProject(ctx, id)
	return err
}

/* --------------- Rating --------------- */

// CreateRating creates a new Rating
func (s *service) CreateRating(ctx context.Context, m *models.Rating) error {
	err := s.repository.CreateRating(ctx, m)
	return err
}

// DeleteRating deletes a Rating by Candidate ID and Project ID
func (s *service) DeleteRating(ctx context.Context, id uint64) error {
	err := s.repository.DeleteRating(ctx, id)
	return err
}
