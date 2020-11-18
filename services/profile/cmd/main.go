package main

import (
	"fmt"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/database"
	"in-backend/services/profile/endpoints"
	"in-backend/services/profile/pb"
	"in-backend/services/profile/providers"
	"in-backend/services/profile/service"
	"in-backend/services/profile/transport"
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
			"svc", "profile",
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
	auth0 := providers.NewAuth0(cfg, client)
	svc := service.New(repo, logger, auth0)
	endpoints := endpoints.MakeEndpoints(svc)

	// set-up grpc transport
	var (
		ocTracing               = kitoc.GRPCServerTrace()
		serverOptions           = []kitgrpc.ServerOption{ocTracing}
		profileService          = transport.NewGRPCServer(endpoints, serverOptions, logger)
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
			pb.RegisterProfileServiceServer(grpcServer, profileService)
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
