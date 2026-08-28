package fs_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/adapters/out/fs"
	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/testutils"
)

func TestFSScanner(t *testing.T) {
	t.Run("Given a valid root directory When NewFSScanner is initialized Then it succeeds", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md": "# Hello",
		})

		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)
		assert.NotNil(t, scanner)
	})

	t.Run("Given an invalid root directory When NewFSScanner is initialized Then it returns error", func(t *testing.T) {
		scanner, err := fs.NewFSScanner("/non/existent/directory/path/12345")
		require.Error(t, err)
		assert.Nil(t, scanner)
	})

	t.Run("Given a file path instead of directory When NewFSScanner is initialized Then it returns ErrDirectoryScan", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"file.txt": "not a dir",
		})
		filePath := filepath.Join(tempDir, "file.txt")

		scanner, err := fs.NewFSScanner(filePath)
		require.Error(t, err)
		assert.Nil(t, scanner)
		assert.ErrorIs(t, err, domain.ErrDirectoryScan)
	})

	t.Run("Given relative and clean paths When SanitizePath is called Then it prevents path traversal", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"docs/guide.md": "# Guide",
		})
		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)

		safePath, err := scanner.SanitizePath("docs/guide.md")
		require.NoError(t, err)
		assert.Equal(t, "docs/guide.md", safePath)

		// Test path escape attempts
		_, errEscape := scanner.SanitizePath("../../etc/passwd")
		require.Error(t, errEscape)
		assert.ErrorIs(t, errEscape, domain.ErrPathEscape)

		_, errEscape2 := scanner.SanitizePath("..")
		require.Error(t, errEscape2)
		assert.ErrorIs(t, errEscape2, domain.ErrPathEscape)
	})

	t.Run("Given existing files When ReadFile is called Then it returns file bytes", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md":     "# Root Readme",
			"docs/intro.md": "# Introduction",
		})
		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)

		ctx := context.Background()

		content, err := scanner.ReadFile(ctx, "README.md")
		require.NoError(t, err)
		assert.Equal(t, []byte("# Root Readme"), content)

		contentIntro, err := scanner.ReadFile(ctx, "docs/intro.md")
		require.NoError(t, err)
		assert.Equal(t, []byte("# Introduction"), contentIntro)
	})

	t.Run("Given non-existent or directory path When ReadFile is called Then it returns ErrFileNotFound", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"docs/intro.md": "# Intro",
		})
		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)

		ctx := context.Background()

		_, errNotFound := scanner.ReadFile(ctx, "missing.md")
		require.Error(t, errNotFound)
		assert.ErrorIs(t, errNotFound, domain.ErrFileNotFound)

		// Reading directory as file should return ErrFileNotFound
		_, errDir := scanner.ReadFile(ctx, "docs")
		require.Error(t, errDir)
		assert.ErrorIs(t, errDir, domain.ErrFileNotFound)
	})

	t.Run("Given a complex directory structure When ScanDirectory is called Then it returns structured tree", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md":            "# Root",
			"docs/architecture.md": "# Arch",
			"docs/api/v1.md":       "# V1",
			"assets/logo.png":      "binary-content",
			".git/config":          "git-config",
		})
		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)

		tree, err := scanner.ScanDirectory(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, tree)
		assert.Equal(t, domain.NodeTypeDirectory, tree.Type)

		// Verify that hidden folder .git was skipped
		for _, child := range tree.Children {
			assert.NotEqual(t, ".git", child.Name)
		}

		// Verify children hierarchy
		var docsNode *domain.FileNode
		var readmeNode *domain.FileNode
		for _, child := range tree.Children {
			if child.Name == "docs" {
				docsNode = child
			}
			if child.Name == "README.md" {
				readmeNode = child
			}
		}

		require.NotNil(t, docsNode)
		assert.Equal(t, domain.NodeTypeDirectory, docsNode.Type)
		assert.Equal(t, "docs", docsNode.RelativePath)

		require.NotNil(t, readmeNode)
		assert.Equal(t, domain.NodeTypeFile, readmeNode.Type)
		assert.True(t, readmeNode.IsMarkdown)
	})

	t.Run("Given a cancelled context When ReadFile or ScanDirectory is called Then it returns context error", func(t *testing.T) {
		tempDir := testutils.CreateTempDirWithFiles(t, map[string]string{
			"README.md": "# Root",
		})
		scanner, err := fs.NewFSScanner(tempDir)
		require.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, errRead := scanner.ReadFile(ctx, "README.md")
		require.Error(t, errRead)
		assert.ErrorIs(t, errRead, context.Canceled)

		_, errScan := scanner.ScanDirectory(ctx)
		require.Error(t, errScan)
		assert.ErrorIs(t, errScan, context.Canceled)
	})

	t.Run("Given default empty string rootDir When NewFSScanner is initialized Then it defaults to current directory", func(t *testing.T) {
		scanner, err := fs.NewFSScanner("")
		require.NoError(t, err)
		assert.NotNil(t, scanner)
	})
}
