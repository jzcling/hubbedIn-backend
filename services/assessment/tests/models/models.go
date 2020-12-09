package models

import (
	"in-backend/services/assessment/models"
	"time"
)

var (
	now time.Time = time.Now()

	// AssessmentNoName is a mock Assessment with no Name
	AssessmentNoName = models.Assessment{
		TimeAllowed:  3600,
		Type:         "Multiple Choice",
		Randomise:    true,
		NumQuestions: 10,
	}

	// AssessmentStatusNoAssessmentID is a mock AssessmentStatus with no AssessmentID
	AssessmentStatusNoAssessmentID = models.AssessmentStatus{
		CandidateID: 1,
		Status:      "Completed",
	}

	// QuestionNoType is a mock Question with no Type
	QuestionNoType = models.Question{
		Text:    "What is javascript?",
		Options: []string{"a programming language", "an ancient manuscript"},
		Answer:  0,
	}

	// TagNoName is a mock Tag with no Name
	TagNoName = models.Tag{
		ID: 999,
	}

	// ResponseNoQuestionID is a mock Response with no QuestionID
	ResponseNoQuestionID = models.Response{
		CandidateID: 1,
		Text:        "code",
		Score:       0,
		TimeTaken:   60,
		CreatedAt:   &now,
	}
)
