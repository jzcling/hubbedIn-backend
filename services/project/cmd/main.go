package main

import (
	"fmt"
	"in-backend/services/project/configs"
	"in-backend/services/project/database"
	"in-backend/services/project/endpoints"
	"in-backend/services/project/pb"
	"in-backend/services/project/service"
	"in-backend/services/project/service/middlewares"
	"in-backend/services/project/transport"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load configs
	cfg, err := configs.LoadConfig(configs.FileName)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	// initialize our structured logger for the service
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "project",
			"ts", log.DefaultTimestampUTC,
			"clr", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	opt := database.GetPgConnectionOptions(cfg)
	db := database.NewDatabase(opt)
	defer db.Close()

	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters

	repo := database.NewRepository(db)
	client := &http.Client{}
	svc := service.New(repo, client, logger)
	svc = middlewares.NewAuthMiddleware(svc)
	svc = middlewares.NewLogMiddleware(logger, svc)
	endpoints := endpoints.MakeEndpoints(svc)

	// set-up grpc transport
	var (
		ocTracing               = kitoc.GRPCServerTrace()
		serverOptions           = []kitgrpc.ServerOption{ocTracing}
		projectService          = transport.NewGRPCServer(endpoints, serverOptions, logger)
		grpcListener, listenErr = net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
		grpcServer              = grpc.NewServer()
	)

	if listenErr != nil {
		level.Error(logger).Log("GRPCListener", listenErr)
	}

	var g group.Group
	{
		/*
			Add an actor (function) to the group.
			Each actor must be pre-emptable by an interrupt function.
			That is, if interrupt is invoked, execute should return.
			Also, it must be safe to call interrupt even after execute has returned.
			The first actor (function) to return interrupts all running actors.
			The error is passed to the interrupt functions, and is returned by Run.
		*/
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", cfg.Server.Port)
			pb.RegisterProjectServiceServer(grpcServer, projectService)
			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

	{
		// Set-up our signal handler.
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 1)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	/*
		Run all actors (functions) concurrently. When the first actor returns,
		all others are interrupted. Run only returns when all actors have exited.
		Run returns the error returned by the first exiting actor
	*/
	level.Error(logger).Log("exit", g.Run())
}
