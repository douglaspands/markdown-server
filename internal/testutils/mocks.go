package testutils

import (
	"context"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/ports"
)

// MockFileScanner é um mock determinístico para a porta ports.FileScannerPort.
type MockFileScanner struct {
	ReadFileFunc      func(ctx context.Context, relativePath string) ([]byte, error)
	ScanDirectoryFunc func(ctx context.Context) (*domain.FileNode, error)
	SanitizePathFunc  func(relativePath string) (string, error)
}

func (m *MockFileScanner) ReadFile(ctx context.Context, relativePath string) ([]byte, error) {
	if m.ReadFileFunc != nil {
		return m.ReadFileFunc(ctx, relativePath)
	}
	return nil, domain.ErrFileNotFound
}

func (m *MockFileScanner) ScanDirectory(ctx context.Context) (*domain.FileNode, error) {
	if m.ScanDirectoryFunc != nil {
		return m.ScanDirectoryFunc(ctx)
	}
	return &domain.FileNode{Name: "root", Type: domain.NodeTypeDirectory}, nil
}

func (m *MockFileScanner) SanitizePath(relativePath string) (string, error) {
	if m.SanitizePathFunc != nil {
		return m.SanitizePathFunc(relativePath)
	}
	return relativePath, nil
}

var _ ports.FileScannerPort = (*MockFileScanner)(nil)

// MockMarkdownRenderer é um mock determinístico para a porta ports.MarkdownRendererPort.
type MockMarkdownRenderer struct {
	RenderHTMLFunc func(ctx context.Context, markdown []byte, currentPath string) (string, string, error)
}

func (m *MockMarkdownRenderer) RenderHTML(ctx context.Context, markdown []byte, currentPath string) (string, string, error) {
	if m.RenderHTMLFunc != nil {
		return m.RenderHTMLFunc(ctx, markdown, currentPath)
	}
	return "<p>mock html</p>", "Mock Title", nil
}

var _ ports.MarkdownRendererPort = (*MockMarkdownRenderer)(nil)
