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

	// AssessmentAttemptNoAssessmentID is a mock AssessmentAttempt with no AssessmentID
	AssessmentAttemptNoAssessmentID = models.AssessmentAttempt{
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

	// AttemptQuestionNoQuestionID is a mock AttemptQuestion with no QuestionID
	AttemptQuestionNoQuestionID = models.AttemptQuestion{
		AttemptID:   1,
		CandidateID: 1,
		Text:        "code",
		Score:       0,
		TimeTaken:   60,
		CreatedAt:   &now,
	}
)
