package main

import (
	"context"
	"fmt"
	assessmentPb "in-backend/services/assessment/pb"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/gocraft/work"
	"github.com/golang/protobuf/ptypes"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Context struct {
}

const (
	appName           string = "hubbedin"
	assessmentSvcAddr string = "assessment-service:50053"
)

func main() {
	pool := work.NewWorkerPool(Context{}, 10, appName, redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)

	// Map the name of jobs to handler functions
	pool.Job("end_assessment_attempt", (*Context).EndAssessmentAttempt)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

// Log defines the middleware for logging
func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

// EndAssessmentAttempt sets an AssessmentAttempt to Complete
func (c *Context) EndAssessmentAttempt(job *work.Job) error {
	// Extract arguments:
	argID := job.ArgString("id")
	if err := job.ArgError(); err != nil {
		return err
	}

	attemptID, err := strconv.ParseUint(argID, 10, 64)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(assessmentSvcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial Failed: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := assessmentPb.NewAssessmentServiceClient(conn)
	getReq := assessmentPb.GetAssessmentAttemptByIDRequest{Id: attemptID}
	aa, err := client.GetAssessmentAttemptByID(ctx, &getReq)
	if err != nil {
		return err
	}

	if aa.Status != "Completed" {
		aa.CompletedAt = ptypes.TimestampNow()
		aa.Status = "Completed"
		updateReq := assessmentPb.UpdateAssessmentAttemptRequest{
			Id:                attemptID,
			AssessmentAttempt: aa,
		}
		_, err := client.UpdateAssessmentAttempt(ctx, &updateReq)
		return err
	}
	return nil
}
