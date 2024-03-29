// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateSample implements createSample operation.
	//
	// Create a new sample item.
	//
	// POST /samples
	CreateSample(ctx context.Context, req *CreateSampleRequestSchema) (CreateSampleRes, error)
	// DeleteSample implements deleteSample operation.
	//
	// Delete a specific sample item.
	//
	// DELETE /samples/{id}
	DeleteSample(ctx context.Context, params DeleteSampleParams) (DeleteSampleRes, error)
	// Health implements health operation.
	//
	// Health.
	//
	// GET /health
	Health(ctx context.Context) (HealthRes, error)
	// ListSample implements listSample operation.
	//
	// Get all sample items.
	//
	// GET /samples
	ListSample(ctx context.Context) (ListSampleRes, error)
	// ReadSample implements readSample operation.
	//
	// Read a specific sample item.
	//
	// GET /samples/{id}
	ReadSample(ctx context.Context, params ReadSampleParams) (ReadSampleRes, error)
	// UpdateSample implements updateSample operation.
	//
	// Update a specific sample item.
	//
	// PUT /samples/{id}
	UpdateSample(ctx context.Context, req *UpdateSampleRequestSchema, params UpdateSampleParams) (UpdateSampleRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
