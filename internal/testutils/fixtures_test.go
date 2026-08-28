package testutils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/testutils"
)

func TestFixtures(t *testing.T) {
	t.Run("Given CreateTempDirWithFiles When files are provided Then directory is created with files", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"test.md":       testutils.SampleMarkdownWithMermaid,
			"sub/nested.md": "# Nested",
		})

		assert.DirExists(t, tempDir)

		content, err := os.ReadFile(filepath.Join(tempDir, "test.md"))
		require.NoError(t, err)
		assert.Equal(t, testutils.SampleMarkdownWithMermaid, string(content))

		nested, err := os.ReadFile(filepath.Join(tempDir, "sub/nested.md"))
		require.NoError(t, err)
		assert.Equal(t, "# Nested", string(nested))
	})
}
