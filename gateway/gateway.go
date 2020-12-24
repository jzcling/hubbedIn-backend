package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	assessmentgw "in-backend/services/assessment/pb"
	joblistinggw "in-backend/services/joblisting/pb"
	profilegw "in-backend/services/profile/pb"
	projectgw "in-backend/services/project/pb"
)

// New creates a new instance of a GRPC gateway
func New(ctx context.Context, profileEndpoint, projectEndpoint, assessmentEndpoint, joblistingEndpoint string) (http.Handler, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := profilegw.RegisterProfileServiceHandlerFromEndpoint(ctx, mux, profileEndpoint, opts)
	if err != nil {
		return nil, err
	}
	err = projectgw.RegisterProjectServiceHandlerFromEndpoint(ctx, mux, projectEndpoint, opts)
	if err != nil {
		return nil, err
	}
	err = joblistinggw.RegisterJoblistingServiceHandlerFromEndpoint(ctx, mux, joblistingEndpoint, opts)
	if err != nil {
		return nil, err
	}
	err = assessmentgw.RegisterAssessmentServiceHandlerFromEndpoint(ctx, mux, assessmentEndpoint, opts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
