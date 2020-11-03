package pb

import (
	"in-backend/services/profile/models"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// ToORM maps the grpc Candidate model to the ORM model
func (c *Candidate) ToORM() models.Candidate {
	var skills []*models.Skill
	s := c.Skills
	for _, skill := range s {
		skills = append(skills, skill.ToORM())
	}

	var academics []*models.AcademicHistory
	a := c.Academics
	for _, academic := range a {
		academics = append(academics, academic.ToORM())
	}

	var jobs []*models.JobHistory
	j := c.Jobs
	for _, job := range j {
		jobs = append(jobs, job.ToORM())
	}

	birthday, err := ptypes.Timestamp(c.Birthday)
	if err != nil {
		birthday = time.Time{}
	}

	createdAt, err := ptypes.Timestamp(c.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	updatedAt, err := ptypes.Timestamp(c.UpdatedAt)
	if err != nil {
		updatedAt = time.Time{}
	}

	deletedAt, err := ptypes.Timestamp(c.DeletedAt)
	if err != nil {
		deletedAt = time.Time{}
	}

	return models.Candidate{
		ID:                     c.Id,
		FirstName:              c.FirstName,
		LastName:               c.LastName,
		Email:                  c.Email,
		ContactNumber:          c.ContactNumber,
		Gender:                 c.Gender,
		Nationality:            c.Nationality,
		ResidenceCity:          c.ResidenceCity,
		ExpectedSalaryCurrency: c.ExpectedSalaryCurrency,
		ExpectedSalary:         c.ExpectedSalary,
		LinkedInURL:            c.LinkedInUrl,
		SCMURL:                 c.ScmUrl,
		EducationLevel:         c.EducationLevel,
		Birthday:               birthday,
		NoticePeriod:           c.NoticePeriod,
		Skills:                 skills,
		Academics:              academics,
		Jobs:                   jobs,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
		DeletedAt:              deletedAt,
	}
}

// ToORM maps the grpc Skill model to the ORM model
func (s *Skill) ToORM() *models.Skill {
	return &models.Skill{
		ID:   s.Id,
		Name: s.Name,
	}
}

// ToORM maps the grpc AcademicHistory model to the ORM model
func (a *AcademicHistory) ToORM() *models.AcademicHistory {
	createdAt, err := ptypes.Timestamp(a.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	updatedAt, err := ptypes.Timestamp(a.UpdatedAt)
	if err != nil {
		updatedAt = time.Time{}
	}

	deletedAt, err := ptypes.Timestamp(a.DeletedAt)
	if err != nil {
		deletedAt = time.Time{}
	}

	return &models.AcademicHistory{
		ID:            a.Id,
		CandidateID:   a.CandidateId,
		InstitutionID: a.InstitutionId,
		CourseID:      a.CourseId,
		YearObtained:  a.YearObtained,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}

// ToORM maps the grpc JobHistory model to the ORM model
func (j *JobHistory) ToORM() *models.JobHistory {
	startDate, err := ptypes.Timestamp(j.StartDate)
	if err != nil {
		startDate = time.Time{}
	}

	endDate, err := ptypes.Timestamp(j.EndDate)
	if err != nil {
		endDate = time.Time{}
	}

	createdAt, err := ptypes.Timestamp(j.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	updatedAt, err := ptypes.Timestamp(j.UpdatedAt)
	if err != nil {
		updatedAt = time.Time{}
	}

	deletedAt, err := ptypes.Timestamp(j.DeletedAt)
	if err != nil {
		deletedAt = time.Time{}
	}
	return &models.JobHistory{
		ID:             j.Id,
		CandidateID:    j.CandidateId,
		CompanyID:      j.CompanyId,
		DepartmentID:   j.DepartmentId,
		Country:        j.Country,
		City:           j.City,
		Title:          j.Title,
		StartDate:      startDate,
		EndDate:        endDate,
		SalaryCurrency: j.SalaryCurrency,
		Salary:         j.Salary,
		Description:    j.Description,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}
