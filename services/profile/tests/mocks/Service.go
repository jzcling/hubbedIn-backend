// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "in-backend/services/profile/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateAcademicHistory provides a mock function with given fields: ctx, a
func (_m *Service) CreateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	ret := _m.Called(ctx, a)

	var r0 *models.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, *models.AcademicHistory) *models.AcademicHistory); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.AcademicHistory) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCandidate provides a mock function with given fields: ctx, c
func (_m *Service) CreateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	ret := _m.Called(ctx, c)

	var r0 *models.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, *models.Candidate) *models.Candidate); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Candidate) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCompany provides a mock function with given fields: ctx, c
func (_m *Service) CreateCompany(ctx context.Context, c *models.Company) (*models.Company, error) {
	ret := _m.Called(ctx, c)

	var r0 *models.Company
	if rf, ok := ret.Get(0).(func(context.Context, *models.Company) *models.Company); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Company)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Company) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCourse provides a mock function with given fields: ctx, c
func (_m *Service) CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error) {
	ret := _m.Called(ctx, c)

	var r0 *models.Course
	if rf, ok := ret.Get(0).(func(context.Context, *models.Course) *models.Course); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Course) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDepartment provides a mock function with given fields: ctx, c
func (_m *Service) CreateDepartment(ctx context.Context, c *models.Department) (*models.Department, error) {
	ret := _m.Called(ctx, c)

	var r0 *models.Department
	if rf, ok := ret.Get(0).(func(context.Context, *models.Department) *models.Department); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Department)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Department) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInstitution provides a mock function with given fields: ctx, i
func (_m *Service) CreateInstitution(ctx context.Context, i *models.Institution) (*models.Institution, error) {
	ret := _m.Called(ctx, i)

	var r0 *models.Institution
	if rf, ok := ret.Get(0).(func(context.Context, *models.Institution) *models.Institution); ok {
		r0 = rf(ctx, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Institution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Institution) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJobHistory provides a mock function with given fields: ctx, a
func (_m *Service) CreateJobHistory(ctx context.Context, a *models.JobHistory) (*models.JobHistory, error) {
	ret := _m.Called(ctx, a)

	var r0 *models.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, *models.JobHistory) *models.JobHistory); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.JobHistory) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSkill provides a mock function with given fields: ctx, s
func (_m *Service) CreateSkill(ctx context.Context, s *models.Skill) (*models.Skill, error) {
	ret := _m.Called(ctx, s)

	var r0 *models.Skill
	if rf, ok := ret.Get(0).(func(context.Context, *models.Skill) *models.Skill); ok {
		r0 = rf(ctx, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Skill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Skill) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, m
func (_m *Service) CreateUser(ctx context.Context, m *models.User) (*models.User, error) {
	ret := _m.Called(ctx, m)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.User); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserSkill provides a mock function with given fields: ctx, us
func (_m *Service) CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error) {
	ret := _m.Called(ctx, us)

	var r0 *models.UserSkill
	if rf, ok := ret.Get(0).(func(context.Context, *models.UserSkill) *models.UserSkill); ok {
		r0 = rf(ctx, us)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserSkill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.UserSkill) error); ok {
		r1 = rf(ctx, us)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcademicHistory provides a mock function with given fields: ctx, cid, ahid
func (_m *Service) DeleteAcademicHistory(ctx context.Context, cid uint64, ahid uint64) error {
	ret := _m.Called(ctx, cid, ahid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, cid, ahid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCandidate provides a mock function with given fields: ctx, id
func (_m *Service) DeleteCandidate(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJobHistory provides a mock function with given fields: ctx, cid, jhid
func (_m *Service) DeleteJobHistory(ctx context.Context, cid uint64, jhid uint64) error {
	ret := _m.Called(ctx, cid, jhid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, cid, jhid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *Service) DeleteUser(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserSkill provides a mock function with given fields: ctx, cid, sid
func (_m *Service) DeleteUserSkill(ctx context.Context, cid uint64, sid uint64) error {
	ret := _m.Called(ctx, cid, sid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, cid, sid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcademicHistory provides a mock function with given fields: ctx, id
func (_m *Service) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.AcademicHistory); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCandidates provides a mock function with given fields: ctx, f
func (_m *Service) GetAllCandidates(ctx context.Context, f models.CandidateFilters) ([]*models.User, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.CandidateFilters) []*models.User); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CandidateFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCompanies provides a mock function with given fields: ctx, f
func (_m *Service) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Company
	if rf, ok := ret.Get(0).(func(context.Context, models.CompanyFilters) []*models.Company); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Company)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CompanyFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCourses provides a mock function with given fields: ctx, f
func (_m *Service) GetAllCourses(ctx context.Context, f models.CourseFilters) ([]*models.Course, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Course
	if rf, ok := ret.Get(0).(func(context.Context, models.CourseFilters) []*models.Course); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CourseFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDepartments provides a mock function with given fields: ctx, f
func (_m *Service) GetAllDepartments(ctx context.Context, f models.DepartmentFilters) ([]*models.Department, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Department
	if rf, ok := ret.Get(0).(func(context.Context, models.DepartmentFilters) []*models.Department); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Department)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.DepartmentFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllInstitutions provides a mock function with given fields: ctx, f
func (_m *Service) GetAllInstitutions(ctx context.Context, f models.InstitutionFilters) ([]*models.Institution, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Institution
	if rf, ok := ret.Get(0).(func(context.Context, models.InstitutionFilters) []*models.Institution); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Institution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.InstitutionFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSkills provides a mock function with given fields: ctx, f
func (_m *Service) GetAllSkills(ctx context.Context, f models.SkillFilters) ([]*models.Skill, error) {
	ret := _m.Called(ctx, f)

	var r0 []*models.Skill
	if rf, ok := ret.Get(0).(func(context.Context, models.SkillFilters) []*models.Skill); ok {
		r0 = rf(ctx, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Skill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SkillFilters) error); ok {
		r1 = rf(ctx, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCandidateByID provides a mock function with given fields: ctx, id
func (_m *Service) GetCandidateByID(ctx context.Context, id uint64) (*models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCompany provides a mock function with given fields: ctx, id
func (_m *Service) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Company
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Company); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Company)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCourse provides a mock function with given fields: ctx, id
func (_m *Service) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Course
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Course); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDepartment provides a mock function with given fields: ctx, id
func (_m *Service) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Department
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Department); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Department)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstitution provides a mock function with given fields: ctx, id
func (_m *Service) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Institution
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Institution); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Institution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobHistory provides a mock function with given fields: ctx, id
func (_m *Service) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.JobHistory); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSkill provides a mock function with given fields: ctx, id
func (_m *Service) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Skill
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Skill); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Skill)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcademicHistory provides a mock function with given fields: ctx, a
func (_m *Service) UpdateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error) {
	ret := _m.Called(ctx, a)

	var r0 *models.AcademicHistory
	if rf, ok := ret.Get(0).(func(context.Context, *models.AcademicHistory) *models.AcademicHistory); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AcademicHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.AcademicHistory) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCandidate provides a mock function with given fields: ctx, c
func (_m *Service) UpdateCandidate(ctx context.Context, c *models.Candidate) (*models.Candidate, error) {
	ret := _m.Called(ctx, c)

	var r0 *models.Candidate
	if rf, ok := ret.Get(0).(func(context.Context, *models.Candidate) *models.Candidate); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Candidate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Candidate) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobHistory provides a mock function with given fields: ctx, j
func (_m *Service) UpdateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error) {
	ret := _m.Called(ctx, j)

	var r0 *models.JobHistory
	if rf, ok := ret.Get(0).(func(context.Context, *models.JobHistory) *models.JobHistory); ok {
		r0 = rf(ctx, j)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.JobHistory) error); ok {
		r1 = rf(ctx, j)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, m
func (_m *Service) UpdateUser(ctx context.Context, m *models.User) (*models.User, error) {
	ret := _m.Called(ctx, m)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.User); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
