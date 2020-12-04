package models

import (
	"in-backend/helpers"
	"in-backend/services/profile/pb"
)

// CandidateToORM maps the proto Candidate model to the ORM model
func CandidateToORM(m *pb.Candidate) *Candidate {
	if m == nil {
		return nil
	}

	var skills []*Skill
	s := m.Skills
	for _, skill := range s {
		skills = append(skills, SkillToORM(skill))
	}

	var academics []*AcademicHistory
	a := m.Academics
	for _, academic := range a {
		academics = append(academics, AcademicHistoryToORM(academic))
	}

	var jobs []*JobHistory
	j := m.Jobs
	for _, job := range j {
		jobs = append(jobs, JobHistoryToORM(job))
	}

	birthday := helpers.ProtoTimeToTime(m.Birthday)
	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)
	deletedAt := helpers.ProtoTimeToTime(m.DeletedAt)

	return &Candidate{
		ID:                     m.Id,
		AuthID:                 m.AuthId,
		FirstName:              m.FirstName,
		LastName:               m.LastName,
		Email:                  m.Email,
		ContactNumber:          m.ContactNumber,
		Picture:                m.Picture,
		Gender:                 m.Gender,
		Nationality:            m.Nationality,
		ResidenceCity:          m.ResidenceCity,
		ExpectedSalaryCurrency: m.ExpectedSalaryCurrency,
		ExpectedSalary:         m.ExpectedSalary,
		LinkedInURL:            m.LinkedInUrl,
		SCMURL:                 m.ScmUrl,
		WebsiteURL:             m.WebsiteUrl,
		EducationLevel:         m.EducationLevel,
		Summary:                m.Summary,
		Birthday:               birthday,
		NoticePeriod:           m.NoticePeriod,
		Skills:                 skills,
		Academics:              academics,
		Jobs:                   jobs,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
		DeletedAt:              deletedAt,
	}
}

// SkillToORM maps the proto Skill model to the ORM model
func SkillToORM(m *pb.Skill) *Skill {
	if m == nil {
		return nil
	}
	return &Skill{
		ID:   m.Id,
		Name: m.Name,
	}
}

// UserSkillToORM maps the proto Skill model to the ORM model
func UserSkillToORM(m *pb.UserSkill) *UserSkill {
	if m == nil {
		return nil
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)

	return &UserSkill{
		ID:          m.Id,
		CandidateID: m.CandidateId,
		SkillID:     m.SkillId,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// InstitutionToORM maps the proto Institution model to the ORM model
func InstitutionToORM(m *pb.Institution) *Institution {
	if m == nil {
		return nil
	}
	return &Institution{
		ID:      m.Id,
		Country: m.Country,
		Name:    m.Name,
	}
}

// CourseToORM maps the proto Course model to the ORM model
func CourseToORM(m *pb.Course) *Course {
	if m == nil {
		return nil
	}
	return &Course{
		ID:    m.Id,
		Level: m.Level,
		Name:  m.Name,
	}
}

// AcademicHistoryToORM maps the proto AcademicHistory model to the ORM model
func AcademicHistoryToORM(m *pb.AcademicHistory) *AcademicHistory {
	if m == nil {
		return nil
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)
	deletedAt := helpers.ProtoTimeToTime(m.DeletedAt)

	return &AcademicHistory{
		ID:            m.Id,
		CandidateID:   m.CandidateId,
		InstitutionID: m.InstitutionId,
		Institution:   InstitutionToORM(m.Institution),
		CourseID:      m.CourseId,
		Course:        CourseToORM(m.Course),
		YearObtained:  m.YearObtained,
		Grade:         m.Grade,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}

// CompanyToORM maps the proto Company model to the ORM model
func CompanyToORM(m *pb.Company) *Company {
	if m == nil {
		return nil
	}
	return &Company{
		ID:   m.Id,
		Name: m.Name,
	}
}

// DepartmentToORM maps the proto Department model to the ORM model
func DepartmentToORM(m *pb.Department) *Department {
	if m == nil {
		return nil
	}
	return &Department{
		ID:   m.Id,
		Name: m.Name,
	}
}

// JobHistoryToORM maps the proto JobHistory model to the ORM model
func JobHistoryToORM(m *pb.JobHistory) *JobHistory {
	if m == nil {
		return nil
	}

	startDate := helpers.ProtoTimeToTime(m.StartDate)
	endDate := helpers.ProtoTimeToTime(m.EndDate)
	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)
	deletedAt := helpers.ProtoTimeToTime(m.DeletedAt)

	return &JobHistory{
		ID:             m.Id,
		CandidateID:    m.CandidateId,
		CompanyID:      m.CompanyId,
		Company:        CompanyToORM(m.Company),
		DepartmentID:   m.DepartmentId,
		Department:     DepartmentToORM(m.Department),
		Country:        m.Country,
		City:           m.City,
		Title:          m.Title,
		StartDate:      startDate,
		EndDate:        endDate,
		SalaryCurrency: m.SalaryCurrency,
		Salary:         m.Salary,
		Description:    m.Description,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}
