//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"lessons/testing/repository"
	service "lessons/testing/service"
)

func setupPostgresContainer(t *testing.T) (testcontainers.Container, *pgx.Conn) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(5 * time.Minute),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	host, err := postgresC.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://postgres:password@%s:%s/testdb", host, port.Port())
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "CREATE TABLE test_table (id SERIAL PRIMARY KEY, data TEXT)")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "INSERT INTO test_table (data) VALUES ('integration test data')")
	if err != nil {
		t.Fatal(err)
	}

	return postgresC, conn
}

func TestServiceIntegration(t *testing.T) {
	postgresC, conn := setupPostgresContainer(t)
	defer func(postgresC testcontainers.Container, ctx context.Context) {
		err := postgresC.Terminate(ctx)
		if err != nil {
			t.Error(err)
		}
	}(postgresC, context.Background())
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			t.Error(err)
		}
	}(conn, context.Background())

	repo := repository.New(conn)
	s := service.New(repo)

	ctx := context.Background()
	result, err := s.ProcessData(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, "Processed: integration test data", result)
}
