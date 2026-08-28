package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/douglas/markdown-server/internal/core/domain"
)

func TestDomainEntities(t *testing.T) {
	t.Run("Given a MarkdownDocument When initialized Then properties are correctly set", func(t *testing.T) {
		doc := domain.MarkdownDocument{
			RelativePath: "docs/readme.md",
			Title:        "Readme",
			RawContent:   []byte("# Readme"),
			HTMLContent:  "<h1>Readme</h1>",
			SizeBytes:    8,
		}

		assert.Equal(t, "docs/readme.md", doc.RelativePath)
		assert.Equal(t, "Readme", doc.Title)
		assert.Equal(t, "<h1>Readme</h1>", doc.HTMLContent)
		assert.Equal(t, int64(8), doc.SizeBytes)
	})

	t.Run("Given a FileNode tree When initialized Then hierarchy is valid", func(t *testing.T) {
		root := &domain.FileNode{
			Name:         "root",
			RelativePath: "",
			Type:         domain.NodeTypeDirectory,
			Children: []*domain.FileNode{
				{
					Name:         "doc.md",
					RelativePath: "doc.md",
					Type:         domain.NodeTypeFile,
					IsMarkdown:   true,
				},
			},
		}

		assert.Equal(t, domain.NodeTypeDirectory, root.Type)
		assert.Len(t, root.Children, 1)
		assert.True(t, root.Children[0].IsMarkdown)
	})

	t.Run("Given a HealthStatus When created Then fields match", func(t *testing.T) {
		now := time.Now()
		status := domain.HealthStatus{
			Status:    "UP",
			Version:   "1.0.0",
			Commit:    "abc1234",
			Uptime:    "5m",
			StartTime: now,
		}

		assert.Equal(t, "UP", status.Status)
		assert.Equal(t, "1.0.0", status.Version)
		assert.Equal(t, "abc1234", status.Commit)
		assert.Equal(t, "5m", status.Uptime)
		assert.Equal(t, now, status.StartTime)
	})

	t.Run("Given sentinel errors When checked Then descriptions are informative", func(t *testing.T) {
		assert.Equal(t, "arquivo não encontrado", domain.ErrFileNotFound.Error())
		assert.Equal(t, "acesso negado: caminho fora do diretório raiz permitido", domain.ErrPathEscape.Error())
		assert.Equal(t, "conteúdo markdown inválido", domain.ErrInvalidMarkdown.Error())
		assert.Equal(t, "falha ao escanear diretório", domain.ErrDirectoryScan.Error())
	})
}
