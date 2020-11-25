package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestProjectIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)

	m1 := (*Project)(nil)
	m2 := &Project{
		Name:      "name",
		RepoURL:   "repo",
		CreatedAt: &timeAt,
		UpdatedAt: &timeAt,
		DeletedAt: &timeAt,
	}

	testIsEqual(t, m1, m2)
}

func TestRatingIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*Rating)(nil)
	m2 := &Rating{
		ProjectID:             1,
		ReliabilityRating:     1,
		MaintainabilityRating: 1,
		SecurityRating:        1,
		SecurityReviewRating:  1,
		Coverage:              1.0,
		Duplications:          1.0,
		Lines:                 1,
		CreatedAt:             &timeAt,
	}

	testIsEqual(t, m1, m2)
}

func TestCandidateProjectIsEqual(t *testing.T) {
	m1 := (*CandidateProject)(nil)
	m2 := &CandidateProject{
		CandidateID: 1,
		ProjectID:   2,
	}

	testIsEqual(t, m1, m2)
}

func testIsEqual(t *testing.T, m1, m2 Comparator) {
	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	var emptyStruct struct{}
	m3 := &emptyStruct
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
				fmt.Println("id")
				assert.Condition(t, func() bool { return m2.IsEqual(m3) })
			}

			copier.Copy(m3, m2)
		}
	}
}
