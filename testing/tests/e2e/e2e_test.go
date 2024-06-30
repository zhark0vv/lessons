// api_e2e_test.go
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"lessons/testing/api"
	"lessons/testing/repository"
	"lessons/testing/service"
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

	_, err = conn.Exec(ctx, "INSERT INTO test_table (data) VALUES ('e2e test data')")
	if err != nil {
		t.Fatal(err)
	}

	return postgresC, conn
}

func TestAPI_E2E(t *testing.T) {
	postgresC, conn := setupPostgresContainer(t)
	defer postgresC.Terminate(context.Background())
	defer conn.Close(context.Background())

	repo := repository.New(conn)
	svc := service.New(repo)
	a := api.New(svc)

	server := httptest.NewServer(a.Router())
	defer server.Close()

	client := &http.Client{Timeout: 10 * time.Second}

	t.Run("GetDataHandler should return data successfully", func(t *testing.T) {
		resp, err := client.Get(fmt.Sprintf("%s/data/1", server.URL))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result map[string]string
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Processed: e2e test data", result["data"])
	})

	t.Run("GetDataHandler should return 404 for missing data", func(t *testing.T) {
		resp, err := client.Get(fmt.Sprintf("%s/data/999", server.URL))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
