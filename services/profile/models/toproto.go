package models

import (
	"in-backend/services/profile/pb"

	"github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
)

// ToProto maps the ORM Candidate model to the proto model
func (c *Candidate) ToProto() *pb.Candidate {
	if c == nil {
		return nil
	}

	var skills []*pb.Skill
	s := c.Skills
	for _, skill := range s {
		skills = append(skills, skill.ToProto())
	}

	var academics []*pb.AcademicHistory
	a := c.Academics
	for _, academic := range a {
		academics = append(academics, academic.ToProto())
	}

	var jobs []*pb.JobHistory
	j := c.Jobs
	for _, job := range j {
		jobs = append(jobs, job.ToProto())
	}

	var err error
	var birthday *timestamppb.Timestamp
	if c.Birthday != nil {
		birthday, err = ptypes.TimestampProto(*c.Birthday)
		if err != nil {
			birthday = nil
		}
	}

	var createdAt *timestamppb.Timestamp
	if c.CreatedAt != nil {
		createdAt, err = ptypes.TimestampProto(*c.CreatedAt)
		if err != nil {
			createdAt = nil
		}
	}

	var updatedAt *timestamppb.Timestamp
	if c.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*c.UpdatedAt)
		if err != nil {
			updatedAt = nil
		}
	}

	var deletedAt *timestamppb.Timestamp
	if c.DeletedAt != nil {
		deletedAt, err = ptypes.TimestampProto(*c.DeletedAt)
		if err != nil {
			deletedAt = nil
		}
	}

	return &pb.Candidate{
		Id:                     c.ID,
		AuthId:                 c.AuthID,
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
		LinkedInUrl:            c.LinkedInURL,
		ScmUrl:                 c.SCMURL,
		WebsiteUrl:             c.WebsiteURL,
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

// ToProto maps the ORM Skill model to the proto model
func (s *Skill) ToProto() *pb.Skill {
	if s == nil {
		return nil
	}
	return &pb.Skill{
		Id:   s.ID,
		Name: s.Name,
	}
}

// ToProto maps the ORM UserSkill model to the proto model
func (us *UserSkill) ToProto() *pb.UserSkill {
	if us == nil {
		return nil
	}

	var err error
	var createdAt *timestamppb.Timestamp
	if us.CreatedAt != nil {
		createdAt, err = ptypes.TimestampProto(*us.CreatedAt)
		if err != nil {
			createdAt = nil
		}
	}

	var updatedAt *timestamppb.Timestamp
	if us.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*us.UpdatedAt)
		if err != nil {
			updatedAt = nil
		}
	}

	return &pb.UserSkill{
		Id:          us.ID,
		CandidateId: us.CandidateID,
		SkillId:     us.SkillID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// ToProto maps the ORM Institution model to the proto model
func (i *Institution) ToProto() *pb.Institution {
	if i == nil {
		return nil
	}
	return &pb.Institution{
		Id:      i.ID,
		Country: i.Country,
		Name:    i.Name,
	}
}

// ToProto maps the ORM Course model to the proto model
func (c *Course) ToProto() *pb.Course {
	if c == nil {
		return nil
	}
	return &pb.Course{
		Id:    c.ID,
		Level: c.Level,
		Name:  c.Name,
	}
}

// ToProto maps the ORM AcademicHistory model to the proto model
func (a *AcademicHistory) ToProto() *pb.AcademicHistory {
	if a == nil {
		return nil
	}

	var err error
	var createdAt *timestamppb.Timestamp
	if a.CreatedAt != nil {
		createdAt, err = ptypes.TimestampProto(*a.CreatedAt)
		if err != nil {
			createdAt = nil
		}
	}

	var updatedAt *timestamppb.Timestamp
	if a.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*a.UpdatedAt)
		if err != nil {
			updatedAt = nil
		}
	}

	var deletedAt *timestamppb.Timestamp
	if a.DeletedAt != nil {
		deletedAt, err = ptypes.TimestampProto(*a.DeletedAt)
		if err != nil {
			deletedAt = nil
		}
	}

	return &pb.AcademicHistory{
		Id:            a.ID,
		CandidateId:   a.CandidateID,
		InstitutionId: a.InstitutionID,
		Institution:   a.Institution.ToProto(),
		CourseId:      a.CourseID,
		Course:        a.Course.ToProto(),
		YearObtained:  a.YearObtained,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}
}

// ToProto maps the ORM Company model to the proto model
func (c *Company) ToProto() *pb.Company {
	if c == nil {
		return nil
	}
	return &pb.Company{
		Id:   c.ID,
		Name: c.Name,
	}
}

// ToProto maps the ORM Department model to the proto model
func (d *Department) ToProto() *pb.Department {
	if d == nil {
		return nil
	}
	return &pb.Department{
		Id:   d.ID,
		Name: d.Name,
	}
}

// ToProto maps the ORM JobHistory model to the proto model
func (j *JobHistory) ToProto() *pb.JobHistory {
	if j == nil {
		return nil
	}
	var err error
	var startDate *timestamppb.Timestamp
	if j.StartDate != nil {
		startDate, err = ptypes.TimestampProto(*j.StartDate)
		if err != nil {
			startDate = nil
		}
	}

	var endDate *timestamppb.Timestamp
	if j.EndDate != nil {
		endDate, err = ptypes.TimestampProto(*j.EndDate)
		if err != nil {
			endDate = nil
		}
	}

	var createdAt *timestamppb.Timestamp
	if j.CreatedAt != nil {
		createdAt, err = ptypes.TimestampProto(*j.CreatedAt)
		if err != nil {
			createdAt = nil
		}
	}

	var updatedAt *timestamppb.Timestamp
	if j.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*j.UpdatedAt)
		if err != nil {
			updatedAt = nil
		}
	}

	var deletedAt *timestamppb.Timestamp
	if j.DeletedAt != nil {
		deletedAt, err = ptypes.TimestampProto(*j.DeletedAt)
		if err != nil {
			deletedAt = nil
		}
	}

	return &pb.JobHistory{
		Id:             j.ID,
		CandidateId:    j.CandidateID,
		CompanyId:      j.CompanyID,
		Company:        j.Company.ToProto(),
		DepartmentId:   j.DepartmentID,
		Department:     j.Department.ToProto(),
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
