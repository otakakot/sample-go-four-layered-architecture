package controller

import (
	"context"

	"github.com/otakakot/sample-go-four-layered-architecture/internal/application/usecase"
	"github.com/otakakot/sample-go-four-layered-architecture/pkg/api"
)

type Sample struct {
	uc usecase.Sample
}

func NewSample(
	uc usecase.Sample,
) *Sample {
	return &Sample{
		uc: uc,
	}
}

// CreateSample implements api.Handler.
func (ctl *Sample) CreateSample(
	ctx context.Context,
	req *api.CreateSampleRequestSchema,
) (api.CreateSampleRes, error) {
	output, err := ctl.uc.Create(ctx, usecase.SampleCreateInput{
		Message: req.Message,
	})
	if err != nil {
		return &api.ErrorResponseSchema{
			Message: err.Error(),
		}, nil
	}

	return &api.CreateSampleResponseSchema{
		Sample: api.Sample{
			ID:        output.Sample.ID,
			Message:   output.Sample.Message,
			CreatedAt: output.Sample.CreatedAt,
			UpdatedAt: output.Sample.UpdatedAt,
		},
	}, nil
}

// DeleteSample implements api.Handler.
func (ctl *Sample) DeleteSample(
	ctx context.Context,
	params api.DeleteSampleParams,
) (api.DeleteSampleRes, error) {
	if _, err := ctl.uc.Delete(ctx, usecase.SampleDeleteInput{
		ID: params.ID,
	}); err != nil {
		return &api.ErrorResponseSchema{
			Message: err.Error(),
		}, nil
	}

	return &api.DeleteSampleNoContent{}, nil
}

// ListSample implements api.Handler.
func (ctl *Sample) ListSample(
	ctx context.Context,
) (api.ListSampleRes, error) {
	output, err := ctl.uc.List(ctx, usecase.SampleListInput{})
	if err != nil {
		return &api.ErrorResponseSchema{
			Message: err.Error(),
		}, nil
	}

	samples := make([]api.Sample, len(output.SampleList))
	for i, sample := range output.SampleList {
		samples[i] = api.Sample{
			ID:        sample.ID,
			Message:   sample.Message,
			CreatedAt: sample.CreatedAt,
			UpdatedAt: sample.UpdatedAt,
		}
	}

	return &api.ListSampleResponseSchema{
		Samples: samples,
	}, nil
}

// ReadSample implements api.Handler.
func (ctl *Sample) ReadSample(
	ctx context.Context,
	params api.ReadSampleParams,
) (api.ReadSampleRes, error) {
	output, err := ctl.uc.Read(ctx, usecase.SampleReadInput{
		ID: params.ID,
	})
	if err != nil {
		return &api.ErrorResponseSchema{
			Message: err.Error(),
		}, nil
	}

	return &api.ReadSampleResponseSchema{
		Sample: api.Sample{
			ID:        output.Sample.ID,
			Message:   output.Sample.Message,
			CreatedAt: output.Sample.CreatedAt,
			UpdatedAt: output.Sample.UpdatedAt,
		},
	}, nil
}

// UpdateSample implements api.Handler.
func (ctl *Sample) UpdateSample(
	ctx context.Context,
	req *api.UpdateSampleRequestSchema,
	params api.UpdateSampleParams,
) (api.UpdateSampleRes, error) {
	output, err := ctl.uc.Update(ctx, usecase.SampleUpdateInput{
		ID:      params.ID,
		Message: req.Message,
	})
	if err != nil {
		return &api.ErrorResponseSchema{
			Message: err.Error(),
		}, nil
	}

	return &api.UpdateSampleResponseSchema{
		Sample: api.Sample{
			ID:        output.Sample.ID,
			Message:   output.Sample.Message,
			CreatedAt: output.Sample.CreatedAt,
			UpdatedAt: output.Sample.UpdatedAt,
		},
	}, nil
}
