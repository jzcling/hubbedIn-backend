package models

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependant models are used.
	orm.RegisterTable((*UserSkill)(nil))
}

// Candidate declares model for Candidate
type Candidate struct {
	tableName struct{} `pg:"candidates,alias:c"`

	ID                     uint64             `json:"id"`
	FirstName              string             `json:"first_name"`
	LastName               string             `json:"last_name"`
	Email                  string             `json:"email" pg:",unique"`
	ContactNumber          string             `json:"contact_number"`
	Gender                 string             `json:"gender,omitempty"`
	Nationality            string             `json:"nationality,omitempty"`
	ResidenceCity          string             `json:"residence_city,omitempty"`
	ExpectedSalaryCurrency string             `json:"expected_salary_currency,omitempty"`
	ExpectedSalary         uint32             `json:"expected_salary,omitempty"`
	LinkedInURL            string             `json:"linked_in_url,omitempty"`
	SCMURL                 string             `json:"scm_url,omitempty" pg:"scm_url"`
	EducationLevel         string             `json:"education_level,omitempty"`
	Birthday               *time.Time         `json:"birthday,omitempty"`
	NoticePeriod           uint32             `json:"notice_period,omitempty"`
	Skills                 []*Skill           `json:"skills,omitempty" pg:"many2many:users_skills"`
	Academics              []*AcademicHistory `json:"academics,omitempty" pg:"rel:has-many"`
	Jobs                   []*JobHistory      `json:"jobs,omitempty" pg:"rel:has-many"`
	CreatedAt              *time.Time         `json:"created_at,omitempty"`
	UpdatedAt              *time.Time         `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt              *time.Time         `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// Skill declares model for Skill
type Skill struct {
	tableName struct{} `pg:"skills,alias:s"`

	ID         uint64       `json:"id"`
	Name       string       `json:"string"`
	Candidates []*Candidate `json:"candidates,omitempty" pg:"many2many:users_skills"`
}

// UserSkill declares model for UserSkill
type UserSkill struct {
	tableName struct{} `pg:"users_skills,alias:us"`

	ID          uint64
	CandidateID uint64
	SkillID     uint64
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// Institution declares model for Institution
type Institution struct {
	tableName struct{} `pg:"institutions,alias:i"`

	ID      uint64    `json:"id"`
	Country string    `json:"country,omitempty"`
	Name    string    `json:"name"`
	Courses []*Course `json:"courses,omitempty" pg:"rel:has-many"`
}

// Course declares model for Course
type Course struct {
	tableName struct{} `pg:"courses,alias:cr"`

	ID            uint64 `json:"id"`
	InstitutionID uint64 `json:"institution_id,omitempty"`
	Level         string `json:"level,omitempty"`
	Name          string `json:"name"`
}

// AcademicHistory declares model for AcademicHistory
type AcademicHistory struct {
	tableName struct{} `pg:"academic_histories,alias:ah"`

	ID            uint64       `json:"id"`
	CandidateID   uint64       `json:"-"`
	Candidate     *Candidate   `json:"candidate,omitempty" pg:"rel:has-one"`
	InstitutionID uint64       `json:"-"`
	Institution   *Institution `json:"institution,omitempty" pg:"rel:has-one"`
	CourseID      uint64       `json:"-"`
	Course        *Course      `json:"course,omitempty" pg:"rel:has-one"`
	YearObtained  uint32       `json:"year_obtained,omitempty"`
	CreatedAt     *time.Time   `json:"created_at,omitempty"`
	UpdatedAt     *time.Time   `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt     *time.Time   `json:"deleted_at,omitempty" pg:",soft_delete"`
}

// Company declares model for Company
type Company struct {
	tableName struct{} `pg:"companies,alias:co"`

	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Departments []*Department `json:"departments,omitempty" pg:"rel:has-many"`
}

// Department declares model for Department
type Department struct {
	tableName struct{} `pg:"departments,alias:d"`

	ID        uint64 `json:"id"`
	CompanyID uint64 `json:"company_id,omitempty"`
	Name      string `json:"name"`
}

// JobHistory declares model for JobHistory
type JobHistory struct {
	tableName struct{} `pg:"job_histories,alias:jh"`

	ID             uint64      `json:"id"`
	CandidateID    uint64      `json:"-"`
	Candidate      *Candidate  `json:"candidate,omitempty" pg:"rel:has-one"`
	CompanyID      uint64      `json:"-"`
	Company        *Company    `json:"company,omitempty" pg:"rel:has-one"`
	DepartmentID   uint64      `json:"-"`
	Department     *Department `json:"department,omitempty" pg:"rel:has-one"`
	Country        string      `json:"country"`
	City           string      `json:"city,omitempty"`
	Title          string      `json:"title"`
	StartDate      *time.Time  `json:"start_date,omitempty"`
	EndDate        *time.Time  `json:"end_date,omitempty"`
	SalaryCurrency string      `json:"salary_currency,omitempty"`
	Salary         uint32      `json:"salary,omitempty"`
	Description    string      `json:"description,omitempty"`
	CreatedAt      *time.Time  `json:"created_at,omitempty"`
	UpdatedAt      *time.Time  `json:"updated_at,omitempty" pg:"default:now()"`
	DeletedAt      *time.Time  `json:"deleted_at,omitempty" pg:",soft_delete"`
}
