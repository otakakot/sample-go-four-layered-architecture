package interactor

import (
	"context"
	"fmt"

	"github.com/otakakot/sample-go-four-layered-architecture/internal/application/usecase"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/model"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/repository"
)

var _ usecase.Sample = (*Sample)(nil)

type Sample struct {
	sampleRepository repository.Sample
}

func NewSample(
	sampleRepository repository.Sample,
) *Sample {
	return &Sample{
		sampleRepository: sampleRepository,
	}
}

// Create implements usecase.Sample.
func (itr *Sample) Create(
	ctx context.Context,
	input usecase.SampleCreateInput,
) (*usecase.SampleCreateOutput, error) {
	sample := model.GenerateSample(input.Message)

	if err := itr.sampleRepository.Save(ctx, *sample); err != nil {
		return nil, fmt.Errorf("failed to save sample: %w", err)
	}

	return &usecase.SampleCreateOutput{
		Sample: *sample,
	}, nil
}

// Read implements usecase.Sample.
func (itr *Sample) Read(
	ctx context.Context,
	input usecase.SampleReadInput,
) (*usecase.SampleReadOutput, error) {
	sample, err := itr.sampleRepository.Find(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find sample: %w", err)
	}

	return &usecase.SampleReadOutput{
		Sample: *sample,
	}, nil
}

// Update implements usecase.Sample.
func (itr *Sample) Update(
	ctx context.Context,
	input usecase.SampleUpdateInput,
) (*usecase.SampleUpdateOutput, error) {
	sample, err := itr.sampleRepository.Find(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find Sample: %w", err)
	}

	sample.Message = input.Message

	if err := itr.sampleRepository.Save(ctx, *sample); err != nil {
		return nil, fmt.Errorf("failed to save Sample: %w", err)
	}

	return &usecase.SampleUpdateOutput{
		Sample: *sample,
	}, nil
}

// Delete implements usecase.Sample.
func (itr *Sample) Delete(
	ctx context.Context,
	input usecase.SampleDeleteInput,
) (*usecase.SampleDeleteOutput, error) {
	if err := itr.sampleRepository.Delete(ctx, input.ID); err != nil {
		return nil, fmt.Errorf("failed to delete sample: %w", err)
	}

	return &usecase.SampleDeleteOutput{}, nil
}

// List implements usecase.Sample.
func (itr *Sample) List(
	ctx context.Context,
	_ usecase.SampleListInput,
) (*usecase.SampleListOutput, error) {
	sample, err := itr.sampleRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list samples: %w", err)
	}

	return &usecase.SampleListOutput{
		SampleList: sample,
	}, nil
}
