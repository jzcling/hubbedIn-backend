package models

import (
	"in-backend/services/joblisting/models"
	"time"
)

var (
	now time.Time = time.Now()

	// JobPostNoTitle is a mock JobPost with no Title
	JobPostNoTitle = models.JobPost{
		CompanyID:      1,
		Description:    "description",
		EmploymentType: "full-time",
		StartAt:        &now,
		ExpireAt:       &now,
	}
)
