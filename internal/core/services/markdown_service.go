package services

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/ports"
)

// MarkdownService implementa a porta ports.MarkdownServicePort orquestrando a leitura e renderização.
type MarkdownService struct {
	scanner  ports.FileScannerPort
	renderer ports.MarkdownRendererPort
}

// NewMarkdownService constrói uma nova instância do serviço com suas dependências explícitas.
func NewMarkdownService(scanner ports.FileScannerPort, renderer ports.MarkdownRendererPort) ports.MarkdownServicePort {
	return &MarkdownService{
		scanner:  scanner,
		renderer: renderer,
	}
}

// GetDocument busca, sanitiza e processa o arquivo Markdown retornando a entidade de domínio completa.
func (s *MarkdownService) GetDocument(ctx context.Context, relativePath string) (*domain.MarkdownDocument, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}

	trimmedPath := strings.TrimSpace(relativePath)
	trimmedPath = strings.TrimPrefix(trimmedPath, "/")

	// Se a rota for raiz, tenta carregar README.md ou readme.md
	if trimmedPath == "" || trimmedPath == "." {
		trimmedPath = "README.md"
	}

	safePath, err := s.scanner.SanitizePath(trimmedPath)
	if err != nil {
		return nil, err
	}

	content, err := s.scanner.ReadFile(ctx, safePath)
	if err != nil {
		// Se README.md falhar na raiz, tenta readme.md em minúsculas
		if trimmedPath == "README.md" {
			if lowerContent, lowerErr := s.scanner.ReadFile(ctx, "readme.md"); lowerErr == nil {
				content = lowerContent
				safePath = "readme.md"
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	htmlContent, title, err := s.renderer.RenderHTML(ctx, content, safePath)
	if err != nil {
		return nil, err
	}

	if title == "" {
		baseName := filepath.Base(safePath)
		title = strings.TrimSuffix(baseName, filepath.Ext(baseName))
	}

	return &domain.MarkdownDocument{
		RelativePath: safePath,
		Title:        title,
		RawContent:   content,
		HTMLContent:  htmlContent,
		SizeBytes:    int64(len(content)),
	}, nil
}

// GetFileTree delega a construção da árvore hierárquica para o scanner de arquivos.
func (s *MarkdownService) GetFileTree(ctx context.Context) (*domain.FileNode, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return s.scanner.ScanDirectory(ctx)
}
