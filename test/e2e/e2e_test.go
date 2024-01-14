package e2e_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/driver/postgres"
	"github.com/otakakot/sample-go-four-layered-architecture/pkg/api"
)

func TestE2E(t *testing.T) {
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

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	cli, err := api.NewClient(endpoint)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}

	if _, err := cli.Health(context.Background()); err != nil {
		t.Fatalf("failed to health check: %s", err)
	}

	t.Run("保存して更新して取得して削除して一覧取得する", func(t *testing.T) {
		t.Parallel()

		createRes, err := cli.CreateSample(context.Background(), &api.CreateSampleRequestSchema{
			Message: "message",
		})
		if err != nil {
			t.Fatalf("failed to create sample: %s", err)
		}

		checkedCreateRes, ok := createRes.(*api.CreateSampleResponseSchema)
		if !ok {
			t.Fatalf("failed to create sample: unknown response: %T", createRes)
		}

		want := checkedCreateRes.Sample

		readRes, err := cli.ReadSample(context.Background(), api.ReadSampleParams{
			ID: want.ID,
		})
		if err != nil {
			t.Fatalf("failed to read sample: %s", err)
		}

		checkedReadRes, ok := readRes.(*api.ReadSampleResponseSchema)
		if !ok {
			t.Fatalf("failed to read sample: unknown response: %T", readRes)
		}

		got := checkedReadRes.Sample

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("sample mismatch (-want +got):\n%s", diff)
		}

		if _, err := cli.UpdateSample(context.Background(), &api.UpdateSampleRequestSchema{
			Message: "updated",
		}, api.UpdateSampleParams{
			ID: want.ID,
		}); err != nil {
			t.Fatalf("failed to update sample: %s", err)
		}

		readRes2, err := cli.ReadSample(context.Background(), api.ReadSampleParams{
			ID: want.ID,
		})
		if err != nil {
			t.Fatalf("failed to read sample: %s", err)
		}

		checkedReadRes2, ok := readRes2.(*api.ReadSampleResponseSchema)
		if !ok {
			t.Fatalf("failed to read sample: unknown response: %T", readRes2)
		}

		got2 := checkedReadRes2.Sample

		want.Message = "updated"

		if diff := cmp.Diff(want, got2); diff != "" {
			t.Errorf("sample mismatch (-want +got):\n%s", diff)
		}

		deleteRes, err := cli.DeleteSample(context.Background(), api.DeleteSampleParams{
			ID: want.ID,
		})
		if err != nil {
			t.Fatalf("failed to delete sample: %s", err)
		}

		if _, ok := deleteRes.(*api.DeleteSampleNoContent); !ok {
			t.Fatalf("failed to delete sample: unknown response: %T", deleteRes)
		}

		listRes, err := cli.ListSample(context.Background())
		if err != nil {
			t.Fatalf("failed to list sample: %s", err)
		}

		checkedListRes, ok := listRes.(*api.ListSampleResponseSchema)
		if !ok {
			t.Fatalf("failed to list sample: unknown response: %T", listRes)
		}

		gots := checkedListRes.Samples

		if len(gots) != 0 {
			t.Errorf("sample count mismatch: want 0, got %d", len(gots))
		}
	})
}
