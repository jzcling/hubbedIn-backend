package models

import (
	"in-backend/services/profile/pb"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// CandidateToORM maps the proto Candidate model to the ORM model
func CandidateToORM(c *pb.Candidate) *Candidate {
	var skills []*Skill
	s := c.Skills
	for _, skill := range s {
		skills = append(skills, SkillToORM(skill))
	}

	var academics []*AcademicHistory
	a := c.Academics
	for _, academic := range a {
		academics = append(academics, AcademicHistoryToORM(academic))
	}

	var jobs []*JobHistory
	j := c.Jobs
	for _, job := range j {
		jobs = append(jobs, JobHistoryToORM(job))
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

	return &Candidate{
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
		Birthday:               &birthday,
		NoticePeriod:           c.NoticePeriod,
		Skills:                 skills,
		Academics:              academics,
		Jobs:                   jobs,
		CreatedAt:              &createdAt,
		UpdatedAt:              &updatedAt,
		DeletedAt:              &deletedAt,
	}
}

// SkillToORM maps the proto Skill model to the ORM model
func SkillToORM(s *pb.Skill) *Skill {
	return &Skill{
		ID:   s.Id,
		Name: s.Name,
	}
}

// InstitutionToORM maps the proto Institution model to the ORM model
func InstitutionToORM(i *pb.Institution) *Institution {
	return &Institution{
		ID:      i.Id,
		Country: i.Country,
		Name:    i.Name,
	}
}

// CourseToORM maps the proto Course model to the ORM model
func CourseToORM(c *pb.Course) *Course {
	return &Course{
		ID:    c.Id,
		Level: c.Level,
		Name:  c.Name,
	}
}

// AcademicHistoryToORM maps the proto AcademicHistory model to the ORM model
func AcademicHistoryToORM(a *pb.AcademicHistory) *AcademicHistory {
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

	return &AcademicHistory{
		ID:            a.Id,
		CandidateID:   a.CandidateId,
		InstitutionID: a.InstitutionId,
		CourseID:      a.CourseId,
		YearObtained:  a.YearObtained,
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
		DeletedAt:     &deletedAt,
	}
}

// CompanyToORM maps the proto Company model to the ORM model
func CompanyToORM(c *pb.Company) *Company {
	return &Company{
		ID:   c.Id,
		Name: c.Name,
	}
}

// DepartmentToORM maps the proto Department model to the ORM model
func DepartmentToORM(d *pb.Department) *Department {
	return &Department{
		ID:   d.Id,
		Name: d.Name,
	}
}

// JobHistoryToORM maps the proto JobHistory model to the ORM model
func JobHistoryToORM(j *pb.JobHistory) *JobHistory {
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
	return &JobHistory{
		ID:             j.Id,
		CandidateID:    j.CandidateId,
		CompanyID:      j.CompanyId,
		DepartmentID:   j.DepartmentId,
		Country:        j.Country,
		City:           j.City,
		Title:          j.Title,
		StartDate:      &startDate,
		EndDate:        &endDate,
		SalaryCurrency: j.SalaryCurrency,
		Salary:         j.Salary,
		Description:    j.Description,
		CreatedAt:      &createdAt,
		UpdatedAt:      &updatedAt,
		DeletedAt:      &deletedAt,
	}
}
