// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "todo/pkg/endpoint"
	http1 "todo/pkg/http"
	service "todo/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Add":     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Add", logger))},
		"Delete":  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Delete", logger))},
		"Get":     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Get", logger))},
		"GetById": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetById", logger))},
		"Update":  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Update", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Get"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Get")), endpoint.InstrumentingMiddleware(duration.With("method", "Get"))}
	mw["Add"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Add")), endpoint.InstrumentingMiddleware(duration.With("method", "Add"))}
	mw["Update"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Update")), endpoint.InstrumentingMiddleware(duration.With("method", "Update"))}
	mw["Delete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Delete")), endpoint.InstrumentingMiddleware(duration.With("method", "Delete"))}
	mw["GetById"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetById")), endpoint.InstrumentingMiddleware(duration.With("method", "GetById"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Get", "Add", "Update", "Delete", "GetById"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
