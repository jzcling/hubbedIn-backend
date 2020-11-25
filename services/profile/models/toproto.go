package models

import (
	"in-backend/helpers"
	"in-backend/services/profile/pb"
)

// ToProto maps the ORM Candidate model to the proto model
func (m *Candidate) ToProto() *pb.Candidate {
	if m == nil {
		return nil
	}

	var skills []*pb.Skill
	s := m.Skills
	for _, skill := range s {
		skills = append(skills, skill.ToProto())
	}

	var academics []*pb.AcademicHistory
	a := m.Academics
	for _, academic := range a {
		academics = append(academics, academic.ToProto())
	}

	var jobs []*pb.JobHistory
	j := m.Jobs
	for _, job := range j {
		jobs = append(jobs, job.ToProto())
	}

	birthday := helpers.TimeToProto(m.Birthday)
	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)
	deletedAt := helpers.TimeToProto(m.DeletedAt)

	return &pb.Candidate{
		Id:                     m.ID,
		AuthId:                 m.AuthID,
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
		LinkedInUrl:            m.LinkedInURL,
		ScmUrl:                 m.SCMURL,
		WebsiteUrl:             m.WebsiteURL,
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

// ToProto maps the ORM Skill model to the proto model
func (m *Skill) ToProto() *pb.Skill {
	if m == nil {
		return nil
	}
	return &pb.Skill{
		Id:   m.ID,
		Name: m.Name,
	}
}

// ToProto maps the ORM UserSkill model to the proto model
func (m *UserSkill) ToProto() *pb.UserSkill {
	if m == nil {
		return nil
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)

	return &pb.UserSkill{
		Id:          m.ID,
		CandidateId: m.CandidateID,
		SkillId:     m.SkillID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// ToProto maps the ORM Institution model to the proto model
func (m *Institution) ToProto() *pb.Institution {
	if m == nil {
		return nil
	}
	return &pb.Institution{
		Id:      m.ID,
		Country: m.Country,
		Name:    m.Name,
	}
}

// ToProto maps the ORM Course model to the proto model
func (m *Course) ToProto() *pb.Course {
	if m == nil {
		return nil
	}
	return &pb.Course{
		Id:    m.ID,
		Level: m.Level,
		Name:  m.Name,
	}
}

// ToProto maps the ORM AcademicHistory model to the proto model
func (m *AcademicHistory) ToProto() *pb.AcademicHistory {
	if m == nil {
		return nil
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)
	deletedAt := helpers.TimeToProto(m.DeletedAt)

	return &pb.AcademicHistory{
		Id:            m.ID,
		CandidateId:   m.CandidateID,
		InstitutionId: m.InstitutionID,
		Institution:   m.Institution.ToProto(),
		CourseId:      m.CourseID,
		Course:        m.Course.ToProto(),
		YearObtained:  m.YearObtained,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}

// ToProto maps the ORM Company model to the proto model
func (m *Company) ToProto() *pb.Company {
	if m == nil {
		return nil
	}
	return &pb.Company{
		Id:   m.ID,
		Name: m.Name,
	}
}

// ToProto maps the ORM Department model to the proto model
func (m *Department) ToProto() *pb.Department {
	if m == nil {
		return nil
	}
	return &pb.Department{
		Id:   m.ID,
		Name: m.Name,
	}
}

// ToProto maps the ORM JobHistory model to the proto model
func (m *JobHistory) ToProto() *pb.JobHistory {
	if m == nil {
		return nil
	}

	startDate := helpers.TimeToProto(m.StartDate)
	endDate := helpers.TimeToProto(m.EndDate)
	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)
	deletedAt := helpers.TimeToProto(m.DeletedAt)

	return &pb.JobHistory{
		Id:             m.ID,
		CandidateId:    m.CandidateID,
		CompanyId:      m.CompanyID,
		Company:        m.Company.ToProto(),
		DepartmentId:   m.DepartmentID,
		Department:     m.Department.ToProto(),
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
