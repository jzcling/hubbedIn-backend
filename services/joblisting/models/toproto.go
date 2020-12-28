package models

import (
	"in-backend/helpers"
	"in-backend/services/joblisting/pb"
)

// ToProto maps the ORM JobPost model to the proto model
func (m *JobPost) ToProto() *pb.JobPost {
	if m == nil {
		return nil
	}

	var skills []*pb.ProfileSkill
	s := m.Skills
	for _, skill := range s {
		skills = append(skills, skill.ToProto())
	}

	createdAt := helpers.TimeToProto(m.CreatedAt)
	updatedAt := helpers.TimeToProto(m.UpdatedAt)
	startAt := helpers.TimeToProto(m.StartAt)
	expireAt := helpers.TimeToProto(m.ExpireAt)

	return &pb.JobPost{
		Id:              m.ID,
		CompanyId:       m.CompanyID,
		HrContactId:     m.HRContactID,
		HiringManagerId: m.HiringManagerID,
		JobPlatformId:   m.JobPlatformID,
		Title:           m.Title,
		Description:     m.Description,
		SeniorityLevel:  m.SeniorityLevel,
		YearsExperience: m.YearsExperience,
		EmploymentType:  m.EmploymentType,
		FunctionId:      m.FunctionID,
		IndustryId:      m.IndustryID,
		Location:        m.Location,
		Remote:          m.Remote,
		SalaryCurrency:  m.SalaryCurrency,
		MinSalary:       m.MinSalary,
		MaxSalary:       m.MaxSalary,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		StartAt:         startAt,
		ExpireAt:        expireAt,
		SkillId:         m.SkillID,
		Company:         m.Company.ToProto(),
		Function:        m.Function.ToProto(),
		Industry:        m.Industry.ToProto(),
		JobPlatform:     m.JobPlatform.ToProto(),
		HrContact:       m.HRContact.ToProto(),
		HiringManager:   m.HiringManager.ToProto(),
		Skills:          skills,
	}
}

// ToProto maps the ORM Company model to the proto model
func (m *Company) ToProto() *pb.JobCompany {
	if m == nil {
		return nil
	}

	var industries []*pb.Industry
	i := m.Industries
	for _, industry := range i {
		industries = append(industries, industry.ToProto())
	}

	var jobs []*pb.JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, job.ToProto())
	}

	var persons []*pb.KeyPerson
	p := m.KeyPersons
	for _, person := range p {
		persons = append(persons, person.ToProto())
	}

	return &pb.JobCompany{
		Id:         m.ID,
		Name:       m.Name,
		LogoUrl:    m.LogoURL,
		Size:       m.Size,
		Industries: industries,
		JobPosts:   jobs,
		KeyPersons: persons,
	}
}

// ToProto maps the ORM Industry model to the proto model
func (m *Industry) ToProto() *pb.Industry {
	if m == nil {
		return nil
	}

	var companies []*pb.JobCompany
	c := m.Companies
	for _, company := range c {
		companies = append(companies, company.ToProto())
	}

	var jobs []*pb.JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, job.ToProto())
	}

	return &pb.Industry{
		Id:        m.ID,
		Name:      m.Name,
		Companies: companies,
		JobPosts:  jobs,
	}
}

// ToProto maps the ORM KeyPerson model to the proto model
func (m *KeyPerson) ToProto() *pb.KeyPerson {
	if m == nil {
		return nil
	}

	updatedAt := helpers.TimeToProto(m.UpdatedAt)

	return &pb.KeyPerson{
		Id:            m.ID,
		CompanyId:     m.CompanyID,
		Name:          "hubbedin",
		ContactNumber: "+6512345678",
		Email:         "email",
		JobTitle:      "cto",
		UpdatedAt:     updatedAt,
		Company:       m.Company.ToProto(),
	}
}

// ToProto maps the ORM JobPlatform model to the proto model
func (m *JobPlatform) ToProto() *pb.JobPlatform {
	if m == nil {
		return nil
	}

	var jobs []*pb.JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, job.ToProto())
	}

	return &pb.JobPlatform{
		Id:       m.ID,
		Name:     m.Name,
		BaseUrl:  m.BaseURL,
		JobPosts: jobs,
	}
}

// ToProto maps the ORM JobFunction model to the proto model
func (m *JobFunction) ToProto() *pb.JobFunction {
	if m == nil {
		return nil
	}

	return &pb.JobFunction{
		Id:   m.ID,
		Name: m.Name,
	}
}

// ToProto maps the ORM Skill model to the proto model
func (m *Skill) ToProto() *pb.ProfileSkill {
	if m == nil {
		return nil
	}

	return &pb.ProfileSkill{
		Id:   m.ID,
		Name: m.Name,
	}
}
