package models

import (
	"in-backend/helpers"
	"in-backend/services/joblisting/pb"
)

// JobPostToORM maps the proto JobPost model to the ORM model
func JobPostToORM(m *pb.JobPost) *JobPost {
	if m == nil {
		return nil
	}

	var skills []*Skill
	s := m.Skills
	for _, skill := range s {
		skills = append(skills, ProfileSkillToORM(skill))
	}

	createdAt := helpers.ProtoTimeToTime(m.CreatedAt)
	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)
	startAt := helpers.ProtoTimeToTime(m.StartAt)
	expireAt := helpers.ProtoTimeToTime(m.ExpireAt)

	return &JobPost{
		ID:              m.Id,
		CompanyID:       m.CompanyId,
		HRContactID:     m.HrContactId,
		HiringManagerID: m.HiringManagerId,
		JobPlatformID:   m.JobPlatformId,
		Title:           m.Title,
		Description:     m.Description,
		SeniorityLevel:  m.SeniorityLevel,
		YearsExperience: m.YearsExperience,
		EmploymentType:  m.EmploymentType,
		FunctionID:      m.FunctionId,
		IndustryID:      m.IndustryId,
		Location:        m.Location,
		Remote:          m.Remote,
		SalaryCurrency:  m.SalaryCurrency,
		MinSalary:       m.MinSalary,
		MaxSalary:       m.MaxSalary,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		StartAt:         startAt,
		ExpireAt:        expireAt,
		SkillID:         m.SkillId,
		Company:         JobCompanyToORM(m.Company),
		Function:        JobFunctionToORM(m.Function),
		Industry:        IndustryToORM(m.Industry),
		JobPlatform:     JobPlatformToORM(m.JobPlatform),
		HRContact:       KeyPersonToORM(m.HrContact),
		HiringManager:   KeyPersonToORM(m.HiringManager),
		Skills:          skills,
	}
}

// JobCompanyToORM maps the proto JobCompany model to the ORM model
func JobCompanyToORM(m *pb.JobCompany) *Company {
	if m == nil {
		return nil
	}

	var industries []*Industry
	i := m.Industries
	for _, industry := range i {
		industries = append(industries, IndustryToORM(industry))
	}

	var jobs []*JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, JobPostToORM(job))
	}

	var persons []*KeyPerson
	p := m.KeyPersons
	for _, person := range p {
		persons = append(persons, KeyPersonToORM(person))
	}

	return &Company{
		ID:         m.Id,
		Name:       m.Name,
		LogoURL:    m.LogoUrl,
		Size:       m.Size,
		Industries: industries,
		JobPosts:   jobs,
		KeyPersons: persons,
	}
}

// IndustryToORM maps the proto Industry model to the ORM model
func IndustryToORM(m *pb.Industry) *Industry {
	if m == nil {
		return nil
	}

	var companies []*Company
	c := m.Companies
	for _, company := range c {
		companies = append(companies, JobCompanyToORM(company))
	}

	var jobs []*JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, JobPostToORM(job))
	}

	return &Industry{
		ID:        m.Id,
		Name:      m.Name,
		Companies: companies,
		JobPosts:  jobs,
	}
}

// KeyPersonToORM maps the proto KeyPerson model to the ORM model
func KeyPersonToORM(m *pb.KeyPerson) *KeyPerson {
	if m == nil {
		return nil
	}

	updatedAt := helpers.ProtoTimeToTime(m.UpdatedAt)

	return &KeyPerson{
		ID:            m.Id,
		CompanyID:     m.CompanyId,
		Name:          "hubbedin",
		ContactNumber: "+6512345678",
		Email:         "email",
		JobTitle:      "cto",
		UpdatedAt:     updatedAt,
		Company:       JobCompanyToORM(m.Company),
	}
}

// JobPlatformToORM maps the proto JobPlatform model to the ORM model
func JobPlatformToORM(m *pb.JobPlatform) *JobPlatform {
	if m == nil {
		return nil
	}

	var jobs []*JobPost
	j := m.JobPosts
	for _, job := range j {
		jobs = append(jobs, JobPostToORM(job))
	}

	return &JobPlatform{
		ID:       m.Id,
		Name:     m.Name,
		BaseURL:  m.BaseUrl,
		JobPosts: jobs,
	}
}

// JobFunctionToORM maps the proto JobFunction model to the ORM model
func JobFunctionToORM(m *pb.JobFunction) *JobFunction {
	if m == nil {
		return nil
	}

	return &JobFunction{
		ID:   m.Id,
		Name: m.Name,
	}
}

// ProfileSkillToORM maps the proto ProfileSkill model to the ORM model
func ProfileSkillToORM(m *pb.ProfileSkill) *Skill {
	if m == nil {
		return nil
	}

	return &Skill{
		ID:   m.Id,
		Name: m.Name,
	}
}
