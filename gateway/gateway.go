package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	profilegw "in-backend/services/profile/pb"
)

// New creates a new instance of a GRPC gateway
func New(ctx context.Context, profileEndpoint string) (http.Handler, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := profilegw.RegisterProfileServiceHandlerFromEndpoint(ctx, mux, profileEndpoint, opts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
