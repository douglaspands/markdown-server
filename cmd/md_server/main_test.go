package main

import (
	"context"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/testutils"
)

func TestRunServer(t *testing.T) {
	t.Run("Given an invalid directory When runServerWithContext is called Then it returns error", func(t *testing.T) {
		cfg := domain.ServerConfig{
			RootDir:         "/non/existent/path/9999",
			Port:            8080,
			AutoOpenBrowser: false,
		}

		err := runServerWithContext(context.Background(), &cobra.Command{}, []string{}, cfg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "falha ao inicializar scanner")
	})

	t.Run("Given a valid directory and context cancellation When runServerWithContext is invoked Then server starts and shuts down cleanly", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md": "# Test Root",
		})

		cfg := domain.ServerConfig{
			RootDir:         tempDir,
			Port:            0,
			AutoOpenBrowser: false,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		err := runServerWithContext(ctx, &cobra.Command{}, []string{}, cfg)
		require.NoError(t, err)
	})

	t.Run("Given AutoOpenBrowser enabled When runServerWithContext starts and cancels Then it executes cleanly", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md": "# Test Browser",
		})

		cfg := domain.ServerConfig{
			RootDir:         tempDir,
			Port:            0,
			AutoOpenBrowser: true,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		err := runServerWithContext(ctx, &cobra.Command{}, []string{}, cfg)
		require.NoError(t, err)
	})
}
