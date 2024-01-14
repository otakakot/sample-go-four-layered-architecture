package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/adapter/gateway"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/domain/model"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/driver/postgres"
)

func TestIntegration(t *testing.T) {
	t.Parallel()

	dsn := os.Getenv("POSTGRES_URL")

	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	db, err := postgres.New(dsn)
	if err != nil {
		t.Fatalf("failed to initialize postgres: %s", err)
	}

	t.Cleanup(func() {
		if _, err := db.ExecContext(context.Background(), "TRUNCATE TABLE samples"); err != nil {
			t.Errorf("failed to truncate table: %s", err)
		}

		if err := db.Close(); err != nil {
			t.Errorf("failed to close database: %s", err)
		}
	})

	gw := gateway.NewSample(db)

	t.Run("保存して取得して削除して一覧取得する", func(t *testing.T) {
		t.Parallel()

		sample := model.GenerateSample("message")

		if err := gw.Save(context.Background(), *sample); err != nil {
			t.Errorf("failed to create Sample: %s", err)
		}

		got, err := gw.Find(context.Background(), sample.ID)
		if err != nil {
			t.Errorf("failed to find Sample: %s", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Sample{}, "CreatedAt", "UpdatedAt"),
		}

		if diff := cmp.Diff(*sample, *got, opts...); diff != "" {
			t.Errorf("sample mismatch (-want +got):\n%s", diff)
		}

		if err := gw.Delete(context.Background(), sample.ID); err != nil {
			t.Errorf("failed to delete sample: %s", err)
		}

		gots, err := gw.List(context.Background())
		if err != nil {
			t.Errorf("failed to list sample: %s", err)
		}

		if len(gots.Samples()) != 0 {
			t.Errorf("samle count mismatch: want 0, got %d", len(gots.Samples()))
		}
	})
}
