package browser_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/adapters/out/browser"
)

func TestBrowserLauncher(t *testing.T) {
	t.Run("Given Windows OS When OpenURL is called Then it executes rundll32 with url protocol", func(t *testing.T) {
		var executedName string
		var executedArgs []string

		mockExecutor := func(ctx context.Context, name string, args ...string) error {
			executedName = name
			executedArgs = args
			return nil
		}

		launcher := browser.NewCustomBrowserLauncher("windows", mockExecutor)
		err := launcher.OpenURL(context.Background(), "http://localhost:8080")

		require.NoError(t, err)
		assert.Equal(t, "rundll32", executedName)
		assert.Equal(t, []string{"url.dll,FileProtocolHandler", "http://localhost:8080"}, executedArgs)
	})

	t.Run("Given Linux OS When OpenURL is called Then it executes xdg-open", func(t *testing.T) {
		var executedName string
		var executedArgs []string

		mockExecutor := func(ctx context.Context, name string, args ...string) error {
			executedName = name
			executedArgs = args
			return nil
		}

		launcher := browser.NewCustomBrowserLauncher("linux", mockExecutor)
		err := launcher.OpenURL(context.Background(), "http://localhost:8080")

		require.NoError(t, err)
		assert.Equal(t, "xdg-open", executedName)
		assert.Equal(t, []string{"http://localhost:8080"}, executedArgs)
	})

	t.Run("Given Darwin OS When OpenURL is called Then it executes open", func(t *testing.T) {
		var executedName string
		var executedArgs []string

		mockExecutor := func(ctx context.Context, name string, args ...string) error {
			executedName = name
			executedArgs = args
			return nil
		}

		launcher := browser.NewCustomBrowserLauncher("darwin", mockExecutor)
		err := launcher.OpenURL(context.Background(), "http://localhost:8080")

		require.NoError(t, err)
		assert.Equal(t, "open", executedName)
		assert.Equal(t, []string{"http://localhost:8080"}, executedArgs)
	})

	t.Run("Given an empty URL When OpenURL is called Then it returns error", func(t *testing.T) {
		launcher := browser.NewCustomBrowserLauncher("linux", func(ctx context.Context, name string, args ...string) error {
			return nil
		})

		err := launcher.OpenURL(context.Background(), "   ")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "url não pode ser vazia")
	})

	t.Run("Given a cancelled context When OpenURL is called Then it returns context error", func(t *testing.T) {
		launcher := browser.NewCustomBrowserLauncher("linux", func(ctx context.Context, name string, args ...string) error {
			return nil
		})
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := launcher.OpenURL(ctx, "http://localhost:8080")
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("Given executor failure When OpenURL is called Then it propagates error", func(t *testing.T) {
		launcher := browser.NewCustomBrowserLauncher("linux", func(ctx context.Context, name string, args ...string) error {
			return errors.New("exec failure")
		})

		err := launcher.OpenURL(context.Background(), "http://localhost:8080")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "exec failure")
	})

	t.Run("Given default constructor When NewBrowserLauncher is initialized Then it is not nil", func(t *testing.T) {
		launcher := browser.NewBrowserLauncher()
		assert.NotNil(t, launcher)
	})

	t.Run("Given a nil executor When OpenURL is called Then it returns error", func(t *testing.T) {
		launcher := browser.NewCustomBrowserLauncher("linux", nil)
		err := launcher.OpenURL(context.Background(), "http://localhost:8080")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "nenhum executor")
	})
}
