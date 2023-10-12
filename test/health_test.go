package test

import (
	"github.com/RamiroCyber/projetc_golang/router/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/health", handler.HealthCheck)

	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err, "Failed to create request")

	resp, err := app.Test(req)
	require.NoError(t, err, "Failed to get response")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	require.Equal(t, "SERVER UP!", string(body), "Unexpected response body")
}
