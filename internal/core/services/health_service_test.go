package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/core/services"
)

func TestHealthCheckService(t *testing.T) {
	t.Run("Given a HealthCheckService When CheckHealth is called with valid context Then it returns status UP and valid metadata", func(t *testing.T) {
		startTime := time.Now().Add(-10 * time.Minute)
		service := services.NewHealthCheckService(startTime)
		ctx := context.Background()

		status, err := service.CheckHealth(ctx)

		require.NoError(t, err)
		assert.NotNil(t, status)
		assert.Equal(t, "UP", status.Status)
		assert.NotEmpty(t, status.Version)
		assert.NotEmpty(t, status.Uptime)
		assert.Equal(t, startTime, status.StartTime)
	})

	t.Run("Given a cancelled context When CheckHealth is called Then it returns context error", func(t *testing.T) {
		service := services.NewHealthCheckService(time.Now())
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		status, err := service.CheckHealth(ctx)

		require.Error(t, err)
		assert.Nil(t, status)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("Given a zero startTime When NewHealthCheckService is initialized Then it defaults to current time", func(t *testing.T) {
		service := services.NewHealthCheckService(time.Time{})
		status, err := service.CheckHealth(context.Background())

		require.NoError(t, err)
		assert.NotNil(t, status)
		assert.False(t, status.StartTime.IsZero())
	})
}
