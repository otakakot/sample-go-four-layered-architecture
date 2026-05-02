package controller

import (
	"context"

	"github.com/otakakot/sample-go-four-layered-architecture/pkg/api"
)

var _ api.Handler = (*Controller)(nil)

type Controller struct {
	*Sample
}

// Health implements api.Handler.
func (ctl *Controller) Health(
	_ context.Context,
) (api.HealthRes, error) {
	return &api.HealthResponseSchema{
		Message: "OK",
	}, nil
}
