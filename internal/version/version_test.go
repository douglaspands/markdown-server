package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/douglas/markdown-server/internal/version"
)

func TestVersion(t *testing.T) {
	t.Run("Given default build metadata When GetInfo is called Then it returns valid info struct", func(t *testing.T) {
		info := version.GetInfo()

		assert.Equal(t, "dev", info.Version)
		assert.Equal(t, "none", info.Commit)
		assert.Equal(t, "unknown", info.Date)
		assert.NotEmpty(t, info.GoVersion)
		assert.NotEmpty(t, info.OS)
		assert.NotEmpty(t, info.Arch)
	})

	t.Run("Given an application name When Formatted is called Then it returns formatted string", func(t *testing.T) {
		formatted := version.Formatted("md_server")
		assert.Contains(t, formatted, "md_server version: dev")
		assert.Contains(t, formatted, "commit: none")
		assert.Contains(t, formatted, "built at: unknown")
	})

	t.Run("Given an empty application name When Formatted is called Then it defaults to md_server", func(t *testing.T) {
		formatted := version.Formatted("")
		assert.Contains(t, formatted, "md_server version: dev")
	})
}
