package main

import (
	"context"
	"errors"
	"fmt"
	"in-backend/gateway"
	"net/http"
)

// GRPCRegisterer registers the grpc gateway
var GRPCRegisterer = registerer("grpc-gateway")

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), func(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
		cfg := parse(extra)
		if cfg == nil {
			return nil, errors.New("wrong config")
		}
		if cfg.name != string(r) {
			return nil, fmt.Errorf("unknown register %s", cfg.name)
		}
		return gateway.New(ctx, cfg.profileEndpoint, cfg.projectEndpoint, cfg.assessmentEndpoint, cfg.joblistingEndpoint)
	})
}

func parse(extra map[string]interface{}) *opts {
	name, ok := extra["name"].(string)
	if !ok {
		return nil
	}

	rawEs, ok := extra["endpoints"]
	if !ok {
		return nil
	}
	es, ok := rawEs.([]interface{})
	if !ok || len(es) < 2 {
		return nil
	}
	endpoints := make([]string, len(es))
	for i, e := range es {
		endpoints[i] = e.(string)
	}

	return &opts{
		name:               name,
		profileEndpoint:    endpoints[0],
		projectEndpoint:    endpoints[1],
		assessmentEndpoint: endpoints[2],
		joblistingEndpoint: endpoints[3],
	}
}

type opts struct {
	name               string
	profileEndpoint    string
	projectEndpoint    string
	assessmentEndpoint string
	joblistingEndpoint string
}
