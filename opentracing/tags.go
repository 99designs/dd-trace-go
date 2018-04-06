package opentracing

const (
	// SpanType defines the Span type (web, db, cache)
	SpanType = "span.type"
	// ServiceName defines the Service name for this Span
	ServiceName = "service.name"
	// The type of operation being measured. Some examples
	// might be "http.handler", "fileserver.upload" or "video.decompress".
	// Name should be set on every span.
	OperationType = "operation.type"
	// NONSTANDARD USE span.SetOperationName INSTEAD
	ResourceName = "resource.name"
	// Error defines an error.
	Error = "error.error"
)
