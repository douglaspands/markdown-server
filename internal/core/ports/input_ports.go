package ports

import (
	"context"

	"github.com/douglas/markdown-server/internal/core/domain"
)

// MarkdownServicePort define a porta de entrada para os casos de uso de visualização de Markdown e navegação.
type MarkdownServicePort interface {
	// GetDocument busca, sanitiza e processa o documento Markdown no caminho relativo informado.
	GetDocument(ctx context.Context, relativePath string) (*domain.MarkdownDocument, error)

	// GetFileTree constrói a árvore hierárquica de arquivos e diretórios a partir da raiz.
	GetFileTree(ctx context.Context) (*domain.FileNode, error)
}

// HealthServicePort define a porta de entrada para verificação de saúde e diagnóstico da aplicação.
type HealthServicePort interface {
	// CheckHealth retorna o estado operacional corrente da aplicação.
	CheckHealth(ctx context.Context) (*domain.HealthStatus, error)
}
