package ports

import (
	"context"

	"github.com/douglas/markdown-server/internal/core/domain"
)

// FileScannerPort define a porta de saída para operações de leitura e inspeção do sistema de arquivos local.
type FileScannerPort interface {
	// ReadFile lê os bytes de um arquivo relativo à raiz, garantindo que não haja escape da pasta.
	ReadFile(ctx context.Context, relativePath string) ([]byte, error)

	// ScanDirectory percorre recursivamente o diretório raiz e constrói a árvore de FileNode.
	ScanDirectory(ctx context.Context) (*domain.FileNode, error)

	// SanitizePath sanitiza o caminho relativo e valida se está dentro da raiz segura.
	SanitizePath(relativePath string) (string, error)
}

// MarkdownRendererPort define a porta de saída para conversão e estilização de Markdown em HTML.
type MarkdownRendererPort interface {
	// RenderHTML converte o conteúdo bruto Markdown em HTML semântico com highlight de código e reescrita de links.
	RenderHTML(ctx context.Context, markdown []byte, currentPath string) (renderedHTML string, title string, err error)
}

// BrowserLauncherPort define a porta de saída para abertura do navegador padrão do sistema operacional.
type BrowserLauncherPort interface {
	// OpenURL dispara a abertura de uma URL no navegador web padrão do usuário.
	OpenURL(ctx context.Context, url string) error
}
