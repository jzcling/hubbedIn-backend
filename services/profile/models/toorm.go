package models

import (
	"in-backend/services/profile/pb"
	"time"

	"github.com/golang/protobuf/ptypes"
)

// CandidateToORM maps the proto Candidate model to the ORM model
func CandidateToORM(c *pb.Candidate) *Candidate {
	if c == nil {
		return nil
	}

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

	b, err := ptypes.Timestamp(c.Birthday)
	birthday := &b
	if err != nil {
		birthday = (*time.Time)(nil)
	}

	ca, err := ptypes.Timestamp(c.CreatedAt)
	createdAt := &ca
	if err != nil {
		createdAt = (*time.Time)(nil)
	}

	ua, err := ptypes.Timestamp(c.UpdatedAt)
	updatedAt := &ua
	if err != nil {
		updatedAt = (*time.Time)(nil)
	}

	da, err := ptypes.Timestamp(c.DeletedAt)
	deletedAt := &da
	if err != nil {
		deletedAt = (*time.Time)(nil)
	}

	return &Candidate{
		ID:                     c.Id,
		FirstName:              c.FirstName,
		LastName:               c.LastName,
		Email:                  c.Email,
		ContactNumber:          c.ContactNumber,
		Picture:                c.Picture,
		Gender:                 c.Gender,
		Nationality:            c.Nationality,
		ResidenceCity:          c.ResidenceCity,
		ExpectedSalaryCurrency: c.ExpectedSalaryCurrency,
		ExpectedSalary:         c.ExpectedSalary,
		LinkedInURL:            c.LinkedInUrl,
		SCMURL:                 c.ScmUrl,
		WebsiteURL:             c.WebsiteUrl,
		EducationLevel:         c.EducationLevel,
		Summary:                c.Summary,
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

// SkillToORM maps the proto Skill model to the ORM model
func SkillToORM(s *pb.Skill) *Skill {
	if s == nil {
		return nil
	}
	return &Skill{
		ID:   s.Id,
		Name: s.Name,
	}
}

// InstitutionToORM maps the proto Institution model to the ORM model
func InstitutionToORM(i *pb.Institution) *Institution {
	if i == nil {
		return nil
	}
	return &Institution{
		ID:      i.Id,
		Country: i.Country,
		Name:    i.Name,
	}
}

// CourseToORM maps the proto Course model to the ORM model
func CourseToORM(c *pb.Course) *Course {
	if c == nil {
		return nil
	}
	return &Course{
		ID:    c.Id,
		Level: c.Level,
		Name:  c.Name,
	}
}

// AcademicHistoryToORM maps the proto AcademicHistory model to the ORM model
func AcademicHistoryToORM(a *pb.AcademicHistory) *AcademicHistory {
	if a == nil {
		return nil
	}

	ca, err := ptypes.Timestamp(a.CreatedAt)
	createdAt := &ca
	if err != nil {
		createdAt = (*time.Time)(nil)
	}

	ua, err := ptypes.Timestamp(a.UpdatedAt)
	updatedAt := &ua
	if err != nil {
		updatedAt = (*time.Time)(nil)
	}

	da, err := ptypes.Timestamp(a.DeletedAt)
	deletedAt := &da
	if err != nil {
		deletedAt = (*time.Time)(nil)
	}

	return &AcademicHistory{
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

// CompanyToORM maps the proto Company model to the ORM model
func CompanyToORM(c *pb.Company) *Company {
	if c == nil {
		return nil
	}
	return &Company{
		ID:   c.Id,
		Name: c.Name,
	}
}

// DepartmentToORM maps the proto Department model to the ORM model
func DepartmentToORM(d *pb.Department) *Department {
	if d == nil {
		return nil
	}
	return &Department{
		ID:   d.Id,
		Name: d.Name,
	}
}

// JobHistoryToORM maps the proto JobHistory model to the ORM model
func JobHistoryToORM(j *pb.JobHistory) *JobHistory {
	if j == nil {
		return nil
	}

	sd, err := ptypes.Timestamp(j.StartDate)
	startDate := &sd
	if err != nil {
		startDate = (*time.Time)(nil)
	}

	ed, err := ptypes.Timestamp(j.EndDate)
	endDate := &ed
	if err != nil {
		endDate = (*time.Time)(nil)
	}

	ca, err := ptypes.Timestamp(j.CreatedAt)
	createdAt := &ca
	if err != nil {
		createdAt = (*time.Time)(nil)
	}

	ua, err := ptypes.Timestamp(j.UpdatedAt)
	updatedAt := &ua
	if err != nil {
		updatedAt = (*time.Time)(nil)
	}

	da, err := ptypes.Timestamp(j.DeletedAt)
	deletedAt := &da
	if err != nil {
		deletedAt = (*time.Time)(nil)
	}

	return &JobHistory{
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
