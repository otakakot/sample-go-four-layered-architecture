package model

import (
	"time"

	"github.com/google/uuid"
)

type Sample struct {
	ID        uuid.UUID
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateSample(
	message string,
) *Sample {
	now := time.Now()

	return &Sample{
		ID:        uuid.New(),
		Message:   message,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type SampleList []Sample

func (tl SampleList) Samples() []Sample {
	return tl
}
