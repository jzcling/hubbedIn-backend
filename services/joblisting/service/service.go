package service

import (
	"context"

	"in-backend/services/joblisting/interfaces"
	"in-backend/services/joblisting/models"
	profilePb "in-backend/services/profile/pb"

	"google.golang.org/grpc"
)

// Service implements the joblisting Service interface
type service struct {
	repository interfaces.Repository
}

var (
	profileSvcAddr string = "profile-service:50051"
)

// New creates and returns a new Service that implements the joblisting Service interface
func New(r interfaces.Repository) interfaces.Service {
	return &service{
		repository: r,
	}
}

/* --------------- Job Post --------------- */

// CreateJobPost creates a new JobPost
func (s *service) CreateJobPost(ctx context.Context, model *models.JobPost) (*models.JobPost, error) {
	m, err := s.repository.CreateJobPost(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// BulkCreateJobPost creates multiple JobPosts
func (s *service) BulkCreateJobPost(ctx context.Context, models []*models.JobPost) ([]*models.JobPost, error) {
	m, err := s.repository.BulkCreateJobPost(ctx, models)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllJobPosts returns all JobPosts that match the filters
func (s *service) GetAllJobPosts(ctx context.Context, f models.JobPostFilters) ([]*models.JobPost, error) {
	m, err := s.repository.GetAllJobPosts(ctx, f)
	if err != nil {
		return nil, err
	}

	skills, err := getAllSkills(ctx)
	if err != nil {
		return nil, err
	}

	for _, jp := range m {
		var jobSkills []*models.Skill
		for _, skid := range jp.SkillID {
			for _, sk := range skills {
				if skid == sk.ID {
					jobSkills = append(jobSkills, sk)
				}
			}
		}
		jp.Skills = jobSkills
	}

	return m, err
}

func getAllSkills(ctx context.Context) ([]*models.Skill, error) {
	conn, err := grpc.Dial(profileSvcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := profilePb.NewProfileServiceClient(conn)
	getReq := profilePb.GetAllSkillsRequest{}
	sk, err := client.GetAllSkills(ctx, &getReq)
	if err != nil {
		return nil, err
	}

	var skills []*models.Skill
	for _, skill := range sk.Skills {
		s := &models.Skill{
			ID:   skill.Id,
			Name: skill.Name,
		}
		skills = append(skills, s)
	}

	return skills, nil
}

// GetJobPostByID finds and returns a JobPost by ID
func (s *service) GetJobPostByID(ctx context.Context, id uint64) (*models.JobPost, error) {
	m, err := s.repository.GetJobPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	skills, err := getAllSkills(ctx)
	if err != nil {
		return nil, err
	}

	var jobSkills []*models.Skill
	for _, skid := range m.SkillID {
		for _, sk := range skills {
			if skid == sk.ID {
				jobSkills = append(jobSkills, sk)
			}
		}
	}
	m.Skills = jobSkills

	return m, err
}

// UpdateJobPost updates a JobPost
func (s *service) UpdateJobPost(ctx context.Context, model *models.JobPost) (*models.JobPost, error) {
	m, err := s.repository.UpdateJobPost(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteJobPost deletes a JobPost by ID
func (s *service) DeleteJobPost(ctx context.Context, id uint64) error {
	err := s.repository.DeleteJobPost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (s *service) CreateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	m, err := s.repository.CreateCompany(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllCompanies returns all Companies that match the filters
func (s *service) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	m, err := s.repository.GetAllCompanies(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// UpdateCompany updates a Company
func (s *service) UpdateCompany(ctx context.Context, model *models.Company) (*models.Company, error) {
	m, err := s.repository.UpdateCompany(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteCompany deletes a Company by ID
func (s *service) DeleteCompany(ctx context.Context, id uint64) error {
	err := s.repository.DeleteCompany(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

/* --------------- Industry --------------- */

// CreateIndustry creates a new Industry
func (s *service) CreateIndustry(ctx context.Context, model *models.Industry) (*models.Industry, error) {
	m, err := s.repository.CreateIndustry(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllIndustries returns all Industries
func (s *service) GetAllIndustries(ctx context.Context, f models.IndustryFilters) ([]*models.Industry, error) {
	m, err := s.repository.GetAllIndustries(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteIndustry deletes a Industry by ID
func (s *service) DeleteIndustry(ctx context.Context, id uint64) error {
	err := s.repository.DeleteIndustry(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

/* --------------- Job Function --------------- */

// CreateJobFunction creates a new JobFunction
func (s *service) CreateJobFunction(ctx context.Context, model *models.JobFunction) (*models.JobFunction, error) {
	m, err := s.repository.CreateJobFunction(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllJobFunctions returns all JobFunctions
func (s *service) GetAllJobFunctions(ctx context.Context, f models.JobFunctionFilters) ([]*models.JobFunction, error) {
	m, err := s.repository.GetAllJobFunctions(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteJobFunction deletes a JobFunction by ID
func (s *service) DeleteJobFunction(ctx context.Context, id uint64) error {
	err := s.repository.DeleteJobFunction(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

/* --------------- Key Person --------------- */

// CreateKeyPerson creates a new KeyPerson
func (s *service) CreateKeyPerson(ctx context.Context, model *models.KeyPerson) (*models.KeyPerson, error) {
	m, err := s.repository.CreateKeyPerson(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// BulkCreateKeyPerson creates multiple KeyPersons
func (s *service) BulkCreateKeyPerson(ctx context.Context, models []*models.KeyPerson) ([]*models.KeyPerson, error) {
	m, err := s.repository.BulkCreateKeyPerson(ctx, models)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllKeyPersons returns all KeyPersons that match the filters
func (s *service) GetAllKeyPersons(ctx context.Context, f models.KeyPersonFilters) ([]*models.KeyPerson, error) {
	m, err := s.repository.GetAllKeyPersons(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetKeyPersonByID finds and returns a KeyPerson by ID
func (s *service) GetKeyPersonByID(ctx context.Context, id uint64) (*models.KeyPerson, error) {
	m, err := s.repository.GetKeyPersonByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return m, err
}

// UpdateKeyPerson updates a KeyPerson
func (s *service) UpdateKeyPerson(ctx context.Context, model *models.KeyPerson) (*models.KeyPerson, error) {
	m, err := s.repository.UpdateKeyPerson(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteKeyPerson deletes a KeyPerson by ID
func (s *service) DeleteKeyPerson(ctx context.Context, id uint64) error {
	err := s.repository.DeleteKeyPerson(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

/* --------------- Job Platform --------------- */

// CreateJobPlatform creates a new JobPlatform
func (s *service) CreateJobPlatform(ctx context.Context, model *models.JobPlatform) (*models.JobPlatform, error) {
	m, err := s.repository.CreateJobPlatform(ctx, model)
	if err != nil {
		return nil, err
	}
	return m, err
}

// GetAllJobPlatforms returns all JobPlatforms
func (s *service) GetAllJobPlatforms(ctx context.Context, f models.JobPlatformFilters) ([]*models.JobPlatform, error) {
	m, err := s.repository.GetAllJobPlatforms(ctx, f)
	if err != nil {
		return nil, err
	}
	return m, err
}

// DeleteJobPlatform deletes a JobPlatform by ID
func (s *service) DeleteJobPlatform(ctx context.Context, id uint64) error {
	err := s.repository.DeleteJobPlatform(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
