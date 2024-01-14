package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/model"
)

type Sample interface {
	Create(ctx context.Context, input SampleCreateInput) (*SampleCreateOutput, error)
	Read(ctx context.Context, input SampleReadInput) (*SampleReadOutput, error)
	Update(ctx context.Context, input SampleUpdateInput) (*SampleUpdateOutput, error)
	Delete(ctx context.Context, input SampleDeleteInput) (*SampleDeleteOutput, error)
	List(ctx context.Context, input SampleListInput) (*SampleListOutput, error)
}

type SampleCreateInput struct {
	Message string
}

type SampleCreateOutput struct {
	Sample model.Sample
}

type SampleReadInput struct {
	ID uuid.UUID
}

type SampleReadOutput struct {
	Sample model.Sample
}

type SampleListInput struct{}

type SampleListOutput struct {
	SampleList model.SampleList
}

type SampleUpdateInput struct {
	ID      uuid.UUID
	Message string
}

type SampleUpdateOutput struct {
	Sample model.Sample
}

type SampleDeleteInput struct {
	ID uuid.UUID
}

type SampleDeleteOutput struct{}
