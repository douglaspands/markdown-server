package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/services"
	"github.com/douglas/markdown-server/internal/testutils"
)

func TestMarkdownService(t *testing.T) {
	t.Run("Given an existing markdown file When GetDocument is called Then it returns rendered document", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) {
				return p, nil
			},
			ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
				return []byte("# Hello World\nContent"), nil
			},
		}
		renderer := &testutils.MockMarkdownRenderer{
			RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
				return "<h1>Hello World</h1><p>Content</p>", "Hello World", nil
			},
		}

		service := services.NewMarkdownService(scanner, renderer)
		ctx := context.Background()

		doc, err := service.GetDocument(ctx, "hello.md")

		require.NoError(t, err)
		assert.NotNil(t, doc)
		assert.Equal(t, "hello.md", doc.RelativePath)
		assert.Equal(t, "Hello World", doc.Title)
		assert.Equal(t, "<h1>Hello World</h1><p>Content</p>", doc.HTMLContent)
		assert.NotEmpty(t, doc.RawContent)
		assert.Equal(t, int64(len("# Hello World\nContent")), doc.SizeBytes)
	})

	t.Run("Given an empty path When GetDocument is called Then it defaults to README.md", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) {
				return p, nil
			},
			ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
				if p == "README.md" {
					return []byte("# Main Readme"), nil
				}
				return nil, domain.ErrFileNotFound
			},
		}
		renderer := &testutils.MockMarkdownRenderer{
			RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
				return "<h1>Main Readme</h1>", "Main Readme", nil
			},
		}

		service := services.NewMarkdownService(scanner, renderer)
		doc, err := service.GetDocument(context.Background(), "")

		require.NoError(t, err)
		assert.Equal(t, "README.md", doc.RelativePath)
		assert.Equal(t, "Main Readme", doc.Title)
	})

	t.Run("Given a non-existent README.md and existing readme.md When GetDocument is called on root Then it falls back to lowercase readme.md", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) {
				return p, nil
			},
			ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
				if p == "readme.md" {
					return []byte("# Lowercase Readme"), nil
				}
				return nil, domain.ErrFileNotFound
			},
		}
		renderer := &testutils.MockMarkdownRenderer{
			RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
				return "<h1>Lowercase Readme</h1>", "", nil
			},
		}

		service := services.NewMarkdownService(scanner, renderer)
		doc, err := service.GetDocument(context.Background(), "/")

		require.NoError(t, err)
		assert.Equal(t, "readme.md", doc.RelativePath)
		assert.Equal(t, "readme", doc.Title)
	})

	t.Run("Given a path outside root When GetDocument is called Then it returns ErrPathEscape", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) {
				return "", domain.ErrPathEscape
			},
		}
		renderer := &testutils.MockMarkdownRenderer{}

		service := services.NewMarkdownService(scanner, renderer)
		doc, err := service.GetDocument(context.Background(), "../../secret.txt")

		require.Error(t, err)
		assert.Nil(t, doc)
		assert.ErrorIs(t, err, domain.ErrPathEscape)
	})

	t.Run("Given a renderer error When GetDocument is called Then it returns error", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) {
				return p, nil
			},
			ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
				return []byte("content"), nil
			},
		}
		renderer := &testutils.MockMarkdownRenderer{
			RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
				return "", "", errors.New("renderer failed")
			},
		}

		service := services.NewMarkdownService(scanner, renderer)
		doc, err := service.GetDocument(context.Background(), "doc.md")

		require.Error(t, err)
		assert.Nil(t, doc)
	})

	t.Run("Given a cancelled context When GetDocument or GetFileTree is called Then it returns context error", func(t *testing.T) {
		service := services.NewMarkdownService(&testutils.MockFileScanner{}, &testutils.MockMarkdownRenderer{})
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, errDoc := service.GetDocument(ctx, "test.md")
		require.Error(t, errDoc)
		assert.ErrorIs(t, errDoc, context.Canceled)

		_, errTree := service.GetFileTree(ctx)
		require.Error(t, errTree)
		assert.ErrorIs(t, errTree, context.Canceled)
	})

	t.Run("Given a valid directory tree When GetFileTree is called Then it delegates to scanner", func(t *testing.T) {
		scanner := &testutils.MockFileScanner{
			ScanDirectoryFunc: func(ctx context.Context) (*domain.FileNode, error) {
				return &domain.FileNode{
					Name: "root",
					Type: domain.NodeTypeDirectory,
					Children: []*domain.FileNode{
						{Name: "index.md", Type: domain.NodeTypeFile, IsMarkdown: true},
					},
				}, nil
			},
		}

		service := services.NewMarkdownService(scanner, &testutils.MockMarkdownRenderer{})
		tree, err := service.GetFileTree(context.Background())

		require.NoError(t, err)
		assert.NotNil(t, tree)
		assert.Equal(t, "root", tree.Name)
		assert.Len(t, tree.Children, 1)
	})
}
