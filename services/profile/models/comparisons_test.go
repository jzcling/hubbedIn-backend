package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCandidateIsEqual(t *testing.T) {
	birthday := time.Date(1990, 1, 5, 0, 0, 0, 0, time.Local)
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)

	m1 := (*Candidate)(nil)
	m2 := &Candidate{
		FirstName:              "first",
		LastName:               "last",
		Email:                  "new@email.com",
		ContactNumber:          "+6563210987",
		Picture:                "picture",
		Gender:                 "male",
		Nationality:            "singapore",
		ResidenceCity:          "singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		LinkedInURL:            "https://www.linkedin.com/in/williamhgates",
		SCMURL:                 "https://github.com/williamhgates",
		WebsiteURL:             "https://billgates.com",
		EducationLevel:         "bachelor",
		Summary:                "summary",
		Birthday:               &birthday,
		NoticePeriod:           2,
		CreatedAt:              &timeAt,
		UpdatedAt:              &timeAt,
		DeletedAt:              &timeAt,
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestSkillIsEqual(t *testing.T) {
	m1 := (*Skill)(nil)
	m2 := &Skill{
		Name: "java",
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestInstitutionIsEqual(t *testing.T) {
	m1 := (*Institution)(nil)
	m2 := &Institution{
		Name:    "national university of singapore",
		Country: "singapore",
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestCourseIsEqual(t *testing.T) {
	m1 := (*Course)(nil)
	m2 := &Course{
		Name:  "computer science",
		Level: "bachelor",
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestAcademicHistoryIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*AcademicHistory)(nil)
	m2 := &AcademicHistory{
		CandidateID:   1,
		InstitutionID: 1,
		CourseID:      1,
		YearObtained:  2020,
		CreatedAt:     &timeAt,
		UpdatedAt:     &timeAt,
		DeletedAt:     &timeAt,
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestCompanyIsEqual(t *testing.T) {
	m1 := (*Company)(nil)
	m2 := &Company{
		Name: "hubbed",
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestDepartmentIsEqual(t *testing.T) {
	m1 := (*Department)(nil)
	m2 := &Department{
		Name: "tech",
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}

func TestJobHistoryIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*JobHistory)(nil)
	m2 := &JobHistory{
		CandidateID:    1,
		CompanyID:      1,
		DepartmentID:   1,
		Country:        "singapore",
		City:           "singapore",
		Title:          "software engineer",
		StartDate:      &timeAt,
		EndDate:        &timeAt,
		SalaryCurrency: "SGD",
		Salary:         1000,
		Description:    "worked hard",
		CreatedAt:      &timeAt,
		UpdatedAt:      &timeAt,
		DeletedAt:      &timeAt,
	}

	assert.Condition(t, func() bool { return m1.IsEqual(m1) })
	assert.Condition(t, func() bool { return !m1.IsEqual(m2) })

	m3 := *m2
	values := reflect.ValueOf(&m3).Elem()
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
				assert.Condition(t, func() bool { return !m2.IsEqual(&m3) })
			}

			if fieldName == "ID" {
				assert.Condition(t, func() bool { return m2.IsEqual(&m3) })
			}

			m3 = *m2
		}
	}
}
