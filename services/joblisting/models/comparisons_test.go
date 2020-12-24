package models

import (
	"in-backend/services/joblisting/interfaces"
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestJobPostIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*JobPost)(nil)
	m2 := &JobPost{
		CompanyID:       1,
		HRContactID:     1,
		HiringManagerID: 1,
		JobPlatformID:   1,
		Title:           "software engineer",
		Description:     "description",
		SeniorityLevel:  "entry",
		YearsExperience: 2,
		EmploymentType:  "full-time",
		FunctionID:      1,
		IndustryID:      1,
		Location:        "singapore",
		Remote:          true,
		SalaryCurrency:  "SGD",
		MinSalary:       1000,
		MaxSalary:       2000,
		SkillID:         []uint64{1, 2},
		CreatedAt:       &timeAt,
		UpdatedAt:       &timeAt,
		StartAt:         &timeAt,
		ExpireAt:        &timeAt,
	}
	m3 := &JobPost{}

	testIsEqual(t, m1, m2, m3)
}

func TestCompanyIsEqual(t *testing.T) {
	m1 := (*Company)(nil)
	m2 := &Company{
		Name:    "hubbedin",
		LogoURL: "https://logo.jpg",
		Size:    50,
	}
	m3 := &Company{}

	testIsEqual(t, m1, m2, m3)
}

func TestIndustryIsEqual(t *testing.T) {
	m1 := (*Industry)(nil)
	m2 := &Industry{
		Name: "tech",
	}
	m3 := &Industry{}

	testIsEqual(t, m1, m2, m3)
}

func TestJobFunctionIsEqual(t *testing.T) {
	m1 := (*JobFunction)(nil)
	m2 := &JobFunction{
		Name: "tech",
	}
	m3 := &JobFunction{}

	testIsEqual(t, m1, m2, m3)
}

func TestCompanyIndustryIsEqual(t *testing.T) {
	m1 := (*CompanyIndustry)(nil)
	m2 := &CompanyIndustry{
		CompanyID:  1,
		IndustryID: 1,
	}
	m3 := &Industry{}

	testIsEqual(t, m1, m2, m3)
}

func TestKeyPersonIsEqual(t *testing.T) {
	timeAt := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)
	m1 := (*KeyPerson)(nil)
	m2 := &KeyPerson{
		CompanyID:     1,
		Name:          "name",
		ContactNumber: "+6512345678",
		Email:         "email",
		JobTitle:      "cto",
		UpdatedAt:     &timeAt,
	}
	m3 := &KeyPerson{}

	testIsEqual(t, m1, m2, m3)
}

func TestJobPlatformIsEqual(t *testing.T) {
	m1 := (*JobPlatform)(nil)
	m2 := &JobPlatform{
		Name:    "indeed",
		BaseURL: "https://indeed.com",
	}
	m3 := &JobPlatform{}

	testIsEqual(t, m1, m2, m3)
}

func testIsEqual(t *testing.T, m1, m2 interfaces.Comparator, m3 interface{}) {
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
