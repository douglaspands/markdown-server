package fs

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/ports"
)

// FSScanner implementa a porta ports.FileScannerPort para o sistema de arquivos local.
type FSScanner struct {
	absRootDir string
}

// NewFSScanner inicializa o scanner validando e resolvendo o caminho absoluto do diretório raiz.
func NewFSScanner(rootDir string) (*FSScanner, error) {
	if rootDir == "" {
		rootDir = "."
	}

	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, domain.ErrDirectoryScan
	}

	return &FSScanner{
		absRootDir: absPath,
	}, nil
}

// SanitizePath higieniza o caminho relativo e assegura que não ocorra Directory Traversal fora da raiz.
func (s *FSScanner) SanitizePath(relativePath string) (string, error) {
	cleaned := filepath.Clean(relativePath)
	cleaned = strings.TrimPrefix(cleaned, string(filepath.Separator))
	cleaned = strings.TrimPrefix(cleaned, "/")
	cleaned = strings.TrimPrefix(cleaned, "\\")

	targetPath := filepath.Join(s.absRootDir, cleaned)
	rel, err := filepath.Rel(s.absRootDir, targetPath)
	if err != nil || strings.HasPrefix(rel, "..") || rel == ".." {
		return "", domain.ErrPathEscape
	}

	return filepath.ToSlash(rel), nil
}

// ReadFile lê os bytes de um arquivo seguro no sistema de arquivos.
func (s *FSScanner) ReadFile(ctx context.Context, relativePath string) ([]byte, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}

	safePath, err := s.SanitizePath(relativePath)
	if err != nil {
		return nil, err
	}

	fullPath := filepath.Join(s.absRootDir, filepath.FromSlash(safePath))
	info, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrFileNotFound
		}
		return nil, err
	}
	if info.IsDir() {
		return nil, domain.ErrFileNotFound
	}

	return os.ReadFile(fullPath)
}

// ScanDirectory percorre recursivamente o diretório raiz e constrói a árvore hierárquica de FileNode.
func (s *FSScanner) ScanDirectory(ctx context.Context) (*domain.FileNode, error) {
	if ctx != nil && ctx.Err() != nil {
		return nil, ctx.Err()
	}

	rootNode := &domain.FileNode{
		Name:         filepath.Base(s.absRootDir),
		RelativePath: "",
		Type:         domain.NodeTypeDirectory,
		Children:     []*domain.FileNode{},
	}

	nodeMap := map[string]*domain.FileNode{
		".": rootNode,
		"":  rootNode,
	}

	err := filepath.WalkDir(s.absRootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}

		if path == s.absRootDir {
			return nil
		}

		// Ignora diretórios ocultos (ex: .git, .agent)
		if strings.HasPrefix(d.Name(), ".") && d.Name() != "." && d.Name() != ".." {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(s.absRootDir, path)
		if err != nil {
			return err
		}
		relPathSlash := filepath.ToSlash(relPath)

		isDir := d.IsDir()
		isMarkdown := false
		if !isDir {
			ext := strings.ToLower(filepath.Ext(d.Name()))
			isMarkdown = ext == ".md" || ext == ".markdown"
		}

		node := &domain.FileNode{
			Name:         d.Name(),
			RelativePath: relPathSlash,
			Type:         domain.NodeTypeFile,
			IsMarkdown:   isMarkdown,
		}
		if isDir {
			node.Type = domain.NodeTypeDirectory
			node.Children = []*domain.FileNode{}
			nodeMap[relPathSlash] = node
		}

		parentRel := filepath.ToSlash(filepath.Dir(relPath))
		if parentRel == "." || parentRel == "/" {
			parentRel = ""
		}

		if parentNode, exists := nodeMap[parentRel]; exists {
			parentNode.Children = append(parentNode.Children, node)
		} else {
			rootNode.Children = append(rootNode.Children, node)
		}

		return nil
	})

	if err != nil {
		return nil, domain.ErrDirectoryScan
	}

	sortFileTree(rootNode)
	return rootNode, nil
}

// sortFileTree ordena pastas primeiro e arquivos alfabeticamente.
func sortFileTree(node *domain.FileNode) {
	if node == nil || len(node.Children) == 0 {
		return
	}

	sort.Slice(node.Children, func(i, j int) bool {
		a, b := node.Children[i], node.Children[j]
		if a.Type != b.Type {
			return a.Type == domain.NodeTypeDirectory
		}
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	})

	for _, child := range node.Children {
		if child.Type == domain.NodeTypeDirectory {
			sortFileTree(child)
		}
	}
}

var _ ports.FileScannerPort = (*FSScanner)(nil)
