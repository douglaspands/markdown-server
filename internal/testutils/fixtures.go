package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// SampleMarkdownWithMermaid é uma fixture de Markdown contendo texto, tabela GFM, código Go e diagrama Mermaid.
const SampleMarkdownWithMermaid = `# Arquitetura do Sistema

Bem-vindo à documentação do **md_server**.

## Tabela de Componentes

| Componente | Tipo | Responsabilidade |
| :--- | :--- | :--- |
| CLI | Entrada | Ponto de entrada e parse de flags |
| HTTP Server | Entrada | Atende requisições web locais |
| Markdown Renderer | Saída | Converte Markdown para HTML |

## Exemplo de Código Go

` + "```go" + `
package main

import "fmt"

func main() {
    fmt.Println("Olá, md_server!")
}
` + "```" + `

## Diagrama de Arquitetura

` + "```mermaid" + `
flowchart TD
    User([Usuário]) --> Browser[Navegador Web]
    Browser --> Server[HTTP Server md_server]
    Server --> Renderer[Goldmark + Chroma]
    Server --> FS[(Sistema de Arquivos)]
` + "```" + `

Para mais informações, consulte o [Guia de Instalação](./installation.md).
`

// CreateTempDirWithFiles cria um diretório temporário para testes e popula com a árvore de arquivos passada.
func CreateTempDirWithFiles(t *testing.T, files map[string]string) string {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "md_server_test_*")
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = os.RemoveAll(tempDir)
	})

	for relPath, content := range files {
		fullPath := filepath.Join(tempDir, relPath)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		require.NoError(t, err)
		err = os.WriteFile(fullPath, []byte(content), 0644)
		require.NoError(t, err)
	}

	return tempDir
}
