package cli_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/adapters/in/cli"
	"github.com/douglas/markdown-server/internal/core/domain"
)

func TestCLI(t *testing.T) {
	t.Run("Given default flags When root command is executed Then runner receives default config", func(t *testing.T) {
		var receivedCfg domain.ServerConfig
		runner := func(cmd *cobra.Command, args []string, cfg domain.ServerConfig) error {
			receivedCfg = cfg
			return nil
		}

		cmd := cli.NewRootCommand(runner)
		cmd.SetArgs([]string{})

		err := cmd.Execute()
		require.NoError(t, err)

		assert.Equal(t, ".", receivedCfg.RootDir)
		assert.Equal(t, 8080, receivedCfg.Port)
		assert.True(t, receivedCfg.AutoOpenBrowser)
	})

	t.Run("Given custom flags When root command is executed Then runner receives custom config", func(t *testing.T) {
		var receivedCfg domain.ServerConfig
		runner := func(cmd *cobra.Command, args []string, cfg domain.ServerConfig) error {
			receivedCfg = cfg
			return nil
		}

		cmd := cli.NewRootCommand(runner)
		cmd.SetArgs([]string{"--dir", "/tmp/docs", "--port", "9090", "--open=false"})

		err := cmd.Execute()
		require.NoError(t, err)

		assert.Equal(t, "/tmp/docs", receivedCfg.RootDir)
		assert.Equal(t, 9090, receivedCfg.Port)
		assert.False(t, receivedCfg.AutoOpenBrowser)
	})

	t.Run("Given a positional path argument When executed Then rootDir is updated", func(t *testing.T) {
		var receivedCfg domain.ServerConfig
		runner := func(cmd *cobra.Command, args []string, cfg domain.ServerConfig) error {
			receivedCfg = cfg
			return nil
		}

		cmd := cli.NewRootCommand(runner)
		cmd.SetArgs([]string{"./my-docs"})

		err := cmd.Execute()
		require.NoError(t, err)

		assert.Equal(t, "./my-docs", receivedCfg.RootDir)
	})

	t.Run("Given --version flag When executed Then version output is printed", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd := cli.NewRootCommand(nil)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{"--version"})

		err := cmd.Execute()
		require.NoError(t, err)
	})

	t.Run("Given version sub-command When executed Then it runs cleanly", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd := cli.NewRootCommand(nil)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{"version"})

		err := cmd.Execute()
		require.NoError(t, err)
	})
}
