// Package http provides functions to trace the net/http package (https://golang.org/pkg/net/http).
package http

import (
	"net/http"

	"github.com/DataDog/dd-trace-go/contrib/internal"
	"github.com/DataDog/dd-trace-go/ddtrace/ext"
	"github.com/DataDog/dd-trace-go/ddtrace/tracer"
)

// ServeMux is an HTTP request multiplexer that traces all the incoming requests.
type ServeMux struct {
	*http.ServeMux
	config *muxConfig
}

// NewServeMux allocates and returns an http.ServeMux augmented with the
// global tracer.
func NewServeMux(opts ...MuxOption) *ServeMux {
	cfg := new(muxConfig)
	defaults(cfg)
	for _, fn := range opts {
		fn(cfg)
	}
	tracer.SetServiceInfo(cfg.serviceName, "net/http", ext.AppTypeWeb)
	return &ServeMux{
		ServeMux: http.NewServeMux(),
		config:   cfg,
	}
}

// ServeHTTP dispatches the request to the handler
// whose pattern most closely matches the request URL.
// We only need to rewrite this function to be able to trace
// all the incoming requests to the underlying multiplexer
func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the resource associated to this request
	_, route := mux.Handler(r)
	resource := r.Method + " " + route
	internal.TraceAndServe(mux.ServeMux, w, r, mux.config.serviceName, resource)
}

// WrapHandler wraps an http.Handler with tracing using the given service and resource.
func WrapHandler(h http.Handler, service, resource string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		internal.TraceAndServe(h, w, req, service, resource)
	})
}
