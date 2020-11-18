package models

import (
	"in-backend/services/profile/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestCandidateToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Candidate{
		ID:                     1,
		AuthID:                 "authId",
		FirstName:              "first",
		LastName:               "last",
		Email:                  "email",
		ContactNumber:          "contact",
		Picture:                "picture",
		Gender:                 "male",
		Nationality:            "singapore",
		ResidenceCity:          "singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		LinkedInURL:            "linkedin",
		SCMURL:                 "github",
		WebsiteURL:             "website",
		EducationLevel:         "bachelor",
		Summary:                "summary",
		Birthday:               &testTime,
		NoticePeriod:           1,
		Skills: []*Skill{
			{
				ID:   1,
				Name: "java",
			},
			{
				ID:   2,
				Name: "javascript",
			},
		},
		Academics: []*AcademicHistory{
			{
				ID:            1,
				CandidateID:   1,
				InstitutionID: 1,
				CourseID:      1,
				YearObtained:  2020,
				CreatedAt:     &testTime,
				UpdatedAt:     &testTime,
				DeletedAt:     &testTime,
			},
			{
				ID:            2,
				CandidateID:   1,
				InstitutionID: 1,
				CourseID:      2,
				YearObtained:  2020,
				CreatedAt:     &testTime,
				UpdatedAt:     &testTime,
				DeletedAt:     &testTime,
			},
		},
		Jobs: []*JobHistory{
			{
				ID:             1,
				CandidateID:    1,
				CompanyID:      1,
				DepartmentID:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "software engineer",
				StartDate:      &testTime,
				EndDate:        &testTime,
				SalaryCurrency: "SGD",
				Salary:         1000,
				Description:    "worked hard",
				CreatedAt:      &testTime,
				UpdatedAt:      &testTime,
				DeletedAt:      &testTime,
			},
			{
				ID:             2,
				CandidateID:    1,
				CompanyID:      1,
				DepartmentID:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "senior software engineer",
				StartDate:      &testTime,
				EndDate:        &testTime,
				SalaryCurrency: "SGD",
				Salary:         2000,
				Description:    "worked hard",
				CreatedAt:      &testTime,
				UpdatedAt:      &testTime,
				DeletedAt:      &testTime,
			},
		},
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
		DeletedAt: &testTime,
	}

	expect := &pb.Candidate{
		Id:                     1,
		FirstName:              "first",
		LastName:               "last",
		Email:                  "email",
		ContactNumber:          "contact",
		Gender:                 "male",
		Nationality:            "singapore",
		ResidenceCity:          "singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		LinkedInUrl:            "linkedin",
		ScmUrl:                 "github",
		EducationLevel:         "bachelor",
		Birthday:               testPbTime,
		NoticePeriod:           1,
		Skills: []*pb.Skill{
			{
				Id:   1,
				Name: "java",
			},
			{
				Id:   2,
				Name: "javascript",
			},
		},
		Academics: []*pb.AcademicHistory{
			{
				Id:            1,
				CandidateId:   1,
				InstitutionId: 1,
				CourseId:      1,
				YearObtained:  2020,
				CreatedAt:     testPbTime,
				UpdatedAt:     testPbTime,
				DeletedAt:     testPbTime,
			},
			{
				Id:            2,
				CandidateId:   1,
				InstitutionId: 1,
				CourseId:      2,
				YearObtained:  2020,
				CreatedAt:     testPbTime,
				UpdatedAt:     testPbTime,
				DeletedAt:     testPbTime,
			},
		},
		Jobs: []*pb.JobHistory{
			{
				Id:             1,
				CandidateId:    1,
				CompanyId:      1,
				DepartmentId:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "software engineer",
				StartDate:      testPbTime,
				EndDate:        testPbTime,
				SalaryCurrency: "SGD",
				Salary:         1000,
				Description:    "worked hard",
				CreatedAt:      testPbTime,
				UpdatedAt:      testPbTime,
				DeletedAt:      testPbTime,
			},
			{
				Id:             2,
				CandidateId:    1,
				CompanyId:      1,
				DepartmentId:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "senior software engineer",
				StartDate:      testPbTime,
				EndDate:        testPbTime,
				SalaryCurrency: "SGD",
				Salary:         2000,
				Description:    "worked hard",
				CreatedAt:      testPbTime,
				UpdatedAt:      testPbTime,
				DeletedAt:      testPbTime,
			},
		},
		CreatedAt: testPbTime,
		UpdatedAt: testPbTime,
		DeletedAt: testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestSkillToProto(t *testing.T) {
	input := &Skill{
		ID:   1,
		Name: "skill",
	}

	expect := &pb.Skill{
		Id:   1,
		Name: "skill",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestUserSkillToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &UserSkill{
		ID:          1,
		CandidateID: 1,
		SkillID:     1,
		CreatedAt:   &testTime,
		UpdatedAt:   &testTime,
		DeletedAt:   &testTime,
	}

	expect := &pb.UserSkill{
		Id:          1,
		CandidateId: 1,
		SkillId:     1,
		CreatedAt:   testPbTime,
		UpdatedAt:   testPbTime,
		DeletedAt:   testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestInstitutionToProto(t *testing.T) {
	input := &Institution{
		ID:      1,
		Name:    "institution",
		Country: "singapore",
	}

	expect := &pb.Institution{
		Id:      1,
		Name:    "institution",
		Country: "singapore",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestCourseToProto(t *testing.T) {
	input := &Course{
		ID:    1,
		Name:  "course",
		Level: "bachelor",
	}

	expect := &pb.Course{
		Id:    1,
		Name:  "course",
		Level: "bachelor",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestAcademicHistoryToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &AcademicHistory{
		ID:            1,
		CandidateID:   1,
		InstitutionID: 1,
		CourseID:      1,
		YearObtained:  2020,
		CreatedAt:     &testTime,
		UpdatedAt:     &testTime,
		DeletedAt:     &testTime,
	}

	expect := &pb.AcademicHistory{
		Id:            1,
		CandidateId:   1,
		InstitutionId: 1,
		CourseId:      1,
		YearObtained:  2020,
		CreatedAt:     testPbTime,
		UpdatedAt:     testPbTime,
		DeletedAt:     testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestCompanyToProto(t *testing.T) {
	input := &Company{
		ID:   1,
		Name: "company",
	}

	expect := &pb.Company{
		Id:   1,
		Name: "company",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestDepartmentToProto(t *testing.T) {
	input := &Department{
		ID:   1,
		Name: "department",
	}

	expect := &pb.Department{
		Id:   1,
		Name: "department",
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestJobHistoryToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &JobHistory{
		ID:             1,
		CandidateID:    1,
		CompanyID:      1,
		DepartmentID:   1,
		Country:        "singapore",
		City:           "singapore",
		Title:          "software engineer",
		StartDate:      &testTime,
		EndDate:        &testTime,
		SalaryCurrency: "SGD",
		Salary:         1000,
		Description:    "worked hard",
		CreatedAt:      &testTime,
		UpdatedAt:      &testTime,
		DeletedAt:      &testTime,
	}

	expect := &pb.JobHistory{
		Id:             1,
		CandidateId:    1,
		CompanyId:      1,
		DepartmentId:   1,
		Country:        "singapore",
		City:           "singapore",
		Title:          "software engineer",
		StartDate:      testPbTime,
		EndDate:        testPbTime,
		SalaryCurrency: "SGD",
		Salary:         1000,
		Description:    "worked hard",
		CreatedAt:      testPbTime,
		UpdatedAt:      testPbTime,
		DeletedAt:      testPbTime,
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}
