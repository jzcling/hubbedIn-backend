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

	"in-backend/services/joblisting"
	"in-backend/services/joblisting/models"
)

// Service implements the joblisting Service interface
type service struct {
	repository joblisting.Repository
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

// New creates and returns a new Service that implements the joblisting Service interface
func New(r joblisting.Repository, l log.Logger, c HTTPClient) joblisting.Service {
	return &service{
		repository: r,
		logger:     l,
		client:     c,
	}
}

/* --------------- Joblisting --------------- */

// CreateJoblisting creates a new Joblisting
func (s *service) CreateJoblisting(ctx context.Context, model *models.Joblisting, cid uint64) (*models.Joblisting, error) {
	logger := log.With(s.logger, "method", "CreateJoblisting")

	f := models.JoblistingFilters{
		RepoURL: model.RepoURL,
	}
	existing, err := s.repository.GetAllJoblistings(ctx, f)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	var m *models.Joblisting
	if existing == nil {
		m, err = s.repository.CreateJoblisting(ctx, model)
		if err != nil {
			level.Error(logger).Log("err", err)
			return nil, err
		}
	} else {
		m = existing[0]
	}

	cp := &models.CandidateJoblisting{
		CandidateID:  cid,
		JoblistingID: m.ID,
	}
	err = s.repository.CreateCandidateJoblisting(ctx, cp)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return m, err
}

// GetAllJoblistings returns all Joblistings
func (s *service) GetAllJoblistings(ctx context.Context, f models.JoblistingFilters) ([]*models.Joblisting, error) {
	logger := log.With(s.logger, "method", "GetAllJoblistings")

	m, err := s.repository.GetAllJoblistings(ctx, f)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// GetJoblistingByID returns a Joblisting by ID
func (s *service) GetJoblistingByID(ctx context.Context, id uint64) (*models.Joblisting, error) {
	logger := log.With(s.logger, "method", "GetJoblistingByID")

	m, err := s.repository.GetJoblistingByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// UpdateJoblisting updates a Joblisting
func (s *service) UpdateJoblisting(ctx context.Context, model *models.Joblisting) (*models.Joblisting, error) {
	logger := log.With(s.logger, "method", "UpdateJoblisting")

	m, err := s.repository.UpdateJoblisting(ctx, model)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return m, err
}

// DeleteJoblisting deletes a Joblisting by ID
func (s *service) DeleteJoblisting(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteJoblisting")

	err := s.repository.DeleteJoblisting(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}

// ScanJoblisting scans a Joblisting using sonarqube
func (s *service) ScanJoblisting(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "ScanJoblisting")

	m, err := s.GetJoblistingByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	go s.scanAndStoreResult(m, logger)

	return nil
}

func (s *service) scanAndStoreResult(m *models.Joblisting, logger log.Logger) error {
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
		JoblistingID: m.ID,
		CreatedAt:    &now,
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

/* --------------- Candidate Joblisting --------------- */

// CreateCandidateJoblisting creates a new CandidateJoblisting
func (s *service) CreateCandidateJoblisting(ctx context.Context, m *models.CandidateJoblisting) error {
	logger := log.With(s.logger, "method", "CreateCandidateJoblisting")

	err := s.repository.CreateCandidateJoblisting(ctx, m)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	return err
}

// DeleteCandidateJoblisting deletes a CandidateJoblisting by ID
func (s *service) DeleteCandidateJoblisting(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteCandidateJoblisting")

	err := s.repository.DeleteCandidateJoblisting(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
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

// DeleteRating deletes a Rating by Candidate ID and Joblisting ID
func (s *service) DeleteRating(ctx context.Context, id uint64) error {
	logger := log.With(s.logger, "method", "DeleteRating")

	err := s.repository.DeleteRating(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
	}
	return err
}
