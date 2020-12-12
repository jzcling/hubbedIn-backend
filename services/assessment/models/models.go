package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependent models are used.
	orm.RegisterTable((*AssessmentQuestion)(nil))
	orm.RegisterTable((*AttemptQuestion)(nil))
	orm.RegisterTable((*QuestionTag)(nil))
}

// Assessment declares the model for Assessment
type Assessment struct {
	tableName struct{} `pg:"assessments,alias:a"`

	ID           uint64               `json:"id"`
	Name         string               `json:"name" pg:",notnull,unique"`
	Description  string               `json:"description"`
	Notes        string               `json:"notes"`
	ImageURL     string               `json:"image_url" pg:"image_url"`
	Difficulty   string               `json:"difficulty"`
	TimeAllowed  uint64               `json:"time_allowed"`
	Type         string               `json:"type"`
	Randomise    bool                 `json:"randomise"`
	NumQuestions uint32               `json:"num_questions"`
	Questions    []*Question          `json:"questions,omitempty" pg:"many2many:assessments_questions"`
	Attempts     []*AssessmentAttempt `json:"assessment_attempts,omitempty" pg:"rel:has-many"`
}

// AssessmentAttempt declares the model for AssessmentAttempt
type AssessmentAttempt struct {
	tableName struct{} `pg:"assessment_attempts,alias:aa"`

	ID               uint64             `json:"id"`
	AssessmentID     uint64             `json:"assessment_id" pg:"assessment_id,notnull"`
	CandidateID      uint64             `json:"candidate_id" pg:"candidate_id,notnull"`
	Status           string             `json:"string" pg:",notnull"`
	StartedAt        *time.Time         `json:"started_at,omitempty"`
	CompletedAt      *time.Time         `json:"completed_at,omitempty"`
	Score            uint32             `json:"score,omitempty" pg:",use_zero"`
	Assessment       *Assessment        `json:"assessment" pg:"rel:has-one"`
	Questions        []*Question        `json:"questions" pg:",many2many:attempts_questions,fk:attempt_id,join_fk:question_id"`
	QuestionAttempts []*AttemptQuestion `json:"question_attempts" pg:"rel:has-many,join_fk:attempt_id"`
}

// Question declares the model for Question
type Question struct {
	tableName struct{} `pg:"questions,alias:q"`

	ID                 uint64               `json:"id"`
	CreatedBy          uint64               `json:"created_by"`
	Type               string               `json:"type" pg:",notnull"`
	Text               string               `json:"text"`
	ImageURL           string               `json:"image_url" pg:"image_url"`
	Options            []string             `json:"options" pg:",array"`
	Answer             uint32               `json:"answer" pg:",use_zero"`
	Tags               []*Tag               `json:"tags" pg:"rel:has-many"`
	Assessments        []*Assessment        `json:"assessments" pg:"many2many:assessments_questions"`
	AssessmentAttempts []*AssessmentAttempt `json:"assessment_attempts" pg:",many2many:attempts_questions,fk:question_id,join_fk:attempt_id"`
	Attempts           []*AttemptQuestion   `json:"attempts" pg:"rel:has-many"`
}

// AttemptQuestion declares the model for AttemptQuestion
type AttemptQuestion struct {
	tableName struct{} `pg:"attempts_questions,alias:aaq"`

	ID          uint64     `json:"id"`
	AttemptID   uint64     `json:"attempt_id" pg:"attempt_id,notnull"`
	QuestionID  uint64     `json:"question_id" pg:"question_id,notnull"`
	CandidateID uint64     `json:"candidate_id" pg:"candidate_id,notnull"`
	Selection   uint32     `json:"selection" pg:",use_zero"`
	Text        string     `json:"text"`
	Score       uint32     `json:"score,omitempty" pg:",use_zero"`
	TimeTaken   uint64     `json:"time_taken,omitempty"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (m *AttemptQuestion) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.CreatedAt = &now
	m.UpdatedAt = &now
	return ctx, nil
}

func (m *AttemptQuestion) BeforeUpdate(ctx context.Context) (context.Context, error) {
	now := time.Now()
	m.UpdatedAt = &now
	return ctx, nil
}

// Tag declares the model for Tag
type Tag struct {
	tableName struct{} `pg:"tags,alias:t"`

	ID   uint64 `json:"id"`
	Name string `json:"name" pg:",notnull,unique"`
}

// QuestionTag declares the model for QuestionTag
type QuestionTag struct {
	tableName struct{} `pg:"questions_tags,alias:qt"`

	ID         uint64 `json:"id"`
	QuestionID uint64 `json:"question_id" pg:"question_id,notnull"`
	TagID      uint64 `json:"tag_id" pg:"tag_id,notnull"`
}

// AssessmentQuestion declares the model for the pivot AssessmentQuestion
type AssessmentQuestion struct {
	tableName struct{} `pg:"assessments_questions,alias:aq"`

	ID           uint64 `json:"id"`
	AssessmentID uint64 `json:"assessment_id" pg:"assessment_id,notnull"`
	QuestionID   uint64 `json:"question_id" pg:"question_id,notnull"`
}
