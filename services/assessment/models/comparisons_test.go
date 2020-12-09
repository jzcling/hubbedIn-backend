package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestAssessmentIsEqual(t *testing.T) {
	m1 := (*Assessment)(nil)
	m2 := &Assessment{
		Name:         "Javascript",
		Description:  "JS Test",
		Notes:        "Notes",
		ImageURL:     "image",
		Difficulty:   "Easy",
		TimeAllowed:  3600,
		Type:         "Multiple Choice",
		Randomise:    true,
		NumQuestions: 10,
	}
	m3 := &Assessment{}

	testIsEqual(t, m1, m2, m3)
}

func TestAssessmentStatusIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*AssessmentStatus)(nil)
	m2 := &AssessmentStatus{
		AssessmentID: 1,
		CandidateID:  1,
		Status:       "Completed",
		StartedAt:    &timeAt,
		CompletedAt & timeAt,
		Score: 5,
	}
	m3 := &AssessmentStatus{}

	testIsEqual(t, m1, m2, m3)
}

func TestQuestionIsEqual(t *testing.T) {
	m1 := (*Question)(nil)
	m2 := &Question{
		CreatedBy: 1,
		Type:      "Open",
		Text:      "text",
		ImageURL:  "image",
		Options:   []string{"test", "test2"},
		Answer:    0,
	}
	m3 := &Question{}

	testIsEqual(t, m1, m2, m3)
}

func TestTagIsEqual(t *testing.T) {
	m1 := (*Tag)(nil)
	m2 := &Tag{
		Name: "javascript",
	}
	m3 := &Tag{}

	testIsEqual(t, m1, m2, m3)
}

func TestQuestionTagIsEqual(t *testing.T) {
	m1 := (*QuestionTag)(nil)
	m2 := &QuestionTag{
		QuestionID: 1,
		TagID:      1,
	}
	m3 := &QuestionTag{}

	testIsEqual(t, m1, m2, m3)
}

func TestResponseIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*Response)(nil)
	m2 := &Response{
		QuestionID:  1,
		CandidateID: 1,
		Selection:   0,
		Text:        "text",
		Score:       0,
		TimeTaken:   10,
		CreatedAt:   &timeAt,
	}
	m3 := &Response{}

	testIsEqual(t, m1, m2, m3)
}

func TestAssessmentQuestionIsEqual(t *testing.T) {
	m1 := (*AssessmentQuestion)(nil)
	m2 := &AssessmentQuestion{
		AssessmentID: 1,
		QuestionID:   1,
	}
	m3 := &AssessmentQuestion{}

	testIsEqual(t, m1, m2, m3)
}

func testIsEqual(t *testing.T, m1, m2 Comparator, m3 interface{}) {
	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	copier.Copy(m3, m2)
	values := reflect.ValueOf(m3).Elem()
	for i := 0; i < values.NumField(); i++ {
		v := values.Field(i)
		if v.CanSet() {
			changed := false
			switch v.Interface().(type) {
			case string:
				v.SetString("string")
				changed = true
			case uint64, uint32:
				v.SetUint(999)
				changed = true
			case *time.Time:
				now := time.Now()
				v.Set(reflect.ValueOf(&now))
				changed = true
			}

			fieldName := values.Type().Field(i).Name
			if fieldName != "ID" && changed {
				assert.Condition(t, func() bool { return !m2.IsEqual(m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(m3) })
			}

			copier.Copy(m3, m2)
		}
	}
}
