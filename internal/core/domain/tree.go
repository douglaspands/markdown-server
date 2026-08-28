package domain

// NodeType indica se o nó na árvore é um arquivo ou diretório.
type NodeType string

const (
	NodeTypeFile      NodeType = "file"
	NodeTypeDirectory NodeType = "directory"
)

// FileNode representa um nó na árvore de navegação de arquivos do diretório raiz.
type FileNode struct {
	Name         string      `json:"name"`
	RelativePath string      `json:"relative_path"`
	Type         NodeType    `json:"type"`
	Children     []*FileNode `json:"children,omitempty"`
	IsMarkdown   bool        `json:"is_markdown"`
}
