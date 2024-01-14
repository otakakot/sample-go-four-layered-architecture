package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/model"
)

type Sample interface {
	List(ctx context.Context) (model.SampleList, error)
	Find(ctx context.Context, id uuid.UUID) (*model.Sample, error)
	Save(ctx context.Context, Sample model.Sample) error
	Delete(ctx context.Context, id uuid.UUID) error
}
