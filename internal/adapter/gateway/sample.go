package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/model"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/repository"
)

var _ repository.Sample = (*Sample)(nil)

type Sample struct {
	db *sql.DB
}

func NewSample(
	db *sql.DB,
) *Sample {
	return &Sample{
		db: db,
	}
}

// Delete implements repository.Sample.
func (gw *Sample) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `DELETE FROM samples WHERE id = $1`

	if _, err := gw.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete row: %w", err)
	}

	return nil
}

// Find implements repository.Sample.
func (gw *Sample) Find(
	ctx context.Context,
	id uuid.UUID,
) (*model.Sample, error) {
	query := `SELECT id, message, created_at, updated_at FROM samples WHERE id = $1`

	row := gw.db.QueryRowContext(ctx, query, id.String())

	var sample model.Sample
	if err := row.Scan(&sample.ID, &sample.Message, &sample.CreatedAt, &sample.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("sample not found: %w", err)
		}

		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &sample, nil
}

// List implements repository.Sample.
func (gw *Sample) List(ctx context.Context) (model.SampleList, error) {
	query := `SELECT id, message, created_at, updated_at FROM samples`

	rows, err := gw.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query rows: %w", err)
	}
	defer rows.Close()

	const size = 100

	samples := make(model.SampleList, 0, size)

	for rows.Next() {
		var sample model.Sample
		if err := rows.Scan(&sample.ID, &sample.Message, &sample.CreatedAt, &sample.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		samples = append(samples, sample)
	}

	return samples, nil
}

// Save implements repository.Sample.
func (gw *Sample) Save(
	ctx context.Context,
	sample model.Sample,
) error {
	query := `
		INSERT INTO samples (id, message, created_at, updated_at) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE 
		SET message = $2, updated_at = $4
		`

	if _, err := gw.db.ExecContext(ctx, query, sample.ID, sample.Message, sample.CreatedAt, sample.UpdatedAt); err != nil {
		return fmt.Errorf("failed to insert row: %w", err)
	}

	return nil
}
