package models

import (
	"time"
)

// Candidate declares model for Candidate
type Candidate struct {
	ID                     uint64             `json:"id"`
	FirstName              string             `json:"first_name"`
	LastName               string             `json:"last_name"`
	Email                  string             `json:"email"`
	ContactNumber          string             `json:"contact_number"`
	Gender                 string             `json:"gender"`
	Nationality            string             `json:"nationality"`
	ResidenceCity          string             `json:"residence_city"`
	ExpectedSalaryCurrency string             `json:"expected_salary_currency"`
	ExpectedSalary         uint32             `json:"expected_salary"`
	LinkedInURL            string             `json:"linked_in_url"`
	SCMURL                 string             `json:"scm_url"`
	EducationLevel         string             `json:"education_level"`
	Birthday               time.Time          `json:"birthday"`
	NoticePeriod           uint32             `json:"notice_period"`
	Skills                 []*Skill           `json:"skills" pg:"many2many:user_skills"`
	Academics              []*AcademicHistory `json:"academics" pg:"rel:has-many"`
	Jobs                   []*JobHistory      `json:"jobs" pg:"rel:has-many"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
	DeletedAt              time.Time          `json:"deleted_at" pg:",soft_delete"`
}

// Skill declares model for Skill
type Skill struct {
	ID   uint64 `json:"id"`
	Name string `json:"string"`
}

// UserSkill declares model for UserSkill
type UserSkill struct {
	ID          uint64
	CandidateID uint64
	SkillID     uint64
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" pg:",soft_delete"`
}

// Institution declares model for Institution
type Institution struct {
	ID      uint64    `json:"id"`
	Country string    `json:"country"`
	Name    string    `json:"name"`
	Courses []*Course `json:"courses" pg:"rel:has-many"`
}

// Course declares model for Course
type Course struct {
	ID            uint64 `json:"id"`
	InstitutionID uint64 `json:"institution_id"`
	Level         string `json:"level"`
	Name          string `json:"name"`
}

// AcademicHistory declares model for AcademicHistory
type AcademicHistory struct {
	ID            uint64       `json:"id"`
	CandidateID   uint64       `json:"-"`
	Candidate     *Candidate   `json:"candidate" pg:"rel:has-one"`
	InstitutionID uint64       `json:"-"`
	Institution   *Institution `json:"institution" pg:"rel:has-one"`
	CourseID      uint64       `json:"-"`
	Course        *Course      `json:"course" pg:"rel:has-one"`
	YearObtained  uint32       `json:"year_obtained"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     time.Time    `json:"deleted_at" pg:",soft_delete"`
}

// Company declares model for Company
type Company struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Departments []*Department `json:"departments" pg:"rel:has-many"`
}

// Department declares model for Department
type Department struct {
	ID        uint64 `json:"id"`
	CompanyID uint64 `json:"company_id"`
	Name      string `json:"name"`
}

// JobHistory declares model for JobHistory
type JobHistory struct {
	ID             uint64      `json:"id"`
	CandidateID    uint64      `json:"-"`
	Candidate      *Candidate  `json:"candidate" pg:"rel:has-one"`
	CompanyID      uint64      `json:"-"`
	Company        *Company    `json:"company" pg:"rel:has-one"`
	DepartmentID   uint64      `json:"-"`
	Department     *Department `json:"department" pg:"rel:has-one"`
	Country        string      `json:"country"`
	City           string      `json:"city"`
	Title          string      `json:"title"`
	StartDate      time.Time   `json:"start_date"`
	EndDate        time.Time   `json:"end_date"`
	SalaryCurrency string      `json:"salary_currency"`
	Salary         uint32      `json:"salary"`
	Description    string      `json:"description"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      time.Time   `json:"deleted_at" pg:",soft_delete"`
}
