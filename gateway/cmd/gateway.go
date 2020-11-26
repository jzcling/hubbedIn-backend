package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"

	"in-backend/gateway"
	"in-backend/gateway/configs"
)

var (
	profileEndpoint = flag.String("profile-endpoint", "profile-service:50051", "profile server endpoint")
	projectEndpoint = flag.String("project-endpoint", "project-service:50052", "project server endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// load configs
	cfg, err := configs.LoadConfig(configs.FileName)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux, err := gateway.New(ctx, *profileEndpoint, *projectEndpoint)
	if err != nil {
		glog.Fatal(err)
	}

	srvAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	s := &http.Server{
		Addr:    srvAddr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Info("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Error("Failed to shutdown http server: %v", err)
		}
	}()

	glog.Info("Starting listening at %s", srvAddr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Fatal("Failed to listen and serve: %v", err)
	}
}
