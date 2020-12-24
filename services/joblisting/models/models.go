package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependent models are used.
	orm.RegisterTable((*CompanyIndustry)(nil))
}

// JobPost declares the model for JobPost
type JobPost struct {
	tableName struct{} `pg:"job_posts,alias:jp"`

	ID              uint64       `json:"id"`
	CompanyID       uint64       `json:"company_id" pg:",notnull"`
	HRContactID     uint64       `json:"hr_contact_id"`
	HiringManagerID uint64       `json:"hiring_manager_id"`
	JobPlatformID   uint64       `json:"job_platform_id"`
	Title           string       `json:"title" pg:",notnull"`
	Description     string       `json:"description" pg:",notnull"`
	SeniorityLevel  string       `json:"seniority_level"`
	YearsExperience uint64       `json:"years_experience"`
	EmploymentType  string       `json:"employment_type" pg:",notnull"`
	FunctionID      uint64       `json:"function_id"`
	IndustryID      uint64       `json:"industry_id"`
	Location        string       `json:"location"`
	Remote          bool         `json:"remote"`
	SalaryCurrency  string       `json:"salary_currency"`
	MinSalary       uint64       `json:"min_salary"`
	MaxSalary       uint64       `json:"max_salary"`
	CreatedAt       *time.Time   `json:"created_at"`
	UpdatedAt       *time.Time   `json:"updated_at"`
	StartAt         *time.Time   `json:"start_at" pg:",notnull"`
	ExpireAt        *time.Time   `json:"expire_at" pg:",notnull"`
	SkillID         []uint64     `json:"skill_id" pg:"skill_id,array"`
	Company         *Company     `json:"company" pg:"rel:has-one"`
	Function        *JobFunction `json:"function" pg:"rel:has-one"`
	Industry        *Industry    `json:"industry" pg:"rel:has-one"`
	JobPlatform     *JobPlatform `json:"job_platform" pg:"rel:has-one"`
	HRContact       *KeyPerson   `json:"hr_contact" pg:"hr_contact,rel:has-one"`
	HiringManager   *KeyPerson   `json:"hiring_manager" pg:"hiring_manager,rel:has-one"`
	Skills          []*Skill     `json:"skills" pg:"-"`
}

// BeforeInsert handles the event before a JobPost is inserted into the DB
func (m *JobPost) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.CreatedAt = &now
	m.UpdatedAt = &now

	if m.ExpireAt == nil {
		expire := (*m.StartAt).AddDate(0, 0, 30)
		m.ExpireAt = &expire
	}

	return ctx, nil
}

// BeforeUpdate handles the event before a JobPost is updated in the DB
func (m *JobPost) BeforeUpdate(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.UpdatedAt = &now
	return ctx, nil
}

// Company declares the model for Company
type Company struct {
	tableName struct{} `pg:"companies,alias:co"`

	ID         uint64       `json:"id"`
	Name       string       `json:"name" pg:",notnull,unique"`
	LogoURL    string       `json:"logo_url" pg:"logo_url,notnull"`
	Size       uint64       `json:"size"`
	Industries []*Industry  `json:"industries" pg:"many2many:companies_industries"`
	JobPosts   []*JobPost   `json:"job_posts" pg:"rel:has-many"`
	KeyPersons []*KeyPerson `json:"key_persons" pg:"rel:has-many"`
}

// Industry declares the model for Industry
type Industry struct {
	tableName struct{} `pg:"industries,alias:id"`

	ID        uint64     `json:"id"`
	Name      string     `json:"name" pg:",notnull,unique"`
	Companies []*Company `json:"companies" pg:"many2many:companies_industries"`
	JobPosts  []*JobPost `json:"job_posts" pg:"rel:has-many"`
}

// CompanyIndustry declares the model for CompanyIndustry
type CompanyIndustry struct {
	tableName struct{} `pg:"companies_industries,alias:ci"`

	ID         uint64 `json:"id"`
	CompanyID  uint64 `json:"company_id" pg:"company_id,notnull"`
	IndustryID uint64 `json:"industry_id" pg:"industry_id,notnull"`
}

// JobFunction declares the model for JobFunction
type JobFunction struct {
	tableName struct{} `pg:"job_functions,alias:jf"`

	ID   uint64 `json:"id"`
	Name string `json:"name" pg:",notnull,unique"`
}

// KeyPerson declares the model for KeyPerson
type KeyPerson struct {
	tableName struct{} `pg:"key_persons,alias:kp"`

	ID            uint64     `json:"id"`
	CompanyID     uint64     `json:"company_id" pg:"company_id,notnull"`
	Name          string     `json:"name" pg:",notnull"`
	ContactNumber string     `json:"contact_number"`
	Email         string     `json:"email"`
	JobTitle      string     `json:"job_title"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Company       *Company   `json:"company" pg:"rel:has-one"`
}

// BeforeInsert handles the event before a KeyPerson is inserted into the DB
func (m *KeyPerson) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.UpdatedAt = &now
	return ctx, nil
}

// BeforeUpdate handles the event before a JobPost is updated in the DB
func (m *KeyPerson) BeforeUpdate(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.UpdatedAt = &now
	return ctx, nil
}

// JobPlatform declares the model for JobPlatform
type JobPlatform struct {
	tableName struct{} `pg:"job_platforms,alias:jpf"`

	ID       uint64     `json:"id"`
	Name     string     `json:"name" pg:",notnull"`
	BaseURL  string     `json:"base_url" pg:"base_url,notnull"`
	JobPosts []*JobPost `json:"job_posts" pg:"rel:has-many"`
}

// Skill replicates the model for Skill in Profile Service
type Skill struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
